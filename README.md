**适合新手的golang后端开发脚手架**

技术栈：golang+gin+gorm
数据库：mysql
已集成中间件功能：输出日志到文件

```markdown
项目结构：
config/:配置文件
controller/:负责处理 HTTP 请求和响应
middleware/:放置封装的中间件
models/:定义数据模型、数据库结构和数据操作
  init.go:链接数据库
request/:负责请求格式
response/:负责响应格式
  init.go:接口返回格式
route/:负责定义路由规则
service/:用于处理业务逻辑和协调数据操作
main.go:程序启动入口
```

