package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"log"
	"strings"
)

const (
	AesKey = "zhangjie--meixin"
)

// 处理数据，对数据进行填充，采用PKCS7（当密钥长度不够时，缺几位补几个几）的方式。
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

// DecodeMsgContent 解密
func DecodeMsgContent(data []byte, key []byte) ([]byte, error) {
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去除填充
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

// EncodeMsgContent 加密数据
func EncodeMsgContent(msgContent string) []byte {
	key := []byte(AesKey)
	// AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("创建加密AES-256算法失败，%#v\n", err.Error())
		return nil
	}
	size := block.BlockSize()
	// 加密字节数据
	encryptBytes := pkcs7Padding([]byte(msgContent), size)
	//初始化加密数据接收切片
	cryptedBytes := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:size])
	//执行加密
	blockMode.CryptBlocks(cryptedBytes, encryptBytes)
	return cryptedBytes
}

// EncodeMsgContentStr 生成AES加密字符串
func EncodeMsgContentStr(content string) string {
	msgContent := EncodeMsgContent(content)
	return base64.StdEncoding.EncodeToString(msgContent)
}

// EncodeMobileBatchLog 批量手机号脱敏
func EncodeMobileBatchLog(mobiles string) string {
	var builder strings.Builder
	if mobiles == "" {
		return ""
	}
	phones := strings.Split(mobiles, ",")
	for _, phone := range phones {
		if length := len(phone); length < 11 {
			// 不足11位不做脱敏,直接返回
			builder.WriteString(phone)
			builder.WriteString(",")
		}
		var tmpBuilder strings.Builder
		for index, char := range phone {
			if index >= 3 && index < 7 {
				tmpBuilder.WriteString("*")
			} else {
				tmpBuilder.WriteString(string(char))
			}
		}
		builder.WriteString(tmpBuilder.String())
		builder.WriteString(",")
	}
	return builder.String()
}
