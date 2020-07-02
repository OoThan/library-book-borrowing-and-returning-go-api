package format

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "email") {
		return errors.New("Invalid Email! ")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password! ")
	}
	return errors.New("Incorrect Details! ")
}
