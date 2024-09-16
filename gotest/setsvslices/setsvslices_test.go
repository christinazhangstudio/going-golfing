package gotest

import (
	"strings"
	"testing"
)

var m = map[rune]bool{'&': true, '$': true, '@': true, '=': true, ';': true, ':': true, '+': true, ' ': true, ',': true, '?': true, '\\': true, '^': true, '`': true, '>': true, '<': true, '{': true, '}': true, '[': true, ']': true, '#': true, '%': true, '"': true, '~': true, '|': true, ' ': true, '°': true}
var sl = []string{"&", "$", "@", "=", ";", ":", "+", " ", ",", "?", "\\", "^", "`", ">", "<",
	"{", "}", "[", "]", "#", "%", "\"", "~", "|", " ", "°"}

const input = "test-&-::;;;#####"

func BenchmarkMapBool(B *testing.B) {
	for n := 0; n < B.N; n++ {
		modified := func(r rune) rune {
			if m[r] {
				return '_'
			}
			return r
		}
		strings.Map(modified, input)

		//fmt.Println(s) 	// BTW fmt.Println takes 60,000 - 70,000 ns!
	}
}

func BenchmarkSlices(B *testing.B) {
	for n := 0; n < B.N; n++ {
		s := input
		for _, v := range sl {
			if strings.Contains(input, v) {
				s = strings.ReplaceAll(
					s,
					v,
					"_")
			}
		}

		//fmt.Println(s)	// BTW fmt.Println takes 60,000 - 70,000 ns!
	}
}

// [number of iterations in time allowed] [ns/op]
// (btw trial just means running same command over again, -count does the same though)
// go test -bench=BenchmarkMapBool -benchtime=10s
//	trial 1: BenchmarkMapBool-16     69301228               177.9 ns/op
//	trial 2: BenchmarkMapBool-16     69708310               172.6 ns/op
//	trial 3: BenchmarkMapBool-16     63817159               188.1 ns/op

// go test -bench=BenchmarkSlices -benchtime=10s
//  trial 1: BenchmarkSlices-16      36230910               331.7 ns/op
//	trial 2: BenchmarkSlices-16      34980547               351.8 ns/op
//	trial 3: BenchmarkSlices-16      35162598               353.6 ns/op

// using benchstat:
// go test -bench=BenchmarkSlices -count=10 > old.txt
// go test -bench=BenchmarkMapBool -count=10 > new.txt
// benchstat slices.txt mapbool.txt
// btw powershell's default encoding for output redirection
// is UTF-16, not UTF-8 (https://stackoverflow.com/questions/40098771/changing-powershells-default-output-encoding-to-utf-8)
// this is unsupported by benchstat
// so in code editor, just changed this to UTF-8
// (also renamed to Benchmark-16 in the output to make things match)
// benchstat old.txt new.txt
// goos: windows
// goarch: amd64
// pkg: github.tesla.com/personal/going-golfing/gotest
// cpu: 11th Gen Intel(R) Core(TM) i7-11850H @ 2.50GHz
//     │   old.txt   │               new.txt                │
//     │   sec/op    │    sec/op     vs base                │
// -16   334.0n ± 5%   187.2n ± 14%  -43.98% (p=0.000 n=10)
