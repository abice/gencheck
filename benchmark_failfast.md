# GENCHECK --failfast
```
BenchmarkCompareGencheck/UUID_Pass-8         	 2000000	       924 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/UUID_Fail-8         	 1000000	      1434 ns/op	     208 B/op	       6 allocs/op
BenchmarkCompareGencheck/Hex_Pass-8          	 3000000	       592 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/Hex_Fail-8          	 1000000	      1035 ns/op	     192 B/op	       6 allocs/op
BenchmarkCompareGencheck/ContainsAny_Pass-8  	 3000000	       554 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/ContainsAny_Fail-8  	 2000000	       904 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestStrings_Pass-8  	200000000	         9.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestStrings_Fail-8  	10000000	       179 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestMap_Pass-8      	200000000	         9.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestMap_Fail-8      	10000000	       192 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestDive_Pass-8     	50000000	        24.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestDive_Fail-8     	 1000000	      1198 ns/op	     664 B/op	      15 allocs/op
BenchmarkCompareGencheck/TestDive_Nil-8      	10000000	       179 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestAll_Pass-8      	 1000000	      1948 ns/op	      72 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestAll_Fail-8      	10000000	       191 ns/op	     128 B/op	       4 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	30.864s
```
