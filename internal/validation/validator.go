package validation

import (
	"fmt"
	"regexp"
	"strconv"
)

// isValidIdExp check whether the given parameter is valid int or not
var isValidIdExp = regexp.MustCompile("^[0-9_]+$").MatchString

// IsValidId validates the given id
func IsValidId(id int64) error {
	if id <= 0 {
		return fmt.Errorf("id must be greater than zero")
	}
	if id > 15 {
		return fmt.Errorf("id must be smaller than or equal to 15")
	}
	if !isValidIdExp(strconv.Itoa(int(id))) {
		return fmt.Errorf("id must only contain digits")
	}
	return nil
}
