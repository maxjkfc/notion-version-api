package notion

import "github.com/maxjkfc/notion-version-api/infra/config"

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
