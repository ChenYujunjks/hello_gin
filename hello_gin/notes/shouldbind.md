### 示例代码
下面的示例代码展示了如何使用`AdditionRequest`结构体来绑定查询参数、表单参数和JSON参数。

```go
package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 定义用于绑定的结构体
type AdditionRequest struct {
	Number1 float64 `form:"number1" json:"number1" binding:"required"`
	Number2 float64 `form:"number2" json:"number2" binding:"required"`
}

func main() {
	gin.SetMode(gin.DebugMode) // 设置Gin的运行模式为DebugMode

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// 查询参数绑定
	r.GET("/multiply", func(c *gin.Context) {
		var query AdditionRequest

		if err := c.ShouldBind(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := query.Number1 * query.Number2
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	// 表单参数绑定
	r.POST("/add", func(c *gin.Context) {
		var form AdditionRequest

		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := form.Number1 + form.Number2
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	// JSON 参数绑定
	r.POST("/json_add", func(c *gin.Context) {
		var json AdditionRequest
		if err := c.ShouldBind(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := json.Number1 + json.Number2
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	// 路径参数绑定
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"user_id": id,
		})
	})

	r.Run(":8080")
}
```

### 解释
1. **查询参数绑定**：
   使用`ShouldBind()`来绑定查询参数，通过在`AdditionRequest`结构体中添加`form`标签来指定参数名。

2. **表单参数绑定**：
   使用`ShouldBind()`来绑定表单参数，通过在`AdditionRequest`结构体中添加`form`标签来指定参数名。

3. **JSON参数绑定**：
   使用`ShouldBind()`来绑定JSON请求体，通过在`AdditionRequest`结构体中添加`json`标签来指定参数名。

### 其他示例
以下是不同类型的请求示例：

#### 查询参数示例
```sh
curl "http://localhost:8080/multiply?number1=5&number2=3"
```

#### 表单参数示例
```sh
curl -X POST -F "number1=5" -F "number2=3" "http://localhost:8080/add"
```

#### JSON 参数示例
```sh
curl -X POST -H "Content-Type: application/json" -d '{"number1":5, "number2":3}' "http://localhost:8080/json_add"
```

#### 路径参数示例
```sh
curl "http://localhost:8080/user/123"
```

通过这种方式，你可以复用同一个结构体来处理各种形式的请求参数，使代码更加简洁和可维护。