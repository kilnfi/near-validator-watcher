package watcher

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
)

func GetStakeFromString(s string) float64 {
	if len(s) == 1 {
		return 0
	}
	l := len(s) - 19 - 5
	v, err := strconv.ParseFloat(s[0:l], 64)
	if err != nil {
		fmt.Println(err)
	}
	return float64(v)
}

func GetFloatFromString(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return v
}

func HashString(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func prettyPrintFloat(f float64) string {
	if f == float64(int(f)) {
		return fmt.Sprintf("%.0f", f)
	}
	return fmt.Sprintf("%.2f", f)
}

func prettyPrintAccountID(accountID string) string {
	return trimPoolSuffix(
		accountID,
		".pool.f863973.m0",
		".poolv1.near",
		".pool.near",
	)
}

func trimPoolSuffix(accountID string, suffixes ...string) string {
	for _, suffix := range suffixes {
		if strings.HasSuffix(accountID, suffix) {
			return strings.TrimSuffix(accountID, suffix)
		}
	}
	return accountID
}
