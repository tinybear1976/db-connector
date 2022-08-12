package dbconnector

import (
	"fmt"
	"testing"
)

func Test_AddFromFiles(t *testing.T) {
	err := AddFromFiles("/Users/wang/program/go/db-connector")
	if err != nil {
		t.Fatalf("AddFromFiles() has error: %v", err)
		return
	}
	expected := 1
	observed1 := len(mariadbs)
	observed2 := len(mariadbs_struct)
	if observed1 != expected || observed2 != expected {
		t.Fatalf("AddFromFiles()  len(mariadbs)= %v,len(mariadbs_struct)= %v, want %v",
			observed1, observed2, expected)
	}
	fmt.Println(">>>>>>", mariadbs, *mariadbs_struct["test"])
}
func Test_AddFromStructs(t *testing.T) {
	ms := []Mariadb_t{
		{
			Key:          "t1",
			Server:       "192.168.1.3",
			Port:         3306,
			Uid:          "root",
			Pwd:          "123",
			DB:           "test",
			Timeout:      "0.5s",
			ReadTimeout:  "5s",
			WriteTimeout: "2s",
		},
		{
			Key:          "t2",
			Server:       "192.168.0.1",
			Port:         3306,
			Uid:          "root",
			Pwd:          "wwww",
			DB:           "mm",
			Timeout:      "",
			ReadTimeout:  "",
			WriteTimeout: "",
		},
	}

	ps := []Postgres_t{
		{
			Key:      "p1",
			Server:   "111",
			Port:     0,
			Username: "",
			Pwd:      "",
			DB:       "",
			Timeout:  0,
		},
		{
			Key:      "p2",
			Server:   "222",
			Port:     0,
			Username: "",
			Pwd:      "",
			DB:       "",
			Timeout:  0,
		},
	}

	rs := []Redis_t{
		{
			Key:    "r1",
			Server: "127.0.0.1",
			Port:   6379,
			Pwd:    "",
			DB:     0,
		},
	}
	err := AddFromStructs(ms, rs, ps)
	if err != nil {
		t.Fatalf("AddFromStructs() has error: %v", err)
		return
	}
	fmt.Println(">>>>>>", mariadbs, *(mariadbs_struct["t1"]), *(mariadbs_struct["t2"]), *(redis_struct["r1"]))
}

func Test_DecryptConnectorFile(t *testing.T) {
	observed, err := DecryptConnectorFile("/Users/wang/program/go/dbconnector/t1" + connector_file_ext)
	if err != nil {
		t.Fatalf("DecryptConnectorFile() has error: %v", err)
		return
	}
	expected := `mariadb{"key":"test","server":"127.0.0.1","port":3306,"uid":"root","pwd":"1234#@\u0026!Keen","db":"city","timeout":10}`
	if observed != expected {
		t.Fatalf("DecryptConnectorFile() = %v, want %v",
			observed, expected)
	}
	// fmt.Println(">>>>>>", mariadbs, *mariadbs_struct["test"])
}

func Test_DecryptConnectorFile2(t *testing.T) {
	observed, err := DecryptConnectorFile("/Users/wang/program/go/dbconnector/t2" + connector_file_ext)
	if err != nil {
		t.Fatalf("DecryptConnectorFile() has error: %v", err)
		return
	}
	expected := `postgres{"key":"test","server":"127.0.0.1","port":5432,"user":"root","pwd":"1234#@\u0026!Keen","db":"city","timeout":20}`
	if observed != expected {
		t.Fatalf("DecryptConnectorFile() = %v, want %v",
			observed, expected)
	}
	// fmt.Println(">>>>>>", mariadbs, *mariadbs_struct["test"])
}
