package main

import (
	"math/rand"
	"time"
)

var chars = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i',
	'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r',
	's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
}

// Increase below in the future
var seeded = false

func GenerateID(n int) string {
	if !seeded {
		rand.Seed(time.Now().UnixNano())
	}
	id := make([]byte, n)
	for i := 0; i < n; i++ {
		id[i] = chars[rand.Intn(len(chars))]
	}
	return string(id)
}
