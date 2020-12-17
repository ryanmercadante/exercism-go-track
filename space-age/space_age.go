package space

// Planet type is a string that should represent a Planet
type Planet string

const earthYearInSeconds = 31557600

// Age is a function that takes seconds and planet as parameters
// and determines your age on each planet
func Age(seconds float64, planet Planet) float64 {
	earthYears := getEarthYears(seconds)

	switch planet {
	case "Mercury":
		return earthYears / 0.2408467
	case "Venus":
		return earthYears / 0.61519726
	case "Earth":
		return earthYears
	case "Mars":
		return earthYears / 1.8808158
	case "Jupiter":
		return earthYears / 11.862615
	case "Saturn":
		return earthYears / 29.447498
	case "Uranus":
		return earthYears / 84.016846
	case "Neptune":
		return earthYears / 164.79132
	default:
		return 0.0
	}
}

// Helper function that takes seconds as a parameter and
// returns the equivalent time in earth years
func getEarthYears(seconds float64) float64 {
	return seconds / 31557600
}
