# logger
--
    import "."


## Usage

#### func  Debugf

```go
func Debugf(format string, i ...interface{})
```
Debugf will output a debug message to the log

#### func  ErrorErr

```go
func ErrorErr(err error)
```
ErrorErr will output an error to the log

#### func  Errorf

```go
func Errorf(format string, i ...interface{})
```
Errorf will output an error to the log

#### func  GetStackTrace

```go
func GetStackTrace(max int) string
```
GetStackTrace returns a stack trace

#### func  Message

```go
func Message(s string)
```
Message will write s to the log

#### func  Messagef

```go
func Messagef(format string, i ...interface{})
```
Messagef will write a formated message to the log
