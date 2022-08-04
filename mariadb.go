package dbconnector

import (
	"encoding/json"
	"fmt"
	"regexp"

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
		Key:          "",
		Server:       "",
		Port:         0,
		Uid:          "",
		Pwd:          "",
		DB:           "",
		Timeout:      "",
		ReadTimeout:  "",
		WriteTimeout: "",
	}
	err := json.Unmarshal([]byte(jsonstr), p)
	if err != nil {
		return fmt.Errorf("Mariadb连接数据JSON格式分析错误:%s", err.Error())
	}
	err = addMariadbByStruct(p)
	// err = sqldb.Ping()
	// if err != nil {
	// 	return err
	// }
	return err
}

func addMariadbByStruct(s *Mariadb_t) error {
	timeout := parseTimeout(s.Timeout)           // 连接超时。Timeout for establishing connections, aka dial timeout. The value must be a decimal number with a unit suffix ("ms", "s", "m", "h"), such as "30s", "0.5m" or "1m30s".
	readTimeout := parseTimeout(s.ReadTimeout)   // I/O读取超时。I/O read timeout. The value must be a decimal number with a unit suffix ("ms", "s", "m", "h"), such as "30s", "0.5m" or "1m30s".
	writeTimeout := parseTimeout(s.WriteTimeout) // I/O写入超时。I/O write timeout. The value must be a decimal number with a unit suffix ("ms", "s", "m", "h"), such as "30s", "0.5m" or "1m30s".
	if len(timeout) > 0 {
		timeout = fmt.Sprintf("&timeout=%s", timeout)
	}
	if len(readTimeout) > 0 {
		readTimeout = fmt.Sprintf("&readTimeout=%s", readTimeout)
	}
	if len(writeTimeout) > 0 {
		writeTimeout = fmt.Sprintf("&writeTimeout=%s", writeTimeout)
	}
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4%s%s%s",
		s.Uid, s.Pwd, s.Server, s.Port, s.DB, timeout, readTimeout, writeTimeout)
	db, err := sqlx.Open("mysql", connstr)
	if err != nil {
		return err
	}

	mariadbs_struct[s.Key] = s
	mariadbs[s.Key] = db
	// err = sqldb.Ping()
	// if err != nil {
	// 	return err
	// }
	return err
}

// 分析超时参数，如果有错误，能按照默认值处理就按照默认值处理，如果不能直接返回空字符串
func parseTimeout(val string) string {
	matched, err := regexp.MatchString("^([0-9]{1,}|[0-9]{1,}[.][0-9]*)(ms|s|h|m)$", val)
	if err != nil || !matched {
		// 忽略错误，返回空字符串
		return ""
	}
	// if matched {
	return val
	// } else {
	// return ""
	// }
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
