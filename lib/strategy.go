package tictacgo2

type strategy struct {
}

func newStrategy() {
	s := &strategy{}
	go s.run()
}

func (s *strategy) run() {

}
