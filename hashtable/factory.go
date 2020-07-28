package hashtable

func NewHashTable(size int64) HashTable {
	return &hashtable{
		size:      size,
		dataArray: make([]*item, size),
	}
}
