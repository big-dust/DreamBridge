package proxy

import (
	"encoding/json"
	"fmt"
	"github.com/big-dust/DreamBridge/internal/pkg/common"
	"io"
	"net/http"
	"net/url"
	"sync"
)

var (
	IP   string
	PORT int
	mu   sync.RWMutex
)

type GenProxyIPResponse struct {
	Code    int    `json:"code"`
	Data    []Data `json:"data"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}
type Data struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

func genProxyIP() (*GenProxyIPResponse, error) {
	resp, err := http.Get(common.CONFIG.String("proxy.link"))
	if err != nil {
		return nil, err
	}
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	IPResp := &GenProxyIPResponse{}
	if err = json.Unmarshal(bodyText, IPResp); err != nil {
		return nil, err
	}
	return IPResp, nil
}

func ChangeHttpProxyIP() {
	resp, _ := genProxyIP()
	mu.Lock()
	IP = resp.Data[0].IP
	PORT = resp.Data[0].Port
	mu.Unlock()
}

func NewHttpClientWithProxy() (*http.Client, error) {
	if !common.CONFIG.Bool("proxy.switchon") {
		return &http.Client{}, nil
	}
	mu.RLock()
	ip := IP
	port := PORT
	mu.RUnlock()

	urli := url.URL{}
	urlproxy, _ := urli.Parse(fmt.Sprintf("http://%s:%d", ip, port))

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		},
	}
	return client, nil
}
