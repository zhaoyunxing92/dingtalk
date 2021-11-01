package cache

//过期的
type Expired interface {
	CreatedAt() int64
	ExpiresIn() int16
}

type Cache interface {
	Set(data Expired) error
	Get(data Expired) error
}
