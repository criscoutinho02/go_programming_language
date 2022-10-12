package main

import "testing"

var args []string = []string{"arg1", "args2", "arg3"}

func BenchmarkPrintByJoin(b *testing.B) {

	for i := 0; i < b.N; i++ {
		printByJoin(args)
	}
}

func BenchmarkPrintByLoop(b *testing.B) {

	for i := 0; i < b.N; i++ {
		printByLoop(args)
	}
}

func BenchmarkPrintByRange(b *testing.B) {

	for i := 0; i < b.N; i++ {
		printByRange(args)
	}
}
