# PLAYGROUND
```
BenchmarkComparePlayground/UUID_Pass-8         	 1000000	      1244 ns/op	      16 B/op	       1 allocs/op
BenchmarkComparePlayground/UUID_Fail-8         	 1000000	      1576 ns/op	     224 B/op	       5 allocs/op
BenchmarkComparePlayground/Hex_Pass-8          	 2000000	       756 ns/op	       0 B/op	       0 allocs/op
BenchmarkComparePlayground/Hex_Fail-8          	 1000000	      1057 ns/op	     208 B/op	       4 allocs/op
BenchmarkComparePlayground/ContainsAny_Pass-8  	 2000000	       716 ns/op	       0 B/op	       0 allocs/op
BenchmarkComparePlayground/ContainsAny_Fail-8  	 1000000	      1174 ns/op	     224 B/op	       4 allocs/op
BenchmarkComparePlayground/TestStrings_Pass-8  	 3000000	       460 ns/op	      16 B/op	       1 allocs/op
BenchmarkComparePlayground/TestStrings_Fail-8  	 1000000	      1359 ns/op	     816 B/op	      13 allocs/op
--- SKIP: BenchmarkComparePlayground/TestMap_Pass
--- SKIP: BenchmarkComparePlayground/TestMap_Fail
BenchmarkComparePlayground/TestDive_Pass-8     	 5000000	       387 ns/op	      32 B/op	       2 allocs/op
BenchmarkComparePlayground/TestDive_Fail-8     	 2000000	       675 ns/op	     272 B/op	       7 allocs/op
BenchmarkComparePlayground/TestDive_Nil-8      	 5000000	       374 ns/op	     208 B/op	       4 allocs/op
BenchmarkComparePlayground/TestAll_Pass-8      	  500000	      3362 ns/op	     152 B/op	       9 allocs/op
BenchmarkComparePlayground/TestAll_Fail-8      	  300000	      4953 ns/op	    2737 B/op	      38 allocs/op
PASS
ok  	github.com/abice/gencheck/internal/benchmark	22.662s
```
