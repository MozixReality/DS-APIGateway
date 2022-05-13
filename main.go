package main

import (
	"APIGateway/constant"
	"APIGateway/routes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	mw := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "./log/log",
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	log.SetOutput(mw)
	constant.ReadConfig(".env")
}

func main() {
	gin.SetMode(viper.GetString("RUN_MODE"))
	port := viper.GetString("PORT")
	routesInit := routes.InitRouter()
	server := &http.Server{
		Addr:           port,
		Handler:        routesInit,
		ReadTimeout:    time.Duration(viper.GetInt("READ_TIMEOUT")) * time.Second,
		WriteTimeout:   time.Duration(viper.GetInt("WRITE_TIMEOUT")) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		log.Println("[info] start http server listening", port)
		return server.ListenAndServe()
	})

	g.Go(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGILL, syscall.SIGFPE)
		select {
		case <-ctx.Done():
			return server.Shutdown(ctx)
		case s := <-c:
			close(c)
			if err := server.Shutdown(ctx); err != nil {
				return err
			}
			return fmt.Errorf("os signal: %v", s)
		}
	})

	if err := g.Wait(); err != nil {
		log.Println("[error]", err)
	}
	log.Println("[info] HTTP Server Exited")
}
