package memcachier

import "github.com/memcachier/mc"

// Memcachier ...
type Memcachier struct {
	server   string
	username string
	password string
}

// NewMemcachier returns instance of memcachier
func NewMemcachier(server, username, password string) *Memcachier {
	return &Memcachier{
		server:   server,
		username: username,
		password: password,
	}
}

// Set sets a value to the cache
// using the specified `key`
func (memcachier *Memcachier) Set(key string, data interface{}) (bool, error) {
	c := mc.NewMC(memcachier.server, memcachier.username, memcachier.password)
	defer c.Quit()
	_, err := c.Set(key, data.(string), uint32(0), uint32(0), uint64(0))

	if err != nil {
		return false, err
	}

	return true, nil
}

// Get returns the `data` saved in cache
// using the specified `key`
func (memcachier *Memcachier) Get(key string) (interface{}, error) {
	c := mc.NewMC(memcachier.server, memcachier.username, memcachier.password)
	defer c.Quit()

	val, _, _, err := c.Get(key)

	if err != nil {
		return nil, err
	}
	return val, nil
}

// Delete returns a boolean value
// if there is a successful deletion
// using the specified `key`,
// returns error otherwise.
func (memcachier *Memcachier) Delete(key string) (bool, error) {
	c := mc.NewMC(memcachier.server, memcachier.username, memcachier.password)
	defer c.Quit()
	err := c.Del(key)

	if err != nil {
		return false, err
	}

	return true, nil
}
