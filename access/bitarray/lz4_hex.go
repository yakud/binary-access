package bitarray

import (
	"encoding/hex"

	"github.com/Workiva/go-datastructures/bitarray"
	"github.com/klauspost/compress/zstd"
)

var lz4Encoder, _ = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedBestCompression), zstd.WithZeroFrames(false))
var lz4Decoder, _ = zstd.NewReader(nil, zstd.WithDecoderConcurrency(4))

func compress(src []byte) []byte {
	//return src
	return lz4Encoder.EncodeAll(src, make([]byte, 0, len(src)))
}

func decompress(src []byte) ([]byte, error) {
	//return src, nil
	return lz4Decoder.DecodeAll(src, nil)
}

type bitArrayLZ4_HEX struct {
	bitarray.BitArray
}

func (c *bitArrayLZ4_HEX) MarshalJSON() ([]byte, error) {
	if chainBin, err := bitarray.Marshal(c.BitArray); err != nil {
		return nil, err
	} else {
		return []byte(`"` + hex.EncodeToString(compress(chainBin)) + `"`), nil // todo efficient hex encoding
	}
}

func (c *bitArrayLZ4_HEX) UnmarshalJSON(data []byte) error {
	chainCompressed, err := hex.DecodeString(string(data[1 : len(data)-1])) // todo efficient hex decoding
	if err != nil {
		return err
	}
	chainBin, err := decompress(chainCompressed)
	if err != nil {
		return err
	}

	chain, err := bitarray.Unmarshal(chainBin)
	if err != nil {
		return err
	}
	c.BitArray = chain

	return nil
}

func NewBitArrayLZ4_HEX(ba bitarray.BitArray) Marshaler {
	return &bitArrayLZ4_HEX{BitArray: ba}
}
