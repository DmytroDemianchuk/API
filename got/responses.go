package got

import "fmt"

type assetsResponse struct {
	Assets    []Asset `json:"data"`
	Timestamp int64   `json:"timestamp"`
}

type assetResponse struct {
	Asset     Asset `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type Asset struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Family    string `json:"family"`
}

func (d Asset) Info() string {
	return fmt.Sprintf("[ID] %s \n[FirstName] %s \n[LastName] %s \n[Family] %s",
		d.ID, d.FirstName, d.LastName, d.Family)
}
