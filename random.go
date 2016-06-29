package main

import (
	"math/rand"
	"time"
)

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890 abcdefghijklmnopqrstuvwxyz" +
	"~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"

const maxlen = 10

var known map[string]bool

func initRandStrings() {
	known = make(map[string]bool)
	rand.Seed(time.Now().UTC().UnixNano())
}

func removefromKnownRandStrings(s string) {
	known[s] = false
}

func randStrings(N int) []string {
	if known == nil {
		initRandStrings()
	}
	r := make([]string, N)
	ri := 0
	buf := make([]byte, maxlen)

	for i := 0; i < N; i++ {
	retry:
		//l := rand.Intn(Maxlen)
		for j := 0; j < maxlen; j++ {
			buf[j] = chars[rand.Intn(len(chars))]
		}
		s := string(buf[0:maxlen])
		if known[s] {
			goto retry
		}
		known[s] = true
		r[ri] = s
		ri++
	}
	return r
}
