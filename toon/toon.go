package toon

import (
	"fmt"

	"github.com/tebben/toon-go-sdk/auth"
)

var (
	apiEndpoint                   = "https://api.toon.eu/toon/v3"
	agreementEndpoint             = "/agreements"
	statusEndpoint                = "/status"
	gasFlowsEndpoint              = "/consumption/gas/flows"
	electricityGraphDataEndpoint  = "/consumption/electricity/data"
	districtHeatGraphDataEndpoint = "/consumption/districtheat/data"
	electricityFlowDataEndpoint   = "/consumption/electricity/flows"
	gasGraphDataEndpoint          = "/consumption/gas/data"
)

// GetAgreements returns the agreementID(s) that are associated with the utility customer.
// The agreementID is used in subsequent calls to access the data of one particular Toon.
func GetAgreements(auth *auth.ToonAuthenticator) (*Agreements, *ErrorResponse) {
	agreements := &Agreements{}
	err := get(constructEndpointURI(agreementEndpoint, nil, ""), auth, agreements, false)
	return agreements, err
}

// GetStatus returns returns information about current power usage, gas usage,
// thermostat information and thermostat programs aswell as connected devices.
func GetStatus(auth *auth.ToonAuthenticator, agreementID string) (*Status, *ErrorResponse) {
	status := &Status{}
	err := get(constructEndpointURI(statusEndpoint, nil, agreementID), auth, status, false)
	return status, err
}

// Consumption

// GetGasFlowData returns the gas consumption for a given time period in 5 minute intervals.
// The data is given for the time period between the given from- and toTime parameters.
// If no parameters are specified, the default value will be used, which is the last 24 hours
// start and end = unix timestamp in milliseconds, supply 0 for start and end when not using
func GetGasFlowData(auth *auth.ToonAuthenticator, agreementID string, start, end int64) (*FlowData, *ErrorResponse) {
	flowData := &FlowData{}
	err := get(constructEndpointURI(gasFlowsEndpoint, constructTimeParams(start, end, IntervalNone), agreementID), auth, flowData, false)
	return flowData, err
}

// GetElectricityGraphData returns the electricity consumption for a given time period. The data is given
// in the interval you specify and for the time period between the given from- and toTime parameters.
// If no parameters are specified, the default values will be used. The default time period is the last
// 24 hours and the default interval is hourly. Supply 0 for start and end when not using
func GetElectricityGraphData(auth *auth.ToonAuthenticator, agreementID string, start, end int64, interval Interval) (*ElectricityGraphData, *ErrorResponse) {
	graphData := &ElectricityGraphData{}
	err := get(constructEndpointURI(electricityGraphDataEndpoint, constructTimeParams(start, end, interval), agreementID), auth, graphData, false)
	return graphData, err
}

// GetDistrictHeatGraphData returns the districtheat consumption for a given time period. The data is given
// in the interval you specify and for the time period between the given from- and toTime parameters.
// If no parameters are specified, the default values will be used. The default time period is the last 24 hours
// and the default interval is hourly. Supply 0 for start and end when not using
func GetDistrictHeatGraphData(auth *auth.ToonAuthenticator, agreementID string, start, end int64, interval Interval) (*FlowData, *ErrorResponse) {
	flowData := &FlowData{}
	err := get(constructEndpointURI(districtHeatGraphDataEndpoint, constructTimeParams(start, end, interval), agreementID), auth, flowData, false)
	return flowData, err
}

// GetElectricityFlowData returns the electricity consumption for a given time period in 5 minute intervals.
// The data is given for the time period between the given from- and toTime parameters. The data will be
// returned in an array appended under the field hours. If no parameters are specified,
// the default time period will be used, which is the last 24 hours.
func GetElectricityFlowData(auth *auth.ToonAuthenticator, agreementID string, start, end int64) (*FlowData, *ErrorResponse) {
	flowData := &FlowData{}
	err := get(constructEndpointURI(electricityFlowDataEndpoint, constructTimeParams(start, end, IntervalNone), agreementID), auth, flowData, false)
	return flowData, err
}

// GetGasGraphData returns the gas consumption for a given time period. The data is given in the interval you specify
// and for the time period between the given from- and toTime parameters. If no parameters are specified,
// the default values will be used. The default time period is the last 24 hours and the default interval is hourly.
func GetGasGraphData(auth *auth.ToonAuthenticator, agreementID string, start, end int64, interval Interval) (*FlowData, *ErrorResponse) {
	flowData := &FlowData{}
	err := get(constructEndpointURI(gasGraphDataEndpoint, constructTimeParams(start, end, interval), agreementID), auth, flowData, false)
	return flowData, err
}

// ToDo: Consumption
// unsubscribePushEvent
// getWebhooks
// subscribeToPushEvent

func constructTimeParams(start, end int64, interval Interval) map[string]string {
	params := map[string]string{}
	if start != 0 {
		params["fromTime"] = fmt.Sprintf("%v", start)
	}

	if end != 0 {
		params["toTime"] = fmt.Sprintf("%v", start)
	}

	if interval.String() != "" && interval != IntervalNone {
		params["interval"] = fmt.Sprintf("%v", interval)
	}

	return params
}
