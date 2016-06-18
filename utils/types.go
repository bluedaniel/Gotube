package utils

// Tube contains ftl data
type Tube struct {
	Name         string
	LineStatuses []struct {
		StatusSeverity            int
		StatusSeverityDescription string
		Reason                    string
	}
}

// StopPointSearchResp contains stopPoint data
type StopPointSearchResp struct {
	Total   int
	Matches []struct {
		IcsID string
		Name  string
		ID    string
	}
}

// StopPointDataResp gets the station ID
type StopPointDataResp struct {
	Lines []struct {
		ID string
	}
}
