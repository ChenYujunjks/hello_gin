### GET请求的HTTP示例

#### 请求行
```
GET /submit-form?name=John+Doe&email=johndoe%40example.com HTTP/1.1
```

#### 请求头部
```
Host: example.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36
```

#### 请求正文
GET请求通常不包含请求正文。


1. **GET请求**：GET请求将参数直接附加到URL中，不包含请求正文。常用于从服务器获取数据。
2. **POST请求**：通过JavaScript阻止表单默认提交行为，使用Fetch API发送JSON格式的POST请求。请求头中设置`Content-Type`为`application/json`，请求正文为JSON字符串。


### 普通表单提交与JSON提交的代码对比

#### 普通表单提交（默认行为）

**HTML部分**
```html
<!DOCTYPE html>
<html>
<head>
    <title>Form Submission Example</title>
</head>
<body>
    <form action="/submit-form" method="post">
        <label for="name">Name:</label>
        <input type="text" id="name" name="name"><br><br>
        <label for="email">Email:</label>
        <input type="email" id="email" name="email"><br><br>
        <input type="submit" value="Submit">
    </form>
</body>
</html>
```

**请求示例**
```
POST /submit-form HTTP/1.1
Host: example.com
Content-Type: application/x-www-form-urlencoded
Content-Length: 27

name=John+Doe&email=johndoe%40example.com
```

#### JSON提交（使用JavaScript阻止默认提交）

**HTML部分**
```html
<!DOCTYPE html>
<html>
<head>
    <title>JSON Form Submission Example</title>
</head>
<body>
    <form id="jsonForm">
        <label for="name">Name:</label>
        <input type="text" id="name" name="name"><br><br>
        <label for="email">Email:</label>
        <input type="email" id="email" name="email"><br><br>
        <input type="submit" value="Submit">
    </form>

    <script>
        document.getElementById('jsonForm').addEventListener('submit', function(event) {
            event.preventDefault(); // 阻止表单默认提交

            const formData = {
                name: document.getElementById('name').value,
                email: document.getElementById('email').value
            };

            fetch('/submit-json', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(formData)
            })
            .then(response => response.json())
            .then(data => console.log(data))
            .catch(error => console.error('Error:', error));
        });
    </script>
</body>
</html>
```

**请求示例**
```
POST /submit-json HTTP/1.1
Host: example.com
Content-Type: application/json
Content-Length: 44

{
    "name": "John Doe",
    "email": "johndoe@example.com"
}
```

### 对比总结

1. **HTML部分**：
   - 普通表单提交直接通过表单的`action`属性和`method`属性指定提交地址和方法。
   - JSON提交使用JavaScript监听表单提交事件，阻止默认提交行为，通过`fetch`方法发送AJAX请求。

2. **请求类型**：
   - 普通表单提交的数据使用`application/x-www-form-urlencoded`格式。
   - JSON提交的数据使用`application/json`格式。

3. **请求正文**：
   - 普通表单提交将数据编码为键值对的形式。
   - JSON提交将数据序列化为JSON字符串。