package dbconnector

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	postgres        map[string]*sqlx.DB // 全部的mariadb数据库连接器
	postgres_struct map[string]*Postgres_t
)

func init() {
	postgres = make(map[string]*sqlx.DB)
	postgres_struct = make(map[string]*Postgres_t)
}

// 清除所有的连接器
func CleanPostgres() {
	for pg := range postgres {
		delete(postgres, pg)
	}
	for k := range postgres_struct {
		delete(postgres_struct, k)
	}
}

func addPostgresByStruct(s *Postgres_t) error {
	connstr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable connect_timeout=%d client_encoding=UTF8",
		s.Server, s.Port, s.Username, s.Pwd, s.DB, s.Timeout)
	db, err := sqlx.Open("postgres", connstr)
	sqlx.Connect("postgres", connstr)
	if err != nil {
		return err
	}

	postgres_struct[s.Key] = s
	postgres[s.Key] = db
	// err = sqldb.Ping()
	// if err != nil {
	// 	return err
	// }
	return err
}

// 通过json字符串原型增加一个Postgres连接器
func addPostgresByJsonString(jsonstr string) error {
	p := &Postgres_t{
		Key:      "",
		Server:   "",
		Port:     5432,
		Username: "",
		Pwd:      "",
		DB:       "",
		Timeout:  20,
	}
	err := json.Unmarshal([]byte(jsonstr), p)
	if err != nil {
		return fmt.Errorf("Mariadb连接数据JSON格式分析错误:%s", err.Error())
	}
	err = addPostgresByStruct(p)
	// err = sqldb.Ping()
	// if err != nil {
	// 	return err
	// }
	return err
}
