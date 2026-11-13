<picture>
    <source media="(prefers-color-scheme: dark)" srcset="./res/github-topper-dark.jpg" />
    <source media="(prefers-color-scheme: light)" srcset="./res/github-topper-light.jpg" />
    <img src="./res/github-topper-light.jpg" />
</picture>

![Static Badge](https://img.shields.io/badge/Release-v1.0.0-green)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

## Overview

Logit is a quick module to simplify basic logging with some extra flair.

## Installation

```bash
go get github.com/krakentech/logit@latest
```

## Usage

### Debug

**Example**

```golang
logit.Debug("This is an %s line", "debug")
```

**Gets**

```shell
00.00.00-00:00:00 🐛 - This is an debug line
```

### Debug Data

**Example (not formatted)**

```golang
logit.DebugData("data as json", false, myTestItem)
```

**Gets**

```shell
00.00.00-00:00:00  🤖 - =================================================================[data as json]=
{"name":"john","age":21}
```

**Example (formatted)**

```golang
logit.DebugData("data as json formated", true, myTestItem)
```

**Gets**

```shell
00.00.00-00:00:00  🤖 - =======================================================[data as json formatted]=
{
        "name": "john",
        "age": 21
}
```

### Info

**Example**

```golang
logit.Info("This is an %s line", "info")
```

**Gets**

```shell
00.00.00-00:00:00 🧠 - This is an info line
```

### Warning

**Example**

```golang
logit.Warning("This is an %s line", "warning")
```

**Gets**

```shell
00.00.00-00:00:00 ⚠️ - This is an warning line
```

### Error

**Example (Error Message)**

```golang
logit.Error("This is an %s line", "error")
```

**Gets**

```shell
00.00.00-00:00:00 ❌ - This is an error line
```

**Example (Error Message with Error)**

```golang
logit.Error(fmt.Errorf("this is the error message"), "This is an %s line", "error")
```

**Gets** 

```shell
00.00.00-00:00:00 ❌ - This is an error line: this is the error message
```





## Settings

### SetWriter

```golang
func SetWriter(newWriter io.Writer)
```

This will update where the output is written. Default is ```os.Stdout```

**Example**

```golang
logit.SetWriter(os.Stderr)
```

### SetMessageTemplate

```golang
func SetMessageTemplate(format string) 
```

The format string can be any string you want the following elements will be replaced when generating output

* **{{time}}**: This is where the date/time will set
* **{{typeIcon}}**: Based on the message type this will be one of the emojis that represent the message type
* **{{typeText}}**: Based on the message type this will be one of the text representations that represent the message type
* **{{message}}**: This is where the passed message will be set.

### SetIsDebug

```golang
func SetDebug(debugMode bool) 
```

If set to true debug messages will print otherwise they will not

**Example**

```golang
logit.SetDebug(true)
```



