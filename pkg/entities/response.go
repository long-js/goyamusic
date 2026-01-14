package entities

type (
	ResponseBase struct {
		Result           interface{}    `json:"result"`
		Error            string         `json:"error,omitempty"`
		ErrorDescription string         `json:"errorDescription,omitempty"`
		InvocationInfo   InvocationInfo `json:"invocationInfo,omitempty"`
	}

	InvocationInfo struct {
		Hostname           string `json:"hostname,omitempty"`
		ReqId              string `json:"reqId,omitempty"`
		ExecDurationMillis int    `json:"execDurationMillis,omitempty"`
	}
)
