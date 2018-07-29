package memcacher

import (
	"bytes"
	"compress/gzip"

	"github.com/bradfitz/gomemcache/memcache"
)

// Memcacher defines the basic methods
// of caching
type Memcacher interface {
	Set(string, interface{}) (bool, error)
	Get(string) (interface{}, error)
	Delete(string) (bool, error)
}

// Memcached struct for our concrete
// implemenation of memcached
type Memcached struct {
	memcachedHost   string // localhost
	memcachedPort   string // 11211
	memcachedServer string // localhost:11211
	isCompressed    bool
	client          *memcache.Client
}

// NewMemcached constructor for our concrete
// imlementation of memcacher
func NewMemcached(host, port, server string) *Memcached {
	return &Memcached{
		memcachedHost:   host,
		memcachedPort:   port,
		memcachedServer: server,
		isCompressed:    true,
		client:          memcache.New(server),
	}
}

// Set sets value to memcache
// and returns boolean if successfully saved
// returns error otherwise
func (m *Memcached) Set(suffix string, val interface{}) (bool, error) {
	var e error
	var key string
	if m.isCompressed {
		key = "mycache." + ".c." + suffix
		e = m.client.Set(&memcache.Item{
			Key:        key,
			Value:      gzcompress(val.(string)),
			Expiration: 0,
		})
	} else {
		key = "mycache." + suffix
		e = m.client.Set(&memcache.Item{
			Key:        key,
			Value:      []byte(val.(string)),
			Expiration: 0,
		})
	}

	if e != nil {
		return false, e
	}
	return true, nil
}

func gzcompress(val string) []byte {
	var b bytes.Buffer

	gz := gzip.NewWriter(&b)

	if _, err := gz.Write([]byte(val)); err != nil {
		return []byte("")
	}
	if err := gz.Flush(); err != nil {
		return []byte("")
	}
	if err := gz.Close(); err != nil {
		return []byte("")
	}
	return b.Bytes()
}
