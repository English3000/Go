// Function: AddGigasecond, receives a time & returns the time a gigagsecond later
package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	//OR just: return t.add(time.Second * time.Duration(1E9))

	timeToAdd := 1000000000 //seconds
	// months, days, hours, minutes
	years := timeToAdd / 365 / 24 / 60 / 60
	timeToAdd -= years * 365 * 24 * 60 * 60
	months := timeToAdd / 30 / 24 / 60 / 60
	timeToAdd -= months * 30 * 24 * 60 * 60
	days := timeToAdd / 24 / 60 / 60
	timeToAdd -= days * 24 * 60 * 60
	hours := timeToAdd / 60 / 60
	timeToAdd -= hours * 60 * 60
	minutes := timeToAdd / 60
	timeToAdd -= minutes * 60

	t = t.AddDate(years, months, days)
	return t.Add(time.Hour*time.Duration(hours) +
		time.Minute*time.Duration(minutes) +
		time.Second*time.Duration(timeToAdd)) //close enough
}
