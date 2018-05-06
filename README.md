# sample-go-api template

You can quickly create an echo-based api project 

# supports
- behavior log
- swagger 2.0
- jwt
- context parameter
- latest ci script

## Getting Started

Get source
```
$ git clone https://gitlab.p2shop.cn:8443/sample/sample-go-api.git
```
Rename the outermost folder to your project name  
Rename sample-api-go to your project name

Run
```
$ cd $GOPATH/src/sample-go-api
$ go run main.go
```

Visit           http://127.0.0.1:8080/  
Visit swagger   http://127.0.0.1:8080/docs/index.html

## Sample
You can see the latest sample here
```
https://gitlab.p2shop.cn:8443/sample/sample-go-api.git
```


## References

- web framework: [echo framework](https://echo.labstack.com/)
- orm tool: [xorm](http://xorm.io/)
- logger : [logrus](https://github.com/sirupsen/logrus)
- configuration tool: [viper](https://github.com/spf13/viper)
- validator: [govalidator](github.com/asaskevich/govalidator)
- utils: https://github.com/pangpanglabs/goutils
