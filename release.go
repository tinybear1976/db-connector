package dbconnector

import (
	"fmt"
	"os"
	"path"

	"github.com/jmoiron/sqlx"
)

type MariadbConnectors string

// 返回连接器，如果没有找到则连接器返回nil
func (ctr MariadbConnectors) Connector() *sqlx.DB {
	key := string(ctr)
	return mariadbs[key]
}

// 返回连接信息，如果没有找到则连接器返回nil
func (ctr MariadbConnectors) Info() *Mariadb_t {
	key := string(ctr)
	return mariadbs_struct[key]
}

// 根据指定路径搜索所有连接器文件，并进行解析，自动添加
func AddFromFiles(currentPath string) error {
	dir, contrFiles, err := searchContrFiles(currentPath)
	if err != nil {
		return fmt.Errorf("路径:%s 获取连接器文件列表失败. %s", dir, err)
	}
	for _, fn := range contrFiles {
		f, err := os.Open(path.Join(dir, fn))
		if err != nil {
			return fmt.Errorf("解析连接文件时出现错误:%s", err.Error())
		}
		defer f.Close()
		plainText, err := getPlaintext(f)
		if err != nil {
			return fmt.Errorf("文件:%s 密文解析失败. %s", f.Name(), err)
		}
		kind, jsonstr, err := getConnectorKind(plainText)
		if err != nil {
			return fmt.Errorf("文件:%s 明文解析失败. %s", f.Name(), err)
		}
		switch kind {
		case kind_mariadb:
			err = addMariadbByJsonString(jsonstr)
			if err != nil {
				err = fmt.Errorf("尝试增加mariadb连接器失败. %s", err)
			}
		case kind_redis:
			err = addRedisByJsonString(jsonstr)
			if err != nil {
				err = fmt.Errorf("尝试增加redis连接器失败. %s", err)
			}
		default:
			err = fmt.Errorf("未知的数据库类型(%s),无法使用", kind)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// 从连接配置结构体切片中进行添加，一般改函数用于第二步，从配置文件中读取后，进行添加，如果与之前的文件连接器key冲突，这里将覆盖之前的内容
func AddFromStructs(ms []Mariadb_t, rs []Redis_t) (err error) {
	for i := 0; i < len(ms); i++ {
		err = addMariadbByStruct(&(ms[i]))
		if err != nil {
			return
		}
	}

	for i := 0; i < len(rs); i++ {
		addRedisByStruct(&(rs[i]))
	}
	return
}

// 解密一个指定的连接器文件，只还原字符串原型,不做进一步解析
func DecryptConnectorFile(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("打开文件错误:%s", err)
	}
	defer f.Close()

	plainText, err := getPlaintext(f)
	if err != nil {
		return "", fmt.Errorf("文件:%s 密文解析失败. %s", f.Name(), err)
	}
	return plainText, nil
}