package logit

import "fmt"

func ErrDecorator(err error, desc string) error {
    return fmt.Errorf("%s: %s", desc, err.Error())
}
