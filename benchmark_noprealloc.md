# GENCHECK --noprealloc
```
BenchmarkCompareGencheck/UUID_Pass-8         	 1000000	      1006 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/UUID_Fail-8         	 1000000	      1383 ns/op	     208 B/op	       6 allocs/op
BenchmarkCompareGencheck/Hex_Pass-8          	 2000000	       585 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/Hex_Fail-8          	 1000000	      1034 ns/op	     192 B/op	       6 allocs/op
BenchmarkCompareGencheck/ContainsAny_Pass-8  	 3000000	       557 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/ContainsAny_Fail-8  	 2000000	       915 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestStrings_Pass-8  	200000000	        10.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestStrings_Fail-8  	 3000000	       594 ns/op	     464 B/op	      12 allocs/op
BenchmarkCompareGencheck/TestMap_Pass-8      	200000000	         9.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestMap_Fail-8      	10000000	       185 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestDive_Pass-8     	50000000	        26.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestDive_Fail-8     	 1000000	      1213 ns/op	     664 B/op	      15 allocs/op
BenchmarkCompareGencheck/TestDive_Nil-8      	10000000	       177 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestAll_Pass-8      	 1000000	      1956 ns/op	      72 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestAll_Fail-8      	  300000	      4963 ns/op	    2313 B/op	      48 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	28.574s
```
