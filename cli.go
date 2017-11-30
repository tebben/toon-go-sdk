package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tebben/toon-go-sdk/auth"
	"github.com/tebben/toon-go-sdk/toon"
)

var (
	clientIDPtr         *string
	clientSecretPtr     *string
	usernamePtr         *string
	passwordPtr         *string
	providerPtr         *string
	callbackHostPtr     *string
	callbackPortPtr     *int
	callbackEndpointPtr *string
	commandPtr          *string
	startPtr            *int64
	endPtr              *int64
	intervalPtr         *string
)

var commands = []string{
	"GetAgreements",
	"GetStatus",
	"GetGasFlowData",
	"GetGasGraphData",
	"GetElectricityFlowData",
	"GetElectricityGraphData",
	"GetDistrictHeatGraphData",
}

func main() {
	// define custom command for -help
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Check https://developer.toon.eu/toonapi/apis for a better understanding of the Toon API")
		flag.PrintDefaults()
		printCommands()
	}

	clientIDPtr = flag.String("clientid", "", "App client id (https://developer.toon.eu/user/me/apps)")
	clientSecretPtr = flag.String("clientsecret", "", "App client secret (https://developer.toon.eu/user/me/apps)")
	usernamePtr = flag.String("username", "", "Username used to login to your provider services")
	passwordPtr = flag.String("password", "", "Password used to login to your provider services")
	providerPtr = flag.String("provider", "eneco", "Name of your provider, e.g eneco, viesgo")
	callbackHostPtr = flag.String("callbackhost", "127.0.0.1", "Host address to run the callback server on")
	callbackPortPtr = flag.Int("callbackport", 8080, "Host port to run the callback server on")
	callbackEndpointPtr = flag.String("callbackendpoint", "/oauthcallback", "Host endpoint")
	commandPtr = flag.String("command", "", "Command to run, check available command with -help")
	startPtr = flag.Int64("start", 0, "Start time for requested data: Unix timestamp in milliseconds")
	endPtr = flag.Int64("end", 0, "End time for requested data: Unix timestamp in milliseconds")
	intervalPtr = flag.String("interval", "", "Interval for requested data, possible values: hours, days, weeks, months, years")
	flag.Parse()

	checkFlags()
	fireCommand()
}

func printCommands() {
	fmt.Println("\nAvailable commands for -command flag")
	fmt.Println("------------------------------")
	for i, c := range commands {
		trail := " - "
		if i == len(commands)-1 {
			trail = "\n"
		}

		fmt.Printf("%s%s", c, trail)
	}
	fmt.Println("------------------------------")
}

func checkFlags() {
	required := []string{"clientid", "clientsecret", "username", "password", "command"}

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			fmt.Fprintf(os.Stderr, "missing required -%s argument/flag\n", req)
			os.Exit(2)
		}
	}
}

func fireCommand() {
	command := strings.ToLower(*commandPtr)
	found := false
	for _, c := range commands {
		if strings.ToLower(c) == command {
			found = true
			break
		}
	}

	if !found {
		log.Fatalf("Command %s not found, use -help to check supported command", command)
	}

	authenticator := auth.NewToonAuthenticator(
		*clientIDPtr,
		*clientSecretPtr,
		*providerPtr,
		fmt.Sprintf("http://%s:%v%s", *callbackHostPtr, *callbackPortPtr, *callbackEndpointPtr),
		"0.0.0.0",
		*callbackEndpointPtr,
		*callbackPortPtr)

	authenticator.StartGetToken(*usernamePtr, *passwordPtr)
	agreements, err := toon.GetAgreements(authenticator)
	if err != nil {
		printResponse(agreements, err)
	}

	ag := *agreements
	switch command {
	case "getagreements":
		{
			printResponse(agreements, err)
			break
		}
	case "getstatus":
		{
			data, err := toon.GetStatus(authenticator, ag[0].AgreementID)
			printResponse(data, err)
			break
		}
	case "getgasflowdata":
		{
			data, err := toon.GetGasFlowData(authenticator, ag[0].AgreementID, *startPtr, *endPtr)
			printResponse(data, err)
			break
		}
	case "getgasgraphdata":
		{
			interval, _ := stringToInterval(*intervalPtr)
			data, err := toon.GetGasGraphData(authenticator, ag[0].AgreementID, *startPtr, *endPtr, interval)
			printResponse(data, err)
			break
		}
	case "getelectricityflowdata":
		{
			data, err := toon.GetElectricityFlowData(authenticator, ag[0].AgreementID, *startPtr, *endPtr)
			printResponse(data, err)
			break
		}
	case "getelectricitygraphdata":
		{
			interval, _ := stringToInterval(*intervalPtr)
			data, err := toon.GetElectricityGraphData(authenticator, ag[0].AgreementID, *startPtr, *endPtr, interval)
			printResponse(data, err)
			break
		}
	case "getdistrictheatgraphdata":
		{
			interval, _ := stringToInterval(*intervalPtr)
			data, err := toon.GetDistrictHeatGraphData(authenticator, ag[0].AgreementID, *startPtr, *endPtr, interval)
			printResponse(data, err)
			break
		}
	}
}

func stringToInterval(intervalString string) (toon.Interval, error) {
	lowerIntervalString := strings.ToLower(intervalString)
	for _, interval := range toon.IntervalEnum {
		if interval.String() == lowerIntervalString {
			return interval, nil
		}
	}

	return 0, fmt.Errorf("Interval %v not supported", intervalString)
}

func printResponse(data interface{}, err *toon.ErrorResponse) {
	if err != nil {
		errString, _ := json.Marshal(err)
		fmt.Println(string(errString))
		return
	}

	dataString, _ := json.Marshal(data)
	fmt.Println(string(dataString))
}
