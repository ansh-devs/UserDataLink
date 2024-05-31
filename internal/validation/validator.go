package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

func ValidateIdsList(ids []int64, maxIdLen int) []error {
	var errsx []error
	for _, userID := range ids {
		if userID == 0 {
			errsx = append(errsx, fmt.Errorf("id must be greater than 0"))
		}
		if userID > int64(maxIdLen) {
			errsx = append(errsx, fmt.Errorf("id must be smaller than %d", maxIdLen))

		}
	}
	return errsx
}

func IsValidCriteria(criteria string) error {
	lowerCaseCriteria := strings.ToLower(criteria)
	if lowerCaseCriteria == "phone" || lowerCaseCriteria == "city " || lowerCaseCriteria == "married" || lowerCaseCriteria == "height" {
		return nil
	} else {
		return fmt.Errorf("incorrect enum value passed : must only containe { PHONE, CITY, MARRIED, HEIGHT }")
	}

}
