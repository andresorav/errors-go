package package1

import (
	"github.com/andresorav/errors-go/package2"
	"github.com/pkg/errors"
)

func RunFunction() error {
	err := package2.RunFunction()

	if err != nil {
		return errors.Wrap(err, "Failed to run function")
	}
}
