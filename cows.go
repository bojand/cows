package cows

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// All returns an array of cow strings
func All() []string {
	file, err := Asset("data/cows.txt")
	if err != nil {
		return nil
	}
	str := string(file)
	re := regexp.MustCompile("\n$")
	str = re.ReplaceAllLiteralString(str, "")
	return strings.Split(str, "\n\n\n")
}

// Random returns a random cow string
func Random() string {
	all := All()
	count := len(all)
	if count == 0 {
		return ""
	}
	return all[randomInt(0, count)]
}
