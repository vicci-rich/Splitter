package math

import (
	"fmt"
	"strconv"
)

func Float64ToInt64(f float64) int64 {
	floatStr := fmt.Sprintf("%.0f", f)
	inst, _ := strconv.ParseInt(floatStr, 10, 64)
	return inst
}

func Float64ToUint64(f float64) uint64 {
	floatStr := fmt.Sprintf("%.0f", f)
	inst, _ := strconv.ParseUint(floatStr, 10, 64)
	return inst
}
