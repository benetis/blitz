# blitz

## Work in progress

## Constraints

- Collection fits into memory

### Benchmarks

#### Map for slices
```
BenchmarkParallelFilterNonIntensive-11      	  224220	      5237 ns/op
BenchmarkSequentialFilterNonIntensive-11    	  942145	      1273 ns/op

BenchmarkParallelFilterIntensive-11         	   12465	     96407 ns/op
BenchmarkSequentialFilterIntensive-11       	    4354	    266093 ns/op

BenchmarkParallelMapNonIntensive-11         	  305452	      3875 ns/op
BenchmarkSequentialMapNonIntensive-11       	  725184	      1613 ns/op

BenchmarkParallelMapIntensive-11            	   12486	     95744 ns/op
BenchmarkSequentialMapIntensive-11          	    4564	    261641 ns/op
```
