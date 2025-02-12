package id

import (
	"crypto/rand"
)

func Generate() string {
	return rand.Text()
}
