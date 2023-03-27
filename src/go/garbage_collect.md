## 几种垃圾回收算法

- ######  标记-清除算法 Mark-Sweep GC

为两个阶段：

1. 标记阶段：从根集合出发，将所有活动对象及其子对象打上标记

2. 清除阶段：遍历堆，将非活动对象（未打上标记）的连接到空闲链表上

 

- ######  标记-压缩 Mark-Compact

和“标记－清除”相似，不过在标记阶段后它将所有活动对象紧密的排在堆的一侧（压缩），消除了内存碎片， 不过压缩是需要花费计算成本的


- ###### 引用计数 Reference Counting

- ###### GC 复制算法

- 将堆分为两个大小相同的空间 From 和 To， 利用 From 空间进行分配，当 From 空间满的时候，GC将其中的活动对象复制到 To 空间，之后将两个空间互换即完成GC。

- ###### 分代回收

- ###### 增量式GC

## 三色标记

- 白色对象 — 潜在的垃圾，其内存可能会被垃圾收集器回收；

- 黑色对象 — 活跃的对象，包括不存在任何引用外部指针的对象以及从根对象可达的对象；

- 灰色对象 — 活跃的对象，因为存在指向白色对象的外部指针，垃圾收集器会扫描这些对象的子对象；

### 过程

* 第一步 , 就是只要是新创建的对象,默认的颜色都是标记为“白色
* 第二步, 每次GC回收开始, 然后从根节点开始遍历所有对象，把遍历到的对象从白色集合放入“灰色”集合。
* 第三步, 遍历灰色集合，将灰色对象引用的对象从白色集合放入灰色集合，之后将此灰色对象放入黑色集合
* 第四步, 重复第三步, 直到灰色中无任何对象.
* 第五步: 回收所有的白色标记表的对象. 也就是回收垃圾.
以上的过程不允许并发，要想并发安全，需要引入内存屏障
### 内存屏障

内存屏障技术是一种屏障指令，它可以让 CPU 或者编译器在执行内存相关操作时遵循特定的约束

### 触发位置

* [`runtime.mallocgc`](https://draveness.me/golang/tree/runtime.mallocgc)  分配内存

* [`runtime.GC`](https://draveness.me/golang/tree/runtime.GC)  开发主动调用

*  后台触发

运行时会在应用程序启动时在后台开启一个用于强制触发垃圾收集的 Goroutine，该 Goroutine 的职责非常简单 — 调用 [`runtime.gcStart`](https://draveness.me/golang/tree/runtime.gcStart) 尝试启动新一轮的垃圾收集：

```go
func init() {
	go forcegchelper()
}

func forcegchelper() {
	forcegc.g = getg()
	for {
		lock(&forcegc.lock)
		atomic.Store(&forcegc.idle, 1)
		goparkunlock(&forcegc.lock, waitReasonForceGGIdle, traceEvGoBlock, 1)
		gcStart(gcTrigger{kind: gcTriggerTime, now: nanotime()})
	}
}
```

```go
func sysmon() {
	...
	for {
		...
		if t := (gcTrigger{kind: gcTriggerTime, now: now}); t.test() && atomic.Load(&forcegc.idle) != 0 {
			lock(&forcegc.lock)
			forcegc.idle = 0
			var list gList
			list.push(forcegc.g)
			injectglist(&list)
			unlock(&forcegc.lock)
		}
	}
}
```

## GC调优

**控制、减少、复用**。

1. 控制变量或者数据的创建频率
2. 减少不必要的数据分配
3. sync.Pool 数据复用（href)[[深度分析 Golang sync.Pool 底层原理 | 编程沉思录 (cyhone.com)](https://www.cyhone.com/articles/think-in-sync-pool/)]
4. 调整GOGC参数