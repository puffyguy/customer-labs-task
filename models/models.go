package models

type Input struct {
	Data map[string]interface{} `json:"data" binding:"required"`
}

type Output struct {
	Event           string                 `json:"event"`
	EventType       string                 `json:"event_type"`
	AppId           string                 `json:"app_id"`
	UserId          string                 `json:"user_id"`
	MessageId       string                 `json:"message_id"`
	PageTitle       string                 `json:"page_title"`
	PageURL         string                 `json:"page_url"`
	BrowserLanguage string                 `json:"browser_language"`
	ScreenSize      string                 `json:"screen_size"`
	Attributes      map[string]interface{} `json:"attributes"`
	Traits          map[string]interface{} `json:"traits"`
}
