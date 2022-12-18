package models

type GtTranslateRes struct {
	Data data `json:"data"`
}

type data struct {
	Translations []translation `json:"translations"`
}

type translation struct {
	TranslatedText string `json:"translatedText"`
}
