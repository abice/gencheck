# GENCHECK --failfast
```
BenchmarkCompareGencheck/UUID_Pass-8         	 2000000	       846 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/UUID_Fail-8         	 1000000	      1233 ns/op	     208 B/op	       6 allocs/op
BenchmarkCompareGencheck/Hex_Pass-8          	 2000000	       611 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/Hex_Fail-8          	 1000000	      1002 ns/op	     192 B/op	       6 allocs/op
BenchmarkCompareGencheck/ContainsAny_Pass-8  	20000000	        74.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/ContainsAny_Fail-8  	 5000000	       285 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestStrings_Pass-8  	200000000	         8.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestStrings_Fail-8  	10000000	       182 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestMap_Pass-8      	200000000	         9.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestMap_Fail-8      	10000000	       186 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestDive_Pass-8     	100000000	        21.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestDive_Fail-8     	 1000000	      1194 ns/op	     800 B/op	      15 allocs/op
BenchmarkCompareGencheck/TestDive_Nil-8      	10000000	       176 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestAll_Pass-8      	 1000000	      1856 ns/op	     104 B/op	       6 allocs/op
BenchmarkCompareGencheck/TestAll_Fail-8      	10000000	       186 ns/op	     128 B/op	       4 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	28.629s
```
