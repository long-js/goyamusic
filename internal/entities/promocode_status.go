package entities

type (
	PromocodeStatus struct {
		Status        string
		Status_desc   string
		AccountStatus AccountStatus
	}

	PromocodeStatusResult struct {
		ResponseBase
		Result PromocodeStatus `json:"result"`
	}
)
