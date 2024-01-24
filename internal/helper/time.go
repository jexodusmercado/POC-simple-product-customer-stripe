package helper

import "time"

func ConvertEpochToTime(epochTime int64) string {
	// Convert Unix timestamp to time.Time
	timeObj := time.Unix(epochTime, 0)

	// Format the time to a readable date and time
	formattedTime := timeObj.Format("2006-01-02 15:04:05")

	return formattedTime
}
