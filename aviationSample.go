package cctools

import "gopkg.in/mgo.v2/bson"

// Create Aviation Document Struct
type TransactionAviation struct {
	Id                    bson.ObjectId `bson:"_id,omitempty"`
	CreatedBy             string        `json:"createdBy"`
	DealName              string        `json:"dealName"`
	SerialNumber          string        `json:"serialNumber"`
	EngineModelno1        string        `json:"engineModelno1"`
	EngineModelno2        string        `json:"engineModelno2"`
	EntityNumber          string        `json:"entityNumber"`
	Date                  string        `json:"date"`
	ConveyanceNumber      string        `json:"conveyanceNumber"`
	Payment               string        `json:"payment"`
	PurchasePrice         string        `json:"purchasePrice"`
	DepositAmountPP       string        `json:"depositAmountPP"`
	LiabilityAmount       string        `json:"liabilityAmount"`
	RecordingLease        string        `json:"recordingLease"`
	DateLeasesuppno1      string        `json:"dateLeasesuppno1"`
	RecordingLeasesuppno1 string        `json:"recordingLeasesuppno1"`
	DateLeaseext          string        `json:"dateLeaseext"`
	RecordingLeaseext     string        `json:"recordingLeaseext"`
	ConveyanceLeaseext    string        `json:"conveyanceLeaseext"`
	EngineModel           string        `json:"engineModel"`
	Term                  string        `json:"term"`
	Delivery              string        `json:"delivery"`
	DepositAmount         string        `json:"depositAmount"`
	EngineAdjustment6     string        `json:"engineAdjustment6"`
	EngineAdjustment7     string        `json:"engineAdjustment7"`
	EngineAdjustment8     string        `json:"engineAdjustment8"`
	EngineAdjustment9     string        `json:"engineAdjustment9"`
	EngineAdjustment10    string        `json:"engineAdjustment10"`
	EngineAdjustment11    string        `json:"engineAdjustment11"`
	EngineAdjustment12    string        `json:"engineAdjustment12"`
	EngineAdjustment13    string        `json:"engineAdjustment13"`
	EngineAdjustment14    string        `json:"engineAdjustment14"`
	EngineRent            string        `json:"engineRent"`
	EngineLLP             string        `json:"engineLLP"`
	AirframeRent4c6       string        `json:"airframeRent4c6"`
	AirframeRent8c12      string        `json:"airframeRent8c12"`
	LandingRent           string        `json:"landingRent"`
	InsuranceValue        string        `json:"insuranceValue"`
	RegistrationNumber    string        `json:"registrationNumber"`
	ClosingDate           string        `json:"closingDate"`
	Fname                 string        `json:"fname"`
	Newfname              string        `json:"newfname"`
	Data                  string        `json:"data"`
}
