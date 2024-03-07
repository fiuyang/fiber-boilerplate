package utils

import (
	"context"
	"fiber-boilerplate/data/response"
)

func ResponseInterceptor(ctx context.Context, resp *response.Response) {
	traceIdInf := ctx.Value("requestid")
	traceId := ""
	if traceIdInf != nil {
		traceId = traceIdInf.(string)
	}
	resp.TraceID = traceId
}
