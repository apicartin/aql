package aql

type Criteria struct {
	Operator string      `json:"o"`
	Value    interface{} `json:"v"`
}

type Filter map[string][]Criteria
type Sort map[string]string

type Parser interface {
	Parse(s string) interface{}
	Sort(s string) interface{}
}
