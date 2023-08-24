package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetIpProvince(key string) (string, error) {
	r, err := http.Get(fmt.Sprintf("https://restapi.amap.com/v3/ip?key=%v", key))
	if err != nil {
		return "", err
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	var res resData
	json.Unmarshal(b, &res)
	if strings.Trim(res.Province, " ") == "" {
		res.Province = "火星"
	}
	return res.Province, nil
}

type resData struct {
	Status   string `json:"status"`
	Province string `json:"province"`
	Info     string `json:"info"`
}
