# PLAYGROUND
```
goos: darwin
goarch: amd64
pkg: github.com/abice/gencheck/internal/benchmark
BenchmarkComparePlayground/UUID_Pass-8         	 1000000	      1081 ns/op	      16 B/op	       1 allocs/op
BenchmarkComparePlayground/UUID_Fail-8         	 1000000	      1307 ns/op	     224 B/op	       5 allocs/op
BenchmarkComparePlayground/Hex_Pass-8          	 2000000	       712 ns/op	       0 B/op	       0 allocs/op
BenchmarkComparePlayground/Hex_Fail-8          	 2000000	       971 ns/op	     208 B/op	       4 allocs/op
BenchmarkComparePlayground/ContainsAny_Pass-8  	10000000	       206 ns/op	       0 B/op	       0 allocs/op
BenchmarkComparePlayground/ContainsAny_Fail-8  	 3000000	       426 ns/op	     224 B/op	       4 allocs/op
BenchmarkComparePlayground/TestStrings_Pass-8  	 3000000	       443 ns/op	      16 B/op	       1 allocs/op
BenchmarkComparePlayground/TestStrings_Fail-8  	 1000000	      1209 ns/op	     816 B/op	      13 allocs/op
--- SKIP: BenchmarkComparePlayground/TestMap_Pass
--- SKIP: BenchmarkComparePlayground/TestMap_Fail
BenchmarkComparePlayground/TestDive_Pass-8     	 5000000	       324 ns/op	      32 B/op	       2 allocs/op
BenchmarkComparePlayground/TestDive_Fail-8     	 3000000	       568 ns/op	     272 B/op	       7 allocs/op
BenchmarkComparePlayground/TestDive_Nil-8      	 5000000	       321 ns/op	     208 B/op	       4 allocs/op
BenchmarkComparePlayground/TestAll_Pass-8      	  500000	      3886 ns/op	     408 B/op	      11 allocs/op
BenchmarkComparePlayground/TestAll_Fail-8      	  300000	      5855 ns/op	    4209 B/op	      51 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	24.440s
```
