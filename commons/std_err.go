package commons

const (
	OKCode          int = 200
	BadRequest          = 4100
	CommonErr           = 5000
	UnAuthorized        = 5103
	NotValidUser        = 5104
	InvalidIdentify     = 5105
	UserExists          = 5106
)

type StdError struct {
	Code    int
	Message string
}

func (e *StdError) Error() string {
	if e == nil {
		return ""
	}

	return e.Message
}

func NewError(code int, msg string) error {
	return &StdError{
		Code:    code,
		Message: msg,
	}
}

func QuickStdError(msg string) error {
	return &StdError{
		Code:    -1,
		Message: msg,
	}
}
