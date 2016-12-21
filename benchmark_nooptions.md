# GENCHECK 
```
BenchmarkCompareGencheck/UUID_Pass-8         	 2000000	       991 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/UUID_Fail-8         	 1000000	      1315 ns/op	     208 B/op	       6 allocs/op
BenchmarkCompareGencheck/Hex_Pass-8          	 2000000	       628 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/Hex_Fail-8          	 2000000	       997 ns/op	     192 B/op	       6 allocs/op
BenchmarkCompareGencheck/ContainsAny_Pass-8  	 3000000	       577 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/ContainsAny_Fail-8  	 2000000	       894 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestStrings_Pass-8  	20000000	        80.2 ns/op	      64 B/op	       1 allocs/op
BenchmarkCompareGencheck/TestStrings_Fail-8  	 3000000	       478 ns/op	     416 B/op	      10 allocs/op
BenchmarkCompareGencheck/TestMap_Pass-8      	30000000	        42.0 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/TestMap_Fail-8      	10000000	       177 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestDive_Pass-8     	20000000	        90.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkCompareGencheck/TestDive_Fail-8     	 1000000	      1210 ns/op	     664 B/op	      15 allocs/op
BenchmarkCompareGencheck/TestDive_Nil-8      	10000000	       171 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestAll_Pass-8      	 1000000	      2053 ns/op	     296 B/op	       6 allocs/op
BenchmarkCompareGencheck/TestAll_Fail-8      	  300000	      4505 ns/op	    2025 B/op	      44 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	29.578s
```
