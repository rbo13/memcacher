package memcacher

import (
	"bufio"
	"bytes"
	"compress/gzip"

	"github.com/bradfitz/gomemcache/memcache"
)

var prefix = "mycache."

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
// implementation of memcacher
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
		key = prefix + ".c." + suffix
		e = m.client.Set(&memcache.Item{
			Key:        key,
			Value:      gzcompress(val.(string)),
			Expiration: 0,
		})
	} else {
		key = prefix + suffix
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

func (m *Memcached) Get(suffix string) (interface{}, error) {
	var key string
	if m.isCompressed {
		key = prefix + ".c." + suffix
	} else {
		key = prefix + suffix
	}

	it, err := m.client.Get(key)
	if err != nil {
		return "", err
	}
	if m.isCompressed {
		return gzuncompress(it.Value)
	}
	return string(it.Value), nil
}

func (m *Memcached) Delete(suffix string) (bool, error) {
	key := prefix + suffix
	if m.isCompressed {
		key = prefix + ".c." + suffix
	}

	e := m.client.Delete(key)

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

func gzuncompress(b []byte) (string, error) {
	bb := bytes.NewBuffer(b)
	zipread, _ := gzip.NewReader(bb)

	defer zipread.Close()
	reader := bufio.NewReader(zipread)

	var (
		part []byte
		err  error
	)
	ret := ""

	for {
		if part, _, err = reader.ReadLine(); err != nil {
			break
		}

		ret += string(part)

	}
	return ret, nil

}

// GetVal returns the value of cache
func GetVal(key string, mem Memcacher) (interface{}, error) {
	return mem.Get(key)
}
