package app

import "context"

type (
	IBundle interface {
		GetName() string
		Run(context.Context)
		Stop(context.Context)
	}
)
