# PLAYGROUND
```
BenchmarkComparePlayground/UUID_Pass-8         	 1000000	      1088 ns/op	      32 B/op	       2 allocs/op
BenchmarkComparePlayground/UUID_Fail-8         	 1000000	      1326 ns/op	     224 B/op	       5 allocs/op
BenchmarkComparePlayground/Hex_Pass-8          	 2000000	       757 ns/op	      16 B/op	       1 allocs/op
BenchmarkComparePlayground/Hex_Fail-8          	 2000000	       977 ns/op	     208 B/op	       4 allocs/op
BenchmarkComparePlayground/ContainsAny_Pass-8  	10000000	       222 ns/op	       0 B/op	       0 allocs/op
BenchmarkComparePlayground/ContainsAny_Fail-8  	 3000000	       483 ns/op	     224 B/op	       4 allocs/op
BenchmarkComparePlayground/TestStrings_Pass-8  	 3000000	       446 ns/op	      16 B/op	       1 allocs/op
BenchmarkComparePlayground/TestStrings_Fail-8  	 1000000	      1277 ns/op	     816 B/op	      13 allocs/op
--- SKIP: BenchmarkComparePlayground/TestMap_Pass
--- SKIP: BenchmarkComparePlayground/TestMap_Fail
BenchmarkComparePlayground/TestDive_Pass-8     	 5000000	       366 ns/op	      32 B/op	       2 allocs/op
BenchmarkComparePlayground/TestDive_Fail-8     	 2000000	       631 ns/op	     272 B/op	       7 allocs/op
BenchmarkComparePlayground/TestDive_Nil-8      	 5000000	       361 ns/op	     208 B/op	       4 allocs/op
BenchmarkComparePlayground/TestAll_Pass-8      	  500000	      3416 ns/op	     184 B/op	      11 allocs/op
BenchmarkComparePlayground/TestAll_Fail-8      	  300000	      4970 ns/op	    3057 B/op	      42 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	24.713s
```
