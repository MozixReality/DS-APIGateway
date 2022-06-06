# API Getway
## Quick Start
### Select the ENV
```
export ENV=compliance
```
swagger document can only start in local or compliance ENV <br>

### Compile
```
make install
```
Will generate APIGetway in the directory. <br>
If you encounter /bin/bash: swag: command not found, try the following method
1. To go/bin directory
2. type ```PATH=$(go env GOPATH)/bin:$PATH```


To start server, it should only execute the APIGetway like
```
./APIGetway
```

# API
## Swagger
After starting the server, the swagger document can see at ```https://mozixreality.ebg.tw/swagger/index.html``` or ```https://localhost:55688/swagger/index.html```