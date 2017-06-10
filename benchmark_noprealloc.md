# GENCHECK --noprealloc
```
BenchmarkCompareGencheck/UUID_Pass-8         	 2000000	       864 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/UUID_Fail-8         	 1000000	      1248 ns/op	     208 B/op	       6 allocs/op
BenchmarkCompareGencheck/Hex_Pass-8          	 2000000	       618 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/Hex_Fail-8          	 2000000	       982 ns/op	     192 B/op	       6 allocs/op
BenchmarkCompareGencheck/ContainsAny_Pass-8  	20000000	        73.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/ContainsAny_Fail-8  	 5000000	       287 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestStrings_Pass-8  	100000000	        10.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestStrings_Fail-8  	 2000000	       603 ns/op	     464 B/op	      12 allocs/op
BenchmarkCompareGencheck/TestMap_Pass-8      	200000000	         9.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestMap_Fail-8      	10000000	       188 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestDive_Pass-8     	100000000	        22.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestDive_Fail-8     	 1000000	      1229 ns/op	     800 B/op	      15 allocs/op
BenchmarkCompareGencheck/TestDive_Nil-8      	10000000	       182 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestAll_Pass-8      	 1000000	      1903 ns/op	     104 B/op	       6 allocs/op
BenchmarkCompareGencheck/TestAll_Fail-8      	  300000	      4811 ns/op	    2609 B/op	      52 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	28.599s
```
