package gotest

import (
	"log/slog"
	"testing"

	"github.com/sirupsen/logrus"
)

// go test -bench=BenchmarkSLog -benchtime=5s
//
// 2024/09/16 14:13:55 INFO printing using slog n=91672
// 2024/09/16 14:13:55 INFO printing using slog n=91673
// 2024/09/16 14:13:55 INFO printing using slog n=91674
// 2024/09/16 14:13:55 INFO printing using slog n=91675
//
//	91676             74197 ns/op
func BenchmarkSLog(B *testing.B) {
	for n := 0; n < B.N; n++ {
		slog.Info("printing using slog", "n", n)
	}
}

// go test -bench=BenchmarkLogrus -benchtime=5s
//
// time="2024-09-16T14:13:06-05:00" level=info msg="printing using logrus84239"
// time="2024-09-16T14:13:06-05:00" level=info msg="printing using logrus84240"
// time="2024-09-16T14:13:06-05:00" level=info msg="printing using logrus84241"
// time="2024-09-16T14:13:06-05:00" level=info msg="printing using logrus84242"
//
//	84243             71056 ns/op
func BenchmarkLogrus(B *testing.B) {
	for n := 0; n < B.N; n++ {
		logrus.Info("printing using logrus", n)
	}
}

// slog is a 4% reduction ???
