package scraper

import (
	"encoding/json"
	"github.com/big-dust/DreamBridge/internal/crawler/model"
	"io"
	"net/http"
	"strings"
)

func SchoolList() (*model.SchoolListResponse, error) {
	client := &http.Client{}
	var data = strings.NewReader(`{"keyword":"","page":1,"province_id":"","ranktype":"","request_type":1,"signsafe":"a6beb63405f371aece65cadfb263f006","size":100000,"top_school_id":"[2461]","type":"","uri":"apidata/api/gkv3/school/lists"}`)
	req, err := http.NewRequest("POST", "https://api.zjzw.cn/web/api/?keyword=&page=6&province_id=&ranktype=&request_type=1&size=20&top_school_id=\\[2461\\]&type=&uri=apidata/api/gkv3/school/lists&signsafe=a6beb63405f371aece65cadfb263f006", data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("authority", "api.zjzw.cn")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", "https://www.gaokao.cn")
	req.Header.Set("referer", "https://www.gaokao.cn/")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Microsoft Edge";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 Edg/121.0.0.0")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	schoolList := &model.SchoolListResponse{}
	if err = json.Unmarshal(bodyText, schoolList); err != nil {
		return nil, err
	}
	return schoolList, nil
}
