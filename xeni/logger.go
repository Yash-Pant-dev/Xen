package xeni

import (
	"fmt"
	"os"
)

const logs_file_name = "xeni_logs.log"
const metrics_file_name = "xeni_metrics.log"

var logs *os.File = nil
var metrics *os.File = nil

/*
Severity description:
0 - Debug, Log file only.
1 - Debug, Log file + Terminal
2 - Metrics
3 - Warnings
4 - Errors
5 - Errors with immediate fatality
*/
func Log(severity int, messages ...any) {

	joinedMessage := fmt.Sprint(messages, "\n")

	switch severity {
	case 0:
		if debugMode && logs != nil {
			logs.WriteString(joinedMessage)
		}
	case 1:
		if debugMode && logs != nil {
			logs.WriteString(joinedMessage)
			fmt.Print(joinedMessage)
		}
	case 2:
		if metrics != nil {
			metrics.WriteString(joinedMessage)
		}
	case 3:
		if logs != nil {
			message := fmt.Sprint("[W] ", joinedMessage)
			logs.WriteString(message)
			fmt.Print(message)
		}
	case 4:
		if logs != nil {
			message := fmt.Sprint("[E]", joinedMessage)
			logs.WriteString(message)
			fmt.Print(message)
		}
	case 5:
		if logs != nil {
			message := fmt.Sprint("[E]", joinedMessage)
			logs.WriteString(message)
			fmt.Print(message)
			panic(message)
		}
	}
}

func initializeLogging() {

	var errLogger, errMetrics error

	logs, errLogger = os.OpenFile(logs_file_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errLogger != nil {
		fmt.Println("Cannot open logs file [", logs_file_name, "]")
		logs = nil
	} else {
		Log(0, "")
		Log(1, "Logger initialized.")
	}

	metrics, errMetrics = os.OpenFile(metrics_file_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errMetrics != nil {
		Log(3, "Cannot open metrics file [", metrics_file_name, "]")
		metrics = nil
	}
}
