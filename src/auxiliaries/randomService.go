package auxiliaries

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"

// RandomService works with random values
type RandomService struct {
	Auxiliary
	SeededRand *rand.Rand
}

// UpdateSeededRand updates seeds for rand on which the random value is decided
func (rs RandomService) UpdateSeededRand() {
	rs.SeededRand = rand.New(rand.NewSource(time.Now().UnixNano() + rs.Configuration.RandomSeedOffset))
}

// RandString will a random string of specified length
func (rs RandomService) RandString(length int) string {
	return rs.stringWithCharset(length, charset)
}

// RandInt will a random interger
func (rs RandomService) RandInt() int {
	return rs.SeededRand.Int()
}

// RandIntRange will a random interger with a specified length
func (rs RandomService) RandIntRange(min int, max int) int {
	return rs.SeededRand.Intn(max-min+1) + min
}

func (rs RandomService) stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rs.SeededRand.Intn(len(charset))]
	}
	return string(b)
}
