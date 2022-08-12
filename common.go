package dbconnector

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	kind_postgres      = "postgres"
	kind_mariadb       = "mariadb"
	kind_redis         = "redis"
	connector_file_ext = ".contr"
)

// 获得明文,ct 的输入一定是密文
func getPlaintext(ct io.Reader) (string, error) {
	if ct == nil {
		return "", fmt.Errorf("ct is nil")
	}
	buf, err := io.ReadAll(ct)
	if err != nil {
		return "", fmt.Errorf("getPlaintext io.ReadAll过程失败. %s", err)
	}
	s := strings.ReplaceAll(strings.ReplaceAll(string(buf), "\r", ""), "\n", "")
	//fmt.Println(">>>>s=", len(s))
	plaintext, err := decrypt(s)
	// fmt.Println(">>>>plaintext=", plaintext)
	// s此时应该是从字符串或文件中读取的密文，下面对其解密
	return plaintext, err
}

// 按照指定路径搜索所有后缀匹配的连接器文件名列表,返回 路径,文件名列表,error
func searchContrFiles(currentPath string) (string, []string, error) {
	var err error
	if currentPath == "" {
		currentPath, err = os.Getwd()
		if err != nil {
			return "", nil, err
		}
	}
	fis, err := os.ReadDir(currentPath)
	if err != nil {
		return "", nil, err
	}
	contr_files := []string{}
	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}
		if filepath.Ext(fi.Name()) == connector_file_ext {
			contr_files = append(contr_files, fi.Name())
		}
	}
	return currentPath, contr_files, nil
}

// 分析获得连接器类型，并返回json的字符串原型
func getConnectorKind(plainText string) (kind string, jsonstr string, err error) {
	runes := []rune(plainText)
	endPos := 0
	for index, u := range runes {
		if u == '{' {
			endPos = index
			break
		}
	}
	if endPos == 0 {
		err = errors.New("连接器数据基础格式错误")
		return "", "", err
	}
	return string(runes[0:endPos]), string(runes[endPos:]), err
}
