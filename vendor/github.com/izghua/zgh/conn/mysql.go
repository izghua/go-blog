/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-11-29
 * Time: 23:42
 */
package conn

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/conf"
	"github.com/izghua/zgh/utils/alarm"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var mysql *xorm.Engine

type SqlParam struct {
	Host string
	Port string
	DataBase string
	UserName string
	Password string
}

type Sp  func(*SqlParam) interface{}

func (p *Sp)SetDbHost(host string) Sp {
	return func(p *SqlParam) interface{} {
		h := p.Host
		p.Host = host
		return h
	}
}

func (p *Sp)SetDbPort(port string) Sp {
	return func(p *SqlParam) interface{} {
		pt := p.Port
		p.Port = port
		return pt
	}
}

func (p *Sp)SetDbDataBase(dataBase string) Sp {
	return func(p *SqlParam) interface{} {
		db := p.DataBase
		p.DataBase = dataBase
		return db
	}
}


func (p *Sp)SetDbPassword(pwd string) Sp {
	return  func(p *SqlParam) interface{} {
		password := p.Password
		p.Password = pwd
		return password
	}
}


func (p *Sp)SetDbUserName(u string) Sp {
	return func(p *SqlParam) interface{} {
		name := p.UserName
		p.UserName = u
		return name
	}
}

func InitMysql(options ...Sp) (*xorm.Engine,error){
	q := &SqlParam{
		Host:conf.DBHOST,
		Port:conf.DBPORT,
		Password:conf.DBPASSWORD,
		DataBase:conf.DBDATABASE,
		UserName:conf.DBUSERNAME,
	}
	for _,option := range options {
		option(q)
	}

	dataSourceName := q.UserName + ":" + q.Password + "@/" + q.DataBase + "?charset=utf8"
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		zgh.ZLog().Error("mysql","初始化数据库，创建连接异常:"+err.Error())
		return nil,err
	}
	engine.TZLocation,_ = time.LoadLocation("Asia/Chongqing")
	engine.SetMaxIdleConns(3)
	engine.SetMaxOpenConns(20)
	engine.SetConnMaxLifetime(0)
	engine.ShowExecTime(true)
	engine.ShowSQL(true)
	mysql = engine
	timer := time.NewTicker(time.Minute * 10)
	go func(conn *xorm.Engine) {
		for _ = range timer.C {
			if err = mysql.Ping(); err != nil {
				MySQLAutoConnect()
			}
		}
	}(mysql)
	return mysql,nil
}

func autoConnectMySQL(tryTimes int, maxTryTimes int) int {
	tryTimes++
	if tryTimes <= maxTryTimes {
		if mysql.Ping() != nil {
			message := fmt.Sprintf("数据库连接失败,已重连%d次", tryTimes)
			zgh.ZLog().Error("mysql",message)
			go alarm.Alarm(message)
		}
		tryTimes = autoConnectMySQL(tryTimes, maxTryTimes)
	}
	return tryTimes
}

func MySQLAutoConnect() {
	autoConnectMySQL(0, 5)
}


func SqlDump(fileName string,filePath string) error {
	zgh.ZLog().Info("message","sql dump file","name",fileName,"path",filePath)
	err := os.Remove(filePath+fileName)
	if err != nil {
		zgh.ZLog().Error("message","sql dump has error","error",err.Error())
	}

	err = mysql.DumpAllToFile(filePath+fileName)

	if err != nil {
		zgh.ZLog().Error("message","sql dump all to file has error","error",err.Error())
		return err
	}
	return nil
}