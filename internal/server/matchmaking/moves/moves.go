package moves

type Move struct {
	X, Y int
}

type Open struct {
	Move
}

type Flag struct {
	Move
}

type Chord struct {
	Move
}
