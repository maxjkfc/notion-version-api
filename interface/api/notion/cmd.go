package notion

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// CheckDatabase -
func (a *API) CheckDatabase() error {
	req := a.newRequest()
	req.SetHeader("Content-Type", "application/json")
	resp, err := req.Get(fmt.Sprintf("%s/%s/%s", a.host, "v1/databases/", a.databaseID))
	if err != nil {
		return err
	}

	spew.Dump(resp.RawResponse)
	return nil
}
