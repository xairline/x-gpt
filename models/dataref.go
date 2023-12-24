package models

type Dataref struct {
	Name         string `yaml:"name"`
	DatarefStr   string `yaml:"value"`
	Precision    int8   `yaml:"precision,omitempty"`
	IsBytesArray bool   `yaml:"isBytesArray,omitempty"`
}
type DatarefValue struct {
	Name  string      `json:"name" `
	Value interface{} `json:"value" `
}
type SetDatarefValue struct {
	Dataref string      `json:"dataref" `
	Value   interface{} `json:"value" `
}
type SetDatarefValueReq struct {
	Request SetDatarefValue `json:"request" `
}
type DatarefValues map[string]DatarefValue
