package response

type Key struct {
	Key string `json:"key"`
}

func NewKey(key string) *Key {
	return &Key{Key: key}
}
