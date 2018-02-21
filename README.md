numino
=

Numino is a tetris-clone that uses numbers instead of shapes. 

[numino screenshot](./numino.png)

## Quickstart
```sh
go install github.com/kharland/numino/cmd/numino
numino
```

## Concepts

### Cells
A cell is a space on the screen.  Numinos that fall from the top of the screen occupy one cell each.

### Numino
Numino are blocks that the player must place at the bottom of the game window. A numino is either _live_ or _dead_.
A live numino's value is between -10 and 10 and a dead numino's value is outside this range.  A numino with a value of
zero is removed from the game.

## Objective
Numinoes will continuously fall from the top of the screen. Your goal is to place as many numinos as possible without letting them reach the top of the screen. 
The game ends as soon as a cell in the top row contains a dead numino.  A numino's value changes when you place another numino on top of it. For example, landing a numino with a value of 3
on a numino with a value of 5 will change that numino's value to 8.

## Controls

### Shifting
You can shift the falling numinos left or right using the _a_ and _d_ keys, respectively.

### Slamming
You can _slam_ the numinoes to the bottom of the screen using the _s_ key.  This immediately 
places the numino at the lowest cell that it can occupy, merging it with another numino if
possible.

## Scoring
Your score is equal to the number of numinos that you successfuly land on the game board. Even if a numino
dies as a result of landing, your score increases by one.

