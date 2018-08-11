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

// Set ...
func (memcachier *Memcachier) Set(key string, data interface{}) (bool, error) {
	c := mc.NewMC(memcachier.server, memcachier.username, memcachier.password)
	defer c.Quit()
	_, err := c.Set(key, data.(string), uint32(0), uint32(0), uint64(0))

	if err != nil {
		return false, err
	}

	return true, nil
}
