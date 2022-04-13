package entity

// definition of keyvalue pairs to be p[assed to log function
type KvPair struct {
	Key   string
	Value interface{}
}

func NewKvPair(key string, value interface{}) KvPair {
	return KvPair{
		Key:   key,
		Value: value,
	}
}
