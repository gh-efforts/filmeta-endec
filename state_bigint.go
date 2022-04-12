package endec

import (
	"fmt"
	stateBig "github.com/filecoin-project/go-state-types/big"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/big"
	"reflect"
)

var (
	tStateBigint = reflect.TypeOf(stateBig.Int{})
)

func stateBigintEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tStateBigint {
		return bsoncodec.ValueEncoderError{Name: "StateBigintEncodeValue", Types: []reflect.Type{tStateBigint}, Received: val}
	}
	b := val.Interface().(stateBig.Int)
	var v primitive.Decimal128
	if b.Nil() {
		v, _ = primitive.ParseDecimal128("0")
	} else {
		v, _ = primitive.ParseDecimal128FromBigInt(b.Int, 0)
	}
	return vw.WriteDecimal128(v)
}

func stateBigintDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tStateBigint {
		return bsoncodec.ValueDecoderError{Name: "StateBigintDecodeValue", Types: []reflect.Type{tStateBigint}, Received: val}
	}
	var data primitive.Decimal128
	var err error
	switch vrType := vr.Type(); vrType {
	case bsontype.Decimal128:
		data, err = vr.ReadDecimal128()
	case bsontype.Null:
		err = vr.ReadNull()
	case bsontype.Undefined:
		err = vr.ReadUndefined()
	case bsontype.Int32:
		v, err := vr.ReadInt32()
		if err != nil {
			val.Set(reflect.ValueOf(stateBig.NewInt(int64(v))))
			return nil
		}

	default:
		return fmt.Errorf("cannot decode %v into a State Bigint", vrType)
	}

	if err != nil {
		return err
	}
	b, exp, err := data.BigInt()
	if err != nil {
		return err
	}
	if exp != 0 {
		b.Div(b, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(-exp)), nil))
	}
	val.Set(reflect.ValueOf(stateBig.NewFromGo(b)))
	return nil
}
