package box

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"reflect"
	"sync"
	"time"
)

type Box struct {
	context.Context

	mu       sync.RWMutex
	items    map[string]any
	deadline time.Time
	err      error
}

func New() *Box {
	return &Box{
		mu:    sync.RWMutex{},
		items: map[string]any{},
	}
}

func (self *Box) Len() int {
	return len(self.items)
}

func (self *Box) Put(values ...any) {
	self.mu.Lock()
	defer self.mu.Unlock()

	for _, value := range values {
		t := reflect.ValueOf(value).Type()
		self.items[t.String()] = value
	}
}

func (self *Box) PutByKey(key string, value any) {
	self.mu.Lock()
	defer self.mu.Unlock()

	self.items[key] = value
}

func (self *Box) Deadline() (deadline time.Time, ok bool) {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.deadline, !self.deadline.IsZero()
}

func (self *Box) Done() <-chan struct{} {
	return nil
}

func (self *Box) Err() error {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.err
}

func (self *Box) Value(key any) any {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.items[reflect.ValueOf(key).String()]
}

func (self *Box) Fork() *Box {
	box := New()
	box.items = maps.Clone(self.items)
	return box
}

func (self *Box) Merge(box *Box) {
	self.mu.Lock()
	defer self.mu.Unlock()
	maps.Copy(self.items, box.items)
}

func (self *Box) Inject(handler any) (func(), error) {
	self.mu.RLock()
	defer self.mu.RUnlock()

	value := reflect.ValueOf(handler)

	if !value.IsValid() {
		return nil, errors.New("handler must be a function")
	}

	t := value.Type()

	if t.Kind() != reflect.Func {
		return nil, errors.New("handler must be a function")
	}

	args := make([]reflect.Value, t.NumIn())

	for i := 0; i < t.NumIn(); i++ {
		paramType := t.In(i)
		item, exists := self.items[paramType.String()]

		if !exists {
			return nil, fmt.Errorf(
				"no injectable value found for '%s'",
				paramType.Name(),
			)
		}

		args[i] = reflect.ValueOf(item)
	}

	return func() {
		value.Call(args)
	}, nil
}

func (self *Box) String() string {
	self.mu.RLock()
	defer self.mu.RUnlock()

	b, _ := json.Marshal(self.items)
	return string(b)
}

func (self *Box) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.items)
}
