package commons

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
