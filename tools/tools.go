package tools

import (
	"fmt"
	"time"
)

func MySQLDate() string {
	_time := time.Now()

	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", _time.Year(), _time.Month(), _time.Day(), _time.Hour(), _time.Minute(), _time.Second())
}
