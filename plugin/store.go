package plugin

import "sync"

var (
	stores   map[string]*store
	curStore *store
	hmux     sync.Mutex
)

// Use Switch to the specified holder
func Use(holderName string) {
	mux.Lock()
	defer mux.Unlock()

	for k, v := range stores {
		if k == holderName {
			curStore = v
			return
		}
	}
	panic("unregistered holder name")
}

type store struct {
	holderName string
	token      string
}
