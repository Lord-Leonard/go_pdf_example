package date

import (
	"strings"
	"time"
)

type Date time.Time

func (c *Date) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("02.04.2006", value) //parse time
	if err != nil {
		return err
	}
	*c = Date(t) //set result using the pointer
	return nil
}

func (c Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("02.04.2006") + `"`), nil
}
