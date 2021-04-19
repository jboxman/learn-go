package main

import (
	"os"
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	os.Args = []string{"/dev/null", "some", "args"}
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	os.Args = []string{"/dev/null", "some", "args"}
	for i := 0; i < b.N; i++ {
		echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	os.Args = []string{"/dev/null", "some", "args"}
	for i := 0; i < b.N; i++ {
		echo3()
	}
}
