package repository

type CacheRepository struct {
	Get func(key string) string
}

func NewCacheRepository() *CacheRepository {
	return &CacheRepository{}
}

func (r *CacheRepository) Get(key string) string {
	return "value"
}
