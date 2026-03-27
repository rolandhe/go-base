package swiss_kit

import "time"

func OfTimeRFC3339Nano(v string) (time.Time, error) {
	return time.Parse(time.RFC3339Nano, v)
}

func OnlyDateTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

// DiffDays returns the number of calendar days between two times.
// It compares dates in each time's own location, unaffected by DST.
func DiffDays(small time.Time, big time.Time) int64 {
	sy, sm, sd := small.Date()
	by, bm, bd := big.Date()
	smallDay := time.Date(sy, sm, sd, 0, 0, 0, 0, time.UTC)
	bigDay := time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	return int64(bigDay.Sub(smallDay) / (24 * time.Hour))
}

func IntDateFromUnixMilli(millis int64) int64 {
	t := time.UnixMilli(millis)
	return IntDate(t)
}

func IntDate(t time.Time) int64 {
	v := t.Year() * 10000
	v += int(t.Month()) * 100
	v += t.Day()
	return int64(v)
}
