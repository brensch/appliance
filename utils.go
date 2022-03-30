package smarthome

func SameLocation(loc1, loc2 Location) bool {
	return loc1.X == loc2.X && loc1.Y == loc2.Y
}

func LocationIsHouse(height, team int8, l Location) bool {
	return (team == 1 && l.Y >= height) || (team == -1 && l.Y < 0)
}

func LocationValid(width, height int8, l Location) bool {
	return l.X >= 0 && l.Y >= 0 && l.X < width && l.Y < height
}
