# GENCHECK 
```
BenchmarkCompareGencheck/UUID_Pass-8         	 2000000	       933 ns/op	      32 B/op	       2 allocs/op
BenchmarkCompareGencheck/UUID_Fail-8         	 1000000	      1202 ns/op	     208 B/op	       6 allocs/op
BenchmarkCompareGencheck/Hex_Pass-8          	 2000000	       664 ns/op	      32 B/op	       2 allocs/op
BenchmarkCompareGencheck/Hex_Fail-8          	 2000000	       969 ns/op	     192 B/op	       6 allocs/op
BenchmarkCompareGencheck/ContainsAny_Pass-8  	20000000	       108 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/ContainsAny_Fail-8  	 5000000	       263 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestStrings_Pass-8  	20000000	        76.9 ns/op	      64 B/op	       1 allocs/op
BenchmarkCompareGencheck/TestStrings_Fail-8  	 3000000	       452 ns/op	     416 B/op	      10 allocs/op
BenchmarkCompareGencheck/TestMap_Pass-8      	50000000	        40.5 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/TestMap_Fail-8      	10000000	       172 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestDive_Pass-8     	20000000	        85.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkCompareGencheck/TestDive_Fail-8     	 1000000	      1155 ns/op	     800 B/op	      15 allocs/op
BenchmarkCompareGencheck/TestDive_Nil-8      	10000000	       170 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestAll_Pass-8      	 1000000	      2018 ns/op	     360 B/op	       8 allocs/op
BenchmarkCompareGencheck/TestAll_Fail-8      	  300000	      4652 ns/op	    2353 B/op	      48 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	28.522s
```
