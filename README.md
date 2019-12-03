# code
A code library written in golang.


### Benchmark
go test -v -run="none" -bench=. -benchtime=1s
```
goos: darwin
goarch: amd64
pkg: github.com/hslam/code
BenchmarkInterfaceTypeString0-4    	334841894	         3.51 ns/op
BenchmarkInterfaceTypeString-4     	294729032	         4.53 ns/op
BenchmarkStringType-4              	1000000000	         0.311 ns/op
BenchmarkInterfaceTypeInt0-4       	350415357	         3.42 ns/op
BenchmarkInterfaceTypeInt-4        	272255924	         4.70 ns/op
BenchmarkIntType-4                 	1000000000	         0.630 ns/op
BenchmarkInterfaceTypeTrue-4       	388718568	         3.13 ns/op
BenchmarkTrueType-4                	1000000000	         0.651 ns/op
BenchmarkInterfaceTypeFalse-4      	384036106	         3.09 ns/op
BenchmarkFalseType-4               	1000000000	         0.618 ns/op
BenchmarkSizeInt-4                 	1000000000	         0.311 ns/op
BenchmarkCodeInt-4                 	40835382	        27.0 ns/op
BenchmarkCodeIntWithBuffer-4       	100000000	        13.5 ns/op
BenchmarkSizeVarint-4              	1000000000	         0.313 ns/op
BenchmarkCodeVarint-4              	47646476	        25.5 ns/op
BenchmarkCodeVarintWithBuffer-4    	100000000	        13.2 ns/op
BenchmarkCodeFloat32-4             	43381042	        27.0 ns/op
BenchmarkCodeFloat32WithBuffer-4   	80151388	        14.7 ns/op
BenchmarkCodeFloat64-4             	30978212	        45.2 ns/op
BenchmarkCodeFloat64WithBuffer-4   	51835364	        23.7 ns/op
PASS
ok  	github.com/hslam/code	23.934s
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Mort Huang)

### Authors
code was written by Mort Huang.
