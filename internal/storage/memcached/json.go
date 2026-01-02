package memcached

import "github.com/federico-paolillo/mines/pkg/matchmaking"

type MatchstateJSON struct {
	Id        string     `json:"id"`
	Lives     int        `json:"lives"`
	State     string     `json:"state"`
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Cells     []CellJSON `json:"cells"`
	StartTime int64      `json:"startTime"`
}

type CellJSON struct {
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Mined bool   `json:"mined"`
	State string `json:"state"`
}

func MatchstateToJSON(match *matchmaking.Matchstate) *MatchstateJSON {
	cells := make([]CellJSON, 0, match.Height*match.Width)

	for _, row := range match.Cells {
		for _, cell := range row {
			cells = append(cells, CellJSON{
				X:     cell.X,
				Y:     cell.Y,
				Mined: cell.Mined,
				State: cell.State,
			})
		}
	}

	return &MatchstateJSON{
		Id:        match.Id,
		Lives:     match.Lives,
		State:     match.State,
		Width:     match.Width,
		Height:    match.Height,
		Cells:     cells,
		StartTime: match.StartTime,
	}
}

func JSONToMatchstate(json *MatchstateJSON) *matchmaking.Matchstate {
	cells := make(matchmaking.Cells, json.Height)
	for i := range cells {
		cells[i] = make([]matchmaking.Cell, json.Width)
	}

	for i, cellJSON := range json.Cells {
		yIndex := i / json.Width
		xIndex := i % json.Width

		cells[yIndex][xIndex] = matchmaking.Cell{
			X:     cellJSON.X,
			Y:     cellJSON.Y,
			Mined: cellJSON.Mined,
			State: cellJSON.State,
		}
	}

	return &matchmaking.Matchstate{
		Id:        json.Id,
		Lives:     json.Lives,
		State:     json.State,
		Width:     json.Width,
		Height:    json.Height,
		Cells:     cells,
		StartTime: json.StartTime,
	}
}
