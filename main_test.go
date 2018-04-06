package main

import (
	"testing"
)

func BenchmarkHelloDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helloDefault()
	}
}
func BenchmarkHelloGlog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helloGlog()
	}
}
func BenchmarkHelloLogrus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helloLogrus()
	}
}
func BenchmarkHelloZap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helloZap()
	}
}
func BenchmarkHelloSeelog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helloSeelog()
	}
}
