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
	SenderID := "citcall"
	MSISDN := "6281287736665"	
	Text := "Kode OTP Anda 12345 berlaku 5 menit"
	Token := "1000"
	ValidTime := 300 // 5 minutes
	LimitTry := 3 

	var request = &citcallgo.CitcallSmsOtpRequest{
		SenderID: SenderID,
		MSISDN: MSISDN,
		Text: Text,
		Token: Token,
		ValidTime: ValidTime,
		LimitTry: LimitTry,
	}
  
	resp, err := smsGateway.SendSmsOtp(request)

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
	fmt.Println(resp)
	fmt.Println("TrxID: ", resp.TrxID)
	fmt.Println("Token: ", resp.Token)

	fmt.Println("Test Verify Fail: ")

	var verifyRequest = &citcallgo.CitcallSmsOtpVerifyRequest{
		TrxID: resp.TrxID,
		MSISDN: MSISDN,
		Token: "9999",
	}
	respVerify, err := smsGateway.VerifySmsOtp(verifyRequest)

	if err != nil {
		fmt.Println("Error server")
	}

	if respVerify.ErrorStatus {
		// Ada error
		fmt.Println("Error: ", respVerify.Error())
	}

	fmt.Println("Citcall response: ")
	fmt.Println(respVerify)
	fmt.Println("Error Status: ", respVerify.ErrorStatus)
	fmt.Println("Error Code: ", respVerify.ResponseCode)
	fmt.Println("Error Info: ", respVerify.Info)


	fmt.Println("Test Verify Success: ")
	verifyRequest = &citcallgo.CitcallSmsOtpVerifyRequest{
		TrxID: resp.TrxID,
		MSISDN: MSISDN,
		Token: Token,
	}
	respVerifySuccess, err := smsGateway.VerifySmsOtp(verifyRequest)

	if err != nil {
		fmt.Println("Error server")
	}

	if respVerifySuccess.ErrorStatus {
		// Ada error
		fmt.Println("Error: ", respVerifySuccess.Error())
	}

	fmt.Println("Citcall response: ")
	fmt.Println(respVerifySuccess)
	fmt.Println("Error Status: ", respVerifySuccess.ErrorStatus)
	fmt.Println("Error Code: ", respVerifySuccess.ResponseCode)
	fmt.Println("Error Info: ", respVerifySuccess.Info)

	// 
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
