# 跳表：为什么Redis一定要用跳表来实现有序集合？

## 跳表的概念

#### 概念

跳表这种数据结构对你来说，可能会比较陌生，因为一般的数据结构和算法书籍里都不怎么会讲。但是它确实是一种各方面性能都比较优秀的动态数据结构，可以支持快速地插入、删除、查找操作，写起来也不复杂，甚至可以替代红黑树（Red-black tree）。

### 如何理解“跳表”？


