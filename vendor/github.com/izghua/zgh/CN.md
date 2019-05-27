# 这是一个给我自己用的公共包

## 主要功能

- 数据库设置
- redis设置
- 日志设置
- 报警设置
- 邮件设置
- 混淆ID设置
- 常用函数设置
  - 各种字符串加星
  - 随机字符串
  - 返回当前时间格式
  - etc
  
> 说明:大多都是有默认值,可以不设置,如需设置,请调用 相关的`set`方法

## demo

### 数据库设置

```
db := new(conn.Sp)
dbUser := db.SetDbUserName("root")
dbPwd := db.SetDbPassword("123456")
dbPort := db.SetDbPort("3306")
dbHost := db.SetDbHost("127.0.0.1")
dbdb := db.SetDbDataBase("izghua")
_,err := conn.InitMysql(dbUser,dbPwd,dbPort,dbHost,dbdb)
```

其他详细请看 `单元测试文件`