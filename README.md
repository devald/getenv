# getenv
## Usage
```
go get github.com/devald/getenv
```
## func Variable
// Variable retrieves the value of the environment variable with error handling
```go
func Variable(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}
```
