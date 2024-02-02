package scraper

import (
	"encoding/json"
	"fmt"
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
	schoolListResp := &model.SchoolListResponse{}
	if err = json.Unmarshal(bodyText, schoolListResp); err != nil {
		return nil, err
	}
	return schoolListResp, nil
}

func SchoolInfo(schoolId int) (*model.SchoolInfoResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://static-data.gaokao.cn/www/2.0/school/%d/info.json", schoolId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Microsoft Edge";v="121", "Chromium";v="121"`)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://www.gaokao.cn/")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 Edg/121.0.0.0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	schoolInfoResp := &model.SchoolInfoResponse{}
	if err = json.Unmarshal(bodyText, schoolInfoResp); err != nil {
		return nil, err
	}
	return schoolInfoResp, nil
}

func JobDetail(schoolId int) (*model.JobDetailResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://static-data.gaokao.cn/www/2.0/school/%d/pc_jobdetail.json", schoolId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Microsoft Edge";v="121", "Chromium";v="121"`)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://www.gaokao.cn/")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 Edg/121.0.0.0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	jobDetailResp := &model.JobDetailResponse{}
	if err = json.Unmarshal(bodyText, jobDetailResp); err != nil {
		return nil, err
	}
	return jobDetailResp, nil
}

func SpecialInfo(schoolId int) (*model.SpecialResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://static-data.gaokao.cn/www/2.0/school/%d/pc_special.json", schoolId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Microsoft Edge";v="121", "Chromium";v="121"`)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://www.gaokao.cn/")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 Edg/121.0.0.0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	nationFeatureResp := &model.SpecialResponse{}
	if err = json.Unmarshal(bodyText, nationFeatureResp); err != nil {
		return nil, err
	}
	return nationFeatureResp, nil
}

// 湖北locationId = 42
func HistoryRecruit(schoolId int, locationId int) (*model.HistoryRecruitResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://static-data.gaokao.cn/www/2.0/history_recruit/%d/%d.json", schoolId, locationId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://www.gaokao.cn/")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	historyRecruitResp := &model.HistoryRecruitResponse{}
	if err := json.Unmarshal(bodyText, historyRecruitResp); err != nil {
		return nil, err
	}
	return historyRecruitResp, nil
}

func HistoryAdmission(schoolId int, locationId int) (*model.HistoryAdmissionResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://static-data.gaokao.cn/www/2.0/history_admission/%d/%d.json", schoolId, locationId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://www.gaokao.cn/")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	historyAdmissionResp := &model.HistoryAdmissionResponse{}
	if err = json.Unmarshal(bodyText, historyAdmissionResp); err != nil {
		return nil, err
	}
	return historyAdmissionResp, nil
}
