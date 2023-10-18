package watcher

import (
	"fmt"
	"strings"
)

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
