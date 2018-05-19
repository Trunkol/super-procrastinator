package helper

import (
	"strconv"
	"time"
)

func CastingUnix(date int) time.Time {

	i, err := strconv.ParseInt(string(date), 10, 64)
	if err != nil {
		panic(err)
	}

	tm := time.Unix(i, 0)

	return tm
}
