package logger

import (
	"fmt"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

// Example usage: `logger.LogInfo("Hello, %s", "world")`
func LogInfo(message string, args ...interface{}) {
	fmt.Printf(Green+message+Reset+"\n", args...)
}

// Example usage: `logger.LogError("Hello, %s", "world")`
func LogError(message string, args ...interface{}) {
	fmt.Printf(Red+message+Reset+"\n", args...)
}

// Example usage: `logger.LogWarning("Hello, %s", "world")`
func LogWarning(message string, args ...interface{}) {
	fmt.Printf(Yellow+message+Reset+"\n", args...)
}
