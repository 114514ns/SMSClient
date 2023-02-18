package internel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ZhihuClient struct {
	Cookie string
}
type Rates struct {
	Base   string `json:"base currency"`
	Symbol string `json:"destination currency"`
}

func (client *ZhihuClient) GetRecommend() {
	request, _ := http.NewRequest("GET", "https://www.zhihu.com/api/v3/feed/topstory/recommend?action=down&ad_interval=-10&after_id=5&desktop=true&page_number=2", nil)
	request.Header.Set("Cookie", client.Cookie)
	h := http.Client{}
	resp, err := h.Do(request)
	var result map[string]interface{}
	all, _ := ioutil.ReadAll(resp.Body)
	//println("返回内容：  " + string(all))
	json.Unmarshal(all, &result)
	if err == nil {
		fmt.Println(result["fresh_text"])
	}
}
