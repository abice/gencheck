# GENCHECK --failfast
```
goos: darwin
goarch: amd64
pkg: github.com/abice/gencheck/internal/benchmark
BenchmarkCompareGencheck/UUID_Pass-12         	 3019042	       413 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/UUID_Fail-12         	 1584471	       756 ns/op	     208 B/op	       6 allocs/op
BenchmarkCompareGencheck/Hex_Pass-12          	 3540170	       350 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/Hex_Fail-12          	 1829170	       651 ns/op	     192 B/op	       6 allocs/op
BenchmarkCompareGencheck/ContainsAny_Pass-12  	31153759	        33.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/ContainsAny_Fail-12  	 6937201	       172 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestStrings_Pass-12  	169096743	         7.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestStrings_Fail-12  	 8132941	       143 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestMap_Pass-12      	197221968	         6.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestMap_Fail-12      	 8839633	       140 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestDive_Pass-12     	80319295	        15.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestDive_Fail-12     	 1629925	       751 ns/op	     592 B/op	      13 allocs/op
BenchmarkCompareGencheck/TestDive_Nil-12      	 9048465	       137 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestAll_Pass-12      	  714092	      1758 ns/op	     328 B/op	       6 allocs/op
BenchmarkCompareGencheck/TestAll_Fail-12      	 8258488	       154 ns/op	     128 B/op	       4 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	23.385s
```
