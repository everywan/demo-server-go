package app

import "time"

const DefaultStopTimeout = time.Minute

type appConfig struct {
	stopTimeout time.Duration // app stop timeoit
	profilePort int           // go profile port
}
