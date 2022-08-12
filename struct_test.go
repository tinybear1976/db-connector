package dbconnector

import (
	"fmt"
	"os"
	"testing"
)

func Test_Mariadbt_SaveFile(t *testing.T) {

	m := Mariadb_t{
		Key:          "test",
		Server:       "127.0.0.1",
		Port:         3306,
		Uid:          "root",
		Pwd:          "1234#@&!Keen",
		DB:           "city",
		Timeout:      "2s",
		ReadTimeout:  "2s",
		WriteTimeout: "",
	}

	err := m.SaveConnectorFile("t1")
	fmt.Println("error: ", err)
	realFilename := "t1" + connector_file_ext
	_, err = os.Stat(realFilename)
	if os.IsNotExist(err) {
		t.Fatal("SaveConnectorFile() = file not exist, want exist")
	}
}

func Test_Postgrest_SaveFile(t *testing.T) {

	m := Postgres_t{
		Key:      "test",
		Server:   "127.0.0.1",
		Port:     5432,
		Username: "root",
		Pwd:      "1234#@&!Keen",
		DB:       "city",
		Timeout:  20,
	}

	err := m.SaveConnectorFile("t2")
	fmt.Println("error: ", err)
	realFilename := "t2" + connector_file_ext
	_, err = os.Stat(realFilename)
	if os.IsNotExist(err) {
		t.Fatal("SaveConnectorFile() = file not exist, want exist")
	}
}

func Test_Redist_SaveFile(t *testing.T) {

	m := Redis_t{
		Key:           "local",
		Server:        "127.0.0.1",
		Port:          6379,
		Pwd:           "",
		DB:            0,
		PoolMaxActive: 0,
	}
	main_name := "redis_local"
	err := m.SaveConnectorFile(main_name)
	fmt.Println("error: ", err)
	realFilename := main_name + connector_file_ext
	_, err = os.Stat(realFilename)
	if os.IsNotExist(err) {
		t.Fatal("SaveConnectorFile() = file not exist, want exist")
	}
}
