package entities

type (
	Settings struct {
		InAppProducts               []Product `json:"inAppProducts,omitempty"`
		NativeProducts              []Product `json:"nativeProducts,omitempty"`
		WebPaymentUrl               string    `json:"webPaymentUrl,omitempty"`
		PromoCodesEnabled           bool      `json:"promoCodesEnabled,omitempty"`
		WebPaymentMonthProductPrice Price     `json:"webPaymentMonthProductPrice"`
	}

	SettingsResult struct {
		ResponseBase
		Result Settings
	}
)
