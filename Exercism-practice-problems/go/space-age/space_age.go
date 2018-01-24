package space

//type Planet string

/* var EarthYearFrac = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
} */

// func Age(age time.Duration planet Planet) {
func Age(age float64, planet string) float64 {
	//var years string
	switch {
	case planet == "Earth":
		return age / 31557600
	case planet == "Mercury":
		return age / 7600544
	case planet == "Venus":
		return age / 19414149
	case planet == "Mars":
		return age / 59354033
	case planet == "Jupiter":
		return age / 374355659
	case planet == "Saturn":
		return age / 929292363
	case planet == "Uranus":
		return age / 2651370019
	case planet == "Neptune":
		return age / 5200418560
	default:
		return 0
	}
	//most performant concatenation in Go:
	// http://herman.asia/efficient-string-concatenation-in-go#and-the-winner-is_1
	// var result bytes.Buffer
	// result.WriteString(string(int(years)))
	// result.WriteString(" ")
	// result.WriteString(planet)
	// result.WriteString("-years old")
	// return result.String()
}
