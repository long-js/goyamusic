package entities

type User struct {
	Login       string `json:"login,omitempty"`
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	FullName    string `json:"fullName,omitempty"`
	Regions     []int  `json:"regions,omitempty"`
	Uid         uint32 `json:"uid,omitempty"`
	Verified    bool   `json:"verified,omitempty"`
}
