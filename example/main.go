package main

import (
    "fmt"
    "github.com/krakentech/logitv2"
)

type TestItem struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {

    myTestItem := TestItem{Name: "john", Age: 21}

    logitv2.Debug("you should not see this")
    logitv2.SetIsDebug(true)
    logitv2.Debug("you should see this")

    logitv2.Info("This is an %s line", "info")
    logitv2.Warn("This is an %s line", "warning")
    logitv2.Err("This is an %s line", "error")
    logitv2.Error(fmt.Errorf("this is the error message"), "This is an %s line", "error")
    logitv2.DebugData("data as json", myTestItem)
    logitv2.DebugDataFormated("data as json formated", myTestItem)

}
