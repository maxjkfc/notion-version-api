package notion

// Object -
type Object struct {
	Object         string `json:"object"`
	ID             string `json:"id"`
	CreatedTime    string `json:"created_time"`
	LastEditedTime string `json:"last_edited_time"`
	Parent         Parent `json:"parent"`
	Icon           Icon   `json:"icon"`
	Cover          Cover  `json:"cover"`
	URL            string `json:"url"`
}

// Parent - 母頁
type Parent struct {
	Type   string `json:"type"`
	PageID string `json:"page_id"`
}

// Icon - Notion Page Icon
type Icon struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}

// Cover - Notion Page Cover Image
type Cover struct {
	Type     string           `json:"type"`
	External CoverExternalURL `json:"external"`
}

// CoverExternalURL -
type CoverExternalURL struct {
	URL string `json:"url"`
}

// Title -
type Title struct {
	Type      string    `json:"type"`
	Text      TitleText `json:"text"`
	PlainText string    `json:"plain_text"`
	Herf      string    `json:"herf"`
}

// Annotations -
type Annotations struct {
	Bold          bool   `json:"bold"`          // 粗體
	Italic        bool   `json:"italic"`        // 斜體
	Strikethrough bool   `json:"strikethrough"` // 刪除線
	Underline     bool   `json:"underline"`     // 底線
	Code          bool   `json:"code"`          // 是否為代碼樣式
	Color         string `json:"color"`         // 字體顏色
}

// TitleText -
type TitleText struct {
	Content string `json:"content"`
	Link    Link   `json:"link"`
}

// Link -
type Link struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// Properties -
type Properties struct {
}
