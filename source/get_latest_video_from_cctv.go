package source

import (
	"fmt"
	"net/http"
	"os"
	"weeklyNews/models"
)

func GetLatestVideoFromCCTV(w http.ResponseWriter, r *http.Request) {
	// 1. 获取视频详情
	guid, err := GetLatestVideoGuid()
	if err != nil {
		models.JSONErrResponse(r, w, "900001", err)
		return
	}
	// 生成文件夹
	//os.Mkdir(fmt.Sprintf("/%s"))
	os.Mkdir(guid, 0777)

	// 保存切片文件，生成对应的m3u8
	uri := fmt.Sprintf(hlsURIFormat, "main", guid) + "main.m3u8"
	GetHLSFile(uri, guid, map[string]interface{}{
		"maxbr": 2048,
	})

}
