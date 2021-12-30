package datetime

import "time"

const DefFormat = time.RFC3339

// ToEpochMS return epoch time in ms by normalize time to UTC first
func ToEpochMS(t time.Time) int64 {
	return t.UTC().UnixMilli()
}

// ToEpoch return epoch time in second by normalize time to UTC first
func ToEpoch(t time.Time) int64 {
	return t.Truncate(time.Second).UTC().Unix()
}

// Now return UTC time with format RFC3339Nano by truncate it to ms
func Now() time.Time {
	return time.Now().UTC().Truncate(time.Second)
}

func ToString(t time.Time) string {
	return t.Format(DefFormat)
}

func FromString(t string) (time.Time, error) {
	return time.Parse(DefFormat, t)
}
