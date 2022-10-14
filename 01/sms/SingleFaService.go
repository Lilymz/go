package main

import (
	"log"
	"strconv"
	"strings"
)

/***
相同内容单发流程的设计
1.接收用户那边传来的数据，并对相关数据进行校验，更正，以及返回相关错误码。
2.拆分传来的手机号，创建对应的短信切片，遍历手机号set相关属性
*/

// CommonRequest 下发参数接收
type CommonRequest struct {
	RequestIp     string `json:"requestIp"`     // 请求ip
	UserId        string `json:"corpId"`        //账号
	Password      string `json:"password"`      // 密码
	CorpService   string `json:"corpService"`   //业务号
	mobiles       string `json:"mobiles"`       //手机号（多个）
	MsgContent    string `json:"msgContent"`    // 内容
	MsgId         string `json:"MsgId"`         //短信Id
	ExtCode       string `json:"extCode"`       // 扩展号
	InterfaceType string `json:"interfaceType"` // 上游接口
}

// HSResponse 返回容器
type HSResponse struct {
	code  string                 `json:"code"`  // 响应码
	text  string                 `json:"text"`  // 描述
	other map[string]interface{} `json:"other"` // 其他数据
}
type SingleFaService struct {
	logger *log.Logger
}

// ISmsService 短信下发接口
type ISmsService interface {
	// DoSomething 下发操作
	DoSomething(request *CommonRequest) HSResponse
}

// DoSomething 短信下发
func (s *SingleFaService) DoSomething(request *CommonRequest) HSResponse {
	corpId := request.UserId
	ip := request.RequestIp
	password := request.Password
	mobiles := request.mobiles
	msgId := request.MsgId
	msgContent := request.MsgContent
	corpService := request.CorpService
	interfaceType := request.InterfaceType
	extCode := request.ExtCode
	// 手机号脱敏
	encodeMobiles := EncodeMobileBatchLog(request.mobiles)
	// 内容加密（AES）
	encodeContent := EncodeMsgContentStr(msgContent)
	s.logger.Printf("SingleFaService user Parameters corpId:%v,ip:%v,password:%v,mobiles:%v,msgId:%v,msgContent:%v,corpService:%v,interfaceType:%v,extCode:%v",
		corpId, ip, password, encodeMobiles, msgId, encodeContent, corpService, interfaceType, extCode)
	if corpId == "" {
		s.logger.Println("user_id: " + corpId + "; user_ip: " + ip + " ; mobile: " + encodeMobiles + "; return_code: " + CommonUserIdError + "; explain: " + "参数:corp_id未填写")
		return HSResponse{
			CommonUserIdError,
			"账号参数填写不合法",
			nil,
		}
	}
	if msgContent == "" || len(msgContent) == 0 || len(msgContent) > 1000 {
		if len(msgContent) == 0 || len(msgContent) > 1000 {
			s.logger.Println("user_id: " + corpId + "; user_ip: " + ip + " ; mobile: " + encodeMobiles + "; return_code: " + CommonMsgContentError + "; explain: " + "短信内容长度不合法,msg_content_length:" + strconv.Itoa(len(msgContent)))
		} else {
			s.logger.Println("user_id: " + corpId + "; user_ip: " + ip + " ; mobile: " + encodeMobiles + "; return_code: " + CommonMsgContentError + "; explain: " + "参数:msg_content未填写")
		}
		return HSResponse{
			CommonMsgContentError,
			"短信内容参数填写不合法",
			nil,
		}
	}
	if mobiles == "" {
		s.logger.Println("user_id: " + corpId + "; user_ip: " + ip + " ; mobile: " + encodeMobiles + "; return_code: " + CommonMobileInfoError + "; explain: " + "参数:mobiles未填写")
		return HSResponse{
			CommonMobileInfoError,
			"手机号参数填写不合法",
			nil,
		}
	}
	_, ok := userMap[corpId]
	if !ok {
		s.logger.Println("user_id: " + corpId + "; user_ip: " + ip + " ; mobile: " + encodeMobiles + "; return_code: " + CommonMobileInfoError + "; explain: " + "参数:mobiles未填写")
		return HSResponse{
			CommonMobileInfoError,
			"账户" + corpId + "不存在或者已经关闭",
			nil,
		}
	}
	//msg_id校验
	if msgId == "" || len(msgId) == 0 || len(msgId) > 50 {
		msgId = genMsgId()
	}
	// 处理extCode,只能填写数字，保护程序运行
	if extCode != "" {
		_, err := strconv.Atoi(extCode)
		if err != nil {
			extCode = ""
		}
	}
	// 手机号切片
	destMobiles := strings.Split(mobiles, ",")
	phoneLength := len(destMobiles)
	if phoneLength > 1000 {
		s.logger.Println("user_id: " + corpId + "; user_ip: " + ip + " ; mobile: " + encodeMobiles + "; return_code: " + CommonMobileInfoError + "; explain: " + "提交的手机号码数量[" + strconv.Itoa(phoneLength) + "]大于1000个")
	}

	return HSResponse{
		"0",
		"success",
		nil,
	}
}
