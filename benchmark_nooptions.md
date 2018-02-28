# GENCHECK 
```
goos: darwin
goarch: amd64
pkg: github.com/abice/gencheck/internal/benchmark
BenchmarkCompareGencheck/UUID_Pass-8         	 2000000	       826 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/UUID_Fail-8         	 1000000	      1096 ns/op	     208 B/op	       6 allocs/op
BenchmarkCompareGencheck/Hex_Pass-8          	 3000000	       579 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/Hex_Fail-8          	 2000000	       881 ns/op	     192 B/op	       6 allocs/op
BenchmarkCompareGencheck/ContainsAny_Pass-8  	20000000	        74.4 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/ContainsAny_Fail-8  	10000000	       199 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestStrings_Pass-8  	20000000	        67.0 ns/op	      64 B/op	       1 allocs/op
BenchmarkCompareGencheck/TestStrings_Fail-8  	 3000000	       408 ns/op	     416 B/op	      10 allocs/op
BenchmarkCompareGencheck/TestMap_Pass-8      	50000000	        37.8 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/TestMap_Fail-8      	10000000	       151 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestDive_Pass-8     	20000000	        80.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkCompareGencheck/TestDive_Fail-8     	 1000000	      1016 ns/op	     800 B/op	      15 allocs/op
BenchmarkCompareGencheck/TestDive_Nil-8      	10000000	       142 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestAll_Pass-8      	  500000	      2424 ns/op	     632 B/op	       8 allocs/op
BenchmarkCompareGencheck/TestAll_Fail-8      	  300000	      4669 ns/op	    2929 B/op	      56 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	25.966s
```
