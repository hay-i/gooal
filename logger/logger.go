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

func LogInfo(message string, args ...interface{}) {
	fmt.Printf(Green+message+Reset+"\n", args...)
}

func LogError(message string, args ...interface{}) {
	fmt.Printf(Red+message+Reset+"\n", args...)
}

func LogWarning(message string, args ...interface{}) {
	fmt.Printf(Yellow+message+Reset+"\n", args...)
}
