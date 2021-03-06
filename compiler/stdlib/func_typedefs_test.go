package stdlib_test

import (
	"errors"
	"testing"

	"github.com/d5/tengo/assert"
	"github.com/d5/tengo/compiler/stdlib"
	"github.com/d5/tengo/objects"
)

func TestFuncAIR(t *testing.T) {
	uf := stdlib.FuncAIR(func(int) {})
	ret, err := uf.Call(&objects.Int{Value: 10})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Undefined{}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncAR(t *testing.T) {
	uf := stdlib.FuncAR(func() {})
	ret, err := uf.Call()
	assert.NoError(t, err)
	assert.Equal(t, &objects.Undefined{}, ret)
	ret, err = uf.Call(objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncARI(t *testing.T) {
	uf := stdlib.FuncARI(func() int { return 10 })
	ret, err := uf.Call()
	assert.NoError(t, err)
	assert.Equal(t, &objects.Int{Value: 10}, ret)
	ret, err = uf.Call(objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncARE(t *testing.T) {
	uf := stdlib.FuncARE(func() error { return nil })
	ret, err := uf.Call()
	assert.NoError(t, err)
	assert.Equal(t, objects.TrueValue, ret)
	uf = stdlib.FuncARE(func() error { return errors.New("some error") })
	ret, err = uf.Call()
	assert.NoError(t, err)
	assert.Equal(t, &objects.Error{Value: &objects.String{Value: "some error"}}, ret)
	ret, err = uf.Call(objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncARIsE(t *testing.T) {
	uf := stdlib.FuncARIsE(func() ([]int, error) { return []int{1, 2, 3}, nil })
	ret, err := uf.Call()
	assert.NoError(t, err)
	assert.Equal(t, array(&objects.Int{Value: 1}, &objects.Int{Value: 2}, &objects.Int{Value: 3}), ret)
	uf = stdlib.FuncARIsE(func() ([]int, error) { return nil, errors.New("some error") })
	ret, err = uf.Call()
	assert.NoError(t, err)
	assert.Equal(t, &objects.Error{Value: &objects.String{Value: "some error"}}, ret)
	ret, err = uf.Call(objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncARS(t *testing.T) {
	uf := stdlib.FuncARS(func() string { return "foo" })
	ret, err := uf.Call()
	assert.NoError(t, err)
	assert.Equal(t, &objects.String{Value: "foo"}, ret)
	ret, err = uf.Call(objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncARSE(t *testing.T) {
	uf := stdlib.FuncARSE(func() (string, error) { return "foo", nil })
	ret, err := uf.Call()
	assert.NoError(t, err)
	assert.Equal(t, &objects.String{Value: "foo"}, ret)
	uf = stdlib.FuncARSE(func() (string, error) { return "", errors.New("some error") })
	ret, err = uf.Call()
	assert.NoError(t, err)
	assert.Equal(t, &objects.Error{Value: &objects.String{Value: "some error"}}, ret)
	ret, err = uf.Call(objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncARSs(t *testing.T) {
	uf := stdlib.FuncARSs(func() []string { return []string{"foo", "bar"} })
	ret, err := uf.Call()
	assert.NoError(t, err)
	assert.Equal(t, array(&objects.String{Value: "foo"}, &objects.String{Value: "bar"}), ret)
	ret, err = uf.Call(objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncASRE(t *testing.T) {
	uf := stdlib.FuncASRE(func(a string) error { return nil })
	ret, err := uf.Call(&objects.String{Value: "foo"})
	assert.NoError(t, err)
	assert.Equal(t, objects.TrueValue, ret)
	uf = stdlib.FuncASRE(func(a string) error { return errors.New("some error") })
	ret, err = uf.Call(&objects.String{Value: "foo"})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Error{Value: &objects.String{Value: "some error"}}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncASRS(t *testing.T) {
	uf := stdlib.FuncASRS(func(a string) string { return a })
	ret, err := uf.Call(&objects.String{Value: "foo"})
	assert.NoError(t, err)
	assert.Equal(t, &objects.String{Value: "foo"}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncASI64RE(t *testing.T) {
	uf := stdlib.FuncASI64RE(func(a string, b int64) error { return nil })
	ret, err := uf.Call(&objects.String{Value: "foo"}, &objects.Int{Value: 5})
	assert.NoError(t, err)
	assert.Equal(t, objects.TrueValue, ret)
	uf = stdlib.FuncASI64RE(func(a string, b int64) error { return errors.New("some error") })
	ret, err = uf.Call(&objects.String{Value: "foo"}, &objects.Int{Value: 5})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Error{Value: &objects.String{Value: "some error"}}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncAIIRE(t *testing.T) {
	uf := stdlib.FuncAIIRE(func(a, b int) error { return nil })
	ret, err := uf.Call(&objects.Int{Value: 5}, &objects.Int{Value: 7})
	assert.NoError(t, err)
	assert.Equal(t, objects.TrueValue, ret)
	uf = stdlib.FuncAIIRE(func(a, b int) error { return errors.New("some error") })
	ret, err = uf.Call(&objects.Int{Value: 5}, &objects.Int{Value: 7})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Error{Value: &objects.String{Value: "some error"}}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncASIIRE(t *testing.T) {
	uf := stdlib.FuncASIIRE(func(a string, b, c int) error { return nil })
	ret, err := uf.Call(&objects.String{Value: "foo"}, &objects.Int{Value: 5}, &objects.Int{Value: 7})
	assert.NoError(t, err)
	assert.Equal(t, objects.TrueValue, ret)
	uf = stdlib.FuncASIIRE(func(a string, b, c int) error { return errors.New("some error") })
	ret, err = uf.Call(&objects.String{Value: "foo"}, &objects.Int{Value: 5}, &objects.Int{Value: 7})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Error{Value: &objects.String{Value: "some error"}}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncASRSE(t *testing.T) {
	uf := stdlib.FuncASRSE(func(a string) (string, error) { return a, nil })
	ret, err := uf.Call(&objects.String{Value: "foo"})
	assert.NoError(t, err)
	assert.Equal(t, &objects.String{Value: "foo"}, ret)
	uf = stdlib.FuncASRSE(func(a string) (string, error) { return a, errors.New("some error") })
	ret, err = uf.Call(&objects.String{Value: "foo"})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Error{Value: &objects.String{Value: "some error"}}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncASSRE(t *testing.T) {

}

func TestFuncARF(t *testing.T) {
	uf := stdlib.FuncARF(func() float64 {
		return 10.0
	})
	ret, err := uf.Call()
	assert.NoError(t, err)
	assert.Equal(t, &objects.Float{Value: 10.0}, ret)
	ret, err = uf.Call(objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncAFRF(t *testing.T) {
	uf := stdlib.FuncAFRF(func(a float64) float64 {
		return a
	})
	ret, err := uf.Call(&objects.Float{Value: 10.0})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Float{Value: 10.0}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
	ret, err = uf.Call(objects.TrueValue, objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncAIRF(t *testing.T) {
	uf := stdlib.FuncAIRF(func(a int) float64 {
		return float64(a)
	})
	ret, err := uf.Call(&objects.Int{Value: 10.0})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Float{Value: 10.0}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
	ret, err = uf.Call(objects.TrueValue, objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncAFRI(t *testing.T) {
	uf := stdlib.FuncAFRI(func(a float64) int {
		return int(a)
	})
	ret, err := uf.Call(&objects.Float{Value: 10.5})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Int{Value: 10}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
	ret, err = uf.Call(objects.TrueValue, objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncAFRB(t *testing.T) {
	uf := stdlib.FuncAFRB(func(a float64) bool {
		return a > 0.0
	})
	ret, err := uf.Call(&objects.Float{Value: 0.1})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Bool{Value: true}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
	ret, err = uf.Call(objects.TrueValue, objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncAFFRF(t *testing.T) {
	uf := stdlib.FuncAFFRF(func(a, b float64) float64 {
		return a + b
	})
	ret, err := uf.Call(&objects.Float{Value: 10.0}, &objects.Float{Value: 20.0})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Float{Value: 30.0}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
	ret, err = uf.Call(objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncAIFRF(t *testing.T) {
	uf := stdlib.FuncAIFRF(func(a int, b float64) float64 {
		return float64(a) + b
	})
	ret, err := uf.Call(&objects.Int{Value: 10}, &objects.Float{Value: 20.0})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Float{Value: 30.0}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
	ret, err = uf.Call(objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncAFIRF(t *testing.T) {
	uf := stdlib.FuncAFIRF(func(a float64, b int) float64 {
		return a + float64(b)
	})
	ret, err := uf.Call(&objects.Float{Value: 10.0}, &objects.Int{Value: 20})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Float{Value: 30.0}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
	ret, err = uf.Call(objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncAFIRB(t *testing.T) {
	uf := stdlib.FuncAFIRB(func(a float64, b int) bool {
		return a < float64(b)
	})
	ret, err := uf.Call(&objects.Float{Value: 10.0}, &objects.Int{Value: 20})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Bool{Value: true}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
	ret, err = uf.Call(objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncAIRSsE(t *testing.T) {
	uf := stdlib.FuncAIRSsE(func(a int) ([]string, error) {
		return []string{"foo", "bar"}, nil
	})
	ret, err := uf.Call(&objects.Int{Value: 10})
	assert.NoError(t, err)
	assert.Equal(t, array(&objects.String{Value: "foo"}, &objects.String{Value: "bar"}), ret)
	uf = stdlib.FuncAIRSsE(func(a int) ([]string, error) {
		return nil, errors.New("some error")
	})
	ret, err = uf.Call(&objects.Int{Value: 10})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Error{Value: &objects.String{Value: "some error"}}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncARB(t *testing.T) {
	uf := stdlib.FuncARB(func() bool { return true })
	ret, err := uf.Call()
	assert.NoError(t, err)
	assert.Equal(t, objects.TrueValue, ret)
	ret, err = uf.Call(objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncARYE(t *testing.T) {
	uf := stdlib.FuncARYE(func() ([]byte, error) {
		return []byte("foo bar"), nil
	})
	ret, err := uf.Call()
	assert.NoError(t, err)
	assert.Equal(t, &objects.Bytes{Value: []byte("foo bar")}, ret)
	uf = stdlib.FuncARYE(func() ([]byte, error) {
		return nil, errors.New("some error")
	})
	ret, err = uf.Call()
	assert.NoError(t, err)
	assert.Equal(t, &objects.Error{Value: &objects.String{Value: "some error"}}, ret)
	ret, err = uf.Call(objects.TrueValue)
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncASRIE(t *testing.T) {
	uf := stdlib.FuncASRIE(func(a string) (int, error) { return 5, nil })
	ret, err := uf.Call(&objects.String{Value: "foo"})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Int{Value: 5}, ret)
	uf = stdlib.FuncASRIE(func(a string) (int, error) { return 0, errors.New("some error") })
	ret, err = uf.Call(&objects.String{Value: "foo"})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Error{Value: &objects.String{Value: "some error"}}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func TestFuncAYRIE(t *testing.T) {
	uf := stdlib.FuncAYRIE(func(a []byte) (int, error) { return 5, nil })
	ret, err := uf.Call(&objects.Bytes{Value: []byte("foo")})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Int{Value: 5}, ret)
	uf = stdlib.FuncAYRIE(func(a []byte) (int, error) { return 0, errors.New("some error") })
	ret, err = uf.Call(&objects.Bytes{Value: []byte("foo")})
	assert.NoError(t, err)
	assert.Equal(t, &objects.Error{Value: &objects.String{Value: "some error"}}, ret)
	ret, err = uf.Call()
	assert.Equal(t, objects.ErrWrongNumArguments, err)
}

func array(elements ...objects.Object) *objects.Array {
	return &objects.Array{Value: elements}
}
