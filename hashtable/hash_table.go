package hashtable

import "log"

type hashtable struct {
	size      int64
	dataArray []*item
}

type item struct {
	Key   int64
	Value int64
}

func (m *hashtable) hashcode(key int64) int64 {
	return key % m.size
}

func (m *hashtable) Get(key int64) (int64, bool) {
	hashIndex := m.hashcode(key)
	// log.Println("hashIndex ", hashIndex, "of key", key)
	for i := hashIndex; i < m.size; i++ {
		if m.dataArray[i] != nil && m.dataArray[i].Key == key {
			return m.dataArray[i].Value, true
		}
	}
	return -1, false
}

func (m *hashtable) Put(key, value int64) bool {
	hashIndex := m.hashcode(key)
	log.Println("hashIndex ", hashIndex, "of key", key)

	mvleft := hashIndex - 1
	mvright := hashIndex
	for {
		if m.dataArray[mvleft] == nil || m.dataArray[mvleft].Key == -1 {
			m.dataArray[mvleft] = &item{
				Key:   key,
				Value: value,
			}
			return true
		}

		if mvleft >= 0 && (m.dataArray[mvright] == nil || m.dataArray[mvright].Key == -1) {
			m.dataArray[mvright] = &item{
				Key:   key,
				Value: value,
			}
		}
		if mvleft == 0 && mvright == m.size-1 {
			return false
		}
		if mvleft > 0 {
			mvleft--
		}
		if mvright < m.size-1 {
			mvright++
		}

	}

}
func (m *hashtable) Remove(key int64) bool {
	hashIndex := m.hashcode(key)
	for i := hashIndex; i < m.size; i++ {
		if m.dataArray[i] == nil || m.dataArray[i].Key != -1 {
			m.dataArray[i].Key = -1
			return true
		}
	}
	return false
}
