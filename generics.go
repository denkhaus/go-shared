package shared

import (
	"sort"

	"github.com/denkhaus/containers"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// from https://bitfieldconsulting.com/golang/functional

func Sort[E constraints.Ordered](s []E) []E {
	result := make([]E, len(s))
	copy(result, s)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

type mapFunc1[E any] func(E) E

func Map[E any](s []E, f mapFunc1[E]) []E {
	result := make([]E, len(s))
	for i := range s {
		result[i] = f(s[i])
	}
	return result
}

type mapFunc2[T1, T2 any] func(T1) T2

func Map2[T1, T2 any](s []T1, f mapFunc2[T1, T2]) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

type forEachFunc[E any] func(E) error

func ForEach[E any](s []E, f forEachFunc[E]) error {
	for i := range s {
		if err := f(s[i]); err != nil {
			return err
		}
	}

	return nil
}

type keepFunc[E any] func(E) bool

func Filter[E any](s []E, f keepFunc[E]) []E {
	result := []E{}
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func FilterIsEven[T constraints.Integer](v T) bool {
	return v%2 == 0
}

func FilterIsOdd[T constraints.Integer](v T) bool {
	return v%2 != 0
}

type reduceFunc[E any] func(E, E) E

func Reduce[E any](s []E, init E, f reduceFunc[E]) E {
	cur := init
	for _, v := range s {
		cur = f(cur, v)
	}
	return cur
}

type Normalizer[T Number] struct {
	minValue T
	maxValue T
}

func (p *Normalizer[T]) Update(v T) {
	p.minValue = containers.Min(p.minValue, v)
	p.maxValue = containers.Max(p.maxValue, v)
}

func (p *Normalizer[T]) Get(v T) T {
	return (v - p.minValue) / (p.maxValue - p.minValue)
}

func OneOf[T comparable](value T, coll ...T) bool {
	for _, val := range coll {
		if val == value {
			return true
		}
	}

	return false
}

// Defaultify returns defaultValue if v is equal to its default value.
func Defaultify[T comparable](v T, defaultValue T) T {
	var def T
	if v == def {
		return defaultValue
	}
	return v
}
