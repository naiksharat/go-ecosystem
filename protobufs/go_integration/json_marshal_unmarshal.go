package main

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func toJSON(p proto.Message) (string, error) {
	out, err := protojson.Marshal(p)
	if err != nil {
		return "", err
	}
	//fmt.Println(out)
	return string(out), nil
}

func fromJSON(json string, p proto.Message) (proto.Message, error) {
	if err := protojson.Unmarshal([]byte(json), p); err != nil {
		return nil, err
	}

	return p, nil
}
