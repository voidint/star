package holder

import "sync"

var pool = make(map[string]Stargazer)
var mux sync.Mutex

// Register Register stargazer with name.
func Register(name string, sg Stargazer) {
	mux.Lock()
	defer mux.Unlock()
	if sg == nil {
		panic("Register stargazer is nil")
	}
	if _, dup := pool[name]; dup {
		panic("Register called twice for stargazer " + name)
	}
	pool[name] = sg
}

// Registered Return registered stargazer names.
func Registered() (names []string) {
	mux.Lock()
	defer mux.Unlock()

	for key := range pool {
		names = append(names, key)
	}
	return names
}

// PickStargazer Pick up stargazer by name.
func PickStargazer(name string) Stargazer {
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

// Stargazer Stargazer
type Stargazer interface {
	Whoami() (holder string)
}
