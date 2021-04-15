package aql

import (
	"encoding/json"
	"strconv"

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

		cond1 := bson.M{mongoOpMap[cri1.Operator]: handleInt64(cri1.Value)}
		if len(v) == 2 {
			// and operator
			cri2 := v[1]
			cond2 := bson.M{mongoOpMap[cri2.Operator]: handleInt64(cri2.Value)}
			andAr = append(andAr, bson.M{k: cond1})
			andAr = append(andAr, bson.M{k: cond2})
		} else if len(v) == 1 {
			andAr = append(andAr, bson.M{k: cond1})
		}
	}
	sqlMap["$and"] = andAr

	return sqlMap
}

func handleInt64(v interface{}) interface{} {
	switch v.(type) {
	case string:
		n, err := strconv.ParseInt(v.(string), 10, 64)
		if err == nil {
			if len(v.(string)) == 19 {
				return n
			}
		}
	}
	return v
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
