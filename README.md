# wordchain-breadth-first-search

Wordchain is an implementation of a breadth-first-search algorithm that given 2 words, start and end word, will
find the shortest chain (if possible to find) that connects this two words. In each step the next word can differ 
from previous with one letter. See more detailed describtion of problem here : [Word Chains](http://codekata.com/kata/kata19-word-chains/)

Examples:
Start: 'ruby' 
End:   'code' 
Result: [ruby rube robe rode code] 

# run the program and tests
Please make sure you have go installed 
Clone the repo and go to wordchain-breadth-first-search/src directory

```bash
$ git clone https://github.com/gohargasparyan/wordchain-breadth-first-search
$ cd wordchain-breadth-first-search/src
```

I have added 2 cases in main.go, so run the main go to see the results run
```bash
$ go run .
```

to run the tests
```bash
$  go test -test.v
```