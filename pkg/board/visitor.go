package board

type Visitor interface {
	Visit(cell *Cell)
}
