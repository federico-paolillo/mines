package matchmaking

import "time"

func NextVersion() Matchversion {
	return time.Now().Unix()
}
