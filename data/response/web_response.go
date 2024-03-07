package response

type Response struct {
	Code     int         `json:"code"`
	Status   string      `json:"status"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
	TraceID  string     `json:"trace_id"`
}

type Error struct {
	Code     int         `json:"code"`
	Status   string      `json:"status"`
	Errors   interface{} `json:"errors,omitempty"`
	// TraceID  string     `json:"trace_id"`
}