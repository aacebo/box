package box

import "reflect"

func Get[T any](box *Box) T {
	box.mu.RLock()
	defer box.mu.RUnlock()

	t := reflect.TypeFor[T]()
	value := box.items[t.String()]
	return value.(T)
}

func GetByKey[T any](box *Box, key string) T {
	box.mu.RLock()
	defer box.mu.RUnlock()
	return box.items[key].(T)
}
