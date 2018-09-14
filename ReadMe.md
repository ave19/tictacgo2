# Design

## Actor Model
I am going to try the Actor Model.  I think go lends itself to it.

After much reading and youtubing, I've decided that things like Software
Transaction Models aren't really doing what they think they're doing, and if
go routines aren't working for you in an actor model, you're not using enough
of them.

So, we're not going to do big data structures.  No hash maps.  We're going to
give every strategy the ability to spin up a new handler for each new board it
sees in a go routine and route messages to that board.  Concurrent reads on
different boards are fine, and each actor will adjust strategy for their
board independently.  No Mutex locks or Syncs if we're lucky.

## Structured Go Routines
They live in structs!
 * https://www.youtube.com/watch?v=2HOO5gIgyMg
 * Evan Huus

 ```
 type foo struct {

 }

 func newFoo() {
 	f := &foo{}
 	go f.run()
 }

 func (f *foo) run() {

 }
 ```

# Design

## Game
A game is going to be a structure that we can pass around between players.

## Board
Or maybe we can pass a board.  It depends on whether the game can manage it.

## Player
There's only two.  They have names and use strategies to determine a move when
faced with a board.

### Move, a Choice given a Board
A Move is the Choice you make when a Player gets a Board (or Game?).  You have
to remember your Moves if you want to know about winning and losing. A Move is
an operation that changes a Board's state.
