package av1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var casesLEB128 = []struct {
	name string
	dec  uint
	enc  []byte
}{
	{
		"a",
		1234567,
		[]byte{0x87, 0xad, 0x4b},
	},
	{
		"b",
		127,
		[]byte{0x7f},
	},
	{
		"c",
		87651321342,
		[]byte{0xfe, 0x8b, 0xb4, 0xc3, 0xc6, 0x2},
	},
}

func TestLEB128Unmarshal(t *testing.T) {
	for _, ca := range casesLEB128 {
		t.Run(ca.name, func(t *testing.T) {
			dec, n, err := LEB128Unmarshal(ca.enc)
			require.NoError(t, err)
			require.Equal(t, len(ca.enc), n)
			require.Equal(t, ca.dec, dec)
		})
	}
}

func TestLEB128Marshal(t *testing.T) {
	for _, ca := range casesLEB128 {
		t.Run(ca.name, func(t *testing.T) {
			enc := LEB128Marshal(ca.dec)
			require.Equal(t, ca.enc, enc)
		})
	}
}

func FuzzLEB128Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		LEB128Unmarshal(b)
	})
}