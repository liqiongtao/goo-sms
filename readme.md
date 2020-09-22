# 发送

```
conf := goo_sms.AliyunConfig{
    Region:       "",
    Appid:        "",
    Secret:       "",
    SignName:     "",
    TemplateCode: "",
}
code, err := goo_sms.New(goo_sms.Aliyun(conf)).Send("18512345678", "mob-login")
if err != nil {
    log.Println(err.Error())
    return
}
log.Println(code)
```

# 验证

```
conf := goo_sms.AliyunConfig{
    Region:       "",
    Appid:        "",
    Secret:       "",
    SignName:     "",
    TemplateCode: "",
}
err := goo_sms.New(goo_sms.Aliyun(conf)).Verify("18512345678", "mob-login", "1234")
if err != nil {
    log.Println(err.Error())
    return
}
```