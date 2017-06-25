package ukvd

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type UKVD struct {
	apiKey string
	log    bool
}

func NewUKVD(key string) *UKVD {
	ukvd := &UKVD{
		apiKey: key,
		log:    false,
	}
	return ukvd
}

func (this *UKVD) SetLog(log bool) {
	this.log = log
}

func (this *UKVD) search(vrm string, dataPackage string, result interface{}) error {
	url := "https://uk1.ukvehicledata.co.uk/api/datapackage/" + dataPackage
	params := make(map[string]string)
	params["v"] = "2"
	params["api_nullitems"] = "0"
	params["key_vrm"] = vrm

	b, err := json.Marshal(&params)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(b)

	req, err := http.NewRequest("GET", url, buf)
	if err != nil {
		return err
	}
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", "ukvd-ipwhitelist "+this.apiKey)

	timeout := time.Duration(time.Duration(10) * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	// Debug log request
	if this.log {
		log.Println("--------------------------------------------------------------------------------")
		log.Println("REQUEST")
		log.Println("--------------------------------------------------------------------------------")
		log.Println("Method:", req.Method)
		log.Println("URL:", req.URL)
		log.Println("Header:", req.Header)
		if buf != nil {
			log.Println("Payload:")
			log.Println(buf.String())
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

  // Debug log response
	if this.log {
    log.Println("Response:")
    log.Println(string(body))
  }

	if string(body) != "" {
		err = json.Unmarshal(body, result)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Upstream error")
	}
	return nil
}

func (this *UKVD) VDICheckBasic(vrm string) (*CheckBasicDataItems, error) {
	result := CheckBasic{}
	err := this.search(vrm, "VDICheckBasic", &result)
	if err != nil {
		return nil, err
	}

	err = parseResultStatus(result.Response.StatusCode)
	if err != nil {
		return nil, err
	}
	return &result.Response.DataItems, nil
}

func (this *UKVD) VDICheckFinance(vrm string) (*CheckFinanceDataItems, error) {
	result := CheckFinance{}
	err := this.search(vrm, "VDICheckFinance", &result)
	if err != nil {
		return nil, err
	}

	err = parseResultStatus(result.Response.StatusCode)
	if err != nil {
		return nil, err
	}
	return &result.Response.DataItems, nil
}

func (this *UKVD) VDICheckFull(vrm string) (*CheckFullDataItems, error) {
	result := CheckFull{}
	err := this.search(vrm, "VDICheckFull", &result)
	if err != nil {
		return nil, err
	}

	err = parseResultStatus(result.Response.StatusCode)
	if err != nil {
		return nil, err
	}
	return &result.Response.DataItems, nil
}

func parseResultStatus(statusCode string) error {
	switch statusCode {
	case "Failure":
		return APIFailureError
	case "InvalidPackageId":
		return InvalidPackageIdError
	case "ItemNotFound":
		return ItemNotFoundError
	case "KeyInvalid":
		return KeyInvalidError
	case "ServiceUnavailable":
		return ServiceUnavailableError
	case "InvalidPackageAccessClass":
		return InvalidPackageAccessClassError
	default:
		return nil
	}
}
