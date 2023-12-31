package main

// filter returns a new slice containing only the elements for which the callback function returns true.
func filter[T any](slice []T, callback func(elem T) bool) []T {
	var filtered []T
	for _, s := range slice {
		if callback(s) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

// every returns true if every element in the slice returns true for the callback function.
func every[T any](slice []T, callback func(elem T) bool) bool {
	for _, s := range slice {
		if !callback(s) {
			return false
		}
	}
	return true
}

// some returns true if any element in the slice returns true for the callback function.
func some[T any](slice []T, callback func(elem T) bool) bool {
	for _, s := range slice {
		if callback(s) {
			return true
		}
	}
	return false
}

// none returns true if no element in the slice returns true for the callback function.
func none[T any](slice []T, callback func(elem T) bool) bool {
	return !some(slice, callback)
}

// includes returns true if the slice contains the element.
func includes[T comparable](slice []T, elem T) bool {
	return some(slice, func(s T) bool {
		return s == elem
	})
}
