package serializer

import (
	"reflect"
)

type ISerializer interface {
	SetToOut(data interface{})
	GetOut() interface{}
	GetType() interface{}
	Serialize(data interface{}) string
	Deserialize(data string) error
	DeserializeType(data string) error
}

type ISerialize interface {
	Serialize(data interface{}) string
	Deserialize(data string, out interface{}) error
	DeserializeType(data string, t interface{}) (interface{}, error)
}

type Options struct {
	Serializer ISerialize
	Out        interface{}
	Type       interface{}
}

type Serializer struct {
	ISerializer
	Options *Options
}

func (s Serializer) GetOut() interface{} {
	return s.Options.Out
}
func (s Serializer) SetToOut(data interface{}) {
	s.Options.Out = data
}

func (s Serializer) GetType() interface{} {
	return s.Options.Type
}

func (s Serializer) Serialize(data interface{}) string {
	return s.Options.Serializer.Serialize(data)
}

func (s Serializer) Deserialize(data string) error {
	return s.Options.Serializer.Deserialize(data, &s.Options.Out)
}

func (s Serializer) DeserializeType(data string) error {
	t := CloneStruct(s.Options.Type)
	deserialize, err := s.Options.Serializer.DeserializeType(data, t)
	s.Options.Out = s.appendToSlice(s.Options.Out, deserialize)
	return err
}

func (s Serializer) appendToSlice(slicePtr interface{}, newItem interface{}) interface{} {
	sliceValue := reflect.ValueOf(slicePtr)

	if sliceValue.Kind() == reflect.Ptr && sliceValue.Elem().Kind() == reflect.Slice {
		slice := sliceValue.Elem()
		newItemValue := reflect.ValueOf(newItem)
		resultSlice := reflect.Append(slice, newItemValue)
		return resultSlice.Interface()
	} else {
		slice := sliceValue
		newItemValue := reflect.ValueOf(newItem)
		resultSlice := reflect.Append(slice, newItemValue)
		return resultSlice.Interface()
	}

	return nil
}

func CloneStruct(src interface{}) interface{} {
	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() != reflect.Ptr || srcValue.IsNil() {
		return nil
	}

	srcElem := srcValue.Elem()
	dest := reflect.New(srcElem.Type()).Interface()
	destValue := reflect.ValueOf(dest).Elem()

	destValue.Set(srcElem)

	return dest
}

func NewSerializer(Options *Options) ISerializer {
	return &Serializer{Options: Options}
}
