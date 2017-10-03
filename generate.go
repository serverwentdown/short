package main

import (
	"time"
	"math/rand"
)

var chars = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i',
	'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r',
	's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'1', '2', '3', '4',	'5', '6', '7', '8', '9', '0',
}

// Increase below in the future
var length = 4
var seeded = false

func GenerateID() string {
	if (!seeded) {
		rand.Seed(time.Now().UnixNano())
	}
	id := make([]byte, length)
	for i := 0; i < length; i++ {
		id[i] = chars[rand.Intn(len(chars))]
	}
	return string(id)
}
