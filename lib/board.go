package tictacgo2

type board struct {
}

func newBoard() {
	b := &board{}
	go b.run()
}

func (b *board) run() {

}
