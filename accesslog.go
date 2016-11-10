package nginsort

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type AccessLog struct {
	Address   string
	User      string
	Date      time.Time
	Request   string
	Status    uint
	BytesSent uint64
	Refer     string
	UserAgent string

	Origin string
}

func Parse(line string) (*AccessLog, error) {
	delimiters := [8]string{` - `, ` [`, `] "`, `" `, ` `, ` "`, `" "`, `"`}
	fields := [8]string{}

	splitted := make([]string, 2)
	splitted[1] = line
	for i, d := range delimiters {
		splitted = strings.SplitN(splitted[1], d, 2)
		if len(splitted) < 2 {
			return nil, errors.New("Format error: '" + line + "'")
		}
		fields[i] = splitted[0]
	}

	layout := "02/Jan/2006:15:04:05 -0700"
	status, err := strconv.ParseUint(fields[4], 10, 64)
	if err != nil {
		return nil, err
	}
	bytesSent, err := strconv.ParseUint(fields[5], 10, 64)
	if err != nil {
		return nil, err
	}

	time, err := time.Parse(layout, fields[2])
	if err != nil {
		return nil, err
	}

	log := &AccessLog{
		Address:   fields[0],
		User:      fields[1],
		Date:      time,
		Request:   fields[3],
		Status:    uint(status),
		BytesSent: bytesSent,
		Refer:     fields[6],
		UserAgent: fields[7],
		Origin:    line,
	}
	return log, nil
}

type ByDate []AccessLog

func (s ByDate) Len() int {
	return len(s)
}
func (s ByDate) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByDate) Less(i, j int) bool {
	return s[i].Date.Unix() < s[j].Date.Unix()
}
