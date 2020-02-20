package main

import (
	"fmt"

	citcallgo "github.com/grosenia/citcall-client"
	viper "github.com/spf13/viper"
)

var citcallClient citcallgo.Client
var smsGateway citcallgo.SmsGateway

func main() {
	fmt.Println("Load Config...")

	viper.SetConfigType("props")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	fmt.Println("Load Config success")
	fmt.Println("Setup client")

	setupClient()

	// Example
	MSISDN := "6281287736665"
	// Token := "1000"
	ValidTime := 300 // 5 minutes
	LimitTry := 3

	var request = &citcallgo.CitcallAsyncMiscallOtpRequest{
		Gateway:   0,
		MSISDN:    MSISDN,
		ValidTime: ValidTime,
		LimitTry:  LimitTry,
	}

	resp, err := smsGateway.RequestMiscallOtp(request)

	if err != nil {
		fmt.Println("Error server")
		return
	}

	if resp.ErrorStatus {
		// Ada error
		fmt.Println("Error: ", resp.Error())
		return
	}

	fmt.Println("Citcall response: ")
	// fmt.Println(resp)
	fmt.Println("TrxID: ", resp.TrxID)
	fmt.Println("Token: ", resp.Token)
}

func setupClient() {
	citcallClient = citcallgo.NewClient()

	// TODO: should put in config file
	citcallClient.SecretAPIKey = viper.GetString("CITCALL_KEY")
	citcallClient.APIEnvType = citcallgo.Sandbox
	citcallClient.LogLevel = 3

	smsGateway = citcallgo.SmsGateway{
		Client: citcallClient,
	}
}
