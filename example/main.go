package main

import (
    "fmt"
    "github.com/krakentech/logit"
)

type TestItem struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    myTestItem := TestItem{Name: "john", Age: 21}
    logit.Debug("you should not see this")
    logit.SetIsDebug(true)
    logit.Debug("This is an %s line", "debug")
    logit.Info("This is an %s line", "info")
    logit.Warn("This is an %s line", "warning")
    logit.Err("This is an %s line", "error")
    logit.Error(fmt.Errorf("this is the error message"), "This is an %s line", "error")
    logit.DebugData("data as json", myTestItem)
    logit.DebugDataFormated("data as json formated", myTestItem)
}
