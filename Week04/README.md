# Week04 作业题目
按照自己的构想，写一个项目, 满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

### 项目目录
```shell script
homework
├── api
│   └── article
│       └── v1
│           ├── article.pb.go
│           └── article.proto
├── cmd
│   └── week04
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
├── go.mod
├── go.sum
├── internal
│   ├── biz
│   │   └── article.go
│   ├── data
│   │   └── article.go
│   ├── pkg
│   │   └── server
│   │       └── server.go
│   └── service
│       └── article.go
└── test
    └── main.go

12 directories, 12 files

```

### 参考文档:
- [quick start](https://github.com/google/wire/blob/master/_tutorial/README.md)
- [Golang依赖注入框架wire使用详解](https://blog.csdn.net/uisoul/article/details/108776073)
- [example](https://github.com/google/go-cloud/tree/master/samples/guestbook) 