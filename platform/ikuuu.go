package platform

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type IKuuuVPN struct {
}

func (IKuuuVPN *IKuuuVPN) signIn(cookie string) (string, error) {
	url := "https://ikuuu.me/user/checkin"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Cookie", cookie)
	req.Header.Add("Content-Length", "0")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Referer", "https://ikuuu.me/user")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 解析响应的 JSON 数据并提取签到结果
	var resMap map[string]interface{}
	err = json.Unmarshal(body, &resMap)
	if err != nil {
		return "", err
	}

	msg, ok := resMap["msg"].(string)
	if !ok {
		return "", errors.New("无法获取签到结果")
	}

	return msg, nil
}

func (IKuuuVPN *IKuuuVPN) Run(pushPlusToken string, cookie string) {
	var title = "ikuuu签到"
	pushPlus := PushPlus{}

	msg, err := ik.signIn(cookie)
	if err != nil {
		pushPlus.Run(pushPlusToken, title, err.Error())
	} else {
		pushPlus.Run(pushPlusToken, title, msg)
	}
}
