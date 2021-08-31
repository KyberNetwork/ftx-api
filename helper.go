package ftxapi

import (
	"fmt"
	"time"
)

func timeToTimestampMS(t time.Time) int64 {
	return t.Unix() * 1000
}

func Int64ToString(i int64) string {
	return fmt.Sprintf("%d", i)
}

func IntToString(i int) string {
	return fmt.Sprintf("%d", i)
}

func Float64ToString(f float64) string {
	return fmt.Sprintf("%f", f)
}

func BoolToString(b bool) string {
	return fmt.Sprintf("%v", b)
}

func IntPointer(i int) *int {
	return &i
}

func StringPointer(s string) *string {
	return &s
}

func Int64Pointer(i int64) *int64 {
	return &i
}

func endPointWithFormat(template string, params ...interface{}) string {
	return fmt.Sprintf(template, params...)
}
