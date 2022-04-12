package endec

import (
	"fmt"
	"github.com/ipfs/go-cid"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"reflect"
)

var (
	tCid = reflect.TypeOf(cid.Cid{})
)

func cidEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tCid {
		return bsoncodec.ValueEncoderError{Name: "cidEncodeValue", Types: []reflect.Type{tCid}, Received: val}
	}
	b := val.Interface().(cid.Cid)
	return vw.WriteString(b.String())
}

func cidDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tCid {
		return bsoncodec.ValueDecoderError{Name: "cidDecodeValue", Types: []reflect.Type{tCid}, Received: val}
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
		return fmt.Errorf("cannot decode %v into a Cid", vrType)
	}

	if err != nil {
		return err
	}
	_cid, err := cid.Parse(data)
	if err != nil {
		return err
	}
	val.Set(reflect.ValueOf(_cid))
	return nil
}
