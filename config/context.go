package config

import "context"

var Context context.Context

func initContext() {
	Context = context.Background()
}
