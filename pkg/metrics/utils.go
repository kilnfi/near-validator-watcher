package metrics

import "hash/fnv"

func BoolToFloat64(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

func StringToFloat64(s string) float64 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return float64(h.Sum32())
}
