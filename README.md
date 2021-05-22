# Console tic-tac-toe with minimax algorithm on Golang

Coded for the competition https://ozondev.ru/ (I won a small cash gift)

## What is Minimax?

Minimax is a artificial intelligence applied in two player games, such as tic-tac-toe, checkers, chess and go. This games are known as zero-sum games, because in a mathematical representation: one player wins (+1) and other player loses (-1) or both of anyone not to win (0).

## How does it works?

The algorithm search, recursively, the best move that leads the Max player to win or not lose (draw). It consider the current state of the game and the available moves at that state, then for each valid move it plays (alternating min and max) until it finds a terminal state (win, draw or lose).

## How to play?

Start app:
```
$ go run main.go
```
