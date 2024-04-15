package conf_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/carlmjohnson/be"
	"github.com/lucafmarques/conf"
	"github.com/lucafmarques/env"
)

var TestURL, _ = url.Parse("https://github.com/lucafmarques")

const ENV_KEY = "ENV_KEY"

type test[T any] struct {
	env  string
	want T
	err  error

	cfg struct {
		V T `env:"ENV_KEY"`
	}
}

func (t test[T]) Error() string {
	return fmt.Errorf("%w: %s", t.err, ENV_KEY).Error()
}

var TestBuildTable = map[string]func(t *testing.T){
	"NOT_STRUCT_REFERENCE": func(t *testing.T) {
		tc := test[int]{
			env: "10",
			err: conf.ErrNotStructReference,
		}

		err := conf.Build(tc.cfg)

		be.Equal(t, tc.err.Error(), err.Error())
	},
	"NOT_STRUCT": func(t *testing.T) {
		tc := test[int]{
			env: "10",
			err: conf.ErrNotStructReference,
		}

		err := conf.Build(&tc.env)

		be.Equal(t, tc.err.Error(), err.Error())
	},
	"UNSETABLE": func(t *testing.T) {
		type cfg struct {
			i int `env:"ENV_KEY"`
		}

		tc := struct {
			t   test[int]
			cfg cfg
		}{
			t: test[int]{
				env: "10",
				err: conf.ErrCantSetField,
			},
			cfg: cfg{},
		}

		t.Setenv(ENV_KEY, tc.t.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.t.Error(), err.Error())
	},
	"NO_KEY": func(t *testing.T) {
		type cfg struct {
			I int
		}

		tc := struct {
			t   test[int]
			cfg cfg
		}{
			t: test[int]{
				env:  "10",
				want: 0,
			},
			cfg: cfg{},
		}

		t.Setenv(ENV_KEY, tc.t.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.t.want, tc.cfg.I)
		be.NilErr(t, err)
	},
	"STRUCT_WITHOUT_KEY": func(t *testing.T) {
		type s struct {
			V int `env:"ENV_KEY"`
		}
		type cfg struct {
			V s
		}

		tc := struct {
			t   test[int]
			cfg cfg
		}{
			t: test[int]{
				env:  "10",
				want: 10,
			},
			cfg: cfg{},
		}

		t.Setenv(ENV_KEY, tc.t.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.t.want, tc.cfg.V.V)
		be.NilErr(t, err)
	},
	"BOOL": func(t *testing.T) {
		tc := test[bool]{
			env:  "true",
			want: true,
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"STRING": func(t *testing.T) {
		tc := test[string]{
			env:  "text",
			want: "text",
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"INT": func(t *testing.T) {
		tc := test[int]{
			env:  "10",
			want: 10,
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"INT8": func(t *testing.T) {
		tc := test[int8]{
			env:  "10",
			want: 10,
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"INT16": func(t *testing.T) {
		tc := test[int16]{
			env:  "10",
			want: 10,
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"INT32": func(t *testing.T) {
		tc := test[int32]{
			env:  "10",
			want: 10,
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"INT64": func(t *testing.T) {
		tc := test[int64]{
			env:  "10",
			want: 10,
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"UINT": func(t *testing.T) {
		tc := test[uint]{
			env:  "10",
			want: 10,
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"UINT8": func(t *testing.T) {
		tc := test[uint8]{
			env:  "10",
			want: 10,
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"UINT16": func(t *testing.T) {
		tc := test[uint16]{
			env:  "10",
			want: 10,
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"UINT32": func(t *testing.T) {
		tc := test[uint32]{
			env:  "10",
			want: 10,
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"UINT64": func(t *testing.T) {
		tc := test[uint64]{
			env:  "10",
			want: 10,
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"FLOAT32": func(t *testing.T) {
		tc := test[float32]{
			env:  "10",
			want: 10.,
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"FLOAT64": func(t *testing.T) {
		tc := test[float64]{
			env:  "10",
			want: 10.,
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"COMPLEX64": func(t *testing.T) {
		tc := test[complex64]{
			env:  "420+69i",
			want: complex64(complex(420, 69)),
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"COMPLEX128": func(t *testing.T) {
		tc := test[complex128]{
			env:  "420+69i",
			want: complex(420, 69),
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"CUSTOM_TEXT_UNMARSHALER": func(t *testing.T) {
		url, _ := url.Parse("https://github.com/lucafmarques")
		tc := test[user]{
			env:  `{"name":"Luca Marques","github":"https://github.com/lucafmarques"}`,
			want: user{Name: "Luca Marques", Github: *url},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"CUSTOM_TEXT_UNMARSHALER_REFERENCE": func(t *testing.T) {
		url, _ := url.Parse("https://github.com/lucafmarques")
		tc := test[*user]{
			env:  `{"name":"Luca Marques","github":"https://github.com/lucafmarques"}`,
			want: &user{Name: "Luca Marques", Github: *url},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, *tc.want, *tc.cfg.V)
		be.NilErr(t, err)
	},
	"CUSTOM_TEXT_UNMARSHALER_UNSET": func(t *testing.T) {
		tc := test[user]{
			env: `{"name":"Luca Marques","github":"https://github.com/lucafmarques"}`,
			err: env.ErrUnset,
		}

		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.Error(), err.Error())
	},
	"UNIMPLEMENTED_TEXT_UNMARSHALER": func(t *testing.T) {
		tc := test[fail]{
			env:  `{"name":"Luca Marques","github":"https://github.com/lucafmarques"}`,
			want: fail{},
			err:  conf.ErrInvalidType,
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.Equal(t, tc.want, tc.cfg.V)
		be.Equal(t, tc.Error(), err.Error())
	},
	"BOOL_SLICE": func(t *testing.T) {
		tc := test[[]bool]{
			env:  "true,false",
			want: []bool{true, false},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.AllEqual(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"STRING_SLICE": func(t *testing.T) {
		tc := test[[]string]{
			env:  "split,text",
			want: []string{"split", "text"},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.AllEqual(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"INT_SLICE": func(t *testing.T) {
		tc := test[[]int]{
			env:  "10,20",
			want: []int{10, 20},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.AllEqual(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"INT8_SLICE": func(t *testing.T) {
		tc := test[[]int8]{
			env:  "10,20",
			want: []int8{10, 20},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.AllEqual(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"INT16_SLICE": func(t *testing.T) {
		tc := test[[]int16]{
			env:  "10,20",
			want: []int16{10, 20},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.AllEqual(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"INT32_SLICE": func(t *testing.T) {
		tc := test[[]int32]{
			env:  "10,20",
			want: []int32{10, 20},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.AllEqual(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"INT64_SLICE": func(t *testing.T) {
		tc := test[[]int64]{
			env:  "10,20",
			want: []int64{10, 20},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.AllEqual(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"UINT_SLICE": func(t *testing.T) {
		tc := test[[]uint]{
			env:  "10,20",
			want: []uint{10, 20},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.AllEqual(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"UINT8_SLICE": func(t *testing.T) {
		tc := test[[]uint8]{
			env:  "10,20",
			want: []uint8{10, 20},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.AllEqual(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"UINT16_SLICE": func(t *testing.T) {
		tc := test[[]uint16]{
			env:  "10,20",
			want: []uint16{10, 20},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.AllEqual(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"UINT32_SLICE": func(t *testing.T) {
		tc := test[[]uint32]{
			env:  "10,20",
			want: []uint32{10, 20},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.AllEqual(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"UINT64_SLICE": func(t *testing.T) {
		tc := test[[]uint64]{
			env:  "10,20",
			want: []uint64{10, 20},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.AllEqual(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"FLOAT32_SLICE": func(t *testing.T) {
		tc := test[[]float32]{
			env:  "420.0,69.",
			want: []float32{420., 69.},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.AllEqual(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"FLOAT64_SLICE": func(t *testing.T) {
		tc := test[[]float64]{
			env:  "420.0,69.",
			want: []float64{420., 69.},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.AllEqual(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"COMPLEX64_SLICE": func(t *testing.T) {
		tc := test[[]complex64]{
			env:  "420+69i,69+420i",
			want: []complex64{complex64(complex(420, 69)), complex64(complex(69, 420))},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.AllEqual(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
	"COMPLEX128_SLICE": func(t *testing.T) {
		tc := test[[]complex128]{
			env:  "420+69i,69+420i",
			want: []complex128{complex(420, 69), complex(69, 420)},
		}

		t.Setenv(ENV_KEY, tc.env)
		err := conf.Build(&tc.cfg)

		be.AllEqual(t, tc.want, tc.cfg.V)
		be.NilErr(t, err)
	},
}

func TestBuild(t *testing.T) {
	for test, f := range TestBuildTable {
		t.Run(test, f)
	}
}

type fail struct{}

type user struct {
	Name   string  `json:"name"`
	Github url.URL `json:"github"`
	error  error
}

func (u user) MarshalText() ([]byte, error) {
	if u.error != nil {
		return []byte{}, u.error
	}
	return []byte(fmt.Sprintf("Name: %s - GitHub: %s", u.Name, u.Github.String())), nil
}

func (u *user) UnmarshalText(data []byte) error {
	tmp := struct {
		Github string `json:"github"`
		Name   string `json:"name"`
	}{}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	url, err := url.Parse(tmp.Github)
	if err != nil {
		return err
	}

	u.Github = *url
	u.Name = tmp.Name

	return nil
}
