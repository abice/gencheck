# GENCHECK 
```
goos: darwin
goarch: amd64
pkg: github.com/abice/gencheck/internal/benchmark
BenchmarkCompareGencheck/UUID_Pass-12         	 2869220	       407 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/UUID_Fail-12         	 1858568	       645 ns/op	     208 B/op	       6 allocs/op
BenchmarkCompareGencheck/Hex_Pass-12          	 3280252	       365 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/Hex_Fail-12          	 1954568	       632 ns/op	     192 B/op	       6 allocs/op
BenchmarkCompareGencheck/ContainsAny_Pass-12  	21148444	        56.3 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/ContainsAny_Fail-12  	 7321660	       167 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestStrings_Pass-12  	19691707	        65.2 ns/op	      64 B/op	       1 allocs/op
BenchmarkCompareGencheck/TestStrings_Fail-12  	 3345627	       363 ns/op	     416 B/op	      10 allocs/op
BenchmarkCompareGencheck/TestMap_Pass-12      	39226683	        31.0 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareGencheck/TestMap_Fail-12      	 9257938	       131 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestDive_Pass-12     	18120820	        67.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkCompareGencheck/TestDive_Fail-12     	 1685066	       704 ns/op	     592 B/op	      13 allocs/op
BenchmarkCompareGencheck/TestDive_Nil-12      	 9715989	       130 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestAll_Pass-12      	  644622	      1933 ns/op	     632 B/op	       8 allocs/op
BenchmarkCompareGencheck/TestAll_Fail-12      	  316861	      4088 ns/op	    2722 B/op	      54 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	22.311s
```
