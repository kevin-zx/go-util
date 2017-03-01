package randomUtil

import (
	"math/rand"
	"time"
	"errors"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
func GetRandomInt(start int, end int) (int,error) {
	if end < start {
		return 0,errors.New("start great than end");
	}
	randMax := end - start
	return start+r.Intn(randMax),nil
}