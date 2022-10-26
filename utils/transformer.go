package utils

import "strconv"

func ConvertToUint64(v any) uint64 {
	switch _v := v.(type) {
	case *uint64:
		return *_v
	case float64:
		return uint64(v.(float64))
	case string:
		n, _ := strconv.ParseUint(v.(string), 10, 64)
		return n
	default:
		return v.(uint64)
	}
}
