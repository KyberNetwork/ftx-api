package ftxapi

import (
	"fmt"
	"strconv"
)

func int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func intToString(i int) string {
	return strconv.Itoa(i)
}

func float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func boolToString(b bool) string {
	return strconv.FormatBool(b)
}

func StringToPointer(s string) *string {
	return &s
}

func endPointWithFormat(template string, params ...interface{}) string {
	return fmt.Sprintf(template, params...)
}
