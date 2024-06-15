HTML表单中的 `<input type="file">` 元素提供了一个用户友好的文件上传界面。以下是其工作原理和为什么你会在浏览器中看到“选择文件”按钮的解释：

1. **HTML结构**：你的HTML表单包含一个 `<input type="file">` 元素。这个输入类型专门用于文件选择和上传。你提供的代码正确地创建了一个文件上传表单：
    ```html
    <!DOCTYPE html>
    <html lang="zh-CN">
    <head>
        <title>上传文件示例</title>
    </head>
    <body>
    <form action="/upload" method="post" enctype="multipart/form-data">
        <input type="file" name="f1">
        <input type="submit" value="上传">
    </form>
    </body>
    </html>
    ```

2. **浏览器行为**：“选择文件”按钮是浏览器默认提供的界面元素，用于让用户选择他们本地文件系统中的文件。当你使用 `<input type="file">` 时，浏览器会自动渲染一个标有“选择文件”（或浏览器语言的等效内容）的按钮，旁边还有一个文本输入框显示所选文件的名称。

3. **文件选择**：当你点击“选择文件”按钮时，浏览器会打开一个文件对话框，让你浏览并选择计算机上的文件。这个交互过程完全由浏览器处理。

4. **文件上传**：一旦选择了文件，表单就可以提交。表单标签中的 `enctype="multipart/form-data"` 属性指定表单数据应以 `multipart/form-data` 进行编码。这个编码类型是文件上传所必需的。当你提交表单时，浏览器会将文件数据和其他表单数据打包，通过HTTP POST请求发送到服务器。

5. **后端处理**：在服务器端，后端（使用Gin框架编写）会处理传入的请求。服务器代码需要解析multipart表单数据，以访问上传的文件。