package serializer

import (
	"encoding/json"
)

type jsonSerializer struct {
	ISerialize
}

func NewJsonSerializer() ISerialize {
	return &jsonSerializer{}
}

func (s *jsonSerializer) Serialize(data interface{}) string {
	marshal, _ := json.Marshal(data)
	return string(marshal)
}

func (s *jsonSerializer) Deserialize(data string, out interface{}) error {
	return json.Unmarshal([]byte(data), &out)
}

func (s *jsonSerializer) DeserializeType(data string, t interface{}) (interface{}, error) {
	err := json.Unmarshal([]byte(data), t)
	return t, err
}
