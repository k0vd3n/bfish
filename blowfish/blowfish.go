package blowfish

import (
	"bfish/src"
	"encoding/binary"
)

var N int = 16

var Key = []byte("This is a crypto blowfish 448 bits key and 64 bits text!")

type Blowfish struct {
	P [16 + 2]uint32
	S [4][256]uint32
}

// New
func New(key []byte) *Blowfish {
	bf := &Blowfish{}

	keyLen := len(key)

	for i := 0; i < 4; i++ {
		for j := 0; j < 256; j++ {
			bf.S[i][j] = src.ORIG_S[i][j]
		}
	}

	k := 0
	for i := 0; i < (N + 2); i++ {
		data := uint32(0)
		for j := 0; j < 4; j++ {
			data = (data << 8) | uint32(key[k])
			k += 1
			if k >= keyLen {
				k = 0
			}
		}
		bf.P[i] = src.ORIG_P[i] ^ data
	}

	datal := uint32(0)
	datar := uint32(0)

	for i := 0; i < (N + 2); i += 2 {
		bf.Encrypt(&datal, &datar)
		bf.P[i] = datal
		bf.P[i+1] = datar
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 256; j += 2 {
			bf.Encrypt(&datal, &datar)
			bf.S[i][j] = datal
			bf.S[i][j+1] = datar
		}
	}
	return bf
}

// Encrypt
func (bf *Blowfish) Encrypt(xl, xr *uint32) {
	Xl := *xl
	Xr := *xr

	for i := 0; i < N; i++ {
		Xl = Xl ^ bf.P[i]
		Xr = bf.f(Xl) ^ Xr
		Xl, Xr = Xr, Xl
	}

	Xl, Xr = Xr, Xl
	Xr = Xr ^ bf.P[N]
	Xl = Xl ^ bf.P[N+1]
	*xl = Xl
	*xr = Xr
}

// Encrypt full string
func EncryptLoop(str string, bf Blowfish) []byte {
	bfstruct := bf
	var tmp uint32
	var xl, xr uint32
	var length int
	var result []byte
	b := make([]byte, 4)
	if len(str)%8 == 0 {
		length = len(str)
	} else {
		length += (8 - (len(str) % 8))
	}
	k := 0
	for i := 0; i < int(length); i += 8 {
		split64bitsTo32bits(join8bitsto64bits(str, &k), &xl, &xr)
		tmp = xl
		xl = xr
		xr = tmp
		bfstruct.Encrypt(&xl, &xr)
		binary.BigEndian.PutUint32(b, xl)
		result = append(result, b...)
		binary.BigEndian.PutUint32(b, xr)
		result = append(result, b...)
	}
	return result
}

// Decrypt
func (bf *Blowfish) Decrypt(xl, xr *uint32) {
	Xl := *xl
	Xr := *xr

	for i := (N + 1); i > 1; i-- {
		Xl = Xl ^ bf.P[i]
		Xr = bf.f(Xl) ^ Xr
		Xl, Xr = Xr, Xl
	}

	Xl, Xr = Xr, Xl
	Xr = Xr ^ bf.P[1]
	Xl = Xl ^ bf.P[0]
	*xl = Xl
	*xr = Xr
}

// f - main function
func (bf *Blowfish) f(x uint32) uint32 {
	d := uint16(x & 0xff)
	x >>= 8
	c := uint16(x & 0xff)
	x >>= 8
	b := uint16(x & 0xff)
	x >>= 8
	a := uint16(x & 0xff)

	y := bf.S[0][a] + bf.S[1][b]
	y = y ^ bf.S[2][c]
	y = y + bf.S[3][d]
	return y
}

// splits a 64-bit block into 2 32-bit blocks
func split64bitsTo32bits(block64b uint64, xl *uint32, xr *uint32) {
	// bytestr := binary.BigEndian.Uint64([]byte(str))
	*xr = uint32(block64b >> 32)
	*xl = uint32(block64b)
}

// puts the following 8 bytes of the input string
// in a 64 - bit block in a cycle of 8 iterations
// of 8 bits per iteration
func join8bitsto64bits(str string, k *int) uint64 {
	kindex := *k
	/*
		stores the number of the last
		byte undivided into blocks of a string
	*/
	bytestr := []byte(str)
	/*
		converts a string into a set of bytes,
		one character in Latin is 1 byte,
		1 character in Cyrillic is 2 bytes
	*/
	block64b := uint64(0)
	/*
		initialize the block in 64 bits,
		every 8 bytes of the string
		will be entered into it
	*/
	for p := kindex; kindex < p+8; kindex++ {
		if kindex >= len(bytestr) {
			block64b <<= 8
		} else {
			block64b = (block64b << 8) | uint64(bytestr[kindex])
		}
	}
	*k = kindex // stores the index of the string bit
	return block64b
}
