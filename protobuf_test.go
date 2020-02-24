package decimal

import (
	"fmt"
	"testing"
)

func TestProtobuf(t *testing.T) {
	x := &Big{}
	x.SetString("12312312312123987239871230897234098237450239847523453.1")
	fmt.Println(x)
	fmt.Println(x.toProto())
}

var CONST = "12312312312123987239871230897234098237450239847523453.1"

var DATA = func() []byte {
	x := &Big{}
	x.SetString(CONST)
	data, _ := x.Marshal()
	return data
}()

func TestProtobufMarshal(t *testing.T) {
	x := &Big{}
	x.SetString("123123123123.1")
	fmt.Println(x)
	fmt.Println(x.toProto())

	data, err := x.Marshal()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data)

	y := &Big{}
	if err := y.Unmarshal(data); err != nil {
		t.Error(err)
		return
	}
	fmt.Println(y)

}

func TestProtobufMarshalTo(t *testing.T) {
	x := &Big{}
	x.SetString("120.0")
	fmt.Println(x)
	fmt.Println(x.toProto())

	data := make([]byte, 100)

	n, err := x.MarshalTo(data)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(n, data)
}

func Test2(t *testing.T) {
	fmt.Println(CONST)
	fmt.Println(DATA)
	x := &Big{}
	if err := x.Unmarshal(DATA); err != nil {
		t.Error(err)
		return
	}
	fmt.Println(x)
}

func BenchmarkDecimalFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := &Big{}
		x.SetString(CONST)
	}
}

func BenchmarkUnmarshalProtobuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := &Big{}
		if err := x.Unmarshal(DATA); err != nil {
			panic(err)
			return
		}
	}
}
