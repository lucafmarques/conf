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
	ErrNotStructPointer = errors.New("argument must be pointer to struct")
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

		if !refF.CanSet() {
			errs = append(errs, ErrCantSetField)
			continue
		}

		if refF.Kind() == reflect.Struct {
			errs = append(errs, parse(refF))
			continue
		}

		if err := set(refF, refFT); err != nil {
			errs = append(errs, err)
			continue
		}
	}

	return errors.Join(errs...)
}

func set(field reflect.Value, fieldT reflect.StructField) error {
	var v any
	var err error
	envN := fieldT.Tag.Get("env")

	switch fieldT.Type.Kind() {
	case reflect.String:
		v, err = env.Get[string](envN)
	case reflect.Bool:
		v, err = env.Get[bool](envN)
	case reflect.Int:
		v, err = env.Get[int](envN)
	case reflect.Float64:
		v, err = env.Get[float64](envN)
	default:
		t, ok := field.Interface().(encoding.TextUnmarshaler)
		if !ok {
			return fmt.Errorf("%w: %s", ErrInvalidType, fieldT.Type.Name())
		}
		if v, err = env.Get[string](envN); err == nil {
			err = t.UnmarshalText([]byte(v.(string)))
		}
	}

	if err != nil {
		if errors.Is(err, env.ErrUnset) {
			err = fmt.Errorf("%w: %s", err, envN)
		}
		return err
	}

	field.Set(reflect.ValueOf(v).Convert(fieldT.Type))
	return nil
}
