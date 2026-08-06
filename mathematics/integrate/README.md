# Benchmarks
```
goos: darwin
goarch: arm64
pkg: github.com/ParkhomenkoDV/mathematics/mathematics/mathematics/integrate
cpu: Apple M4
BenchmarkNQuad1D-10                     36066766                32.43 ns/op            8 B/op          1 allocs/op
BenchmarkNQuad2D-10                      8055234               148.5 ns/op            16 B/op          1 allocs/op
BenchmarkNQuad3D-10                      1675933               716.4 ns/op            24 B/op          1 allocs/op
BenchmarkAdaptiveSimpsonPoly-10         92516151                12.99 ns/op            0 B/op          0 allocs/op
BenchmarkAdaptiveSimpsonSin-10           2844236               422.3 ns/op             0 B/op          0 allocs/op
```