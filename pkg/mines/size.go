package mines

type Size struct {
	Width, Height int
}

// This function checks if a location is within the boundaries of the size.
//
// It is assumed that the size has origin 0,0 at bottom-left of the screen,
// as if it was the 1st quadrant of Cartesian Plane.
func (size *Size) Contains(location Location) bool {
	return location.X <= size.Width && location.Y <= size.Height
}
