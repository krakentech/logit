<picture>
    <source media="(prefers-color-scheme: dark)" srcset="./res/github-topper-dark.png" />
    <source media="(prefers-color-scheme: light)" srcset="./res/github-topper-light.png" />
    <img src="./res/github-topper-light.png" />
</picture>

![GitHub Release](https://img.shields.io/github/v/release/krakentech/logit)
![Coverage](https://img.shields.io/badge/Coverage-66.7%25-yellow)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/krakentech/logit)
![GitHub License](https://img.shields.io/github/license/krakentech/logit)

## Installation

```bash
go get github.com/krakentech/logit@latest
```

## Usage

### Setup

When you set a few options to customize your experience. These can be changed at any time.

#### IsDebug

If debug is set to true debug messages will be printed

```golang
logit.IsDebug(true)
```

#### SetOutFormat

You can change the output pattern of the message. The default pattern is ```{{time}} {{prefix}}  {{type}} - {{message}}``` here are a list of current properties you can use in your own pattern.

| key          | description                                       |
|--------------|---------------------------------------------------|
| {{time}}     | The current time based on the time pattern        |
| {{prefix}}   | If a prefix is set this is where it will be shown |
| {{type}}     | The Emoji representing the different msg types    |
| {{message}}  | The message that will be printed                  |

***Note:*** Please create an issue with other suggestions as this is what was needed at the time of writing.

#### SetTimeFormat

You can change the format of the date/time printed. The default is ```06.02.01-15:04:05``` and to make it custom just use standard time formatting for golang.

#### SetWriter

Change where the logs as sent. This takes any io.Writer

### Code Usage

