package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	want := "Code.education Rocks!!!"
	result := greeting()

	if want != result {
		t.Fail()
	}
}

func TestEmptyMessage(t *testing.T) {
	want := "Code.education Rocks!!!"
	result := greeting()

	if want != result {
		t.Fail()
	}
}
