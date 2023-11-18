package main

import (
	"autoSign/platform"
	"os"
	"strings"
)

func main() {
	args := os.Args
	pushPlusToken := args[1]
	refreshTokens := args[2]
	//bilibiliCookies := args[3]
	ikuuuCookies := args[4]
	if refreshTokens != "null" {
		refreshTokenList := strings.Split(refreshTokens, ",")
		aliCloudDisk := platform.AliCloudDisk{}
		for _, refreshToken := range refreshTokenList {
			aliCloudDisk.Run(pushPlusToken, refreshToken)
		}
	}
	//if bilibiliCookies != "null" {
	//	bilibiliCookieList := strings.Split(bilibiliCookies, ",")
	//	bilibili := platform.Bilibili{}
	//	for _, bilibiliCookie := range bilibiliCookieList {
	//		bilibili.Run(pushPlusToken, bilibiliCookie)
	//	}
	//}
	if ikuuuCookies != "null" {
		ikuuuCookieList := strings.Split(ikuuuCookies, ",")
		iku := platform.IKuuuVPN{}
		for _, ikuuuCookie := range ikuuuCookieList {
			iku.Run(pushPlusToken, ikuuuCookie)
		}
	}

}
