package ekko_errors

import "fmt"

type RuntimeError struct {
	code    int32
	message string
}

var (
	InvalidRequest = RuntimeError{code: 10001, message: "Invalid Request: 'product' field is invalid"}
	ClockDrift     = RuntimeError{code: 10002, message: "clock drift"}
	CurrencyExceed = RuntimeError{code: 10003, message: "currency is exceeded"}
)

func (e *RuntimeError) Error() string {
	return fmt.Sprintf("[%d] %s", e.code, e.message)
}
