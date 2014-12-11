package main

import "testing"

func TestLongestString(t *testing.T) {
	shas := []string{
		"e6dc64339f263e878b2294df00851245a3f8fb44",
		"a4bb41115fe96ac9cca229c231e9509044a4de37",
		"deadbeef671f9468df534662928d285fa40a7cbd",
		"72cf0cbf72dee39091598093adb1a4f9408b169e",
	}
	longest := findLongestString(shas)

	if longest != "deadbeef671f9468df534662928d285fa40a7cbd" {
		t.Error("couldn't find the longest sha")
	}
}
