package auxiliaries

import (
	config "imsapi/config"
	"math/rand"
	"time"
)

// RandomGenerator works with random values
type RandomGenerator struct {
	SeededRand *rand.Rand
	Config     *config.Configuration
}

// UpdateSeededRand updates seeds for rand on which the random value is decided
func (rs *RandomGenerator) UpdateSeededRand(offset int64) {
	rs.SeededRand = rand.New(rand.NewSource(time.Now().UnixNano() + rs.Config.RandomSeedOffset + offset))
}

// RandStringWithAlphaNumeric will return a random string of specified length and AlphaNumeric characters
func (rs *RandomGenerator) RandStringWithAlphaNumeric(length int) string {
	return rs.stringWithCharset(length, "abcdefghijklmnopqrstuvwxyz"+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+"0123456789")
}

// RandInt will return a random interger
func (rs *RandomGenerator) RandInt() int {
	return rs.SeededRand.Int()
}

// RandIntRange will return a random interger with a specified length
func (rs *RandomGenerator) RandIntRange(min int, max int) int {
	return rs.SeededRand.Intn(max-min+1) + min
}

func (rs *RandomGenerator) stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rs.SeededRand.Intn(len(charset))]
	}
	return string(b)
}
