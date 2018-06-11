package internal

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenIntSlice will generate a slice of random integers
func GenIntSlice(size int) []int {
	out := make([]int, size)
	for i := 0; i < size; i++ {
		r := rand.Int()
		sign := 1
		if r%2 == 0 {
			sign = -1
		}
		out[i] = rand.Int()*sign
	}

	return out
}

func GenInt8Slice(size int) []int8 {
	out := make([]int8, size)
	for i := 0; i < size; i++ {
		r := rand.Int()
		sign := 1
		if r%2 == 0 {
			sign = -1
		}
		out[i] = int8(rand.Int()*sign)
	}

	return out
}

func GenInt16Slice(size int) []int16 {
	out := make([]int16, size)
	for i := 0; i < size; i++ {
		r := rand.Int()
		sign := 1
		if r%2 == 0 {
			sign = -1
		}
		out[i] = int16(rand.Int()*sign)
	}

	return out
}

func GenInt32Slice(size int) []int32 {
	out := make([]int32, size)
	for i := 0; i < size; i++ {
		r := rand.Int()
		sign := int32(1)
		if r%2 == 0 {
			sign = int32(-1)
		}
		out[i] = rand.Int31()*sign
	}

	return out
}

func GenInt64Slice(size int) []int64 {
	out := make([]int64, size)
	for i := 0; i < size; i++ {
		r := rand.Int()
		sign := int64(1)
		if r%2 == 0 {
			sign = int64(-1)
		}
		out[i] = rand.Int63()*sign
	}

	return out
}

func GenUIntSlice(size int) []uint {
	out := make([]uint, size)
	for i := 0; i < size; i++ {
		out[i] = uint(rand.Uint32())
	}

	return out
}

func GenUInt8Slice(size int) []uint8 {
	out := make([]uint8, size)
	for i := 0; i < size; i++ {
		out[i] = uint8(rand.Uint32())
	}

	return out
}

func GenUInt16Slice(size int) []uint16 {
	out := make([]uint16, size)
	for i := 0; i < size; i++ {
		out[i] = uint16(rand.Uint32())
	}

	return out
}

func GenUInt32Slice(size int) []uint32 {
	out := make([]uint32, size)
	for i := 0; i < size; i++ {
		out[i] = rand.Uint32()
	}

	return out
}

func GenUInt64Slice(size int) []uint64 {
	out := make([]uint64, size)
	for i := 0; i < size; i++ {
		out[i] = rand.Uint64()
	}

	return out
}

func GenFloat32Slice(size int) []float32 {
	out := make([]float32, size)
	for i := 0; i < size; i++ {
		out[i] = rand.Float32()
	}

	return out
}

func GenFloat64Slice(size int) []float64 {
	out := make([]float64, size)
	for i := 0; i < size; i++ {
		out[i] = rand.Float64()
	}

	return out
}

func GenUIntPtrSlice(size int) []uintptr {
	out := make([]uintptr, size)
	for i := 0; i < size; i++ {
		out[i] = uintptr(rand.Uint64())
	}

	return out
}

func GenStringSlice(size int) []string {
	out := make([]string, size)
	charset := `abcdefghijklmnopqrstuvwxyz`
	ll := len(charset)

	for i := 0; i < size; i++ {
		l := rand.Intn(32)
		b := make([]byte, l)
		for k := 0; k < l; k++ {
			b[k] = charset[rand.Intn(ll)]
		}
		out[i] = string(b)
	}

	return out
}
