package ukvd

type CheckBasic struct {
	Request  request `json:"Request"`
	Response struct {
		DataItems         CheckBasicDataItems `json:"DataItems"`
		StatusCode        string              `json:"StatusCode"`
		StatusInformation statusInfo          `json:"StatusInformation"`
		StatusMessage     string              `json:"StatusMessage"`
	} `json:"Response"`
}

type CheckBasicDataItems struct {
	Colour                 string `json:"Colour"`
	DateFirstRegistered    string `json:"DateFirstRegistered"`
	DateFirstRegisteredUk  string `json:"DateFirstRegisteredUk"`
	EngineCapacity         int    `json:"EngineCapacity"`
	Exported               bool   `json:"Exported"`
	FuelType               string `json:"FuelType"`
	GearCount              int    `json:"GearCount"`
	HighRiskRecordCount    int    `json:"HighRiskRecordCount"`
	Imported               bool   `json:"Imported"`
	ImportedFromOutsideEu  bool   `json:"ImportedFromOutsideEu"`
	LatestKeeperChangeDate string `json:"LatestKeeperChangeDate"`
	LatestV5cIssuedDate    string `json:"LatestV5cIssuedDate"`
	Make                   string `json:"Make"`
	MileageRecordCount     int    `json:"MileageRecordCount"`
	Model                  string `json:"Model"`
	PlateChangeCount       int    `json:"PlateChangeCount"`
	PreviousKeeperCount    int    `json:"PreviousKeeperCount"`
	Scrapped               bool   `json:"Scrapped"`
	Stolen                 bool   `json:"Stolen"`
	TransmissionType       string `json:"TransmissionType"`
	VinLast5               string `json:"VinLast5"`
	Vrm                    string `json:"Vrm"`
	WrittenOff             bool   `json:"WrittenOff"`
	YearOfManufacture      string `json:"YearOfManufacture"`
}
