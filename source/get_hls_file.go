package source

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"
	mockClient "weeklyNews/http"

	"github.com/grafov/m3u8"
	"github.com/panjf2000/ants/v2"
)

type taskFunc func()

const hlsURL = "https://hls.cntv.lxdns.com"
const hlsURIFormat = hlsURL + "/asp/hls" + "/%s/" + "0303000a/3/default" + "/%s/"
const tsFilePath = "%s/%s"

func GetHLSFile(url string, guid string, params map[string]interface{}) {
	reply, err := mockClient.Get(url, params, nil)
	if err != nil {
		return
	}

	p, listType, err := m3u8.DecodeFrom(bytes.NewBuffer(reply), true)
	if err != nil {
		return
	}
	switch listType {
	case m3u8.MEDIA:
		re := regexp.MustCompile(`/(\d+)\.m3u8$`)
		bandWidth := re.FindStringSubmatch(url)
		mediapl := p.(*m3u8.MediaPlaylist)
		//fmt.Println(mediapl)
		filePath := fmt.Sprintf(tsFilePath, guid, bandWidth[1])
		os.Mkdir(filePath, 0777)
		downloadSegment(bandWidth[1], guid, filePath, mediapl)
	case m3u8.MASTER:
		masterpl := p.(*m3u8.MasterPlaylist)
		for _, variant := range masterpl.Variants {
			if variant != nil {
				GetHLSFile(hlsURL+variant.URI, guid, nil)
			}
		}
	}
}

func downloadSegment(bandWidth string, guid string, filePath string, mediapl *m3u8.MediaPlaylist) {

	p, _ := ants.NewPool(10)
	defer p.Release()

	var wg sync.WaitGroup

	tsURL := fmt.Sprintf(hlsURIFormat, bandWidth, guid)
	isExist, _ := PathExists(filePath)
	if !isExist {
		os.Mkdir(filePath, 0777)
	}

	for index, segment := range mediapl.Segments {
		if segment != nil {
			wg.Add(1)
			_ = p.Submit(downloadTask(tsURL+segment.URI, index, filePath, &wg))
		}
	}
	wg.Wait()
}

func downloadTask(uri string, index int, filePath string, wg *sync.WaitGroup) taskFunc {
	return func() {
		res, err := mockClient.Response(uri, nil, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()
		out, err := os.Create(fmt.Sprintf("%s/%d.ts", filePath, index))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer out.Close()
		_, err = io.Copy(out, res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		wg.Done()
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
