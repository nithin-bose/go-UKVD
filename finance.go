package ukvd

type CheckFinance struct {
	Request  request `json:"Request"`
	Response struct {
		DataItems         CheckFinanceDataItems `json:"DataItems"`
		StatusCode        string                `json:"StatusCode"`
		StatusInformation statusInfo            `json:"StatusInformation"`
		StatusMessage     string                `json:"StatusMessage"`
	} `json:"Response"`
}

type CheckFinanceDataItems struct {
	Data                interface{}   `json:"Data"`
	FinanceRecordCount  int           `json:"FinanceRecordCount"`
	FinanceRecordList   []interface{} `json:"FinanceRecordList"`
	LookupStatusCode    interface{}   `json:"LookupStatusCode"`
	LookupStatusMessage interface{}   `json:"LookupStatusMessage"`
}
