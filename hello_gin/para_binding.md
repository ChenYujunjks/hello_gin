在 Express 中，处理 POST 请求通常是通过中间件来解析请求体，然后根据请求体的格式来处理数据。Express 本身没有特定于表单或 JSON 的请求处理方式，而是通过中间件来实现的。下面是如何在 Express 中处理这两种类型的数据：

### 1. 处理 JSON 数据

要处理 JSON 数据，需要使用 `express.json()` 中间件。

```javascript
const express = require('express');
const app = express();

app.use(express.json());

app.post('/add', (req, res) => {
    const { number1, number2 } = req.body;
    if (typeof number1 === 'undefined' || typeof number2 === 'undefined') {
        return res.status(400).json({ error: 'number1 and number2 are required' });
    }
    const result = number1 + number2;
    res.json({ result });
});

app.listen(3000, () => {
    console.log('Server is running on port 3000');
});
```

### 2. 处理表单数据

要处理表单数据，需要使用 `express.urlencoded()` 中间件。

```javascript
const express = require('express');
const app = express();

app.use(express.urlencoded({ extended: true }));

app.post('/subtract', (req, res) => {
    const { number1, number2 } = req.body;
    if (!number1 || !number2) {
        return res.status(400).json({ error: 'number1 and number2 are required' });
    }
    const num1 = parseFloat(number1);
    const num2 = parseFloat(number2);
    if (isNaN(num1) || isNaN(num2)) {
        return res.status(400).json({ error: 'invalid number format' });
    }
    const result = num1 - num2;
    res.json({ result });
});

app.listen(3000, () => {
    console.log('Server is running on port 3000');
});
```

### 3. 同时处理 JSON 和表单数据

你可以同时使用 `express.json()` 和 `express.urlencoded()` 来处理这两种类型的数据。然后在路由处理函数中检测请求体的格式并进行相应处理。

```javascript
const express = require('express');
const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.post('/add', (req, res) => {
    let number1, number2;

    if (req.headers['content-type'] === 'application/json') {
        ({ number1, number2 } = req.body);
    } else {
        ({ number1, number2 } = req.body);
        number1 = parseFloat(number1);
        number2 = parseFloat(number2);
    }

    if (typeof number1 === 'undefined' || typeof number2 === 'undefined') {
        return res.status(400).json({ error: 'number1 and number2 are required' });
    }

    if (isNaN(number1) || isNaN(number2)) {
        return res.status(400).json({ error: 'invalid number format' });
    }

    const result = number1 + number2;
    res.json({ result });
});

app.listen(3000, () => {
    console.log('Server is running on port 3000');
});
```

### 总结

- 使用 `express.json()` 中间件来处理 JSON 格式的请求体。
- 使用 `express.urlencoded({ extended: true })` 中间件来处理 URL 编码的表单数据。
- 可以同时使用这两个中间件来处理不同类型的数据，并在路由处理函数中根据请求头或数据格式来进行相应处理。

Express 的灵活性允许你根据需求来配置中间件，从而适应不同的数据提交方式。


***
### post表单提交至 JSON绑定路由

你的问题是由于你的表单使用`method="post"`发送数据，这将以`application/x-www-form-urlencoded`格式发送数据，而不是以JSON格式发送数据。你的处理函数期望的是JSON格式的数据，因此会出现解析错误。

为了使表单数据以JSON格式发送，你需要使用JavaScript来处理表单提交，并通过AJAX请求发送JSON数据。下面是修改后的`add.html`和相应的JavaScript代码：

**templates/add.html**:
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Addition</title>
</head>
<body>
    <h1>Addition</h1>
    <form id="addForm">
        <label for="number1">Number 1:</label>
        <input type="number" id="number1" name="number1" required>
        <br>
        <label for="number2">Number 2:</label>
        <input type="number" id="number2" name="number2" required>
        <br>
        <button type="submit">Add</button>
    </form>

    <script>
        document.getElementById('addForm').addEventListener('submit', function(event) {
            event.preventDefault(); // Prevent the form from submitting the traditional way

            const number1 = document.getElementById('number1').value;
            const number2 = document.getElementById('number2').value;

            fetch('/add', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ number1: parseFloat(number1), number2: parseFloat(number2) }),
            })
            .then(response => response.json())
            .then(data => {
                // Handle the response data
                alert(`Result: ${data.result}`);
            })
            .catch(error => {
                console.error('Error:', error);
            });
        });
    </script>
</body>
</html>
```

使用这个修改后的HTML文件，当你提交表单时，JavaScript代码会拦截表单的默认提交行为，将表单数据转换为JSON格式，并通过AJAX请求发送到服务器。这样，服务器端的处理函数就可以正常处理JSON数据，而不会出现解析错误。

确保服务器端的代码如下：

```go
r.POST("/add", func(c *gin.Context) {
    var request AdditionRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result := request.Number1 + request.Number2
    c.JSON(http.StatusOK, gin.H{"result": result})
})
```

这样，你的`add`路由函数应该可以正确处理表单提交的数据，并返回计算结果。

---
### POST 请求 

好的，深入了解以`application/x-www-form-urlencoded`格式和JSON格式发送数据的区别可以帮助你更好地理解HTTP请求的内部工作原理。

### 1. 发送POST请求

确实，`application/x-www-form-urlencoded`和JSON格式的数据都是通过POST请求发送的，但它们的内容类型和数据格式不同。

### 2. HTTP请求体的区别

#### `application/x-www-form-urlencoded`格式

这种格式的请求体是键值对的形式，并且键和值都是URL编码的。

##### 举例：add表单的HTTP请求体

假设表单的数据如下：

- Number1: 5
- Number2: 3

表单提交时会发送如下HTTP请求：

```
POST /add HTTP/1.1
Host: example.com
Content-Type: application/x-www-form-urlencoded
Content-Length: 19

number1=5&number2=3
```

在这种格式下，数据被编码成键值对的形式，并且多个键值对通过`&`符号连接。

#### JSON格式

这种格式的请求体是JSON字符串。

##### 举例：add表单的HTTP请求体

假设数据如下：

- Number1: 5
- Number2: 3

通过JavaScript发送AJAX请求时，会发送如下HTTP请求：

```
POST /add HTTP/1.1
Host: example.com
Content-Type: application/json
Content-Length: 24

{
    "number1": 5,
    "number2": 3
}
```

在这种格式下，数据被编码成JSON格式的字符串。

### 总结

- **`application/x-www-form-urlencoded`格式**：数据以键值对的形式编码，键和值都经过URL编码。适合传统HTML表单提交。
- **JSON格式**：数据以JSON字符串的形式编码，更加结构化，适合现代Web应用程序通过JavaScript发送的请求。

选择哪种格式取决于你的应用需求。如果你在处理传统HTML表单提交，`application/x-www-form-urlencoded`格式是默认和方便的。但如果你需要传递复杂的结构化数据，JSON格式更加合适。


---
### 阻止表单默认提交
是的，React 和 Express 全栈项目中常常会提到阻止表单的默认提交行为，这是因为通常需要使用JavaScript（比如通过`fetch`或`axios`）来处理表单数据的提交，以便更灵活地处理请求和响应。

### 为什么阻止表单默认提交？

1. **灵活处理数据**：通过AJAX请求，你可以在提交前对表单数据进行验证、处理或转换，例如将数据转换为JSON格式。
2. **提高用户体验**：AJAX请求是异步的，不会刷新页面，这样用户体验更好。例如，提交表单后可以直接在页面上显示响应结果，而不需要重新加载整个页面。
3. **更好的错误处理**：通过AJAX请求，你可以更好地处理错误并在页面上显示具体的错误信息，而不是简单地跳转到一个错误页面。
4. **统一的接口**：在前后端分离的架构中，前端通常通过AJAX请求与后端API接口通信，这样可以保持前端和后端的解耦，便于维护和扩展。

### 示例：React 表单提交

以下是一个使用React和Express的简单示例，展示如何处理表单的AJAX提交：

#### 前端 (React)

```javascript
import React, { useState } from 'react';

function App() {
  const [number1, setNumber1] = useState('');
  const [number2, setNumber2] = useState('');
  const [result, setResult] = useState(null);

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const response = await fetch('/api/add', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ number1: parseFloat(number1), number2: parseFloat(number2) }),
      });
      const data = await response.json();
      setResult(data.result);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (
    <div>
      <h1>Addition</h1>
      <form onSubmit={handleSubmit}>
        <label>
          Number 1:
          <input
            type="number"
            value={number1}
            onChange={(e) => setNumber1(e.target.value)}
            required
          />
        </label>
        <br />
        <label>
          Number 2:
          <input
            type="number"
            value={number2}
            onChange={(e) => setNumber2(e.target.value)}
            required
          />
        </label>
        <br />
        <button type="submit">Add</button>
      </form>
      {result !== null && <p>Result: {result}</p>}
    </div>
  );
}

export default App;
```

#### 后端 (Express)

```javascript
const express = require('express');
const bodyParser = require('body-parser');
const app = express();

app.use(bodyParser.json());

app.post('/api/add', (req, res) => {
  const { number1, number2 } = req.body;
  if (typeof number1 === 'number' && typeof number2 === 'number') {
    const result = number1 + number2;
    res.json({ result });
  } else {
    res.status(400).json({ error: 'Invalid input' });
  }
});

const port = process.env.PORT || 5000;
app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
```

### 总结

通过使用AJAX请求（如`fetch`或`axios`），你可以更灵活地处理表单数据的提交，提高用户体验，并且可以在提交前对数据进行处理和验证。这也是为什么在现代前后端分离的架构中，表单的默认提交行为通常会被阻止并改用AJAX请求来处理数据的原因。