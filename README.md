# Wordchain Breadth First Search

Wordchain is an implementation of a breadth-first-search algorithm that given 2 words, start and end word, will
find the shortest chain (if possible to find) that connects this two words. In each step the next word can differ 
from previous with one letter. See more detailed describtion of problem here : [Word Chains](http://codekata.com/kata/kata19-word-chains/)

Examples:
Start: 'ruby' 
End:   'code' 
Result: [ruby rube robe rode code] 

# Run the program and tests
Please make sure you have go installed 
Clone the repo and go to wordchain-breadth-first-search/src directory

```bash
$ git clone git@github.com:gohargasparyan/wordchain-breadth-first-search.git
$ cd wordchain-breadth-first-search/src
```

I have added 2 cases in main.go, so run to see the results
```bash
$ go run .
```

to run the tests from src dir
```bash
$  go test -test.v  
```
or from the project root
```bash
$  go test ./... -test.v
```
# Benchmarking
In order to run benchmarks from src dir
```bash
$  go test -bench=.
```
or from the project root
```bash
$  go test ./... -bench=.
```

You will see that the word chain runs under a second (aprx. 0.17 second). 
You will also notice that running the word chain forwards and backwards (e.g. “lead” into “gold” and “gold” into “lead”) 
takes take slightly different time (aprx. 0.05 seconds difference)

# Things could be further improved

- Add convenient CLI 