package main

import (
    "fmt"
    logit "github.com/krakentech/logit_prev"
)

type TestItem struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {

    myTestItem := TestItem{Name: "john", Age: 42}

    // Run Methods Without Debug
    fmt.Println("Run The commands without 'Debug' set on")
    logit.Debug("You should not see this")
    logit.Info("This is an %s line", "info")
    logit.Warn("This is an %s line", "warning")
    logit.Err("This is an %s line", "error")
    logit.Error(fmt.Errorf("this is the error message"), "This is an %s line", "error")
    logit.DebugData("You will not see this", myTestItem)

    logit.SetIsDebug(true)
    fmt.Println("Run The commands with 'Debug' set on")
    logit.Debug("You should now see me")
    logit.Info("This is an %s line", "info")
    logit.Warn("This is an %s line", "warning")
    logit.Err("This is an %s line", "error")
    logit.Error(fmt.Errorf("this is the error message"), "This is an %s line", "error")
    logit.DebugData("You will see this", myTestItem)
}
