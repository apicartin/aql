package aql

import (
	"encoding/json"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestParser(*testing.T) {
	filter := "{\r\n  \"a\": [\r\n    {\r\n      \"o\": \"=\",\r\n      \"v\": \"vinay\"\r\n    },\r\n    {\r\n      \"o\": \"=\",\r\n      \"v\": \"kumar\"\r\n    }\r\n  ],\r\n  \"b\": [\r\n    {\r\n      \"o\": \"!=\",\r\n      \"v\": \"12\"\r\n    }\r\n  ],\r\n  \"c\": [\r\n    {\r\n      \"o\": \"in\",\r\n      \"v\": [\r\n        2,\r\n        4\r\n      ]\r\n    }\r\n  ]\r\n}"
	m := ParseFilterToMongo(filter)
	j, err := json.Marshal(m)
	if err == nil {
		logrus.Infoln(string(j))
	}
}
func TestParserFloat(*testing.T) {
	filter := "{\r\n  \"a\": [\r\n    {\r\n      \"o\": \"=\",\r\n      \"v\": \"vinay\"\r\n    },\r\n    {\r\n      \"o\": \"=\",\r\n      \"v\": \"kumar\"\r\n    }\r\n  ],\r\n  \"b\": [\r\n    {\r\n      \"o\": \"!=\",\r\n      \"v\": \"12\"\r\n    }\r\n  ],\r\n  \"c\": [\r\n    {\r\n      \"o\": \"in\",\r\n      \"v\": [\r\n        2.3,\r\n        4.2\r\n      ]\r\n    }\r\n  ]\r\n}"
	m := ParseFilterToMongo(filter)
	j, err := json.Marshal(m)
	if err == nil {
		logrus.Infoln(string(j))
	}
}

func TestSortToMongo(*testing.T) {
	sort := "{\"a\":\"asc\",\"b\":\"desc\"}"
	m := ParseSortToMongo(sort)
	j, err := json.Marshal(m)
	if err == nil {
		logrus.Infoln(string(j))
	}
}
