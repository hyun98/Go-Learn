package errors

import "fmt"

const (
	NotFountUser = iota
	DBError
)

var errMessage = map[int64]string{
	NotFountUser: "not Fount User",
	DBError:      "DB Err",
}

func Errorf(code int64, args ...interface{}) error {
	if message, ok := errMessage[code]; ok {
		return fmt.Errorf("%s : %v", message, args)
	} else {
		return fmt.Errorf("not found err code")
	}
}
