package activity

type Activity struct {
	Activity      string  `json:"activity"`
	Type          string  `json:"type"`
	Participants  int32   `json:"participants"`
	Price         float32 `json:"price"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
	Accessibility float32 `json:"accessibility"`
}
