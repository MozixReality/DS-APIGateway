package constant

import (
	"log"

	"github.com/spf13/viper"
)

func ReadConfig(configPath string) {
	viper.SetConfigFile(configPath)
	viper.AddConfigPath(".")

	viper.SetDefault("PORT", ":55688")
	viper.SetDefault("RUN_MODE", "release")
	viper.SetDefault("READ_TIMEOUT", 180)
	viper.SetDefault("WRITE_TIMEOUT", 60)
	viper.SetDefault("REQUEST_TIMEOUT", 60)

	envs := []string{
		"PORT",
		"RUN_MODE",
		"READ_TIMEOUT",
		"WRITE_TIMEOUT",
		"REQUEST_TIMEOUT",
	}

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			log.Println(err)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}
}
