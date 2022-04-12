package endec

import "go.mongodb.org/mongo-driver/bson/bsoncodec"

func BuildDefaultRegistry() *bsoncodec.Registry {
	rb := bsoncodec.NewRegistryBuilder().
		RegisterTypeEncoder(tAddress, bsoncodec.ValueEncoderFunc(addressEncodeValue)).
		RegisterTypeDecoder(tAddress, bsoncodec.ValueDecoderFunc(addressDecodeValue)).
		RegisterTypeEncoder(tCid, bsoncodec.ValueEncoderFunc(cidEncodeValue)).
		RegisterTypeDecoder(tCid, bsoncodec.ValueDecoderFunc(cidDecodeValue)).
		RegisterTypeEncoder(tAmount, bsoncodec.ValueEncoderFunc(amountEncodeValue)).
		RegisterTypeDecoder(tAmount, bsoncodec.ValueDecoderFunc(amountDecodeValue)).
		RegisterTypeEncoder(tStateBigint, bsoncodec.ValueEncoderFunc(stateBigintEncodeValue)).
		RegisterTypeDecoder(tStateBigint, bsoncodec.ValueDecoderFunc(stateBigintDecodeValue))

	bsoncodec.DefaultValueEncoders{}.RegisterDefaultEncoders(rb)
	bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(rb)
	return rb.Build()
}
