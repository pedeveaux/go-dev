package goalpacegenerator

import "math"

type paceUnit struct {
	meters  int
	minutes int
	seconds int
}

// convertPace takes a time in minutes and seconds and a distance
// in meters and converts it to an equivlanet pace at a shorter ..distance
func (s *paceUnit) convertPace(goalMinutes int, goalSeconds int, goalDistance int) {
	totalSeconds := float64(goalMinutes*60 + goalSeconds)
	goalPace := totalSeconds / float64(goalDistance)

	pace := goalPace * float64(s.meters)
	if pace < 100 {
		s.seconds = int(pace)
	} else {
		s.minutes = int(pace / 60)
		s.seconds = int(math.Mod(pace, 60))
	}
}
