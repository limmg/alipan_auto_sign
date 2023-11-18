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
	req.Header.Add("authority", "ikuuu.me")
	req.Header.Add("method", "POST")
	req.Header.Add("path", "/user/checkin")
	req.Header.Add("scheme", "https")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("Content-Length", "0")
	req.Header.Add("Origin", "https://ikuuu.me")
	req.Header.Add("Referer", "https://ikuuu.me/user")
	req.Header.Add("Sec-Ch-Ua", "\"Microsoft Edge\";v=\"119\", \"Chromium\";v=\"119\", \"Not?A_Brand\";v=\"24\"")
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.0.0")
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

	msg, err := IKuuuVPN.signIn(cookie)
	if err != nil {
		pushPlus.Run(pushPlusToken, title, err.Error())
	} else {
		pushPlus.Run(pushPlusToken, title, msg)
	}
}
