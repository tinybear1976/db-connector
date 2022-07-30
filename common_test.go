package dbconnector

import (
	"fmt"
	"os"
	"testing"
)

func Test_searchContrFiles(t *testing.T) {
	dir, files, err := searchContrFiles("/Users/wang/program/go/db-connector")
	if err != nil {
		t.Fatalf("AddFromFiles() has error: %v", err)
		return
	}
	expected1 := 1
	expected2 := "t1.contr"
	observed1 := len(files)
	observed2 := ""
	if len(files) > 0 {
		observed2 = files[0]
	}
	fmt.Println("return dir", dir)
	if observed1 != expected1 || observed2 != expected2 {
		t.Fatalf("searchContrFiles()  len(files)= %d,filename= %v, want len=%v, filename=%v",
			observed1, observed2, expected1, expected2)
	}
}

func Test_getPlaintext(t *testing.T) {
	file := "/Users/wang/program/go/db-connector/t1.contr"
	f, err := os.Open(file)
	if err != nil {
		t.Fatalf("getPlaintext() has error: %v", err)
		return
	}
	//f := strings.NewReader("nqYNLm/MQ1VdlnoDQhQ4lncNi/D5M4s1V/2h8sidExK+P5FM8WZv3r0Gbihx5PwPwWMTRHDN0QNRFqLVO9NrPpAk3e07UwUr6J1HEaw4GXaJQJx4CK0JzFYyjN7WuUNiu04GoGG4AG8ISxkjQEiNXLyDz2PGuyVLVlIbGsma8q0=")
	observed, err := getPlaintext(f)
	if err != nil {
		t.Fatalf("getPlaintext() has error: %v", err)
		return
	}
	fmt.Println("plaintext=", observed)
	expected := `mariadb{"key":"test","server":"127.0.0.1","port":3306,"uid":"root","pwd":"1234#@\u0026!Keen","db":"city","timeout":10}`
	if observed != expected {
		t.Fatalf("getPlaintext()  text:%v, want:%v",
			observed, expected)
	}
	defer f.Close()

}
