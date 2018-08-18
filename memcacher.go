package memcacher

// Memcacher defines the basic methods
// of caching
type Memcacher interface {
	Set(string, interface{}) (bool, error)
	Get(string) (interface{}, error)
	Delete(string) (bool, error)
}

func Set(key string, data interface{}, memcacher Memcacher) (bool, error) {
	return memcacher.Set(key, data)
}

func Get(key string, memcacher Memcacher) (interface{}, error) {
	return memcacher.Get(key)
}

func Delete(key string, memcacher Memcacher) (bool, error) {
	return memcacher.Delete(key)
}
