# blitz

## Work in progress

## Constraints

- Collection fits into memory

### Benchmarks

#### Map for slices
```
BenchmarkParallelFilterNonIntensive-11      	   99740	     13252 ns/op
BenchmarkSequentialFilterNonIntensive-11    	  841010	      1324 ns/op

BenchmarkParallelFilterIntensive-11         	   10000	    106285 ns/op
BenchmarkSequentialFilterIntensive-11       	    4563	    263706 ns/op

BenchmarkParallelMapNonIntensive-11         	  299680	      3991 ns/op
BenchmarkSequentialMapNonIntensive-11       	  704674	      1634 ns/op

BenchmarkParallelMapIntensive-11            	   12525	     95172 ns/op
BenchmarkSequentialMapIntensive-11          	    4616	    260535 ns/op
```
