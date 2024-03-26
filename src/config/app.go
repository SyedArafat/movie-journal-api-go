package config

import "os"

const (
	DateFormat         = "2006-01-02"
	ServerErrorMessage = "Something went wrong"
)

var FileLog = os.Getenv("WRITE_LOG_IN_FILE")
