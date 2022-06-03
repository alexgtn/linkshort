package encoding

import "math/big"

func ToBase62(str string) string {
	var i big.Int

	i.SetBytes([]byte(str))

	return i.Text(62)
}
