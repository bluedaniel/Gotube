package utils

import "strconv"

// TubeStatus returns all tube statuses
func TubeStatus() string {
	return "https://api.tfl.gov.uk/line/mode/tube/status"
}

// StopPointSearchURL search for station
func StopPointSearchURL(query string) string {
	return "https://api.tfl.gov.uk/StopPoint/Search/" + query + "?modes=tube"
}

// StopPointURL data for station
func StopPointURL(ID string) string {
	return "https://api.tfl.gov.uk/StopPoint/" + ID
}

// StopPointDeadline is HTML for first/last trains from station
func StopPointDeadline(ID string, line string, night bool) string {
	return "https://tfl.gov.uk/Timetables/FirstLastServicesSummaryAjax?fromId=" + ID + "&lines=" + line + "&firstNextDay=" + strconv.FormatBool(night)
}
