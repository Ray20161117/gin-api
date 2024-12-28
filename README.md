# go gin项目

## 项目初始化及安装依赖

### 项目初始化
```
1.创建项目: gin-api
2.进入项目目录在终端命令Terminal中初始化项目：go mod init gin-api
```
### 安装依赖
```
在终端命令Terminal中安装依赖
先安装项目依赖包go.mod: go mod init admin-api
安装gin命令: go get github.com/gin-gonic/gin@v1.8.1
安装gorm命令: go get gorm.io/gorm
安装mysql命令: go get gorm.io/driver/mysql
安装log命令:
    go get github.com/sirupsen/logrus
    go get github.com/lestrrat-go/file-rotatelogs
    go get github.com/rifflock/lfshook
安装go-redis命令: go get github.com/go-redis/redis/v8@v8.11.5
安装base64Captcha命令: go get github.com/mojocn/base64Captcha@v1.3.1
安装jwt-go命令: go get github.com/dgrijalva/jwt-go
安装yaml命令: go get gopkg.in/yaml.v3
安装获取客户端OS和browser命令: go get -u github.com/wenlng/go-user-agent
安装ip地址命令: go get github.com/gogf/gf
安装swagger命令:
go get github.com/swaggo/files 
go get github.com/swaggo/gin-swagger
```
### 项目结构
```
可以根据项目需要设置相应的目录结构，以下是gin-api项目的目录结构：
gin-api
├── common                    公共目录
│   ├── utils                 工具目录
│   └── response              响应结果目录
├── config                    配置目录
│   ├── constant              常量配置目录
│   └── yaml_config           读取yaml配置目录
├── controllers               控制器目录
├── docs                      swagger文档目录
├── middlewares               中间件目录
├── models                    模型目录
├── pkg                       公共包目录
│   ├── jwt                   jwt工具目录
│   ├── log                   log工具目录
│   ├── mysql                 mysql工具目录
│   └── redis                 redis工具目录
├── routes                    路由目录
│   └── routes.go             路由文件
├── services                  服务目录, 通常负责业务逻辑
├── utils                     工具目录
├── config.yaml               基础配置 
├── main.go                   项目入口文件
├── go.mod                    项目依赖包
├── go.sum                    项目依赖包 
└── README.md                 说明文档
```

#### 注意事项
```
1.预使用swagger工具生成API文档，需要安装swagger相关依赖包，具体安装方法请参考上面的安装依赖步骤。
2.项目中使用了gin-swagger，需要在main.go(本项目在router文件中引入)文件中导入相关包，具体使用方法请参考gin-swagger官方文档。

_ "gin-api/docs"
swaggerFiles "github.com/swaggo/files"
ginSwagger "github.com/swaggo/gin-swagger"
3.通过swag init命令生成API文档
```