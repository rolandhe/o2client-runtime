package o2client_runtime

import (
	"strconv"
	"strings"
)

func ConvertIntBoolToString(v bool) string {
	return strconv.FormatBool(v)
}

func ConvertIntBoolListToString(v []bool) string {
	return convertListToString(v, ConvertIntBoolToString)
}

func ConvertIntByteToString(v byte) string {
	return strconv.Itoa(int(v))
}

func ConvertIntByteListToString(v []byte) string {
	return convertListToString(v, ConvertIntByteToString)
}

func ConvertInt16ToString(v int16) string {
	return strconv.Itoa(int(v))
}

func ConvertInt16ListToString(v []int16) string {
	return convertListToString(v, ConvertInt16ToString)
}
func ConvertInt32ToString(v int32) string {
	return strconv.Itoa(int(v))
}

func ConvertInt32ListToString(v []int32) string {
	return convertListToString(v, ConvertInt32ToString)
}

func ConvertInt64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func ConvertInt64ListToString(v []int64) string {
	return convertListToString(v, ConvertInt64ToString)
}

func ConvertDoubleToString(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func ConvertDoubleListToString(v []float64) string {
	return convertListToString(v, ConvertDoubleToString)
}

type baseType interface {
	bool | byte | int16 | int32 | int64 | float64
}

func convertListToString[T baseType](ar []T, fn func(v T) string) string {
	if len(ar) == 0 {
		return ""
	}
	elems := make([]string, 0, len(ar))
	for _, v := range ar {
		elems = append(elems, fn(v))
	}

	return strings.Join(elems, ",")
}


