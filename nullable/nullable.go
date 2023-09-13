package nullable

import (
	"encoding/json"
)

type Nullable[T any] struct {
	hasValue bool
	value    T
}

func Just[T any](v T) Nullable[T] {
	return Nullable[T]{
		hasValue: true,
		value:    v,
	}
}

func Null[T any]() Nullable[T] {
	var t T
	return Nullable[T]{
		hasValue: false,
		value:    t,
	}
}

func (n Nullable[T]) HasValue() bool {
	return n.hasValue
}

func (n Nullable[T]) Value() (T, bool) {
	if !n.hasValue {
		return n.value, false
	}

	return n.value, true
}

func (n Nullable[T]) ValueOrDefault(d T) T {
	if !n.hasValue {
		return d
	}

	return n.value
}

func (n Nullable[T]) AssertValue() T {
	if !n.hasValue {
		panic("null dereference error")
	}
	return n.value
}

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.hasValue {
		return []byte("null"), nil
	}

	return json.Marshal(n.value)
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.hasValue = false
		var t T
		n.value = t
		return nil
	}

	var value T
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	n.hasValue = true
	n.value = value
	return nil
}

func Then[T any, V any](n Nullable[T], fn func(T) Nullable[V]) Nullable[V] {
	if !n.hasValue {
		return Null[V]()
	}

	return fn(n.value)
}
