package entities

type Video struct {
	Title                   string   `json:"title,omitempty"`
	Cover                   string   `json:"cover,omitempty"`
	EmbedUrl                string   `json:"embedUrl,omitempty"`
	Provider                string   `json:"provider,omitempty"`
	YoutubeUrl              string   `json:"youtubeUrl,omitempty"`
	ThumbnailUrl            string   `json:"thumbnailUrl,omitempty"`
	Text                    string   `json:"text,omitempty"`
	HtmlAutoPlayVideoPlayer string   `json:"htmlAutoPlayVideoPlayer,omitempty"`
	Regions                 []string `json:"regions,omitempty"`
	ProviderVideoId         int      `json:"providerVideoId,omitempty"`
	Duration                int      `json:"duration,omitempty"`
}
