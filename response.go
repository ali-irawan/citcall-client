package citcallgo

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// A FlexInt is an int that can be unmarshalled from a JSON field
// that has either a number or a string value.
// E.g. if the json field contains an string "42", the
// FlexInt value will be "42".
type FlexInt int

// UnmarshalJSON implements the json.Unmarshaler interface, which
// allows us to ingest values of any json type as an int and run our custom conversion
func (fi *FlexInt) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return json.Unmarshal(b, (*int)(fi))
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*fi = FlexInt(i)
	return nil
}

// CitcallSmsOtpResponse citcall SMSOTP Response
type CitcallSmsOtpResponse struct {
	ResponseCode FlexInt `json:"rc"`
	Info         string  `json:"info"`
	SmsCount     int     `json:"sms_count"`
	MSISDN       string  `json:"msisdn"`
	Text         string  `json:"text"`
	TrxID        string  `json:"trxid"`
	Currency     string  `json:"currency"`
	Price        int     `json:"price"`
	Token        string     `json:"token"`
	ErrorStatus  bool    `json:"-"`
}

func (e CitcallSmsOtpResponse) Error() string {
	return fmt.Sprintf("[%d] %s", e.ResponseCode, e.Info)
}

// CitcallSmsOtpResponse citcall SMSOTP Response
type CitcallSmsResponse struct {
	ResponseCode FlexInt `json:"rc"`
	Info         string  `json:"info"`
	SmsCount     int     `json:"sms_count"`
	MSISDN       string  `json:"msisdn"`
	Text         string  `json:"text"`
	TrxID        string  `json:"trxid"`
	Currency     string  `json:"currency"`
	Price        int     `json:"price"`
	Token        int     `json:"token"`
	ErrorStatus  bool    `json:"-"`
}

func (e CitcallSmsResponse) Error() string {
	return fmt.Sprintf("[%d] %s", e.ResponseCode, e.Info)
}

// CitcallMiscallOtpResponse citcall MISCALLOTP Response
type CitcallMiscallOtpResponse struct {
	ResponseCode FlexInt `json:"rc"`
	TrxID        string  `json:"trxid"`
	MSISDN       string  `json:"msisdn"`
	Token        string  `json:"token"`
	Gateway      int     `json:"gateway"`
	ErrorStatus  bool    `json:"-"`
}

func (e CitcallMiscallOtpResponse) Error() string {
	return fmt.Sprintf("[%d] %s", e.ResponseCode, "Error in miscall OTP")
}

// CitcallSmsOtpVerifyResponse citcall SMSOTP Verify Response
type CitcallSmsOtpVerifyResponse struct {
	ResponseCode FlexInt `json:"rc"`
	Info         string  `json:"info"`
	TrxID        string  `json:"trxid"`
	TrxIDVerify  string  `json:"trxid_verify"`
	Token        string  `json:"token"`
	TokenVerify  string  `json:"token_verify"`

	ErrorStatus bool `json:"-"`
}

func (e CitcallSmsOtpVerifyResponse) Error() string {
	return fmt.Sprintf("[%d] %s", e.ResponseCode, e.Info)
}
