package aql

import (
	"encoding/json"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

var d JSONDecoder

//MongoParser -
type MongoParser struct {
}

//Parse  -
func (mp MongoParser) Parse(f string) interface{} {
	initMap()
	sqlMap := bson.M{}
	andAr := []bson.M{}
	r := d.Decode(f)

	for k, v := range r {

		if len(v) == 1 {
			cri1 := v[0]
			cond1 := bson.M{mongoOpMap[cri1.Operator]: handleInt64(cri1.Value)}
			andAr = append(andAr, bson.M{k: cond1})
		} else if len(v) > 1 {
			// add or conditions
			orAr := []bson.M{}
			for _, vc := range v {
				cond2 := bson.M{mongoOpMap[vc.Operator]: handleInt64(vc.Value)}
				orAr = append(orAr, bson.M{k: cond2})
			}
			andAr = append(andAr, bson.M{"$or": orAr})
		}

	}
	sqlMap["$and"] = andAr

	return sqlMap
}

//Sort -
func (mp MongoParser) Sort(f string) interface{} {
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

func handleInt64(v interface{}) interface{} {

	switch v.(type) {
	case string:

		n, err := strconv.ParseInt(v.(string), 10, 64)
		if err == nil {

			if len(v.(string)) == 19 {
				return n
			}
		}
		break
	case []interface{}:
		newArr := []interface{}{}
		for _, vv := range v.([]interface{}) {
			newArr = append(newArr, handleInt64(vv))
		}
		v = newArr
	}
	return v
}
