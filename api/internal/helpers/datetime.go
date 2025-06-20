package helpers

import "time"

// ParseDateTime attempts to parse a date string using common date-time formats.
// It is useful for handling date strings from various sources with different formats.
func ParseDateTime(dateTimeString string) (time.Time, error) {
	submittedAt, err := time.Parse(time.RFC3339, dateTimeString)

	if err != nil {
		submittedAt, err = time.Parse("2006-01-02T15:04:05-07:00", dateTimeString)

		if err != nil {
			return time.Time{}, err
		}
	}

	return submittedAt, nil
}