package gotest

import (
	"log/slog"
	"testing"

	"github.com/sirupsen/logrus"
)

// go test -bench=BenchmarkSLog -count=5
// 17044            125824 ns/op
func BenchmarkSLog(B *testing.B) {
	for n := 0; n < B.N; n++ {
		slog.Info("printing using slog", "n", n)
	}
}

// go test -bench=BenchmarkLogrus -count=5
// 18296             75298 ns/op
func BenchmarkLogrus(B *testing.B) {
	for n := 0; n < B.N; n++ {
		logrus.Info("printing using logrus", n)
	}
}

// slog is a 40% reduction compared to logrus
