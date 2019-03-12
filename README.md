# MyfirstSimpleBlockChain-1

1.区块结构

  height           区块高度
  
  PrevBlockHash    前一个区块的hash值
  
  Data             写入的数据，此处为简易实现
  
  TimeStamp        时间戳
  
  Hash             hash值
  
2.计算hash

  此处利用SHA256算法
  
  1）拼接字段：此处运用最简单的方式拼接字段
  
  2）设置hash：Sethash（）
  
3.创建区块 newBlock()

  1）创建创世区块：CreateGenesisBlock（）
  
  2）定义区块链：
  
  type Blockchain struct {
  
	Blocks []*Block//存储有序区块
	
}
