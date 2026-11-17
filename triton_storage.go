package triton_storage

import (
	"context"
	"sync"
	"time"

	go_iterator "github.com/golang-infrastructure/go-iterator"
	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
)

// TritonStorage triton存储实现
// 尚未接入真实服务端，仅为占位实现，不可用于分布式锁。
// 真正接入时需借助服务端原子语义实现 CAS。
type TritonStorage struct {
	mu sync.RWMutex
}

var _ storage.Storage = &TritonStorage{}

func NewTritonStorage(ctx context.Context) (*TritonStorage, error) {
	return &TritonStorage{}, nil
}

const StorageName = "triton-storage"

func (x *TritonStorage) GetName() string { return StorageName }
func (x *TritonStorage) Capabilities() []storage.StorageCapability {
	// Capabilities 如实声明：这个实现是占位壳，GetTime 返回本机时间而非可信服务端时钟，
	// 条件写方法也返回 ErrStorageCapabilityMissing。因此不声明 CapabilityReliableTime。
	// Watch 仍保留：它只是轮询，不构成建锁前提。
	return []storage.StorageCapability{storage.CapabilityWatch}
}
func (x *TritonStorage) Watch(ctx context.Context, lockId string) (<-chan *storage.LockInformation, error) {
	return storage.NewPollingLockWatcher(x, 500*time.Millisecond).Watch(ctx, lockId)
}
func (x *TritonStorage) Init(ctx context.Context) error { return nil }
func (x *TritonStorage) CreateWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return storage_lock.ErrStorageCapabilityMissing
}
func (x *TritonStorage) UpdateWithVersion(ctx context.Context, lockId string, expectedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return storage_lock.ErrStorageCapabilityMissing
}
func (x *TritonStorage) DeleteWithVersion(ctx context.Context, lockId string, expectedVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return storage_lock.ErrStorageCapabilityMissing
}
func (x *TritonStorage) Get(ctx context.Context, lockId string) (string, error) {
	x.mu.RLock()
	defer x.mu.RUnlock()
	return "", storage_lock.ErrLockNotFound
}
func (x *TritonStorage) GetTime(ctx context.Context) (time.Time, error) { return time.Now(), nil }
func (x *TritonStorage) Close(ctx context.Context) error { return nil }
func (x *TritonStorage) List(ctx context.Context) (go_iterator.Iterator[*storage.LockInformation], error) {
	return go_iterator.FromSlice([]*storage.LockInformation{}), nil
}
