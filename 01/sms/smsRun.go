package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var userMap = make(map[string]User, 8)

type User struct {
	id       string
	password string
	status   uint8
}

func initUser() {
	userMap["ca0526"] = User{
		"ca0526",
		"123456",
		0,
	}
	userMap["ca0527"] = User{
		"ca0527",
		"123456",
		0,
	}
	userMap["ca0528"] = User{
		"ca0528",
		"123456",
		0,
	}
}
func main() {
	// 初始化本地数据库
	initUser()
	var logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	service := SingleFaService{
		logger,
	}
	request := &CommonRequest{
		"192.168.1.1",
		"ca0528",
		"14853",
		"single",
		"17689935953,19522821941",
		"【和选商城】\n和选企业福利采购将员工个性化需求、企业人文关怀理念相结合，融合节日场景、福利主题，为企业客户提供一站式指定节日和全年福利解决方案，让企业福利采购变得省心又省力；让员工福利选购变得随心又开心！?\n和选，做更懂企业的综合性服务平台！\n回复T退订（该短信为视频短信测试）",
		"",
		"asgasg",
		"jingdong",
	}
	something := service.DoSomething(request)
	fmt.Println(something)
	// 异常,异步，异步等待，channel
	err1()
	// 异步
	async()
	var wg sync.WaitGroup
	wg.Add(4)
	// 异步等待
	func() {
		go handlerWait("zhangsan", &wg)
		go handlerWait("lisi", &wg)
		go handlerWait("wangwu", &wg)
	}()
	testChannel(&wg)
	wg.Wait()
	files, curPath := listFiles("d:\\zhangjie\\Desktop\\20221031_华融漏洞扫描修复\\T1\\cluster_sender\\lib")
	if len(files) > 0 {
		for _, filePath := range files {
			println(string(filepath.Separator))
			open, err := os.Open(curPath + string(filepath.Separator) + filePath)
			if err != nil {
				fmt.Println("读取文件时发生错误")
			}
			fmt.Println("文件句柄打开成功：" + open.Name())
		}
	}
}

// 异常
func err1() {
	err := err2()
	if err != nil {
		fmt.Println("err1 hanlder error" + err.Error())
	}
}
func err2() error {
	success, err := err3()
	if !success {
		return err
	}
	fmt.Println("err1 call success")
	return nil
}
func err3() (bool, error) {
	_, err := os.Open("asfsafas")
	if err != nil {
		return false, err
	}
	fmt.Println("成功获取文件")
	return true, nil
}
func handler(name string) {
	fmt.Println("当前用户：" + name + ",已成功处理完成！")
}
func handlerWait(name string, wg *sync.WaitGroup) {
	fmt.Println("当前用户：" + name + ",已成功处理完成！")
	wg.Done()
}

// 异步
func async() {
	go handler("zhangjie")
	go handler("meixin")
	go handler("zhangjian")
}

// 通道(通道可以认为是在协程之前用于通信得一种手段)
func testChannel(wg *sync.WaitGroup) {
	in := make(chan int, 10000)
	go producer(in)
	go consumer(in, wg)
}

// 生产者  1.生产产品，当到达十个产品时停止生产
func producer(in chan int) {
	i := 0
	for {
		if i == 100 {
			close(in)
			break
		}
		i++
		in <- i
	}
}

// 消费者 1.消费产品，当消费完所有产品时结束生产
func consumer(out chan int, wg *sync.WaitGroup) {
	for {
		select {
		case data, ok := <-out:
			if !ok {
				wg.Done()
				return
			}
			fmt.Println(data)
		}
	}
}

// 文件操作根据传入的文件路径，返回对应路径下的所有文件
func listFiles(path string) ([]string, string) {
	filePaths := make([]string, 0)
	readDir, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("解析文件目录错误：%v\n", err.Error())
		return filePaths, path
	}
	for _, dir := range readDir {
		if info, _ := dir.Info(); !dir.IsDir() && info.Size() > 0 {
			filePaths = append(filePaths, dir.Name())
		}
	}
	return filePaths, path
}
