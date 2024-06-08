package err

import "fmt"

const (
	BindingFailure = "bind 실패 : "
	ServerErr      = "server 에러"
	NoDocument     = "데이터 없음 : "
)

func ErrorMsg(status string, err error) string {
	return fmt.Sprint(status+"%s", err.Error())
}
