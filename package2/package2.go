package package2

import (
	"github.com/andresorav/errors-go/package3"
	"github.com/pkg/errors"
)

func RunFunction() error {
	err := package3.RunFunction()

	if err != nil {
		return errors.Wrap(err, "Failed to run function")
	}

	return nil
}
