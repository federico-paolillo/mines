package server

// POST /game 						=> 200 (GameState)
// POST /game/<uuid>/move => 200 (GameState), 400 (Validation), 422 (Game Lost/Won)
