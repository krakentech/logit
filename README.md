<picture>
    <source media="(prefers-color-scheme: dark)" srcset="./res/readme-topper-dark.jpg" />
    <source media="(prefers-color-scheme: light)" srcset="./res/readme-topper-light.jpg" />
    <img src="./res/readme-topper-light.jpg" />
</picture>

![Static Badge](https://img.shields.io/badge/Release-v1.0.0-green)
![Coverage](https://img.shields.io/badge/Coverage-76.5%25-brightgreen)

## Overview

Logit is a quick module to simplify basic logging with some extra flair.

## Installation

```bash
go get github.com/krakentech/logit@latest
```

## Usage

### Examples

#### Debug

```golang
logit.Debug("This is an %s line", "debug")
```

Gets
```24.17.12-20:11:44 🐛 - This is an debug line```

```golang
logit.DebugData("data as json", myTestItem)
```

Gets
```
24.17.12-20:25:44 🧶 - =================================================================[data as json]=
{"name":"john","age":21}
24.17.12-20:25:44 🧶 - ================================================================================
```

```golang
logit.DebugDataFormated("data as json formated", myTestItem)
```

Gets
```
24.17.12-20:25:44 🧶 - ========================================================[data as json formated]=
{
        "name": "john",
        "age": 21
}
24.17.12-20:25:44 🧶 - ================================================================================

```


#### Info

```golang
logit.Info("This is an %s line", "info")
```

Gets
```24.17.12-20:25:44 🧠 - This is an info line```


#### Warning

```golang
logit.Warn("This is an %s line", "warning")
```

Gets
```24.17.12-20:25:44 🚧 - This is an warning line```

#### Error

```golang
logit.Err("This is an %s line", "error")
```

Gets
```24.17.12-20:25:44 🛑 - This is an error line```

```golang
logit.Error(fmt.Errorf("this is the error message"), "This is an %s line", "error")
```

Gets
```24.17.12-20:25:44 🛑 - This is an error line: this is the error message```


### Settings

#### SetWriter

```golang
func SetWriter(newWriter io.Writer)
```

This will update where the output is written. Default is ```os.Stdout```

#### SetOutFormat
    
```golang
func SetOutFormat(format string) 
```

The format string can be any string you want the following elements will be replaced when generating output

* **{{time}}**: This is where the date/time will set
* **{{type}}**: This is where the different logging types (error, debug, info, warning) icons will be placed
* **{{message}}**: This is where the passed message will be set.

#### SetTimeFormat

```golang
func SetTimeFormat(format string)
```

This takes any valid golang time format string

#### SetIsDebug

```golang
func SetIsDebug(debugMode bool) 
```

If set to true debug messages will print otherwise they will not

#### SetTestMode

```golang
func SetTestMode(testMode bool)
```

This is used for testing and default the date/time printed (if date/time is in the format) will be "00.00.00-00:00:00". This is useful if you want to test logged data

