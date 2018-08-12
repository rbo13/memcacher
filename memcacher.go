package memcacher

// Memcacher defines the basic methods
// of caching
type Memcacher interface {
	Set(string, interface{}) (bool, error)
	Get(string) (interface{}, error)
	Delete(string) (bool, error)
}
