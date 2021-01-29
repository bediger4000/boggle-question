# Daily Coding Problem: Problem #619 [Easy]  


This problem was asked by Coursera.

Given a 2D board of characters and a word, find if the word exists in the
grid.

The word can be constructed from letters of sequentially adjacent cell,
where "adjacent" cells are those horizontally or vertically neighboring. The
same letter cell may not be used more than once.

For example, given the following board:

    [
      ['A','B','C','E'],
      ['S','F','C','S'],
      ['A','D','E','E']
    ]

exists(board, "ABCCED") returns true,
exists(board, "SEE") returns true,
exists(board, "ABCB") returns false.

---

### Related

Daily Coding Problem: Problem #772 [Easy] seems related.

This problem was asked by Facebook.

Boggle is a game played on a 4 x 4 grid of letters.
The goal is to find as many words as possible that can be formed by
a sequence of adjacent letters in the grid,
using each cell at most once.
Given a game board and a dictionary of valid words,
implement a Boggle solver.

#### Related analysis

The solution to the Facebook problem would use the Coursera
boggle solver to attempt to find each word from the dictionary of valid words.
Iterating over the dictionary of valid words is the obvious possibility,
but there are also obvious optimizations,
like finding all letters in the 4 x 4 grid of letters,
then only iterating over words beginning with those 16 (or fewer) letters.

Oddly, this extra work and optimizations seems to still keep
the Facebook version in the "Easy" rating.
---

### Related #2

Daily Coding Problem: Problem #784 [Easy]

This problem was asked by Microsoft.

Given a 2D matrix of characters and a target word,
write a function that returns whether the word can be found in the matrix by
going left-to-right, or up-to-down.

For example, given the following matrix:

```
[['F', 'A', 'C', 'I'],
 ['O', 'B', 'Q', 'P'],
 ['A', 'N', 'O', 'B'],
 ['M', 'A', 'S', 'S']]
```

and the target word 'FOAM',
you should return true,
since it's the leftmost column.
Similarly, given the target word 'MASS',
you should return true,
since it's the last row.


---

## Build and run

    $ go build solution1.go
    $ ./solution1 -f example.matrix -w ABCCED
    reading matrix from "example.matrix"
    A B C E 
    S F C S 
    A D E E 
    
    Looking for "ABCCED" in matrix
    Found "ABCCED" in matrix

`solution1` and `solution2` are identical in building and running.

## Analysis

This is a slightly generalized "boggle" game single word finder.

This is also what the coding interview books will call a "backtracking" problem,
which means it's recursive, but the program probably has to make more than
a single recurse at any given depth into the call stack:
if a recursive search comes up empty,
the program must choose another search to make based on the partial
solution is has so far.

There's a lot of subtasks in this problem,
even if you don't read in a board from a file.

* Board representation, 2-D matrix or something else
* Looking for the first letter of the word in the open board
* Looking for a letter in a board with some locations previously used
* Finding legal locations to check for a letter
* Account for previously used locations in board
* Decide when the program has found the desired word

I chose to account for "used" locations on the board
with a Go map.
I encoded the `x,y` locations of letters
as `x*len(board[x]) + y`, which is a base len(board[x])
integer.
The value of the used map is true if the location on the board so encoded 
has already been used to compose the word so far.

I made an [alternative](solution2.go) that uses the board itself (a Go `[][]rune` slice-of-slices)
to encode which locations had already been used to compose the word
by setting used locations to some non-letter value that would never
match.
It's 6 lines shorter, but maybe easier to understand.

i don't believe the "easy" rating.
It may be relatively straightforward, requiring only basic algorithm knowledge,
but it's not easy given the number of sub-tasks,
and the oddity of finding moves on the board that don't re-use a location.
There's also the matter of choosing data structures.
This seems like one problem where Rob Pike's advice about figuring out the
correct data structure doesn't necessarily come true.
You can use different data structures with almost identical algorithms.

If the interviewer is interested in whether the candidate can rough in
an algorithm, then fill in the fine details, this might be a good question.
Approach to various sub-tasks can reveal a candidate's organizing skills,
and maybe a little bit of data structure knowledge.
There's really one one algorithm, but it is a backtracking algorithm,
so maybe that's enough for the interviewers who put stock in "big-O"
type knowledge.

It's not a question for an entry-level candidate, and most junior-level
programmers are going to get bogged down in the details.
Even senior-level candidates are going to screw up finding "next letter"
locations that fit on the board, and don't re-use a location.

The interview using this should allot a great deal of time.
The candidate should block out the solution before starting.
