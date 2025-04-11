package types

type ID uint64

func (i ID) Uint64() uint64 {
	return uint64(i)
}
