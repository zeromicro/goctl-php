# goctl-php

![go-zero](https://img.shields.io/badge/Github-go--zero-brightgreen?link=https://github.com/tal-tech/go-zero&logo=github)
![License](https://img.shields.io/badge/License-MIT-blue?link=https://github.com/zeromicro/goctl-android/blob/main/LICENSE)
![Go](https://github.com/zeromicro/goctl-android/workflows/Go/badge.svg)

goctl-php是一款基于goctl的插件，用于生成 php 调用端（服务端） http server请求代码。
本插件特性：
* 仅支持post json
* 支持get query参数
* 支持path路由变量
* 仅支持响应体为json

> 警告：本插件是对goctl plugin开发流程的指引，切勿用于生产环境。

# 插件使用
* 编译goctl-php插件
    ```shell script
    $ GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/zeromicro/goctl-php
    ```
* 将`$GOPATH/bin`中的`goctl-php`添加到环境变量
* 创建api文件
    ```go
    info(
    	title: "type title here"
    	desc: "type desc here"
    	author: "type author here"
    	email: "type email here"
    	version: "type version here"
    )
    
    
    type (
    	RegisterReq {
    		Username string `json:"username"`
    		Password string `json:"password"`
    		Mobile string `json:"mobile"`
    	}
    	
    	LoginReq {
    		Username string `json:"username"`
    		Password string `json:"password"`
    	}
    )
    
    service user-api {
    	@doc(
    		summary: "注册"
    	)
    	@handler register
    	post /api/user/register (RegisterReq)
    	
    	@doc(
    		summary: "登录"
    	)
    	@handler login
    	post /api/user/login (LoginReq)
    }
    ```
* 生成php代码
    
    ```shell script
    $ goctl api plugin -plugin goctl-php="php -package Tal" -api user.api -dir .
    ```
    >说明： 其中`goctl-php`为可执行的二进制文件，`"php -package Tal"`为goctl-plugin自定义的参数，这里需要用引号`""`引起来。

我们来看一下生成php代码后的目录结构
```text
├── bean
│   ├── LoginReq.php
│   ├── RegisterReq.php
├── service
│   ├── IService.php
│   └── Service.php
└── user.api
```

> [点击这里]() 查看php示例源码

composer依赖
```txt
//TODO
```

> 本插件是基于***来实现http请求，因此会用到一些php依赖，composer包管理形式自行处理。

* 编写测试
```php
//todo
```

* 请求结果
    * client log
    ```text
    register success
    login success
    search:{"age":20,"birthday":"1991-01-01","description":"coding now","name":"zeromicro","tag":["Golang","Php"]}
    userInfo:{"age":20,"birthday":"1991-01-01","description":"coding now","name":"zeromicro","tag":["Golang","Php"]}
    ```
    
    * server log
    ```text
    Login: {Username:zeromicro Password:111111}
    Register: {Username:zeromicro Password:1111 Mobile:12311111111}
    ```

# 插件开发流程

* 自定义参数
    ```go
    commands = []*cli.Command{
        {
            Name:   "android",
            Usage:  "generates http client for android",
            Action: action.Android,
            Flags: []cli.Flag{
                &cli.StringFlag{
                    Name:  "package",
                    Usage: "the package of android",
                },
            },
        },
    }
    ```

* 获取goctl传递过来的json信息
    * 利用goctl中提供的方法解析
        ```go
        plugin, err := plugin.NewPlugin()
        if err != nil {
            return err
        }
        ```
  
    * 或者自定义结构体去反序列化
  
        ```go
        var plugin generate.Plugin
        plugin.ParentPackage = pkg
        err = json.Unmarshal(std, &plugin)
        if err != nil {
            return err
        }
        ```

* 实现插件逻辑
    ```go
    generate.Do(plugin)
    ```

>说明：上述摘要代码来自goctl-php,完整信息可浏览源码。

