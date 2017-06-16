package ukvd

import (
  "bytes"
  "net/http"
  "encoding/json"
  "time"
  "io/ioutil"
)

func NewUKVD(key string) *UKVD {
  ukvd := &UKVD{
      apiKey: key,
      dataPackage: "VehicleData",
  }
  return ukvd
}

func (this *UKVD) Search(vrm string) (*VehicleData, error) {
  url := "https://uk1.ukvehicledata.co.uk/api/datapackage/" + this.dataPackage
  params := make(map[string]string)
  params["v"] = "2"
  params["api_nullitems"] = "1"
  params["key_vrm"] = vrm

  b, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}
  buf := bytes.NewBuffer(b)

  req, err := http.NewRequest("GET", url, buf)
  if err != nil {
		return nil, err
  }
  req.Header.Add("cache-control", "no-cache")
  req.Header.Add("content-type", "application/json; charset=utf-8")
  req.Header.Add("Authorization", "ukvd-ipwhitelist " + this.apiKey)

  timeout := time.Duration(time.Duration(10) * time.Second)
  client := &http.Client{
			Timeout: timeout,
	}

  // Debug log request
	if this.log {
		s.log("--------------------------------------------------------------------------------")
		s.log("REQUEST")
		s.log("--------------------------------------------------------------------------------")
		s.log("Method:", req.Method)
		s.log("URL:", req.URL)
		s.log("Header:", req.Header)
		s.log("Form:", req.Form)
		s.log("Payload:")
		if r.RawPayload && s.Log && buf != nil {
			s.log(base64.StdEncoding.EncodeToString(buf.Bytes()))
		} else {
			s.log(pretty(r.Payload))
		}
	}


  resp, err := client.Do(req)
	if err != nil {
		return nil, err
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
  }
  result := &VehicleData{}
  if string(body) != "" {
		err = json.Unmarshal(body, result)
    if err != nil {
  		return nil, err
    }
  } else {
    return nil, errors.New("Upstream error")
  }
  return result, nil
}
