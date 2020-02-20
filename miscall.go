package citcallgo

import (
	"bytes"
	"encoding/json"

	"github.com/nbs-go/clog"
	_ "github.com/nbs-go/clogrus"
)

// RequestMiscallOtp request miscall for OTP
func (gateway *SmsGateway) RequestMiscallOtp(req *CitcallAsyncMiscallOtpRequest) (response *CitcallMiscallOtpResponse, err error) {
	log := clog.Get()
	resp := CitcallMiscallOtpResponse{}
	jsonReq, _ := json.Marshal(req)

	path := gateway.Client.APIEnvType.CreateAsyncMiscallURL()
	httpRequest, err := gateway.Client.NewRequest("POST", path, bytes.NewBuffer(jsonReq))

	if err != nil {
		return nil, err
	}

	httpStatus, err := gateway.Client.ExecuteRequest(httpRequest, &resp)
	if err != nil {
		log.Error("Error charging ", err)
		return nil, err
	}

	if httpStatus != 200 {
		resp.ErrorStatus = true
	} else {
		if resp.ResponseCode != CitcallOK {
			resp.ErrorStatus = true
		} else {
			resp.ErrorStatus = false
		}
	}

	return &resp, nil
}
