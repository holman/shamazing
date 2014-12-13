package main

import "testing"

func TestLongestString(t *testing.T) {
	shas := []string{
		"e6dc64339f263e878b2294df00851245a3f8fb44",
		"a4bb41115fe96ac9cca229c231e9509044a4de37",
		"deadbeef671f9468df534662928d285fa40a7cbd",
		"72cf0cbf72dee39091598093adb1a4f9408b169e",
	}
	commits = shas
	result, chunk := findResult("[a-f]+")

	if result != "deadbeef671f9468df534662928d285fa40a7cbd" {
		t.Error("couldn't find the sha")
	}

	if chunk != "deadbeef" {
		t.Error("couldn't find the chunk")
	}
}

func TestLongestInteger(t *testing.T) {
	shas := []string{
		"e6dc64339f263e878b2294df00851245a3f8fb44",
		"a4bb41115fe96ac9cca229c231e9509044a4de37",
		"deadbeef671f9468df534662928d285fa40a7cbd",
		"72cf0cbf72dee39091598093adb1a4f9408b169e",
	}
	commits = shas
	result, chunk := findResult("[0-9]+")

	if result != "72cf0cbf72dee39091598093adb1a4f9408b169e" {
		t.Error("couldn't find the sha")
	}

	if chunk != "39091598093" {
		t.Error("couldn't find the chunk")
	}
}

func TestLongestRepeating(t *testing.T) {
	shas := []string{
		"e6dc64339f263e878b2294df00851245a3f8fb44",
		"a4bb41115fffffffe96ac9cca229c231e9509044",
		"deadbeef671f9468df534662928d285fa40a7cbd",
		"72cf0cbf72dee39091598093adb1a4f9408b169e",
	}
	commits = shas
	result, chunk := findResult(`(\\w)(\\1+)`)

	if result != "a4bb41115fffffffe96ac9cca229c231e9509044" {
		t.Error("couldn't find the sha")
	}

	if chunk != "fffffff" {
		t.Error("couldn't find the chunk")
	}
}
