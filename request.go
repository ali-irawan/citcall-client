package citcallgo

const (
	CitcallOK                          = 0
	CitcallWrongMethod                 = 99
	CitcallAuthorizationFailed         = 98
	CitcallMissingParameter            = 88
	CitcallUserIDNotFound              = 77
	CitcallSenderIDClosed              = 44
	CitcallUnknownErrorOrFailed        = 6
	CitcallServiceTemporaryUnavailable = 34
	CitcallWrongPassword               = 76
	CitcallMaintenanceInProgress       = 66
	CitcallInsufficientAmount          = 14
)

// CitcallSmsOtpRequest citcall SMSOTP Request
type CitcallSmsOtpRequest struct {
	SenderID  string `json:"senderid"`
	MSISDN    string `json:"msisdn"`
	Text      string `json:"text"`
	Token     string `json:"token"`
	ValidTime int    `json:"valid_time"`
	LimitTry  int    `json:"limit_try"`
}

// CitcallSmsRequest citcall SMSOTP Request
type CitcallSmsRequest struct {
	SenderID string `json:"senderid"`
	MSISDN   string `json:"msisdn"`
	Text     string `json:"text"`
}

// CitcallSmsOtpVerifyRequest citcall SMSOTP Verify Request
type CitcallSmsOtpVerifyRequest struct {
	TrxID  string `json:"trxid"`
	MSISDN string `json:"msisdn"`
	Token  string `json:"token"`
}

// CitcallAsyncMiscallOtpRequest cicall async miscall request
type CitcallAsyncMiscallOtpRequest struct {
	MSISDN    string `json:"msisdn"`
	Gateway   int    `json:"gateway"`
	ValidTime int    `json:"valid_time"`
	LimitTry  int    `json:"limit_try"`
}
