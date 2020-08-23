package main

// Config wraps some configuration parameters that will be used in runtime.
type Config struct {
	Query          string
	ClearData      bool
	SMTPHost       string
	SMTPPass       string
	SMTPPort       int
	SMTPUser       string
	NotifyEmail    string
	WorkDir        string
	SearchInterval uint
}
