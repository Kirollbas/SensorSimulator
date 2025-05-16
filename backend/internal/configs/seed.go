package configs

import (
	"sync"
	"time"

	"github.com/db47h/rand64/v3/splitmix64"
)

var prngInstance splitmix64.Rng
var prngOnce sync.Once

func GetSeed() int64 {
	configOnce.Do(func() {
		prngInstance = splitmix64.Rng{}
		prngInstance.Seed(time.Now().UTC().UnixNano())
	})

	return prngInstance.Int63()
}
