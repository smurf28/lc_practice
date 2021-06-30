package common

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

// 参数传递，只能修改，不能新增或者删除原始数据
// 默认 s=s[0:len(s)]，取下限不取上限，数学表示为：[)
// leetcode 中，全局变量不要当做返回值，否则刷题检查器会报错

func TestStack(t *testing.T) {
	// 栈
	// 创建栈
	stack := make([]int, 0)
	// push压入
	stack = append(stack, 10)
	// pop弹出
	// v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	// 检查栈空
	// len(stack) == 0
}

func TestQueue(t *testing.T) {
	// 创建队列
	queue := make([]int, 0)
	// enqueue入队
	queue = append(queue, 10)
	// dequeue出队
	// v := queue[0]
	queue = queue[1:]
	// 长度0为空
	// len(queue) == 0
}

// map 键需要可比较，不能为 slice、map、function
// map 值都有默认值，可以直接操作默认值，如：m[age]++ 值由 0 变为 1
// 比较两个 map 需要遍历，其中的 kv 是否相同，因为有默认值关系，所以需要检查 val 和 ok 两个值
func TestMap(t *testing.T) {
	// 创建
	m := make(map[string]int)
	// 设置kv
	m["hello"] = 1
	// 删除k
	delete(m, "hello")
	// 遍历
	for k, v := range m {
		println(k, v)
	}
}

func TestSort(t *testing.T) {
	// int排序
	sort.Ints([]int{})
	// 字符串排序
	sort.Strings([]string{})
	// 自定义排序
	s := []int{}
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
}

func TestMath(t *testing.T) {
	// int32 最大最小值
	// math.MaxInt32 // 实际值：1<<31-1
	// math.MinInt32 // 实际值：-1<<31

	// int64 最大最小值（int默认是int64）
	// math.MaxInt64
	// math.MinInt64
}

func TestCopy(t *testing.T) {
	// 删除a[i]，可以用 copy 将i+1到末尾的值覆盖到i,然后末尾-1
	a := make([]int, 0)
	i := 1
	copy(a[i:], a[i+1:])
	a = a[:len(a)-1]

	// make创建长度，则通过索引赋值
	// a := make([]int, n)
	// a[n] = x

	// make长度为0，则通过append()赋值
	// a := make([]int, 0)
	// a = append(a, x)

	// 拷贝数组
	// track := make([]int, len(sub))
	// copy(track, sub)
}

// 类型转换
func TestConvert(t *testing.T) {
	// byte转数字
	s := "12345"                        // s[0] 类型是byte
	num := int(s[0] - '0')              // 1
	str := string(s[0])                 // "1"
	b := byte(num + '0')                // '1'
	fmt.Printf("%d%s%c\n", num, str, b) // 111

	// 字符串转数字
	num, _ = strconv.Atoi(s)
	str = strconv.Itoa(num)
}

// 回溯法模板
func TestRoll(t *testing.T) {
	// 	result = []
	// 	func backtrack(选择列表,路径):
	// 		if 满足结束条件:
	// 			result.add(路径)
	// 			return
	// 		for 选择 in 选择列表:
	// 			做选择
	// 			backtrack(选择列表,路径)
	// 			撤销选择

	// atomic.AddUint64(addr *uint64, delta uint64)

}

// 分治法
// func traversal(root *TreeNode) ResultType  {
//     // nil or leaf
//     if root == nil {
//         // do something and return
//     }

//     // Divide
//     ResultType left = traversal(root.Left)
//     ResultType right = traversal(root.Right)

//     // Conquer
//     ResultType result = Merge from left and right

//     return result
// }

func worker(name string, stopch chan struct{}) {
	for {
		select {
		case <-stopch:
			fmt.Println("receive a stop signal, ", name)
			return
		default:
			fmt.Println("worker", name)
		}
	}

}

const (
	test = 123
)

func TestConst(t *testing.T) {
	t.Log(test)
}

type CustomError struct{}

func (*CustomError) Error() string {
	return ""
}

// 判断一个类型是否实现了某个接口
func TestReflect(t *testing.T) {

	typeOfError := reflect.TypeOf((*error)(nil)).Elem()
	customErrorPtr := reflect.TypeOf(&CustomError{})
	customError := reflect.TypeOf(CustomError{})

	fmt.Println(customErrorPtr.Implements(typeOfError)) // #=> true
	fmt.Println(customError.Implements(typeOfError))    // #=> false

	s := make([]int, 1000, 1000)

	for index, _ := range s {
		s[index] = index
	}
}
