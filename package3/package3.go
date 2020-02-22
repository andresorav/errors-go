package package2

import (
	"fmt"
	"github.com/pkg/errors"
)

type PackageError struct {
	Code string
	Err  error
}

func (e *PackageError) Error() string {
	return fmt.Sprintf("Got an package error: %s", e.Err.Error())
}

func RunFunction() error {
	return &PackageError{
		Code: "123",
		Err:  errors.New("bad package"),
	}
}
