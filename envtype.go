package citcallgo

// EnvironmentType value
type EnvironmentType int8

const (
	_ EnvironmentType = iota

	// Sandbox : represent sandbox environment
	Sandbox

	// Production : represent production environment
	Production
)


// TODO should read from config
var typeString = map[EnvironmentType]string{
	Sandbox:    "http://104.199.196.122/gateway/v3",
	Production: "http://104.199.196.122/gateway/v3",
}

// implement stringer
func (e EnvironmentType) String() string {
	for k, v := range typeString {
		if k == e {
			return v
		}
	}
	return "undefined"
}

// CreateSmsOtpURL : Create SMSOTP URL
func (e EnvironmentType) CreateSmsOtpURL() string {
	return e.String() + "/smsotp"
}

// CreateSmsOtpVerifyURL : Create VERIFY URL
func (e EnvironmentType) CreateSmsOtpVerifyURL() string {
	return e.String() + "/verify"
}

// CreateAsyncMiscallURL : Create Asynchronous Miscall URL
func (e EnvironmentType) CreateAsyncMiscallURL() string {
	return e.String() + "/asynccall"
}

func (e EnvironmentType) CreateSmsURL() string {
	return e.String() + "/sms"
}
