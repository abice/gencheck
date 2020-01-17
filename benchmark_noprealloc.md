# GENCHECK --noprealloc
```
goos: darwin
goarch: amd64
pkg: github.com/abice/gencheck/internal/benchmark
BenchmarkCompareGencheck/UUID_Pass-12         	 3060320	       388 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/UUID_Fail-12         	 1633210	       675 ns/op	     208 B/op	       6 allocs/op
BenchmarkCompareGencheck/Hex_Pass-12          	 3488574	       345 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/Hex_Fail-12          	 1812626	       663 ns/op	     192 B/op	       6 allocs/op
BenchmarkCompareGencheck/ContainsAny_Pass-12  	36120229	        34.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/ContainsAny_Fail-12  	 6942110	       171 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestStrings_Pass-12  	168171955	         7.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestStrings_Fail-12  	 2564859	       461 ns/op	     464 B/op	      12 allocs/op
BenchmarkCompareGencheck/TestMap_Pass-12      	188108667	         6.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestMap_Fail-12      	 8653820	       141 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestDive_Pass-12     	77069458	        15.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareGencheck/TestDive_Fail-12     	 1583066	       727 ns/op	     592 B/op	      13 allocs/op
BenchmarkCompareGencheck/TestDive_Nil-12      	 8740198	       132 ns/op	     128 B/op	       4 allocs/op
BenchmarkCompareGencheck/TestAll_Pass-12      	  695394	      1772 ns/op	     328 B/op	       6 allocs/op
BenchmarkCompareGencheck/TestAll_Fail-12      	  286522	      4457 ns/op	    3443 B/op	      59 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	23.473s
```
