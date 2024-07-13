package utils

import (
	"context"
	"scylla/entity"
)

func ResponseInterceptor(ctx context.Context, resp *entity.Response) {
	traceIdInf := ctx.Value("requestid")
	traceId := ""
	if traceIdInf != nil {
		traceId = traceIdInf.(string)
	}
	resp.TraceID = traceId
}
