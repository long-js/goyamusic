package entities

type Vinyl struct {
	Url       string `json:"url,omitempty"`
	Title     string `json:"title,omitempty"`
	Picture   string `json:"picture,omitempty"`
	Media     string `json:"media,omitempty"`
	ArtistIds []int  `json:"artistIds,omitempty"`
	Year      int    `json:"year,omitempty"`
	Price     int    `json:"price,omitempty"`
	OfferId   int    `json:"offerId,omitempty"`
}
