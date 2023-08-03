nats 底层存储结构

1. dir list

```text
jetstream
└── $G
    └── streams
        └── foo_stream
            ├── meta.inf
            ├── meta.sum
            ├── msgs
            │   ├── 7.blk
            │   └── 7.idx
            └── obs
                └── C1
                    ├── meta.inf
                    ├── meta.sum
                    └── o.dat

```
foo_stream/meta.inf

```json
{
    "Created":"2021-11-07T12:56:13.56432767Z",
    "Name":"",
    "durable_name":"C1",
    "deliver_subject":"_INBOX.J6h6nsYgSi7rm9ryKBaCRh",
    "deliver_policy":"all",
    "ack_policy":"explicit",
    "ack_wait":30000000000,
    "max_deliver":-1,
    "filter_subject":"foo_subject",
    "replay_policy":"instant",
    "max_ack_pending":524288
}
```
foo_stream/meta.sum
```txt
44be9c30336b4ad0
```
meta.inf
```json
{
    "Created":"2021-11-07T12:56:13.562392499Z",
    "name":"foo_stream",
    "subjects":[
        "foo_subject"
    ],
    "retention":"interest",
    "max_consumers":-1,
    "max_msgs":-1,
    "max_bytes":-1,
    "max_age":0,
    "max_msgs_per_subject":-1,
    "max_msg_size":-1,
    "discard":"old",
    "storage":"file",
    "num_replicas":1,
    "duplicate_window":120000000000
}
```

	// Index file for a consumer.
	consumerState = "o.dat"
1. 构建缓存 
通过这个方法重新构建msgblock 的缓存
```go
if ld, err := mb.rebuildState(); err != nil && ld != nil {
```


server 启动时候读取idx 的索引文件内容 来判断每个msg block的文件内包含到sequence范围


```go

type StreamStore interface
} {
	StoreMsg(subject string, hdr, msg []byte) (uint64, int64, error)
	StoreRawMsg(subject string, hdr, msg []byte, seq uint64, ts int64) error
	SkipMsg() uint64
	LoadMsg(seq uint64) (subject string, hdr, msg []byte, ts int64, err error)
	LoadLastMsg(subject string) (subj string, seq uint64, hdr, msg []byte, ts int64, err error)
	RemoveMsg(seq uint64) (bool, error)
	EraseMsg(seq uint64) (bool, error)
	Purge() (uint64, error)
	PurgeEx(subject string, seq, keep uint64) (uint64, error)
	Compact(seq uint64) (uint64, error)
	Truncate(seq uint64) error
	GetSeqFromTime(t time.Time) uint64
	FilteredState(seq uint64, subject string) SimpleState
	SubjectsState(filterSubject string) map[string]SimpleState
	State() StreamState
	FastState(*StreamState)
	Type() StorageType
	RegisterStorageUpdates(StorageUpdateHandler)
	UpdateConfig(cfg *StreamConfig) error
	Delete() error
	Stop() error
	ConsumerStore(name string, cfg *ConsumerConfig) (ConsumerStore, error)
	Snapshot(deadline time.Duration, includeConsumers, checkMsgs bool) (*SnapshotResult, error)
	Utilization() (total, reported uint64, err error)
}

```

3. 获取消息根据sequence

```go
func (mb *msgBlock) cacheLookup(seq uint64) (*fileStoredMsg, error) {
	bi, _, hashChecked, err := mb.slotInfo(int(seq - mb.cache.fseq))
   // 这个msg block的第一个seq的数据 假如要找1500 这个开始1000 那么要找到位置就是1500-1000
   // 因为每个msgblock 的索引默认从0 开始 统计该block 中全部有多少个数量
}
```
4. 底层消息存储到具体结构
22（头部长度）+subject+msg+hash(8位)
```txt 
/消息的结构
	// Formats
	// Format with no header
	// total_len(4) sequence(8) timestamp(8) subj_len(2) subj msg hash(8)
	// With headers, high bit on total length will be set.
	// total_len(4) sequence(8) timestamp(8) subj_len(2) subj hdr_len(4) hdr msg hash(8)
```

5. 获取消息是
```go

// Get next available message from underlying store.
// Is partition aware and redeliver aware.
// Lock should be held.
func (o *consumer) getNextMsg() (subj string, hdr, msg []byte, seq uint64, dc uint64, ts int64, err error) {
	if o.mset == nil || o.mset.store == nil {
		return _EMPTY_, nil, nil, 0, 0, 0, errBadConsumer
	}
	for {
		seq, dc := o.sseq, uint64(1)
		if o.hasSkipListPending() {
			seq = o.lss.seqs[0]
			if len(o.lss.seqs) == 1 {
				o.sseq = o.lss.resume
				o.lss = nil
				o.updateSkipped()
			} else {
				o.lss.seqs = o.lss.seqs[1:]
			}
		} else if o.hasRedeliveries() {
			seq = o.getNextToRedeliver()
			dc = o.incDeliveryCount(seq)
			if o.maxdc > 0 && dc > o.maxdc {
				// Only send once
				if dc == o.maxdc+1 {
					o.notifyDeliveryExceeded(seq, dc-1)
				}
				// Make sure to remove from pending.
				delete(o.pending, seq)
				continue
			}
		} else if o.maxp > 0 && len(o.pending) >= o.maxp {
			// maxp only set when ack policy != AckNone and user set MaxAckPending
			// Stall if we have hit max pending.
			return _EMPTY_, nil, nil, 0, 0, 0, errMaxAckPending
		}

		subj, hdr, msg, ts, err := o.mset.store.LoadMsg(seq)
		if err == nil {
			if dc == 1 { // First delivery.
				o.sseq++
				if o.cfg.FilterSubject != _EMPTY_ && !o.isFilteredMatch(subj) {
					log.Println("跳过该消息", o.sseq-1)
					o.updateSkipped()
					continue
				}
			}
			// We have the message_queue here.
			return subj, hdr, msg, seq, dc, ts, nil
		}
		// We got an error here. If this is an EOF we will return, otherwise
		// we can continue looking.
		if err == ErrStoreEOF || err == ErrStoreClosed || err == errNoCache || err == errPartialCache {
			return _EMPTY_, nil, nil, 0, 0, 0, err
		}
		// Skip since its deleted or expired.
		o.sseq++
	}
}

```
```text
o.mset.store.LoadMsg(seq)
```
通过上面到方式来从开始的sequence 递增来一直线性到查找
