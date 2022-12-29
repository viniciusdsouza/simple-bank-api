package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().Unix())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney(min, max float64) string {
	f := min + rand.Float64()*(max-min)
	s := fmt.Sprintf("%v", f)
	return s
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currecies := []string{"BRL", "EUR", "USD", "CAD"}
	n := len(currecies)
	return currecies[rand.Intn(n)]
}
