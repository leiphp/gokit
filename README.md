# gokit
Go 常用工具类 SDK，可以把它理解为一个「组件库」或者「工具集合」，用于在你自己的项目中快速复用常用的第三方功能，比如七牛云、日志封装、链路追踪、Redis、数据库等。集成简单，功能完善中，持续更新！

## 使用
#### 启用 Go module【推荐】
在您的项目中的go.mod文件内添加这行代码
```go
require github.com/leiphp/gokit v1.0.0
```

并且在项目中使用 "github.com/leiphp/gokit"

例如
```go
import (
    "github.com/leiphp/gokit/pkg/sdk/qiniu"
)
```

## 文件结构如下：  
// gokit/  
// ├── go.mod  
// ├── cmd/            # 可执行程序（可选）  
// │   └── demo/  
// │       └── main.go  
// ├── pkg/            # 可供外部使用的包（工具库核心）  
// │   ├── core/  
// │   │   ├── logger/  
// │   │   │   └── logger.go  
// │   │   └── config/  
// │   │       └── config.go  
// │   ├── sdk/  
// │   │   └── qiniu/  
// │   │       ├── config.go  
// │   │       └── qiniu.go  
// ├── internal/       # 内部使用模块  
// │   └── encryption/  
// │       └── secret.go  
// ├── utils/          # 公共工具函数（字符串、时间等）  
// │   └── string.go  
// └── example/        # 示例代码  
//     └── usage.go  