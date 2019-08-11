# MyfirstSimpleBlockChain-1

## 1.区块结构

  height           区块高度
  
  PrevBlockHash    前一个区块的hash值
  
  Data             写入的数据，此处为简易实现
  
  TimeStamp        时间戳
  
  Hash             hash值
  
## 2.计算hash

  此处利用SHA256算法
  
  1）拼接字段：此处运用最简单的方式拼接字段
  
  2）设置hash：Sethash（）
  
## 3.创建区块 newBlock()

  1）创建创世区块：CreateGenesisBlock（）
  
  2）定义区块链：
  <pre>type Blockchain struct {
  
	Blocks []*Block//存储有序区块	
}</pre>

## 4.bitcoin block structure
完整的比特币的区块头（block header）结构如下：

|Field	|Purpose|Updated when...|Size (Bytes)|
|-------|-------|---------------|------------|
|Version|Block version number|	You upgrade the software and it specifies a new version	|4|
|hashPrevBlock	|256-bit hash of the previous block header	|A new block comes in	|32|
|hashMerkleRoot|	256-bit hash based on all of the transactions in the block|	A transaction is accepted	|32|
|Time|	Current timestamp as seconds since 1970-01-01T00:00 UTC	|Every few seconds	|4|
|Bits|	Current target in compact format	|The difficulty is adjusted	|4|
|Nonce|	32-bit number (starts at 0)	|A hash is tried (increments)	|4|

而我们的 Data, 在比特币中对应的是交易，是另一个单独的数据结构。为了简便起见，目前将这两个数据结构放在了一起。在真正的比特币中，区块的数据结构如下：

|Field|	Description|	Size|
|---|---|---|
|Magic no|	value always 0xD9B4BEF9|	4 bytes|
|Blocksize	|number of bytes following up to end of block|	4 bytes|
|Blockheader	|consists of 6 items|	80 bytes|
Transact|		|           |


[**转向第二个blockchain**](https://github.com/Billy1900/SimpleBlockchain-2)
