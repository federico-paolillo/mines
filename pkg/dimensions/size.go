package dimensions

type Size struct {
	Width, Height int
}

// This function checks if a location is within the boundaries of the size.
// 0, 0 is never within bounds.
func (size *Size) Contains(location Location) bool {
	xIsContained := location.X > 0 && location.X <= size.Width
	yIsContained := location.Y > 0 && location.Y <= size.Height

	return xIsContained && yIsContained
}

func (size *Size) Area() int {
	return size.Width * size.Height
}
