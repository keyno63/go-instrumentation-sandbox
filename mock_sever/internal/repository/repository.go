package repository

type CacheRepository struct {
}

func NewCacheRepository() *CacheRepository {
	return &CacheRepository{}
}

func (r *CacheRepository) Get(key string) string {
	return "value"
}
