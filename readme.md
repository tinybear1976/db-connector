# Readme

## 优先级
如果存在连接器文件，比如 xxx.contr,那么模块会先行解析并加载。如果配置文件中存在配置信息部分(明文结构体)，模块会继续分析，并覆盖相同key的连接信息。如果是在开发模式下，方便明文配置数据库连接，则可以在配置文件中保留配置信息部分，作为正式发行版本，应该删除该配置数据，只保留连接器文件即可

## install
```go
go get github.com/tinybear1976/dbconnector
```

## basic
1. load config
```go
import github.com/tinybear1976/dbconnector

// ...
currentPath:=os.Getwd() // or get from Args
// Automatically scan files(.contr) in a directory and load
dbconnector.AddFromFiles(currentPath)
// (if there is) get from config file. if the keys are the same, the latter overrides the former
dbconnector.AddFromStructs(mariadbSlice,redisSlice) 
```
2. save contr file (reference)
```go
m := Mariadb_t{
		Key:     "test",
		Server:  "127.0.0.1",
		Port:    3306,
		Uid:     "root",
		Pwd:     "1234#@&!Keen",
		DB:      "city",
		Timeout:      "",
		ReadTimeout:  "",
		WriteTimeout: "",
	}
// only main file name
err := m.SaveConnectorFile("t1")

m := Redis_t{
		Key:    "local",
		Server: "127.0.0.1",
		Port:   6379,
		Pwd:    "",
		DB:     0,
	}
main_name := "redis_local"
err := m.SaveConnectorFile(main_name)
```

3. use in project
mariadb:
```go
// ...
const (
    // Mariadb_t{ key:"app",.... }
    APP MariadbConnectors="app"
)

// ...
func function()  {
    //...
    // ok: return *sqlx.DB, else: return nil
    db:=APP.Connector()
    // db.Query... 
    //...
}
```

redis
```go
// ...
const (
    // Redis_t{ key:"redislocal",.... }
    REDISLOCAL RedisConnectors="redislocal"
)

// ...
func function()  {
    //...
    // ok: return *redis.Conn,nil, else: return nil,error
    conn,err:=REDISLOCAL.Connect()
    // (*conn).Do()
    // conn.GET()
    //...
}
```