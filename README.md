# This is a gin-like framework demo

- [Web framework](#web-framework)
    - [router](#router)
    - [middleware](#middleware)
    - [html rendering](#html-rendering)
    
## router
```go
func main() {
    v1 := r.Group("/v1")
    {
        v1.GET("/hello", func(c *gee.Context) {
            c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
        }
    )
}
```

## middleware
```go
func main() {
    // Creates a router without any middleware by default
    r := gin.New()

    // Global middleware
    r.Use(gee.Logger())

    r.Use(gee.Recovery())

    r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
}
```

## html-rendering

Using LoadHTMLGlob() or LoadHTMLFiles()

```go
func main() {
    r := gee.New()
    r.Use(middleware.Logger())
    r.SetFuncMap(template.FuncMap{
        "FormatAsDate": FormatAsDate,
    })
    r.LoadHTMLGlob("templates/*")
    r.Static("/assets", "./static")
}
```

## 仅个人学习
### 实现参考项目
https://github.com/gin-gonic/gin

https://github.com/geektutu/7days-golang/tree/master/gee-web

