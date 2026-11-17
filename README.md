# go-triton-storage

## 能力声明

这个实现只声明了部分能力，**不满足建锁的全部前提**，
`NewStorageLock` 会拒绝在不满足前提的存储上创建锁。

已声明：

- `CapabilityMissing` — 见源码
- `CapabilityReliableTime` — 可信、跨节点可协调的时间源（建锁前提）
- `CapabilityWatch` — 状态变更通知（仅优化等待，不是互斥前提）

`GetName()` 返回 `triton-storage`。

## 构造函数

```go
func NewTritonStorage(ctx context.Context)
```

包名：`triton_storage`

## 分布式锁的前提

锁的正确性依赖两件事：同一个 `lockId` 上有**线性一致的 CAS**，以及
**可靠、单调且跨节点可协调的时间源**。能力标志是接口自述，不是证明；
具体约束与故障模型见
[go-storage-lock 的 README](https://github.com/storage-lock/go-storage-lock#能力边界)。

## 用法

```go
store, err := triton_storage.NewTritonStorage(...)
lock, err := storage_lock.NewStorageLock(store, "my-lock-key")
owner := lock.NewOwnerId()
err = lock.Lock(ctx, owner)
```

同一套 `构造 Storage → NewStorageLock → Lock(ctx, ownerId)` 在
[go-storage-lock-example](https://github.com/storage-lock/go-storage-lock-example) 里有
各后端的完整示例。

## License

MIT
