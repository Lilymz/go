package greeting

import (
	"fmt"
	"runtime"
)

func Hello(name string) string {
	return fmt.Sprintf("hello,%v,welcome to golang !", name)
}
func Version() (version string) {
	version = runtime.Version()
	return
}
