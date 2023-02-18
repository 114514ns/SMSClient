package main

import (
	"awesomeProject/internel"
	_ "embed"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//go:embed cookie.txt
var cookie string

func main() {
	var client = new(SMSClient)
	client.session = "351976e9421a44bc9f94586b1569d214"
	//client.sendMessage("EZ4ENCE")
	var zhihu = new(internel.ZhihuClient)
	zhihu.Cookie = cookie
	zhihu.GetRecommend()
}

type SMSClient struct {
	session string
}

func (smsClient *SMSClient) sendPOST(msg string) {
	println("session: " + smsClient.session)

	url := "https://xiaoyuan.aoyadianzi.cn:7443/v1/parent/miniapp/command/sendMsg"
	method := "POST"

	payload := strings.NewReader(msg)
	fmt.Println(payload)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Host", "xiaoyuan.aoyadianzi.cn:7443")
	req.Header.Add("sessionId", "351976e9421a44bc9f94586b1569d214")
	req.Header.Add("from", "STU_PARENT")
	req.Header.Add("user-agent", "Chrome 114514")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
func (smsClient *SMSClient) sendMessage(msg string) {
	smsClient.sendPOST(`{commandType:14,commandMsg:{"ledFlag":1,"vibration":0,"sound":0,"smsType":5,"context":"` + msg + `","displayNum":1,"displayType":1},imei:862677060127893,cardId:134182}`)
}
