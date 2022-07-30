package dbconnector

import (
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	mariadbs        map[string]*sqlx.DB // 全部的mariadb数据库连接器
	mariadbs_struct map[string]*Mariadb_t
)

func init() {
	mariadbs = make(map[string]*sqlx.DB)
	mariadbs_struct = make(map[string]*Mariadb_t)
}

// 通过json字符串原型增加一个Mariadb连接器
func addMariadbByJsonString(jsonstr string) error {
	p := &Mariadb_t{
		Key:     "",
		Server:  "",
		Port:    0,
		Uid:     "",
		Pwd:     "",
		DB:      "",
		Timeout: 0,
	}
	err := json.Unmarshal([]byte(jsonstr), p)
	if err != nil {
		return fmt.Errorf("Mariadb连接数据JSON格式分析错误:%s", err.Error())
	}
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		p.Uid, p.Pwd, p.Server, p.Port, p.DB)
	db, err := sqlx.Open("mysql", connstr)
	if err != nil {
		return err
	}
	if p.Timeout > 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(p.Timeout))
	}
	mariadbs_struct[p.Key] = p
	mariadbs[p.Key] = db
	// err = sqldb.Ping()
	// if err != nil {
	// 	return err
	// }
	return err
}

func addMariadbByStruct(s *Mariadb_t) error {
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		s.Uid, s.Pwd, s.Server, s.Port, s.DB)
	db, err := sqlx.Open("mysql", connstr)
	if err != nil {
		return err
	}
	if s.Timeout > 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(s.Timeout))
	}
	mariadbs_struct[s.Key] = s
	mariadbs[s.Key] = db
	// err = sqldb.Ping()
	// if err != nil {
	// 	return err
	// }
	return err
}

// 清除所有的连接器
func CleanMariadb() {
	for k := range mariadbs {
		delete(mariadbs, k)
	}
	for k := range mariadbs_struct {
		delete(mariadbs_struct, k)
	}
}
