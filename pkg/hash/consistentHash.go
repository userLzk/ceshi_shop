package hash
//
import (
	"hash/crc32"
	"log"
	"sort"
	"strconv"
	"sync"
)

type HashRing []int

//哈希函数声明
type Hash func(data []byte) uint32
//实现一致性哈希
type Map struct {
	//结构包含哈希函数
	hash Hash
	//结点信息
	nodes     map[int]Node
	//虚拟节点数
	replicas int
	//哈希环
	ring HashRing
	//哈希环对应关系
	hashMap map[int]string

	hashMux sync.Mutex // 添加锁
}

//设置结点信息
type Node struct {
	Ip string //结点地址
	Id int //id
	Port int //端口
	Weight int //权重
	ReplicaId int //副本id
}
//设置结点信息定义
func NewNode(id, weight, port int, ip string)*Node  {
	return &Node{
		Id:id,
		Ip: ip,
		Port: port,
		Weight: weight,
	}
}

func NewHashConsistent() *Map {
	return &Map{
		//节点信息
		nodes: make(map[int]Node),
		//虚拟节点数
		replicas: 2,
		//哈希环
		ring: HashRing{},
		//哈希对应关系
		hashMap:make(map[int]string),
	}
}


//获取对应节点
func (m * Map)GetMapKey(key string) string  {

	//hash := m.hash([]byte(key))
	//log.Println(hash)
	//return hash
	return ""
}
//添加结点
//func (m *Map) Add(keys ...string)  {
//	//分配哈希key
//	for _, key :=range keys{
//		//设置副本结点
//		for i := 0; i < m.replicas; i++ {
//			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
//			//生成哈希环
//			m.ring = append(m.ring,hash)
//			//设置对应关系
//			m.hashMap[hash] =key
//		}
//	}
//	// 哈希环排序
//	sort.Ints(m.ring)
//}
//添加结点
func (m *Map) Add(node *Node) bool  {
	m.hashMux.Lock()
	defer m.hashMux.Unlock()

	if _, ok := m.hashMap[node.Id]; ok {
		log.Fatalln("Id err:",node.Id)
		return false
	}
	//设置副本结点
	for i := 0; i <= m.replicas; i++ {
		//根据每个节点信息组成该节点 once
		hashkey := m.GetHashKey(i,node)
		node.ReplicaId = i
		m.nodes[hashkey] = *node

	}
	m.ring = HashRing{}
	//处理对应关系
	for key := range m.nodes{
		m.ring = append(m.ring,key)
	}
	// 哈希环排序
	sort.Ints(m.ring)
	return true
}

//回去副本节点key
func (m *Map)GetHashKey(i int,node *Node) int {
	return int(m.CheckHashKey(strconv.Itoa(i)+"-"+strconv.Itoa(node.Id) + "-" + strconv.Itoa(node.Weight)+"-"+ strconv.Itoa(node.Weight) + "-" + node.Ip))
}

//计算哈希32
func (m *Map)CheckHashKey(str string) uint32  {
	return crc32.ChecksumIEEE([]byte(str))
}

//检测结点有效性
func (m *Map) Reset(keys ...string)  {
	//先清空hash 环
	m.ring = nil
	m.hashMap = map[int]string{}
	//m.Add(keys...)
}

