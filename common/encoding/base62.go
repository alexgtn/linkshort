package encoding

import "math/big"

func ToBase62(in []byte) string {
	var i big.Int

	i.SetBytes(in)

	return i.Text(62)
}
