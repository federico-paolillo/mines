package board

import "github.com/federico-paolillo/mines/pkg/dimensions"

type Generator interface {
	Generate(size dimensions.Size, mines int) *Board
}
