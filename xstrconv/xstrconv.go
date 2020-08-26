package xstrconv

import "strconv"

var Itoa = strconv.Itoa

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	must(err)
	return i
}

func ParseBool(str string) bool {
	b, err := strconv.ParseBool(str)
	must(err)
	return b
}

var FormatFloat = FormatFloat64

func FormatFloat64(f float64, prec int) string {
	return strconv.FormatFloat(f, 'f', prec, 64)
}

var ParseFloat = ParseFloat64

func ParseFloat32(s string) float64 {
	f, err := strconv.ParseFloat(s, 32)
	must(err)
	return f
}

func ParseFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	must(err)
	return f
}

func ParseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 0)
	must(err)
	return int(i)
}

func ParseInt8(s string) int8 {
	i, err := strconv.ParseInt(s, 10, 8)
	must(err)
	return int8(i)
}

func ParseInt32(s string) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	must(err)
	return int32(i)
}

func ParseInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	must(err)
	return i
}

func ParseUint(s string, base int, bitSize int) uint {
	i, err := strconv.ParseUint(s, 10, 0)
	must(err)
	return uint(i)
}

func ParseUint8(s string, base int, bitSize int) uint8 {
	i, err := strconv.ParseUint(s, 10, 8)
	must(err)
	return uint8(i)
}

func ParseUint64(s string, base int, bitSize int) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	must(err)
	return i
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
