package internal

import (
	"flag"
	"os"
)

type Appstate struct {
	env     map[string]string
	Prefork bool
	Port    int
}

var (
	State Appstate

	prefork = flag.Bool("prefork", false, "enable prefork")
	port    = flag.Int("port", 3000, "port to listen on")
	tz      = flag.String("tz", "Asia/Jakarta", "timezone")
)

func Setup() {
	flag.Parse()
	State.env = make(map[string]string)
	State.Prefork = *prefork
	State.Port = *port

	os.Setenv("TZ", *tz)
}
