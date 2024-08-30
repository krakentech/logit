package logit

import "fmt"

// ErrDecorator will add the passed description to the front of the error message on the passed error
func ErrDecorator(err error, desc string) error {
    if err == nil {
        return err
    }
    return fmt.Errorf("%s: %s", desc, err.Error())
}
