package ukvd

type UKVD struct {
  apiKey string
  dataPackage string
}

type MyJsonName struct {
	Request struct {
		DataKeys struct {
			Vrm string `json:"Vrm"`
		} `json:"DataKeys"`
		PackageID       string `json:"PackageId"`
		PackageVersion  int    `json:"PackageVersion"`
		RequestGUID     string `json:"RequestGuid"`
		ResponseVersion int    `json:"ResponseVersion"`
	} `json:"Request"`
	Response struct {
		DataItems struct {
			CertificateOfDestructionIssued interface{}   `json:"CertificateOfDestructionIssued"`
			Colour                         string        `json:"Colour"`
			ColourChangeCount              interface{}   `json:"ColourChangeCount"`
			Data                           interface{}   `json:"Data"`
			DateFirstRegistered            string        `json:"DateFirstRegistered"`
			DateFirstRegisteredUk          string        `json:"DateFirstRegisteredUk"`
			EngineCapacity                 int           `json:"EngineCapacity"`
			ExportDate                     interface{}   `json:"ExportDate"`
			Exported                       bool          `json:"Exported"`
			FinanceRecordCount             int           `json:"FinanceRecordCount"`
			FinanceRecordList              []interface{} `json:"FinanceRecordList"`
			FuelType                       string        `json:"FuelType"`
			GearCount                      int           `json:"GearCount"`
			HighRiskRecordCount            int           `json:"HighRiskRecordCount"`
			HighRiskRecordList             interface{}   `json:"HighRiskRecordList"`
			ImportDate                     interface{}   `json:"ImportDate"`
			ImportUsedBeforeUkRegistration interface{}   `json:"ImportUsedBeforeUkRegistration"`
			Imported                       bool          `json:"Imported"`
			ImportedFromOutsideEu          bool          `json:"ImportedFromOutsideEu"`
			LatestColourChangeDate         interface{}   `json:"LatestColourChangeDate"`
			LatestKeeperChangeDate         string        `json:"LatestKeeperChangeDate"`
			LatestV5cIssuedDate            string        `json:"LatestV5cIssuedDate"`
			LookupStatusCode               interface{}   `json:"LookupStatusCode"`
			LookupStatusMessage            interface{}   `json:"LookupStatusMessage"`
			Make                           string        `json:"Make"`
			MileageAnomalyDetected         interface{}   `json:"MileageAnomalyDetected"`
			MileageRecordCount             int           `json:"MileageRecordCount"`
			MileageRecordList              interface{}   `json:"MileageRecordList"`
			Model                          string        `json:"Model"`
			PlateChangeCount               int           `json:"PlateChangeCount"`
			PlateChangeList                interface{}   `json:"PlateChangeList"`
			PreviousColour                 interface{}   `json:"PreviousColour"`
			PreviousKeeperCount            int           `json:"PreviousKeeperCount"`
			PreviousKeepers                interface{}   `json:"PreviousKeepers"`
			ScrapDate                      interface{}   `json:"ScrapDate"`
			Scrapped                       bool          `json:"Scrapped"`
			Stolen                         bool          `json:"Stolen"`
			StolenContactNumber            interface{}   `json:"StolenContactNumber"`
			StolenDate                     interface{}   `json:"StolenDate"`
			StolenInfoSource               interface{}   `json:"StolenInfoSource"`
			StolenPoliceForce              interface{}   `json:"StolenPoliceForce"`
			StolenStatus                   interface{}   `json:"StolenStatus"`
			TransmissionType               string        `json:"TransmissionType"`
			VicTestDate                    interface{}   `json:"VicTestDate"`
			VicTestResult                  interface{}   `json:"VicTestResult"`
			VicTested                      interface{}   `json:"VicTested"`
			VinLast5                       string        `json:"VinLast5"`
			Vrm                            string        `json:"Vrm"`
			WriteOffCategory               interface{}   `json:"WriteOffCategory"`
			WriteOffDate                   interface{}   `json:"WriteOffDate"`
			WriteOffRecordCount            interface{}   `json:"WriteOffRecordCount"`
			WriteOffRecordList             interface{}   `json:"WriteOffRecordList"`
			WrittenOff                     bool          `json:"WrittenOff"`
			YearOfManufacture              string        `json:"YearOfManufacture"`
		} `json:"DataItems"`
		StatusCode        string `json:"StatusCode"`
		StatusInformation struct {
			Lookup struct {
				AdviceTextList []interface{} `json:"AdviceTextList"`
				StatusCode     string        `json:"StatusCode"`
				StatusMessage  string        `json:"StatusMessage"`
			} `json:"Lookup"`
		} `json:"StatusInformation"`
		StatusMessage string `json:"StatusMessage"`
	} `json:"Response"`
}
