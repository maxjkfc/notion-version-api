package notion

import (
	"github.com/go-resty/resty/v2"
	"github.com/maxjkfc/notion-version-api/infra/config"
)

// API -
type API struct {
	host       string
	version    string
	token      string
	databaseID string
}

// NewAPI -
func NewAPI(c config.Config) (api API, err error) {
	api = API{
		host:       c.Notion.Host,
		version:    c.Notion.Version,
		token:      c.Notion.Token,
		databaseID: c.Notion.Database.ID,
	}

	return
}

func (a *API) newRequest() *resty.Request {

	cli := resty.New()

	req := cli.R().SetAuthScheme("Bearer")
	req.SetAuthToken(a.token)
	req.SetHeader("Notion-Version", a.version)

	return req
}
