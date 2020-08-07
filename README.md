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

This is also what the coding interview books will call a "bactracking" problem.
Which means it's recursive, but the program probably has to make more than
a single recurse at any given depth into the call stack.

There's a lot of subtasks in this problem,
even if you don't read in a board from a file.

* Look for the first letter of the word in the open board
* Look for a letter in a board with some locations previously used
* Find legal locations to check for a letter
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
