package main

import (
	"testing"
)

func init() {
	Setup()
}
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
func BenchmarkBigvalueDefaut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bigvalueDefault()
	}
}
func BenchmarkBigvalueGlog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bigvalueGlog()
	}
}
func BenchmarkBigvalueLogrus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bigvalueLogrus()
	}
}
func BenchmarkBigvalueZap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bigvalueZap()
	}
}
func BenchmarkBigvalueSeelog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bigvalueSeelog()
	}
}
