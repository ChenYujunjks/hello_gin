`gin.Context` 是 Gin 框架中的核心类型，它提供了许多方法来处理 HTTP 请求和响应。以下是一些常用的方法及其简要介绍：

### 请求相关的方法

1. **Query**: 获取 URL 查询参数
   ```go
   value := c.Query("key")
   ```

2. **DefaultQuery**: 获取 URL 查询参数，如果不存在则返回默认值
   ```go
   value := c.DefaultQuery("key", "defaultValue")
   ```

3. **PostForm**: 获取表单参数
   ```go
   value := c.PostForm("key")
   ```

4. **DefaultPostForm**: 获取表单参数，如果不存在则返回默认值
   ```go
   value := c.DefaultPostForm("key", "defaultValue")
   ```

5. **Param**: 获取路径参数
   ```go
   value := c.Param("name")
   ```

6. **BindJSON**: 绑定 JSON 请求体到结构体
   ```go
   var json struct {
       Key string `json:"key"`
   }
   if err := c.BindJSON(&json); err != nil {
       // 处理错误
   }
   ```

7. **ShouldBind**: 根据请求内容类型自动选择绑定方法
   ```go
   var form struct {
       Key string `form:"key"`
   }
   if err := c.ShouldBind(&form); err != nil {
       // 处理错误
   }
   ```

### 响应相关的方法

1. **JSON**: 返回 JSON 响应
   ```go
   c.JSON(http.StatusOK, gin.H{"message": "pong"})
   ```

2. **XML**: 返回 XML 响应
   ```go
   c.XML(http.StatusOK, gin.H{"message": "pong"})
   ```

3. **String**: 返回纯文本响应
   ```go
   c.String(http.StatusOK, "Hello, %s", name)
   ```

4. **HTML**: 渲染 HTML 模板并返回响应
   ```go
   c.HTML(http.StatusOK, "template.tmpl", gin.H{"title": "Main website"})
   ```

5. **Data**: 返回二进制数据
   ```go
   c.Data(http.StatusOK, "application/octet-stream", data)
   ```

6. **File**: 返回文件
   ```go
   c.File("path/to/file")
   ```

7. **Redirect**: 重定向到指定 URL
   ```go
   c.Redirect(http.StatusMovedPermanently, "http://example.com")
   ```

### 上下文控制和其他方法

1. **Abort**: 终止当前请求的执行链
   ```go
   c.Abort()
   ```

2. **AbortWithStatus**: 终止请求并返回状态码
   ```go
   c.AbortWithStatus(http.StatusUnauthorized)
   ```

3. **AbortWithStatusJSON**: 终止请求并返回状态码和 JSON 响应
   ```go
   c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
   ```

4. **Set**: 设置上下文中的键值对
   ```go
   c.Set("key", "value")
   ```

5. **Get**: 获取上下文中的键值对
   ```go
   if value, exists := c.Get("key"); exists {
       // 处理 value
   }
   ```

6. **MustGet**: 获取上下文中的键值对，如果不存在则抛出错误
   ```go
   value := c.MustGet("key")
   ```

### 示例代码

下面是一个完整的示例，展示了如何使用这些方法：

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

    r.POST("/submit", func(c *gin.Context) {
        var json struct {
            Name string `json:"name"`
        }
        if err := c.BindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"name": json.Name})
    })

    r.GET("/user/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.String(http.StatusOK, "Hello, %s", name)
    })

    r.GET("/search", func(c *gin.Context) {
        query := c.DefaultQuery("q", "default")
        c.JSON(http.StatusOK, gin.H{"result": "Search result for " + query})
    })

    r.Run(":8080")
}
```

通过这些方法，你可以灵活地处理各种类型的 HTTP 请求和响应，满足不同的应用需求。