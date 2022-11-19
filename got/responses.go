package got

import "fmt"

type assetsResponse struct {
	Assets []Asset
}

type assetResponse struct {
	Asset Asset
}

type Asset struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Director     string `json:"director"`
	Release_date string `json:"release_date"`
}

func (d Asset) Info() string {
	return fmt.Sprintf("[ID] %s \n[Title] %s \n[Director] %s \n[ReleaseDate] %s",
		d.ID, d.Title, d.Director, d.Release_date)
}
