package ukvd

import (
	"testing"
)

func TestInvalidRegistration(t *testing.T) {
	vd := NewUKVD("71adfc85-6aa0-488d-b40d-7f2d63a56186")
	details, err := vd.Search("KLM123")
	fmt.Println("details")
}
