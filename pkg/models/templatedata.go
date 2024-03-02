package models

// TemplateData data to send from handlers to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	//for Forms
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
