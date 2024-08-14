package utils

import (
	"context"
	"scylla/dto"
)

func ResponseInterceptor(ctx context.Context, resp *dto.Response) {
	traceIdInf := ctx.Value("requestid")
	traceId := ""
	if traceIdInf != nil {
		traceId = traceIdInf.(string)
	}
	resp.TraceID = traceId
}
