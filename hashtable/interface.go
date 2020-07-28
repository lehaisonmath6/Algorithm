package hashtable

type HashTable interface {
	Put(key, value int64) bool
	Get(key int64) (int64, bool)
	Remove(key int64) bool
}
