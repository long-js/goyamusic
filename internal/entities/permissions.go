package entities

type (
	Permissions struct {
		Until   string   `json:"until,omitempty"`
		Values  []string `json:"values,omitempty"`
		Default []string `json:"default,omitempty"`
	}

	PermissionAlerts struct {
		Alerts []string `json:"alerts"`
	}

	PermissionAlertsResult struct {
		ResponseBase
		Result PermissionAlerts
	}
)
