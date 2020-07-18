package auxiliaries

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"

type RandomService struct {
	Auxiliary
	SeededRand *rand.Rand
}

func (this RandomService) UpdateSeededRand() {
	this.SeededRand = rand.New(rand.NewSource(time.Now().UnixNano() + this.Configuration.RandomSeedOffset))
}

func (this RandomService) stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[this.SeededRand.Intn(len(charset))]
	}
	return string(b)
}

func (this RandomService) RandString(length int) string {
	return this.stringWithCharset(length, charset)
}

func (this RandomService) RandInt() int {
	return this.SeededRand.Int()
}

func (this RandomService) RandIntRange(min int, max int) int {
	return this.SeededRand.Intn(max-min+1) + min
}
