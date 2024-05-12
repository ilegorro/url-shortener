package random

import (
	"math/rand"
	"strings"
	"time"
)

func NewRandomString(length int) string {
	var b strings.Builder
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		b.WriteByte('a' + byte(rnd.Intn('z'-'a'+1)))
	}

	return b.String()
}
