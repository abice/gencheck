# PLAYGROUND
```
goos: darwin
goarch: amd64
pkg: github.com/abice/gencheck/internal/benchmark
BenchmarkComparePlayground/UUID_Pass-12         	 2187891	       594 ns/op	      16 B/op	       1 allocs/op
BenchmarkComparePlayground/UUID_Fail-12         	 1668792	       754 ns/op	     224 B/op	       5 allocs/op
BenchmarkComparePlayground/Hex_Pass-12          	 2723215	       434 ns/op	       0 B/op	       0 allocs/op
BenchmarkComparePlayground/Hex_Fail-12          	 1905054	       644 ns/op	     208 B/op	       4 allocs/op
BenchmarkComparePlayground/ContainsAny_Pass-12  	 9700926	       126 ns/op	       0 B/op	       0 allocs/op
BenchmarkComparePlayground/ContainsAny_Fail-12  	 3581966	       323 ns/op	     224 B/op	       4 allocs/op
BenchmarkComparePlayground/TestStrings_Pass-12  	 3550527	       342 ns/op	      16 B/op	       1 allocs/op
BenchmarkComparePlayground/TestStrings_Fail-12  	 1000000	      1058 ns/op	     816 B/op	      13 allocs/op
--- SKIP: BenchmarkComparePlayground/TestMap_Pass
--- SKIP: BenchmarkComparePlayground/TestMap_Fail
BenchmarkComparePlayground/TestDive_Pass-12     	 4978272	       240 ns/op	      32 B/op	       2 allocs/op
BenchmarkComparePlayground/TestDive_Fail-12     	 2600906	       479 ns/op	     272 B/op	       7 allocs/op
BenchmarkComparePlayground/TestDive_Nil-12      	 4445193	       273 ns/op	     208 B/op	       4 allocs/op
BenchmarkComparePlayground/TestAll_Pass-12      	  339356	      3064 ns/op	     408 B/op	      11 allocs/op
BenchmarkComparePlayground/TestAll_Fail-12      	  223174	      5596 ns/op	    4211 B/op	      51 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	19.918s
```
