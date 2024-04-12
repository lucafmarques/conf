package conf

import (
	"encoding"
	"errors"
	"fmt"
	"reflect"

	"github.com/lucafmarques/env"
)

var (
	ErrInvalidType      = errors.New("invalid type")
	ErrCantSetField     = errors.New("struct field can't be set")
	ErrNotStructPointer = errors.New("cfg argument must be pointer to struct")
)

func Build(cfg any) error { return validate(cfg) }

func validate(v any) error {
	ptrRef := reflect.ValueOf(v)
	if ptrRef.Kind() != reflect.Pointer {
		return ErrNotStructPointer
	}
	ref := ptrRef.Elem()
	if ref.Kind() != reflect.Struct {
		return ErrNotStructPointer
	}

	return parse(ref)
}

func parse(ref reflect.Value) error {
	refT := ref.Type()
	errs := make([]error, 0, refT.NumField())

	for i := range refT.NumField() {
		refF := ref.Field(i)
		refFT := refT.Field(i)
		key := refFT.Tag.Get("env")

		if !refF.CanSet() {
			errs = append(errs, fmt.Errorf("%w: %s", ErrCantSetField, key))
			continue
		}
		if refF.Kind() == reflect.Struct && key == "" {
			errs = append(errs, parse(refF))
			continue
		}

		if err := set(key, refF, refFT); err != nil {
			errs = append(errs, err)
			continue
		}
	}

	return errors.Join(errs...)
}

func set(key string, field reflect.Value, fieldT reflect.StructField) error {
	var v any
	var err error

	if key == "" {
		return nil
	}

	if tm := asTextUnmarshaler(field); tm != nil {
		val, err := env.Get[string](key)
		if err != nil {
			return err
		}

		return tm.UnmarshalText([]byte(val))
	}

	switch fieldT.Type.Kind() {
	case reflect.String:
		v, err = env.Get[string](key)
	case reflect.Bool:
		v, err = env.Get[bool](key)
	case reflect.Int:
		v, err = env.Get[int](key)
	case reflect.Int8:
		v, err = env.Get[int8](key)
	case reflect.Int16:
		v, err = env.Get[int16](key)
	case reflect.Int32:
		v, err = env.Get[int32](key)
	case reflect.Int64:
		v, err = env.Get[int64](key)
	case reflect.Uint:
		v, err = env.Get[uint](key)
	case reflect.Uint8:
		v, err = env.Get[uint8](key)
	case reflect.Uint16:
		v, err = env.Get[uint16](key)
	case reflect.Uint32:
		v, err = env.Get[uint32](key)
	case reflect.Uint64:
		v, err = env.Get[uint64](key)
	case reflect.Float32:
		v, err = env.Get[float32](key)
	case reflect.Float64:
		v, err = env.Get[float64](key)
	case reflect.Complex64:
		v, err = env.Get[complex64](key)
	case reflect.Complex128:
		v, err = env.Get[complex128](key)
	case reflect.Slice:
		switch fieldT.Type.Elem().Kind() {
		case reflect.String:
			v, err = env.Get[[]string](key)
		case reflect.Bool:
			v, err = env.Get[[]bool](key)
		case reflect.Int:
			v, err = env.Get[[]int](key)
		case reflect.Int8:
			v, err = env.Get[[]int8](key)
		case reflect.Int16:
			v, err = env.Get[[]int16](key)
		case reflect.Int32:
			v, err = env.Get[[]int32](key)
		case reflect.Int64:
			v, err = env.Get[[]int64](key)
		case reflect.Uint:
			v, err = env.Get[[]uint](key)
		case reflect.Uint8:
			v, err = env.Get[[]uint8](key)
		case reflect.Uint16:
			v, err = env.Get[[]uint16](key)
		case reflect.Uint32:
			v, err = env.Get[[]uint32](key)
		case reflect.Uint64:
			v, err = env.Get[[]uint64](key)
		case reflect.Float32:
			v, err = env.Get[[]float32](key)
		case reflect.Float64:
			v, err = env.Get[[]float64](key)
		case reflect.Complex64:
			v, err = env.Get[[]complex64](key)
		case reflect.Complex128:
			v, err = env.Get[[]complex128](key)
		}
	default:
		err = fmt.Errorf("%w: %s", ErrInvalidType, key)
	}

	if err != nil {
		return err
	}

	field.Set(reflect.ValueOf(v).Convert(fieldT.Type))
	return nil
}

func asTextUnmarshaler(field reflect.Value) encoding.TextUnmarshaler {
	if reflect.Pointer == field.Kind() {
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
	} else if field.CanAddr() {
		field = field.Addr()
	}

	tm, ok := field.Interface().(encoding.TextUnmarshaler)
	if !ok {
		return nil
	}
	return tm
}
