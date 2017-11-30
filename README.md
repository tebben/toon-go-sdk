toon-go-sdk can be used in your project to interact with the Toon API endpoint at https://api.toon.eu/toon/v3  
There is a cli.go file which can be build and started, however this is used to do some testing, maybe there will be a full cli to interact with the Toon API in the future.

## how-to

Create an account at [https://developer.toon.eu/](https://developer.toon.eu/)  
After registering create an app and use "http://127.0.0.1:8080/oauthcallback" as Callback URL.  

Create a ToonAuthenticator, you can find your client id and secret from the created app at https://developer.toon.eu/. Provider should be something as eneco or viesgo
```
authenticator := auth.NewToonAuthenticator(
		{clientID},
		{clientSecret},
		{provider},
		"http://127.0.0.1:8080/oauthcallback",
		"0.0.0.0",
		"/oauthcallback",
		8080)
```

Get a token using your provider credentials, for instance the credentials u use to login to the Eneco website. 
```
authenticator.StartGetToken(username, password)
```

GetAgreements example
```
agreements, err := toon.GetAgreements(authenticator)
```

Note: An initial agreement API call is needed since all other calls require an AgreementID

GetStatus example
```
ag := *agreements
data, err := toon.GetStatus(authenticator, ag[0].AgreementID)
```

## API implementation status
This project is work in progress

### Production
- [ ] agreementId-production-electricity-flows-get
- [ ] agreementId-production-electricity-data-get
- [ ] getElectricityGraphDataEneco
- [ ] getElectricityProductionAndDelivery

### Agreements
- [x] getAgreements

### Status
- [x] getStatus

### Consumption
- [x] getGasFlowData
- [x] getElectricityGraphData
- [x] getDistrictHeatGraphData
- [x] getElectricityFlowData
- [x] getGasGraphData
- [ ] unsubscribePushEvent
- [ ] getWebhooks
- [ ] subscribeToPushEvent

### Thermostat
- [ ] getThermostatPrograms
- [ ] updateThermostatPrograms
- [ ] setThermostatState
- [ ] getThermostatStates
- [ ] updateCurrentTemperature
- [ ] getCurrentTemperature

### Devices
- [ ] getDeviceConfiguration
- [ ] updateDeviceConfiguration
- [ ] getDevicesGraphData
- [ ] getDevicesConfiguration
- [ ] updateDevicesConfiguration
- [ ] getDevicesFlows