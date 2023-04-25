package models

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Time time.Time

const LAYOUT = "15:04"

func (t *Time) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*t = Time(nullTime.Time)
	return
}

func (t Time) Value() (driver.Value, error) {
	y, m, d := time.Time(t).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Time(t).Location()), nil
}

// GormDataType gorm common data type
func (t Time) GormDataType() string {
	return "date"
}

func (t Time) GobEncode() ([]byte, error) {
	return time.Time(t).GobEncode()
}

func (t *Time) GobDecode(b []byte) error {
	return (*time.Time)(t).GobDecode(b)
}

func (t Time) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, time.Time(t).Format(LAYOUT))), nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	parsedDate, err := time.Parse(LAYOUT, s)
	*t = Time(parsedDate)
	return err
}

func (t Time) String() string {
	return time.Time(t).Format(LAYOUT)
}
