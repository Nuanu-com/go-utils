package slice_utils

func First[T any](s []T) *T {
	if len(s) > 0 {
		result := s[0]
		return &result
	}

	return nil
}

func Map[T any, R any](slice []T, modifier func(i T) R) []R {
	results := []R{}

	for _, s := range slice {
		results = append(results, modifier(s))
	}

	return results
}

func GroupBy[T any, C comparable](slice []T, comparison func(T) C) map[C][]T {
	result := map[C][]T{}

	for _, s := range slice {
		comparison_result := comparison(s)
		result[comparison_result] = append(result[comparison_result], s)
	}

	return result
}

func Reduce[T any, C any](slice []T, initial C, reducer func(C, T) C) C {
	result := initial

	for _, s := range slice {
		result = reducer(result, s)
	}

	return result
}

func Filter[T any](slice []T, filterFn func(T) bool) []T {
	results := []T{}

	for _, s := range slice {
		if filterFn(s) {
			results = append(results, s)
		}
	}

	return results
}

func Get[T any](s []T, idx int) (T, bool) {
	if idx >= len(s) {
		var result T
		return result, false
	}

	return s[idx], true
}

func FindBy[T any](s []T, predicate func(T) bool) (T, bool) {
	var result T
	found := false

	for _, item := range s {
		if predicate(item) {
			result = item
			found = true
			break
		}
	}

	return result, found
}
