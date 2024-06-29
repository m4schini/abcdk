package cas

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func Hash(data []byte) string {
	sum := sha512.Sum512(data)
	h := hex.EncodeToString(sum[:])
	return fmt.Sprintf("sha512:%v", h)
}
