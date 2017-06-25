package ukvd

type request struct {
	DataKeys struct {
		Vrm string `json:"Vrm"`
	} `json:"DataKeys"`
	PackageID       string `json:"PackageId"`
	PackageVersion  int    `json:"PackageVersion"`
	RequestGUID     string `json:"RequestGuid"`
	ResponseVersion int    `json:"ResponseVersion"`
}

type statusInfo struct {
	Lookup struct {
		AdviceTextList []interface{} `json:"AdviceTextList"`
		StatusCode     string        `json:"StatusCode"`
		StatusMessage  string        `json:"StatusMessage"`
	} `json:"Lookup"`
}
