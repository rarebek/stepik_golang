package main

import (
	"fmt"
	"time"
)

func main() {
	const now = 1589570165
	var (
		min int
		sec int
	)
	fmt.Scanf("%d мин. %d сек.", &min, &sec)
	seconds := 60*min + sec
	seconds += now
	t := time.Unix(int64(seconds), 0).UTC()

	fmt.Println(t.Format(time.UnixDate))
}
