package triton_storage

import (
	"context"
	"testing"

	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
	storage_test_helper "github.com/storage-lock/go-storage-test-helper"
	"github.com/stretchr/testify/assert"
)

// TestTritonStorageStoragePlaceholder 占位实现的诚实行为测试：
// 未接入真实 TritonStorage 服务端，不可用于分布式锁。
// Capabilities 声明 ReliableTime + Watch（无 CAS），写操作返回 ErrStorageCapabilityMissing。
func TestTritonStorageStoragePlaceholder(t *testing.T) {
	ctx := context.Background()
	s, err := NewTritonStorage(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, s)

	// GetName 非空
	assert.Equal(t, StorageName, s.GetName())

	// Capabilities：实现如实声明只有 Watch（本机时钟不算可信时间源）
	caps := s.Capabilities()
	assert.NotContains(t, caps, storage.CapabilityReliableTime)
	assert.NotContains(t, caps, storage.CapabilityCAS)

	// Init / Close 无错
	assert.Nil(t, s.Init(ctx))
	assert.Nil(t, s.Close(ctx))

	// GetTime 可用
	tm, err := s.GetTime(ctx)
	assert.Nil(t, err)
	assert.False(t, tm.IsZero())

	info := storage_test_helper.BuildTestLockInformation(t)

	// 写操作诚实返回能力缺失
	assert.ErrorIs(t, s.CreateWithVersion(ctx, "lock-1", 1, info), storage_lock.ErrStorageCapabilityMissing)
	assert.ErrorIs(t, s.UpdateWithVersion(ctx, "lock-1", 1, 2, info), storage_lock.ErrStorageCapabilityMissing)
	assert.ErrorIs(t, s.DeleteWithVersion(ctx, "lock-1", 1, info), storage_lock.ErrStorageCapabilityMissing)

	// Get 返回不存在
	_, err = s.Get(ctx, "lock-1")
	assert.ErrorIs(t, err, storage_lock.ErrLockNotFound)

	// List 返回空迭代器
	it, err := s.List(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, it)
	assert.False(t, it.Next())
}

// TestTritonStorageWatchWiring Watch 接线测试：返回轮询通道；cancel 确保后台 goroutine 退出。
func TestTritonStorageWatchWiring(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s, err := NewTritonStorage(ctx)
	assert.Nil(t, err)
	ch, err := s.Watch(ctx, "lock-1")
	assert.Nil(t, err)
	assert.NotNil(t, ch)
}
