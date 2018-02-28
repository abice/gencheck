# GENCHECK --failfast
```
goos: darwin
goarch: amd64
pkg: github.com/abice/gencheck/internal/benchmark
BenchmarkCompareGencheck/UUID_Pass-8         	 2000000	       845 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/UUID_Fail-8         	 1000000	      1153 ns/op	     208 B/op	       6 allocs/op
BenchmarkCompareGencheck/Hex_Pass-8          	 2000000	       594 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/Hex_Fail-8          	 2000000	       945 ns/op	     192 B/op	       6 allocs/op
BenchmarkCompareGencheck/ContainsAny_Pass-8  	30000000	        47.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/ContainsAny_Fail-8  	10000000	       224 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestStrings_Pass-8  	200000000	         9.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestStrings_Fail-8  	10000000	       173 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestMap_Pass-8      	200000000	         8.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestMap_Fail-8      	10000000	       167 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestDive_Pass-8     	100000000	        23.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestDive_Fail-8     	 1000000	      1081 ns/op	     800 B/op	      15 allocs/op
BenchmarkCompareGencheck/TestDive_Nil-8      	10000000	       166 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestAll_Pass-8      	 1000000	      2371 ns/op	     328 B/op	       6 allocs/op
BenchmarkCompareGencheck/TestAll_Fail-8      	10000000	       175 ns/op	     128 B/op	       4 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	31.191s
```
