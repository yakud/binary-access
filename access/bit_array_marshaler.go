package access

import (
	"encoding/hex"

	"github.com/klauspost/compress/zstd"

	"github.com/Workiva/go-datastructures/bitarray"
)

// todo refactor

var lz4Encoder, _ = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedBestCompression), zstd.WithZeroFrames(false))
var lz4Decoder, _ = zstd.NewReader(nil, zstd.WithDecoderConcurrency(4))

var bitArrayMarshalerDefault BitArrayMarshaler = &bitArrayLZ4_HEX{}

func BitArrayEncode(arr bitarray.BitArray) ([]byte, error) {
	return bitArrayMarshalerDefault.Encode(arr)
}

func BitArrayDecode(hex []byte) (bitarray.BitArray, error) {
	return bitArrayMarshalerDefault.Decode(hex)
}

func compress(src []byte) []byte {
	return lz4Encoder.EncodeAll(src, make([]byte, 0, len(src)))
}

func decompress(src []byte) ([]byte, error) {
	return lz4Decoder.DecodeAll(src, nil)
}

type BitArrayMarshaler interface {
	Encode(bitarray.BitArray) ([]byte, error)
	Decode([]byte) (bitarray.BitArray, error)
}

type bitArrayLZ4_HEX struct {
}

func (c *bitArrayLZ4_HEX) Encode(array bitarray.BitArray) ([]byte, error) {
	if chainBin, err := bitarray.Marshal(array); err != nil {
		return nil, err
	} else {
		return []byte(hex.EncodeToString(compress(chainBin))), nil // todo efficient hex encoding
	}
}

func (c *bitArrayLZ4_HEX) Decode(hexData []byte) (bitarray.BitArray, error) {
	chainCompressed, err := hex.DecodeString(string(hexData)) // todo efficient hex decoding
	if err != nil {
		return nil, err
	}
	chainBin, err := decompress(chainCompressed)
	if err != nil {
		return nil, err
	}

	chain, err := bitarray.Unmarshal(chainBin)
	if err != nil {
		return nil, err
	}

	return chain, nil
}
