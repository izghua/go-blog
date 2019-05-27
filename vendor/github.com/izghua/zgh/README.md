# This is a package just for me

- [中文文档](./CN.md)

## main

- set DB
- set Redis
- set Log
- set Alarm
- set Mail
- set HashId
- set common functions
  - encrypt string
  - rand string
  - return time format
  - etc
  
> PS:Function has default value,so you can not be set it ,if you want to set it.you can use `set` func with it

## demo

### set DB

```
db := new(conn.Sp)
dbUser := db.SetDbUserName("root")
dbPwd := db.SetDbPassword("123456")
dbPort := db.SetDbPort("3306")
dbHost := db.SetDbHost("127.0.0.1")
dbdb := db.SetDbDataBase("izghua")
_,err := conn.InitMysql(dbUser,dbPwd,dbPort,dbHost,dbdb)
```

### if you want watch more, see `unit test file`
