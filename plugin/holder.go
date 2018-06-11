package plugin

import "sync"

var pool = make(map[string]Holder)
var mux sync.Mutex

// Register Register Holder with name.
func Register(name string, h Holder) {
	mux.Lock()
	defer mux.Unlock()
	if h == nil {
		panic("Register Holder is nil")
	}
	if _, dup := pool[name]; dup {
		panic("Register called twice for Holder " + name)
	}
	pool[name] = h
}

// Registered Return registered holder names.
func Registered() (names []string) {
	mux.Lock()
	defer mux.Unlock()

	for key := range pool {
		names = append(names, key)
	}
	return names
}

// PickHolder Pick up holder by name.
func PickHolder(name string) Holder {
	mux.Lock()
	defer mux.Unlock()

	for key := range pool {
		if key == name {
			return pool[key]
		}
	}
	return nil
}

const (
	// DefHolder Default holder name
	DefHolder = "github"
)

// Holder holder
type Holder interface {
	Whoami() (holder string)
	SetToken(token string)
}
