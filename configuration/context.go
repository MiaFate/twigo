package configuration

import "context"

var Ctx context.Context

func SetContext(ctx context.Context) {
	Ctx = context.TODO()
}
