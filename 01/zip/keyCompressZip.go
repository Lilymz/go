package main

import (
	"archive/zip"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//var compress = &Compress{}
	//pwd, err := os.Getwd()
	//if err!=nil{
	//	log.Printf("获取当前路径失败：%s\n",err.Error())
	//	return
	//}
	//flag.StringVar(&SourcePath,"sourcePath",pwd+"/source.zip","指定一个需要压缩的zip文件:！")
	//flag.StringVar(&TargetPath,"targetPath",pwd+"/target.zip","请指定当前需要加密zip文件名和路径！")
	//flag.IntVar(&mode,"mode",0,"请指定一个压缩默认，0：加密压缩！1：解密压缩 ")
	//flag.StringVar(&Key,"key","","请至少使用一个加密密钥！")
	//flag.Parse()
	//compress.SourcePath = SourcePath
	//compress.TargetPath = TargetPath
	//compress.Mode = uint(mode)
	//compress.Key = Key
	//log.Printf("当前结构：%#v\n",compress)
	//if compress.Mode==0 {
	//	log.Printf("使用加密压缩模式开启...")
	//	ZipToKeyZip(compress)
	//}else if compress.Mode==1 {
	//	log.Printf("使用解密压缩模式开启...")
	//	KeyZipToZip(compress)
	//}

	content, err := encrypt([]byte("1234567890123456"), []byte("hello1"), aes.BlockSize)
	if err != nil {
		log.Println("加密错误")
		return
	}
	log.Println("加密内容：" + base64.StdEncoding.EncodeToString(content))
	decodeString, _ := base64.StdEncoding.DecodeString("idCx4mc+xk7eKFQnMoFXLZIZ+UJkL+cvjx1N3coN23I=")
	decrypt, err := decrypt([]byte("1234567890123456"), decodeString)
	if err != nil {
		log.Println("加密错误")
		return
	}
	log.Println("解密内容：" + string(decrypt))
}

// ZipToKeyZip 将原始得zip压缩包转化为每个文件以某个key进行加密得文件，在进行压缩
func ZipToKeyZip(compressBean *Compress) {
	// 读取当前文件得所有压缩文件
	zipFile, err := zip.OpenReader(compressBean.SourcePath)
	if err != nil {
		log.Println("无法打开 ZIP 文件：", err)
		return
	}
	defer zipFile.Close()
	var (
		fileReader io.ReadCloser
		destFile   *os.File
		fileMap    = make(map[string]string, 8)
	)
	getwd, _ := os.Getwd()
	var basePath = getwd
	// 遍历 ZIP 文件中的文件
	for _, file := range zipFile.File {
		// 打开 ZIP 文件中的文件
		fileReader, err = file.Open()
		if err != nil {
			log.Println("无法打开 ZIP 文件中的文件：", err)
			return
		}
		data := make([]byte, file.CompressedSize64)
		fileReader.Read(data)

		ciphertext, err := encrypt([]byte(compressBean.Key), []byte(strings.TrimSpace(string(data))), aes.BlockSize)
		if err != nil {
			log.Println("加密失败:", err)
			return
		}
		if strings.Contains(file.Name, "txt") {
			decrypt, err := decrypt([]byte(compressBean.Key), []byte(base64.StdEncoding.EncodeToString(ciphertext)))
			if err != nil {
				log.Println("验证解密结果：" + string(decrypt))
			}
		}
		fileMap["compress-"+strings.ReplaceAll(file.Name, ".", "&")] = base64.StdEncoding.EncodeToString(ciphertext)
		if "" != basePath && basePath != getwd {
			basePath = filepath.Dir(compressBean.TargetPath)
		}

	}
	for key, value := range fileMap {
		// 将 ZIP 文件中的文件内容复制到目标文件中
		err := os.WriteFile(basePath+"/"+key, []byte(value), 0666)
		if err != nil {
			log.Println("加密文件内容时出错：", err)
			return
		}
	}
	err = compressDirToZip(basePath, compressBean.TargetPath)
	if err != nil {
		log.Println("压缩文件出错：", err)
		return
	}
	defer fileReader.Close()
	defer destFile.Close()
	// 遍历当前压缩文件中得文件，进行加密处理放入到对应得map[name]encryptData中
	// 将当前得map压缩为对应得zip文件
}
func compressDirToZip(sourceDirPath, zipFilePath string) error {
	delFiles := make([]string, 8)
	// 创建输出的 ZIP 文件
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// 创建 zip.Writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 遍历源目录
	err = filepath.Walk(sourceDirPath, func(filePath string, fileInfo os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		// 跳过目录
		if fileInfo.IsDir() {
			return nil
		}

		if strings.HasSuffix(filepath.Ext(filePath), ".go") || strings.HasSuffix(filepath.Ext(filePath), ".exe") {
			return nil
		}
		delFiles = append(delFiles, filePath)
		// 获取相对于源目录的文件路径
		relativePath, err := filepath.Rel(sourceDirPath, filePath)

		if err != nil {
			return err
		}

		// 创建压缩文件的文件头
		header, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			return err
		}

		// 设置压缩文件名（保留相对路径）
		header.Name = filepath.ToSlash(relativePath)
		// 创建压缩文件的写入器
		writer, err := zipWriter.CreateHeader(header)

		if err != nil {
			return err
		}

		// 打开源文件
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// 将源文件内容复制到压缩文件中
		length, err := io.Copy(writer, file)
		if err != nil {
			return err
		}
		fmt.Printf("current file:%v,writeLength:%v\n", filePath, length)
		return nil
	})
	for _, file := range delFiles {
		os.Remove(file)
	}
	if err != nil {
		return err
	}

	return nil
}

// KeyZipToZip 将加密压缩包解密为普通压缩包
func KeyZipToZip(compressBean *Compress) {
	reader, err := zip.OpenReader(compressBean.SourcePath)
	if err != nil {
		log.Println("读取encrypt compress zip error" + err.Error())
	}
	defer reader.Close()
	var zipFile io.ReadCloser
	for _, file := range reader.File {
		zipFile, err = file.Open()
		if err != nil {
			log.Println("Failed to read zip file: " + err.Error())
			continue
		}
		// Read the content of the file
		fileContent, err := ioutil.ReadAll(zipFile)
		if err != nil {
			log.Println("Failed to read zip file content: " + err.Error())
			continue
		}
		// Do something with the file content
		log.Println("File name: " + file.Name)
		log.Println("File size: ", file.UncompressedSize64)
		if strings.Contains(file.Name, "txt") {
			fmt.Println(string(fileContent))
			_, err := base64.StdEncoding.DecodeString(string(fileContent))
			if err != nil {
				log.Println("base解密异常：" + err.Error())
				return
			}
			decrypt, err := decrypt([]byte(compressBean.Key), fileContent)
			if err != nil {
				log.Println("解密异常：" + err.Error())
				return
			}
			log.Println("File decrypt content: ", string(decrypt))
		}
	}
	defer zipFile.Close()
	// 遍历当前压缩文件中得文件，进行加密处理放入到对应得map[name]encryptData中
	// 将当前得map压缩为对应得zip文件
}

func encrypt(key, plaintext []byte, blockSize int) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 使用 AES 加密的 CBC 模式
	paddedPlaintext := pkcs7Pad(plaintext, blockSize) // 填充明文
	ciphertext := make([]byte, blockSize+len(paddedPlaintext))
	//iv := ciphertext[:blockSize]
	//if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	//	return nil, err
	//}

	mode := cipher.NewCBCEncrypter(block, nil)
	mode.CryptBlocks(ciphertext[blockSize:], paddedPlaintext)

	return ciphertext, nil
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}
func decrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	waitDecrypt := ciphertext[aes.BlockSize:]
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(waitDecrypt, waitDecrypt)
	unpad, _ := pkcs7Unpad(waitDecrypt)
	return unpad, nil
}

func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data) - aes.BlockSize
	if length == 0 {
		return data, nil
	}
	return data[:length], nil
}
