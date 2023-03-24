package main

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomCode() string {
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random 6-digit code
	return strconv.Itoa(rand.Intn(900000) + 100000)
}