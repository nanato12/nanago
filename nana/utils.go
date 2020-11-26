package nana

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randString ...
func randString(n int) string {
	rs1Letters := []rune("0123456789abcdef")
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(b)
}

// geneDeviceID ...
func geneDeviceID() string {
	return randString(4) + "-" + randString(2) + "-" + randString(2) + "-" + randString(2) + "-" + randString(6)
}

// geneDeviceID ...
func geneAppsflyerID() string {
	return randString(13) + "-" + randString(7)
}
