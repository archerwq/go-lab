/*
https://golang.org/pkg/testing/

Asking go test to run your benchmarks does not disable the tests in the package.
If you want to skip the tests, you can do so by passing a regex to the -run flag
that will not match anything.
`go test -run=XXX -bench=.`
*/
package basic

import "testing"

// go test -v github.com/archerwq/go-lab/concept/test
func TestAverage(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"NoData", args{[]float64{}}, 0},
		{"NormalCase", args{[]float64{1, 2}}, 1.5},
		{"NegativeCase", args{[]float64{-1, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.data); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

var result int

func benchmarkFib(i int, b *testing.B) {
	/* Benchmark functions are run several times by the testing package.
	The value of b.N will increase each time until the benchmark runner
	is satisfied with the stability of the benchmark.

	Each benchmark is run for a minimum of 1 second by default.
	If the second has not elapsed when the Benchmark function returns,
	the value of b.N is increased in the sequence 1, 2, 5, 10, 20, 50, …
	and the function run again.

	The final BenchmarkFib40 only ran two times with the average was just
	under a second for each run. As the testing package uses a simple
	average (total time to run the benchmark function over b.N)
	this result is statistically weak. You can increase the minimum benchmark
	time using the -benchtime flag to produce a more accurate result.

	https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
	*/
	var r int
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = Fib(i)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

// go test -run=XXX -bench=. github.com/archerwq/go-lab/concept/test
func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }
