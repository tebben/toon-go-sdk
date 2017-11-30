package toon

type Interval int

const (
	IntervalNone Interval = 1 + iota
	IntervalHours
	IntervalDays
	IntervalWeeks
	IntervalMonths
	IntervalYears
)

var intervals = [...]string{
	"none",
	"hours",
	"days",
	"weeks",
	"months",
	"years",
}

var IntervalEnum = []Interval{IntervalNone, IntervalHours, IntervalDays, IntervalWeeks, IntervalMonths, IntervalYears}

// String() function will return the name of an interval
func (i Interval) String() string {
	return intervals[i-1]
}

// ErrorResponse description
type ErrorResponse struct {
	Fault Fault `json:"fault"`
}

// Fault description
type Fault struct {
	Faultstring string      `json:"faultstring"`
	Detail      FaultDetail `json:"detail"`
}

// FaultDetail description
type FaultDetail struct {
	Errorcode string `json:"errorcode"`
}

// Agreements contains an array of agreements for logged in user
type Agreements []Agreement

// Agreement contains general information about a Toon device
type Agreement struct {
	AgreementID            string `json:"agreementId"`
	AgreementIDChecksum    string `json:"agreementIdChecksum"`
	Street                 string `json:"street"`
	HouseNumber            string `json:"houseNumber"`
	PostalCode             string `json:"agrepostalCodeementId"`
	City                   string `json:"city"`
	HeatingType            string `json:"heatingType"`
	DisplayCommonName      string `json:"displayCommonName"`
	DisplayHardwareVersion string `json:"displayHardwareVersion"`
	DisplaySoftwareVersion string `json:"displaySoftwareVersion"`
	IsToonSolar            bool   `json:"isToonSolar"`
	IsToonly               bool   `json:"isToonly"`
}

type FlowData struct {
	Hours  []FlowDataValue `json:"hours"`
	Days   []FlowDataValue `json:"days"`
	Weeks  []FlowDataValue `json:"weeks"`
	Months []FlowDataValue `json:"months"`
	Years  []FlowDataValue `json:"years"`
}

type FlowDataValue struct {
	Timestamp int64   `json:"timestamp"`
	Unit      string  `json:"unit"`
	Value     float64 `json:"value"`
}

type ElectricityGraphData struct {
	Hours  []GraphData `json:"hours"`
	Days   []GraphData `json:"days"`
	Weeks  []GraphData `json:"weeks"`
	Months []GraphData `json:"months"`
	Years  []GraphData `json:"years"`
}

type GraphData struct {
	Timestamp int64   `json:"timestamp"`
	Unit      string  `json:"unit"`
	Peak      float64 `json:"peak"`
	OffPeak   float64 `json:"offPeak"`
}

// Status description
type Status struct {
	ThermostatStates      ThermostatStates `json:"thermostatStates"`
	ThermostatInfo        ThermostatInfo   `json:"thermostatInfo"`
	SmokeDetectors        SmokeDetectors   `json:"smokeDetectors"`
	DeviceConfigInfo      DeviceConfigInfo `json:"deviceConfigInfo"`
	DeviceStatusInfo      DeviceStatusInfo `json:"deviceStatusInfo"`
	PowerUsage            PowerUsage       `json:"powerUsage"`
	GasUage               GasUsage         `json:"gasUsage"`
	LastUpdateFromDisplay int64            `json:"lastUpdateFromDisplay"`
	ServerTime            int64            `json:"serverTime"`
}

// ThermostatStates description
type ThermostatStates struct {
	States []ThermostatState `json:"state"`
}

// ThermostatState description
type ThermostatState struct {
	ID        int `json:"id"`
	TempValue int `json:"tempValue"`
	Dhq       int `json:"dhw"`
}

// ThermostatInfo description
type ThermostatInfo struct {
	CurrentSetpoint        int    `json:"currentSetpoint"`
	CurrentDisplayTemp     int    `json:"currentDisplayTemp"`
	ProgramState           int    `json:"programState"`
	ActiveState            int    `json:"activeState"`
	NextProgram            int    `json:"nextProgram"`
	NextState              int    `json:"nextState"`
	NextTime               int    `json:"nextTime"`
	NextSetpoint           int    `json:"nextSetpoint"`
	ErrorFound             int    `json:"errorFound"`
	BoilerModuleConnected  int    `json:"boilerModuleConnected"`
	RealSetpoint           int    `json:"realSetpoint"`
	BurnerInfo             string `json:"burnerInfo"`
	OtCommError            string `json:"otCommError"`
	CurrentModulationLevel int    `json:"currentModulationLevel"`
	HaveOTBoiler           int    `json:"haveOTBoiler"`
}

// SmokeDetectors description
type SmokeDetectors struct {
	// Have no smoke detector myself yet and unable to find which info returns
	Devices []interface{} `json:"device"`
}

// DeviceConfigInfo description
type DeviceConfigInfo struct {
	Configs []DeviceConfig `json:"device"`
}

// DeviceConfig description
type DeviceConfig struct {
	DevUUID           string `json:"devUUID"`
	DevType           string `json:"devType"`
	Name              string `json:"name"`
	FlowGraphUUID     string `json:"flowGraphUuid"`
	QuantityGraphUUID string `json:"quantityGraphUuid"`
	Position          int    `json:"position"`
	InSwitchAll       int    `json:"inSwitchAll"`
	InSwitchSchedule  int    `json:"inSwitchSchedule"`
	SwitchLocked      string `json:"switchLocked"`
	UsageCapable      string `json:"usageCapable"`
	CurrentState      string `json:"currentState"`
	RgbColor          string `json:"rgbColor"`
	Zwuuid            string `json:"zwuuid"`
}

// DeviceStatusInfo description
type DeviceStatusInfo struct {
	Status           []DeviceStatus   `json:"device"`
	InSwitchAllTotal InSwitchAllTotal `json:"inSwitchAllTotal"`
}

// DeviceStatus description
type DeviceStatus struct {
	DevUUID            string      `json:"devUUID"`
	Name               string      `json:"name"`
	CurrentUsage       int         `json:"currentUsage"`
	DayUsage           float64     `json:"dayUsage"`
	AvgUsage           float64     `json:"avgUsage"`
	CurrentState       int         `json:"currentState"`
	IsConnected        int         `json:"isConnected"`
	NetworkHealthState interface{} `json:"networkHealthState"`
	RgbColor           string      `json:"rgbColor"`
}

// InSwitchAllTotal description
type InSwitchAllTotal struct {
	CurrentState int     `json:"currentState"`
	CurrentUsage float64 `json:"currentUsage"`
	DayUsage     float64 `json:"dayUsage"`
	AvgUsage     float64 `json:"avgUsage"`
}

// PowerUsage description
type PowerUsage struct {
	Value                int     `json:"value"`
	DayCost              float64 `json:"dayCost"`
	ValueProduced        int     `json:"valueProduced"`
	DayCostProduced      float64 `json:"dayCostProduced"`
	ValueSolar           int     `json:"valueSolar"`
	MaxSolar             int     `json:"maxSolar"`
	DayCostSolar         float64 `json:"dayCostSolar"`
	AvgSolarValue        float64 `json:"avgSolarValue"`
	AvgValue             float64 `json:"avgValue"`
	AvgDayValue          float64 `json:"avgDayValue"`
	AvgProduValue        float64 `json:"avgProduValue"`
	AvgDayProduValue     float64 `json:"avgDayProduValue"`
	MeterReading         int     `json:"meterReading"`
	MeterReadingLow      int     `json:"meterReadingLow"`
	MeterReadingProdu    int     `json:"meterReadingProdu"`
	MeterReadingLowProdu int     `json:"meterReadingLowProdu"`
	DayUsage             int     `json:"dayUsage"`
	DayLowUsage          int     `json:"dayLowUsage"`
	TodayLowestUsage     int     `json:"todayLowestUsage"`
	IsSmart              int     `json:"isSmart"`
}

// GasUsage description
type GasUsage struct {
	Value        int     `json:"value"`
	DayCost      float64 `json:"dayCost"`
	AvgValue     float64 `json:"avgValue"`
	MeterReading int     `json:"meterReading"`
	AvgDayValue  float64 `json:"avgDayValue"`
	DayUsage     int     `json:"dayUsage"`
	IsSmart      int     `json:"isSmart"`
}
