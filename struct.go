package dbconnector

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Redis_t struct {
	Key           string `toml:"key" json:"key"`
	Server        string `toml:"server"`
	Port          int    `toml:"port"`
	Pwd           string `toml:"pwd"`
	DB            int    `toml:"db"`
	PoolMaxActive int    `toml:"poolMaxActive"` // 0无限制
	MaxIdle       int
	IdleTimeout   int
	MaxActive     int
}

// 将连接信息按照规范写入流（加密后内容）
func (t Redis_t) CreateConnector(output io.Writer) error {
	buf, err := json.Marshal(t)
	if err != nil {
		return fmt.Errorf("createConnector生成json阶段错误: %s", err)
	}
	data := []byte(kind_mariadb + string(buf))
	fmt.Println(string(data))
	clptext, err := encrypt(data)
	if err != nil {
		return fmt.Errorf("createConnector文本加密阶段错误: %s", err)
	}
	_, err = output.Write([]byte(clptext))
	return err
}

// 提供文件名(不含后缀，模块自动添加)，保存生成连接器文件
func (t Redis_t) SaveConnectorFile(onlyMainFilename string) error {
	filename := onlyMainFilename + connector_file_ext
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("生成连接器文件错误. %s", err)
	}
	defer f.Close()
	err = t.CreateConnector(f)
	return err
}

type Mariadb_t struct {
	Key          string `toml:"key" json:"key"`
	Server       string `toml:"server" json:"server"`
	Port         int    `toml:"port" json:"port"`
	Uid          string `toml:"uid" json:"uid"`
	Pwd          string `toml:"pwd" json:"pwd"`
	DB           string `toml:"db" json:"db"`
	Timeout      string `toml:"timeout" json:"timeout"`           // 字符串，携带单位如果忘记写默认s
	ReadTimeout  string `toml:"readTimeout" json:"readTimeout"`   // 字符串，携带单位如果忘记写默认s
	WriteTimeout string `toml:"writeTimeout" json:"writeTimeout"` // 字符串，携带单位如果忘记写默认s
}

// 将连接信息按照规范写入流（加密后内容）
func (t Mariadb_t) CreateConnector(output io.Writer) error {
	buf, err := json.Marshal(t)
	if err != nil {
		return fmt.Errorf("createConnector生成json阶段错误: %s", err)
	}
	data := []byte(kind_mariadb + string(buf))
	fmt.Println(string(data))
	clptext, err := encrypt(data)
	if err != nil {
		return fmt.Errorf("createConnector文本加密阶段错误: %s", err)
	}
	_, err = output.Write([]byte(clptext))
	return err
}

// 提供文件名(不含后缀，模块自动添加)，保存生成连接器文件
func (t Mariadb_t) SaveConnectorFile(onlyMainFilename string) error {
	filename := onlyMainFilename + connector_file_ext
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("生成连接器文件错误. %s", err)
	}
	defer f.Close()
	err = t.CreateConnector(f)
	return err
}

type Postgres_t struct {
	Key      string `toml:"key" json:"key"`
	Server   string `toml:"server" json:"server"`
	Port     int    `toml:"port" json:"port"`
	Username string `toml:"user" json:"user"`
	Pwd      string `toml:"pwd" json:"pwd"`
	DB       string `toml:"db" json:"db"`
	Timeout  int    `toml:"timeout" json:"timeout"`
}

// 将连接信息按照规范写入流（加密后内容）
func (t Postgres_t) CreateConnector(output io.Writer) error {
	buf, err := json.Marshal(t)
	if err != nil {
		return fmt.Errorf("createConnector生成json阶段错误: %s", err)
	}
	data := []byte(kind_postgres + string(buf))
	fmt.Println(string(data))
	clptext, err := encrypt(data)
	if err != nil {
		return fmt.Errorf("createConnector文本加密阶段错误: %s", err)
	}
	_, err = output.Write([]byte(clptext))
	return err
}

// 提供文件名(不含后缀，模块自动添加)，保存生成连接器文件
func (t Postgres_t) SaveConnectorFile(onlyMainFilename string) error {
	filename := onlyMainFilename + connector_file_ext
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("生成连接器文件错误. %s", err)
	}
	defer f.Close()
	err = t.CreateConnector(f)
	return err
}
