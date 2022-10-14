package main

import (
	"fmt"
	"log"
	"os"
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
}
