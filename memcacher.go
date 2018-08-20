package memcacher

// Memcacher defines the basic methods
// of caching
type Memcacher interface {
	Set(string, interface{}) (bool, error)
	Get(string) (interface{}, error)
	Delete(string) (bool, error)
}

// Set sets data to the cache
// and returns a boolean value
// data is saved successfully,
// returns error otherwise
func Set(key string, data interface{}, memcacher Memcacher) (bool, error) {
	return memcacher.Set(key, data)
}

// Get retrieves data from cache
func Get(key string, memcacher Memcacher) (interface{}, error) {
	return memcacher.Get(key)
}

// Delete deletes an item from the cache
// using the specified key. Returns
// boolean value if succcessful, error otherwise
func Delete(key string, memcacher Memcacher) (bool, error) {
	return memcacher.Delete(key)
}
