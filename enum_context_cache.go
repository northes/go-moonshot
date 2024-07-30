package moonshot

type ContextCacheStatus string

const (
	ContextCacheStatusPending  ContextCacheStatus = "pending"  // 当缓存被初次创建时，其初始状态为 pending
	ContextCacheStatusReady    ContextCacheStatus = "ready"    // 如果参数合法，缓存创建成功，其状态变更为 ready
	ContextCacheStatusError    ContextCacheStatus = "error"    // 如果参数不合法，或因其他原因缓存创建失败，其状态变更为 error
	ContextCacheStatusInactive ContextCacheStatus = "inactive" // 对于已过期的缓存，其状态变更为 inactive
)

func (c ContextCacheStatus) String() string {
	return string(c)
}

type ContextCacheOrder string

const (
	ContextCacheOrderAsc  ContextCacheOrder = "asc"  // 升序
	ContextCacheOrderDesc ContextCacheOrder = "desc" // 降序
)

func (c ContextCacheOrder) String() string {
	return string(c)
}
