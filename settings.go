package logit

const (
    OutFormatDefault  = "{{time}} {{type}} - {{message}}"
    TimeFormatDefault = "06.02.01-15:04:05"
)

type logType string

var (
    logTypeDebug logType = "🐛"
    logTypeInfo  logType = "🧠"
    logTypeWarn  logType = "🚧"
    logTypeError logType = "🛑"
    logTypeData  logType = "🧶"
)
