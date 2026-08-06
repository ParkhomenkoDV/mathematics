# Benchmarks
```
goos: darwin
goarch: arm64
pkg: github.com/ParkhomenkoDV/mathematics/mathematics/mathematics/integrate
cpu: Apple M4
BenchmarkNQuad1D-10                     34089015                35.48 ns/op            8 B/op          1 allocs/op
BenchmarkNQuad2D-10                      8412632               142.2 ns/op            16 B/op          1 allocs/op
BenchmarkNQuad3D-10                      1744362               687.6 ns/op            24 B/op          1 allocs/op
BenchmarkAdaptiveSimpsonPoly-10         100000000               10.43 ns/op            0 B/op          0 allocs/op
BenchmarkAdaptiveSimpsonSin-10           3604192               307.3 ns/op             0 B/op          0 allocs/op
```