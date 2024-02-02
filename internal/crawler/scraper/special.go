package scraper

import (
	"encoding/json"
	"fmt"
	"github.com/big-dust/DreamBridge/internal/crawler/response"
	"io"
	"net/http"
)

func SpecialInfo(schoolId int) (*response.SpecialResponse, error) {
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
	nationFeatureResp := &response.SpecialResponse{}
	if err = json.Unmarshal(bodyText, nationFeatureResp); err != nil {
		return nil, err
	}
	return nationFeatureResp, nil
}

// 湖北locationId = 42
func HistoryRecruit(schoolId int, locationId int) (*response.HistoryRecruitResponse, error) {
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
	historyRecruitResp := &response.HistoryRecruitResponse{}
	if err := json.Unmarshal(bodyText, historyRecruitResp); err != nil {
		return nil, err
	}
	return historyRecruitResp, nil
}

func HistoryAdmission(schoolId int, locationId int) (*response.HistoryAdmissionResponse, error) {
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
	historyAdmissionResp := &response.HistoryAdmissionResponse{}
	if err = json.Unmarshal(bodyText, historyAdmissionResp); err != nil {
		return nil, err
	}
	return historyAdmissionResp, nil
}
