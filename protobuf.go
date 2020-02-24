package decimal

import (
	"github.com/golang/protobuf/proto"
	v1 "github.com/tensoralpha/decimal/protobuf/decimal/v1"
)

// Add support for gogoprotobuf custom type

// Must implement the following

// func (t T) Marshal() ([]byte, error) {}
// func (t *T) MarshalTo(data []byte) (n int, err error) {}
// func (t *T) Unmarshal(data []byte) error {}
// func (t *T) Size() int {}

// func (t T) MarshalJSON() ([]byte, error) {}
// func (t *T) UnmarshalJSON(data []byte) error {}

// // only required if the compare option is set
// func (t T) Compare(other T) int {}
// // only required if the equal option is set
// func (t T) Equal(other T) bool {}
// // only required if populate option is set
// func NewPopulatedT(r randyThetest) *T {}

func (x *Big) toProto() *v1.Decimal {
	form, negative, value, exp := x.Decompose(nil)
	if form != byte(finite) {
		panic("must be finite to be represented as protobuf")
	}
	return &v1.Decimal{
		Value:    value,
		Scale:    -exp,
		Negative: negative,
	}
}

func (a Big) Equal(b Big) bool {
	return a.toProto().Equal(b.toProto())
}

func (x *Big) Size() int {
	return x.toProto().Size()
}

func (x Big) Marshal() ([]byte, error) {
	return proto.Marshal(x.toProto())
}

func (x *Big) MarshalTo(data []byte) (int, error) {
	m := x.toProto()
	size := m.Size()
	return m.MarshalToSizedBuffer(data[:size])
}

func (x *Big) Unmarshal(data []byte) error {
	m := &v1.Decimal{}
	if err := proto.Unmarshal(data, m); err != nil {
		return err
	}
	if err := x.Compose(byte(finite), m.Negative, m.Value, -m.Scale); err != nil {
		return err
	}
	return nil
}
