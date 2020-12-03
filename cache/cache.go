package cache

type CacheDB interface {
	Set(key, val string) error
	Get(key string) (string, error)
}
