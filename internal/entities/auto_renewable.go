package entities

type AutoRenewable struct {
	Expires       string  `json:"expires,omitempty"`
	Vendor        string  `json:"vendor,omitempty"`
	VendorHelpUrl string  `json:"vendorHelpUrl,omitempty"`
	ProductId     string  `json:"productId,omitempty"`
	OrderId       int     `json:"orderId,omitempty"`
	Finished      bool    `json:"finished,omitempty"`
	Product       Product `json:"product"`
	MasterInfo    User    `json:"masterInfo"`
}
