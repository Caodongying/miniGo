package leakybucket //包名要和所在目录一样，并且全部小写，没有大写或者下划线

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct{
	leakRate float64 // 漏出速率，精度为秒
	lastLeakTime time.Time // 上次漏出的时间
	capacity float64 // 漏桶总容量
	currentWater float64 // 漏桶当前水量

	// 多线程安全
	mu sync.Mutex
}

func (lb *LeakyBucket) allowRequest() bool {
  lb.mu.Lock()
  defer lb.mu.Unlock()
  // 首先计算从上次漏水到现在为止，应该漏出去多少水
  now := time.Now()
  elapsed := now.Sub(lb.lastLeakTime).Seconds() // 转换成float64防止丢失精度
  fmt.Println("距离上一次漏水过去了", elapsed, "秒")
  // 更新桶中现存水量和最近一次的漏水时间
  lb.currentWater = max(lb.currentWater - elapsed * lb.leakRate, 0)
  lb.lastLeakTime = now
  // 判断是否可以处理请求
  if (lb.currentWater + 1) <= lb.capacity {
	lb.currentWater += 1
	fmt.Println("可以处理请求。")
	return true
  } else {
	fmt.Println("请求被拒绝。")
	return false
  }
}

func newLeakyBucket(rate, capacity float64) *LeakyBucket {
	// 返回结构体的地址，避免拷贝
	return &LeakyBucket{
		leakRate: rate,
		lastLeakTime: time.Now(),
		capacity: capacity,
		currentWater: 0,
	}
}

func main() {
 lb := newLeakyBucket(3, 10) // 漏水速率为3个/秒，桶的容量为10

 // 模拟流量处理
 for i:=0; i<40; i++ {
	time.Sleep(200*time.Millisecond)
	fmt.Println(i)
	lb.allowRequest()
 }

 time.Sleep(2*time.Second)

 for i:=0; i<40; i++ {
	time.Sleep(200*time.Millisecond)
	fmt.Println(i)
	lb.allowRequest()
 }

}
