package dimensions

type Location struct {
	X, Y int
}

// This function returns all the adjacent locations around an origin.
//
// An adjacent location is a location that is ± 1 away from the origin.
func (origin Location) AdjacentLocations() [8]Location {
	// From 12 o-clock, clockwise
	return [8]Location{
		{origin.X, origin.Y + 1},
		{origin.X + 1, origin.Y + 1},
		{origin.X + 1, origin.Y},
		{origin.X + 1, origin.Y - 1},
		{origin.X, origin.Y - 1},
		{origin.X - 1, origin.Y - 1},
		{origin.X - 1, origin.Y},
		{origin.X - 1, origin.Y + 1},
	}
}
