package endec

import (
	"fmt"
	"github.com/filecoin-project/go-state-types/abi"
	big1 "github.com/filecoin-project/go-state-types/big"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/big"
	"reflect"
)

var (
	tAmount = reflect.TypeOf(abi.TokenAmount{})
)

func amountEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tAmount {
		return bsoncodec.ValueEncoderError{Name: "tokenAmountEncodeValue", Types: []reflect.Type{tAmount}, Received: val}
	}
	b := val.Interface().(abi.TokenAmount)
	var v primitive.Decimal128
	if b.Nil() {
		v, _ = primitive.ParseDecimal128("0")
	} else {
		v, _ = primitive.ParseDecimal128FromBigInt(b.Int, 0)
	}
	return vw.WriteDecimal128(v)
}

func amountDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tAmount {
		return bsoncodec.ValueDecoderError{Name: "tokenAmountDecodeValue", Types: []reflect.Type{tAmount}, Received: val}
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
	default:
		return fmt.Errorf("cannot decode %v into a token amount", vrType)
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
	val.Set(reflect.ValueOf(big1.NewFromGo(b)))
	return nil
}
