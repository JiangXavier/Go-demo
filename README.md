# Go-demo
## blog
- 项目使用Gin+MySQL+Gorm，采用RESTful API
- 配置管理使用viper，使用fsnotify配置热更新
- 日志管理使用lumberjack
- 使用Swagger生成接口文档
- 使用validator进行参数绑定和校验，同时借助语言包进行国际化处理
- 支持文件上传并对外提供静态资源的访问服务
- 使用JWT对接口进行访问控制
- 自定义Recovery并实现邮件报警处理
- 使用Ratelimit进行接口限流控制
- 使用Jaeger分布式链路追踪系统进行链路追踪并同时追踪日志和SQL
- 项目结构

```shell
blog-service
├── configs
├── docs
├── global
├── internal
│   ├── dao
│   ├── middleware
│   ├── model
│   ├── routers
│   └── service
├── pkg
├── storage
├── scripts
└── third_party
```
- configs：配置文件。
- docs：文档集合。
- global：全局变量。
- internal：内部模块。
  - dao：数据访问层（Database Access Object），所有与数据相关的操作都会在 dao 层进行，例如 MySQL、ElasticSearch 等。
  - middleware：HTTP 中间件。
  - model：模型层，用于存放 model 对象。
  - routers：路由相关逻辑处理。
  - service：项目核心业务逻辑。
- pkg：项目相关的模块包。
- storage：项目生成的临时文件。
- scripts：各类构建，安装，分析等操作的脚本。
- third_party：第三方的资源工具，例如 Swagger UI
## spider
- 爬取静态数据--豆瓣电影Top250
- 爬取动态数据--B站电影评分
- goroutine并发爬虫
