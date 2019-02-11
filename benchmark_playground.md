# PLAYGROUND
```
goos: darwin
goarch: amd64
pkg: github.com/abice/gencheck/internal/benchmark
BenchmarkComparePlayground/UUID_Pass-8         	 2000000	       880 ns/op	      16 B/op	       1 allocs/op
BenchmarkComparePlayground/UUID_Fail-8         	 1000000	      1168 ns/op	     224 B/op	       5 allocs/op
BenchmarkComparePlayground/Hex_Pass-8          	 2000000	       621 ns/op	       0 B/op	       0 allocs/op
BenchmarkComparePlayground/Hex_Fail-8          	 2000000	       850 ns/op	     208 B/op	       4 allocs/op
BenchmarkComparePlayground/ContainsAny_Pass-8  	10000000	       170 ns/op	       0 B/op	       0 allocs/op
BenchmarkComparePlayground/ContainsAny_Fail-8  	 3000000	       412 ns/op	     224 B/op	       4 allocs/op
BenchmarkComparePlayground/TestStrings_Pass-8  	 3000000	       425 ns/op	      16 B/op	       1 allocs/op
BenchmarkComparePlayground/TestStrings_Fail-8  	 1000000	      1240 ns/op	     816 B/op	      13 allocs/op
--- SKIP: BenchmarkComparePlayground/TestMap_Pass
--- SKIP: BenchmarkComparePlayground/TestMap_Fail
BenchmarkComparePlayground/TestDive_Pass-8     	 5000000	       327 ns/op	      32 B/op	       2 allocs/op
BenchmarkComparePlayground/TestDive_Fail-8     	 3000000	       591 ns/op	     272 B/op	       7 allocs/op
BenchmarkComparePlayground/TestDive_Nil-8      	 5000000	       338 ns/op	     208 B/op	       4 allocs/op
BenchmarkComparePlayground/TestAll_Pass-8      	  500000	      4019 ns/op	     408 B/op	      11 allocs/op
BenchmarkComparePlayground/TestAll_Fail-8      	  200000	      6275 ns/op	    4209 B/op	      51 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	24.609s
```
