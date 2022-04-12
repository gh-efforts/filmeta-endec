package endec

import (
	"fmt"
	"reflect"

	"github.com/filecoin-project/go-address"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

var (
	tAddress = reflect.TypeOf(address.Address{})
)

func addressEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tAddress {
		return bsoncodec.ValueEncoderError{Name: "addressEncodeValue", Types: []reflect.Type{tAddress}, Received: val}
	}
	b := val.Interface().(address.Address)
	return vw.WriteString(b.String())
}

func addressDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tAddress {
		return bsoncodec.ValueDecoderError{Name: "addressDecodeValue", Types: []reflect.Type{tAddress}, Received: val}
	}
	var data string
	var err error
	switch vrType := vr.Type(); vrType {
	case bsontype.String:
		data, err = vr.ReadString()
	case bsontype.Null:
		err = vr.ReadNull()
	case bsontype.Undefined:
		err = vr.ReadUndefined()
	default:
		return fmt.Errorf("cannot decode %v into a Address", vrType)
	}

	if err != nil {
		return err
	}
	addr, err := address.NewFromString(data)
	if err != nil {
		return err
	}
	val.Set(reflect.ValueOf(addr))
	return nil
}
