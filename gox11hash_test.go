package gox11hash

import (
  "bytes"
  "testing"

  "github.com/blockcypher/gox11hash"
)


func TestX11Hash(t *testing.T) {
  x11h  := gox11hash.Sum(
            []byte("00000000000000000000000000000000000000000000000000000000000000000000000000000000"))
  exp   := [32]byte{84, 183, 251, 186, 152, 73, 8, 109, 107, 58, 125, 219, 129, 97, 234, 243, 241,
                    195, 39, 242, 189, 238, 17, 253, 216, 217, 31, 243, 5, 113, 8, 170}
  if bytes.Compare(x11h, exp[:]) != 0 {
    t.Error("Hashing produced", x11h, "expected", exp)
  }
}
