package serializer

import (
	"encoding/json"
	"github.com/leeqvip/gophp/serialize"
	"github.com/onuragtas/gophp"
	"log"
)

type phpSerializer struct {
	ISerialize
}

func NewPhpSerializer() ISerialize {
	return &phpSerializer{}
}

func (s *phpSerializer) Serialize(data interface{}) string {
	marshal, err := gophp.Serialize(data)
	if err != nil {
		log.Println(err)
	}
	return string(marshal)
}

func (s *phpSerializer) Deserialize(data string, out interface{}) error {
	value, err := serialize.UnMarshal([]byte(data))
	marshal, err := json.Marshal(value)
	err = json.Unmarshal(marshal, &out)
	return err
}

func (s *phpSerializer) DeserializeType(data string, t interface{}) (interface{}, error) {
	value, err := serialize.UnMarshal([]byte(data))
	marshal, err := json.Marshal(value)
	err = json.Unmarshal(marshal, t)
	return t, err
}
