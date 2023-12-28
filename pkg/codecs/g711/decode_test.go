package g711

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeMuLaw(t *testing.T) {
	require.Equal(t,
		DecodeMulaw([]byte{1, 2, 3, 255, 254, 253}),
		[]byte{
			0x86, 0x84, 0x8a, 0x84, 0x8e, 0x84, 0x00, 0x00,
			0x00, 0x08, 0x00, 0x10,
		},
	)
}

func TestDecodeALaw(t *testing.T) {
	require.Equal(t,
		DecodeAlaw([]byte{1, 2, 3, 255, 254, 253}),
		[]byte{
			0xeb, 0x80, 0xe8, 0x80, 0xe9, 0x80, 0x03, 0x50,
			0x03, 0x70, 0x03, 0x10,
		},
	)
}