package source

import (
	"strings"
	mockClient "weeklyNews/http"

	"github.com/valyala/fastjson"
)

const apiURL = "https://api.cntv.cn"

func GetVideoListByColumn(limit int32) (replyString string, err error) {
	uri := "/NewVideo/getVideoListByColumn"
	args := map[string]interface{}{
		"id":        "TOPC1451559180488841",
		"n":         limit,
		"sort":      "desc",
		"mode":      0,
		"serviceId": "tvcctv",
		//"cb":        "Callback",
	}
	head := map[string]interface{}{
		"Accept":             "*/*",
		"Accept-Encoding":    "gzip, deflate, br",
		"Accept-Language":    "zh-CN,zh;q=0.9",
		"Referer":            "https://tv.cctv.com/",
		"Sec-Ch-Ua":          "\"Google Chrome\";v=\"119\", \"Chromium\";v=\"119\", \"Not?A_Brand\";v=\"24\"",
		"Sec-Ch-Ua-Mobile":   "?0",
		"Sec-Ch-Ua-Platform": "\"macOS\"",
		"Sec-Fetch-Dest":     "script",
		"Sec-Fetch-Mode":     "no-cors",
		"Sec-Fetch-Site":     "cross-site",
		"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
	}
	reply, err := mockClient.Get(apiURL+uri, args, head)
	if err != nil {
		return
	}
	return string(reply), err
}

func GetLatestVideoGuid() (string, error) {
	reply, err := GetVideoListByColumn(20)
	if err != nil {
		return "", err
	}
	var jsonParser fastjson.Parser
	v, err := jsonParser.Parse(reply)
	if err != nil {
		return "", err
	}

	return strings.Trim(v.Get("data").GetArray("list")[0].Get("guid").String(), `"`), err
}

type GetVideoListByColumnReply struct {
	Total int64        `json:"total"`
	List  []*VideoInfo `json:"list"`
}

type VideoInfo struct {
	Brief     string `json:"brief"`
	Mode      int    `json:"mode"`
	Image     string `json:"image"`
	FocusDate int64  `json:"focus_date"`
	Length    string `json:"length"`
	Guid      string `json:"guid"`
	ID        string `json:"id"`
	Time      string `json:"time"`
	Title     string `json:"title"`
	URL       string `json:"url"`
}
