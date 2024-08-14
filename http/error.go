package httppinger

import "fmt"

type UnexpectedStatusCodeError struct {
	StatusCode int
}

func (u UnexpectedStatusCodeError) Error() string {
	return fmt.Sprintf("ping failed with %d status code", u.StatusCode)
}
