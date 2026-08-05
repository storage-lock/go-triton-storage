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
	return []storage.StorageCapability{storage.CapabilityCAS, storage.CapabilityReliableTime}
}
func (x *TritonStorage) Init(ctx context.Context) error { return nil }
func (x *TritonStorage) CreateWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return nil
}
func (x *TritonStorage) UpdateWithVersion(ctx context.Context, lockId string, expectedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return nil
}
func (x *TritonStorage) DeleteWithVersion(ctx context.Context, lockId string, expectedVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return nil
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
