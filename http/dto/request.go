package dto

type URLShortenerRequest struct {
	URL string `json:"url"`
}

func (u *URLShortenerRequest) Validate() error {
	if len(u.URL) == 0 {
		return &Error{"URL is required"}
	}
	return nil
}
