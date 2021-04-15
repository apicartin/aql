package aql

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

var d JSONDecoder

type Criteria struct {
	Operator string      `json:"o"`
	Value    interface{} `json:"v"`
}

type Filter map[string][]Criteria
type Sort map[string]string

func ParseFilterToMongo(f string) bson.M {
	initMap()
	sqlMap := bson.M{}
	andAr := []bson.M{}
	r := d.Decode(f)

	for k, v := range r {
		cri1 := v[0]
		cond1 := bson.M{mongoOpMap[cri1.Operator]: cri1.Value}
		if len(v) == 2 {
			// and operator
			cri2 := v[1]
			cond2 := bson.M{mongoOpMap[cri2.Operator]: cri2.Value}
			andAr = append(andAr, bson.M{k: cond1})
			andAr = append(andAr, bson.M{k: cond2})
		} else if len(v) == 1 {
			andAr = append(andAr, bson.M{k: cond1})
		}
	}
	sqlMap["$and"] = andAr

	return sqlMap
}

func ParseSortToMongo(f string) bson.M {
	result := bson.M{}
	r := make(map[string]string)
	err := json.Unmarshal([]byte(f), &r)
	if err != nil {
		return result
	}
	for k, v := range r {
		if v == "asc" {
			result[k] = 1
		} else if v == "desc" {
			result[k] = -1
		}
	}
	return result
}
