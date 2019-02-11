# GENCHECK --noprealloc
```
goos: darwin
goarch: amd64
pkg: github.com/abice/gencheck/internal/benchmark
BenchmarkCompareGencheck/UUID_Pass-8         	 2000000	       855 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/UUID_Fail-8         	 1000000	      1161 ns/op	     208 B/op	       6 allocs/op
BenchmarkCompareGencheck/Hex_Pass-8          	 3000000	       507 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/Hex_Fail-8          	 2000000	       861 ns/op	     192 B/op	       6 allocs/op
BenchmarkCompareGencheck/ContainsAny_Pass-8  	30000000	        43.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/ContainsAny_Fail-8  	10000000	       236 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestStrings_Pass-8  	200000000	         7.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestStrings_Fail-8  	 2000000	       594 ns/op	     464 B/op	      12 allocs/op
BenchmarkCompareGencheck/TestMap_Pass-8      	200000000	         8.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestMap_Fail-8      	10000000	       175 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestDive_Pass-8     	50000000	        24.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestDive_Fail-8     	 1000000	      1159 ns/op	     800 B/op	      15 allocs/op
BenchmarkCompareGencheck/TestDive_Nil-8      	10000000	       167 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestAll_Pass-8      	 1000000	      2300 ns/op	     328 B/op	       6 allocs/op
BenchmarkCompareGencheck/TestAll_Fail-8      	  300000	      5777 ns/op	    3649 B/op	      61 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	29.262s
```
