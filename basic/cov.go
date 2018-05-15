package basic

import (
	"log"

	"github.com/pkg/errors"
)

// SomeEnum is an enumeration example
type SomeEnum int

const (
	// E1 is the first value of this enumaration
	E1 SomeEnum = iota
	// E2 is ...
	E2
	// E3 is ...
	E3
	// E4 is ...
	E4
)

// EnumAsString is a dummy function to used to illustrate code coverage
func EnumAsString(e SomeEnum) (string, error) {
	switch e {
	case E1:
		log.Println("E1 given")
		return "E1", nil
	case E2:
		log.Println("E2 given")
		return "E2", nil
	case E3:
		log.Println("E3 given")
		return "E3", nil
	case E4:
		log.Println("E4 given")
		return "E4", nil
	}

	return "", errors.Errorf("unexpected value %d for SomeEnum", e)
}
