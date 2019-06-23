# Permfect 
A Go tool for creating permutations of alphanumerics.

This tool was written to compare speeds with Dominic's Python [fast-permute](https://github.com/singe/fast-permute) code. tl;dr, the Python implementation is much faster.

## What Permfect does

If you have an input file containing:
```
1
2
3
```

This will produce
```
1
2
3
12
13
21
23
31
32
123
132
213
231
312
321
```

## How to run it

`go run permfect.go -infile input.txt -outbase speedTest`   

This will create multiple files with the length of the strings in the file prepended e.g. if you give an output base name of "foo", then "7foo" will contain all permutations of length 7.

## Generating specific lengths
You can limit the lengths of output strings generated by using the `-min` and `-max` flags. e.g:

`go run permfect.go -infile input.txt -outbase speedTest -min 2 -max 4`

## Debug output
Use the `-debug` flag to include additioanl debugging.

```
$ go run permfect.go -infile input.txt -outbase speedTest -min 2 -max 6 -debug
[+] Calculating permutations from input file 'input.txt' and writing to 2 files (output base of 'speedTest')...
 \---> Finished calculating permutations of length 2
 \---> Finished writing to file 2speedTest
 \---> Finished calculating permutations of length 3
 \---> Finished writing to file 3speedTest
[+] Finished. Took 1.186827ms to calculate 12 permutations.
``` 


## Performance
Below is a comparison between `permfect.go` and [fast-permute](https://github.com/singe/fast-permute). The input consists of the set `{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, a, b}`.        

```
$ go run permfect.go -infile input.txt -outbase speedTest
[+] Calculating permutations from input file 'input.txt' and writing to 12 files (output base of 'speedTest')...
[+] Finished. Took 17m55.72132326s to calculate 1285148319 permutations.
```
*Result for permfect.go: 17m55s minutes.*


```
$ time python permute.py input.txt speedTest
python permute.py input.txt speedTest  769.67s user 14.38s system 252% cpu 5:10.79 total
```
*Result for permute.py: 12m48s minutes.*

*Python version 1.4x faster.*

## Analysis 
The above result was surprising, as it was hoped that Go's concurrency would give a faster result. The problem may however be with the non standard Go permutation library being used, which is an implementation of Python's itertools (fast-permute uses this). Running some comparison tests purely calculating permutations (with no disk I/O etc) results in the following:

```
//Go code to calculte permutations, using github.com/ernestosuarez/itertools
for _ = range itertools.PermutationsStr(input, len(data)) {
		counter++
}
```

```
#Python code to calculate permutations, using standard librray itertools
for p in permutations(data):
    counter = counter + 1
```


```
$ go run comparePerms.go
[+] Running test for github.com/ernestosuarez/itertools...
 \--> Finished. Took 8m32.269753709s to calculate 479001600 permutations.

[+] Running test for Python's itertools...
 \--> Finished. Took 1m6.955504805s to calculate 479001600 permutations.
 ```

 The standard Python [itertools](https://docs.python.org/2/library/itertools.html#itertools.permutations) library is 8x faster. This _seems_ to suggest that the [Go port](github.com/ernestosuarez/itertools) of Python's itertools is not optimized.

 ## Alternative Algorithms

 There are many ways to calculate permutations, and various algoritms exist such as [quickperm](http://www.quickperm.org/), [Knuth's shuffle](https://en.wikipedia.org/wiki/Random_permutation#Knuth_shuffles), and [Heap's algorithm](https://en.wikipedia.org/wiki/Heap%27s_algorithm). 
 
 A quick performance test of an alternative permuation algorithm in Go reveals that the itertools one may indeed be slow, but still lagging behind the very quick Python version:
 ```
 go run comparePerms.go
[+] Running test for github.com/ernestosuarez/itertools...
 \--> Finished. Took 8m32.269753709s to calculate 479001600 permutations.

[+] Running test for github.com/Ramshackle-Jamathon/go-quickPerm...
 \--> Finished. Took 3m54.22257733s to calculate 479001600 permutations.

[+] Running test for Python's itertools...
 \--> Finished. Took 1m6.955504805s to calculate 479001600 permutations.
 ```


 However by definition a permutation is the length of the input string. e.g the permutations of `{1, 2, 3}` are:
 `{123, 132, 213, 231, 312, 321}`. In order to calculate varying permutation lengths (as in the Python implementation) we need to enumerate all subsets of size k (for k = 2, ...,n, where n is the size of the array). i.e calculate all the combinations subsets of the input and for each enumerated subset enumerate the permutation for it.

There are similarly many different algorithms for calculating combinatons. Non-recursive algorithms seem to perform better with Go.

# Todo
 * Test various combination + permutation algorithms to try beat the Python version's speed.


# Links etc
https://stackoverflow.com/questions/41174875/algorithm-to-generate-all-permutations-of-different-sizes-of-a-group-of-integers  
https://stackoverflow.com/questions/5395684/different-length-permutations  
https://stackoverflow.com/questions/127704/algorithm-to-return-all-combinations-of-k-elements-from-n  
https://stackoverflow.com/questions/2779094/what-algorithm-can-calculate-the-power-set-of-a-given-set  
https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go  
https://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order  
https://docs.python.org/2/library/itertools.html#itertools.permutations  
https://github.com/ernestosuarez/itertools/blob/master/permutations.go#L98  


