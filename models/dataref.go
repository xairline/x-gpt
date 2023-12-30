package models

type Dataref struct {
	Name         string `yaml:"name"`
	Value        string `yaml:"value"`
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
type SendCommandReq struct {
	Command string `json:"command" `
}
type DatarefValues map[string]DatarefValue
