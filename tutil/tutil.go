package tutil

import (
	"math/rand"
	"regexp"
	"time"

	"github.com/siguovo/project_S/stools"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandNum() int {
	na := rand.Int()
	nb := rand.Int()
	return stools.SAdd(na, nb)
}

func CheckURL(url string) bool {
	matched, _ := regexp.Match("^(http|https)://", []byte(url))

	return matched
}

// vesrion 1.0.1
func VTestGetVersion() string {
	return "v1.0.1"
}
