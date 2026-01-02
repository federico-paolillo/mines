package memory

import (
	"time"
)

func NextVersion() uint64 {
	//nolint:gosec // An int64 fits in an uint64 and UnixNano is always positive
	return uint64(time.Now().UnixNano())
}
