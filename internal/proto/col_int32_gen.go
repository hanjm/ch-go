// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"github.com/go-faster/errors"
)

// ColInt32 represents Int32 column.
type ColInt32 []int32

// Compile-time assertions for ColInt32.
var (
	_ Input  = ColInt32{}
	_ Result = (*ColInt32)(nil)
)

// Type returns ColumnType of Int32.
func (ColInt32) Type() ColumnType {
	return ColumnTypeInt32
}

// Rows returns count of rows in column.
func (c ColInt32) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColInt32) Reset() {
	*c = (*c)[:0]
}

// EncodeColumn encodes Int32 rows to *Buffer.
func (c ColInt32) EncodeColumn(b *Buffer) {
	const size = 32 / 8
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	var offset int
	for _, v := range c {
		bin.PutUint32(
			b.Buf[offset:offset+size],
			uint32(v),
		)
		offset += size
	}
}

// DecodeColumn decodes Int32 rows from *Reader.
func (c *ColInt32) DecodeColumn(r *Reader, rows int) error {
	const size = 32 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	for i := 0; i < len(data); i += size {
		v = append(v,
			int32(bin.Uint32(data[i:i+size])),
		)
	}
	*c = v
	return nil
}
