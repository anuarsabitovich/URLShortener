package getlink

import (
	"net/http"
)

func GetFullLink(shortLink string) (string, error) {
	resp, err := http.Get(shortLink)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 && resp.StatusCode <= 399 {
		location, err := resp.Location()
		if err != nil {
			return "", err
		}
		return location.String(), nil
	}

	return shortLink, nil
}
