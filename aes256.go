package dbconnector

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var (
	req_key     = "tiny"
	req_iv      = "WangZhaoWang0815"
	key_maxsize = 32
	key         = []byte(req_key)
	iv_maxsize  = 16
	iv          = []byte(req_iv)
)

func init() {
	key = zeroPadding(key, key_maxsize)
	iv = zeroPadding(iv, iv_maxsize)
}

// 加密
func encrypt(text []byte) (string, error) {
	//生成cipher.Block 数据块
	block, err := aes.NewCipher(key)
	if err != nil {
		// log.Println("错误 -" + err.Error())
		return "", err
	}
	//填充内容，如果不足16位字符
	blockSize := block.BlockSize()
	originData := pad(text, blockSize)
	//加密方式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//加密，输出到[]byte数组
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func zeroPadding(k []byte, maxsize int) []byte {
	real_len := len(k)
	var rst []byte
	if real_len < maxsize {
		rst = append(rst, k...)
		add_els := bytes.Repeat([]byte{byte(0)}, maxsize-real_len)
		rst = append(rst, add_els...)
	} else {
		rst = append(rst, k[0:maxsize]...)
	}
	return rst
}

func pad(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 解密
func decrypt(text string) (string, error) {
	decode_data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	// fmt.Println(">>>> ", string(decode_data))
	//生成密码数据块cipher.Block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//解密模式
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//输出到[]byte数组
	origin_data := make([]byte, len(decode_data))
	blockMode.CryptBlocks(origin_data, decode_data)
	//去除填充,并返回
	return string(unpad(origin_data)), nil
}

func unpad(ciphertext []byte) []byte {
	length := len(ciphertext)
	//去掉最后一次的padding
	unpadding := int(ciphertext[length-1])
	return ciphertext[:(length - unpadding)]
}
