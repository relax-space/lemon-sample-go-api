# sample-go-api
an go-api-based sample 

# feature
- url:https://staging.p2shop.com.cn/sample/fruits
- swagger 2.0 :https://staging.p2shop.com.cn/sample/docs/index.html
- jwt url:https://staging.p2shop.com.cn/sample/v1/fruits


# tips

## Before submitting the code, please download the go-api code
```
git remote add a-upstream https://gitlab.p2shop.cn:8443/sample/go-api.git
git fetch a-upstream master
git rebase upstream/master a-upstream/master
```

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
