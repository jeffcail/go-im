package auth

import "fmt"

// emailTemplate 注册邮件模版
func emailTemplate(emailCode string) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>IM注册邮件</title>
</head>
<style>
    .mail{
        margin: 0 auto;
        border-radius: 45px;
        height: 400px;
        padding: 10px;
        background-color: #CC9933;
    }
    .code {
        color: #f6512b;
        font-weight: bold;
        font-size: 30px;
        padding: 2px;
    }
</style>
<body>
<div class="mail">
    <h3>尊敬的用户您好:您正在注册im应用账号!</h3>
    <p>下面是您的验证码:</p>
        <p class="code">%s</p>
        <p>请注意查收!验证码有效期为5分钟。谢谢</p>
</div>
</body>
</html>`, emailCode)
}
