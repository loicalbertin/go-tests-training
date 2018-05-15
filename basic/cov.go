package basic

import "log"

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

// SeeCoverage is a dummy function to used to illustrate code coverage
func SeeCoverage(e SomeEnum) {
	switch e {
	case E1:
		log.Println("E1 given")
	case E2:
		log.Println("E2 given")
	case E3:
		log.Println("E3 given")
	case E4:
		log.Println("E4 given")
	}
}
