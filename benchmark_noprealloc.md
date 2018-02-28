# GENCHECK --noprealloc
```
goos: darwin
goarch: amd64
pkg: github.com/abice/gencheck/internal/benchmark
BenchmarkCompareGencheck/UUID_Pass-8         	 2000000	       791 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/UUID_Fail-8         	 1000000	      1110 ns/op	     208 B/op	       6 allocs/op
BenchmarkCompareGencheck/Hex_Pass-8          	 3000000	       537 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/Hex_Fail-8          	 2000000	       991 ns/op	     192 B/op	       6 allocs/op
BenchmarkCompareGencheck/ContainsAny_Pass-8  	30000000	        46.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/ContainsAny_Fail-8  	10000000	       218 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestStrings_Pass-8  	200000000	         8.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestStrings_Fail-8  	 3000000	       529 ns/op	     464 B/op	      12 allocs/op
BenchmarkCompareGencheck/TestMap_Pass-8      	200000000	         8.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestMap_Fail-8      	10000000	       171 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestDive_Pass-8     	100000000	        23.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestDive_Fail-8     	 1000000	      1067 ns/op	     800 B/op	      15 allocs/op
BenchmarkCompareGencheck/TestDive_Nil-8      	10000000	       155 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestAll_Pass-8      	  500000	      2300 ns/op	     328 B/op	       6 allocs/op
BenchmarkCompareGencheck/TestAll_Fail-8      	  300000	      5588 ns/op	    3649 B/op	      61 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	29.810s
```
