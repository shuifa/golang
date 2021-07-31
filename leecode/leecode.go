package main

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)


 type ListNode struct {
      Val int
      Next *ListNode
 }


 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }


type Node struct {
  Val int
  Next *Node
  Random *Node
}



func main() {
	fmt.Println(titleToNumber("FXSHRXW"))
}

//  Excel表列序号
func titleToNumber(columnTitle string) int {
	var ans int
	for i := 0; i < len(columnTitle); i++ {
		ans += int(columnTitle[i] - 'A' + 1) * int(math.Pow(26, float64(len(columnTitle) - i - 1)))
	}
	return ans
}


// 二叉树寻路
func pathInZigZagTree(label int) []int {

	ret := make([]int, 0)

	for label >= 1 {
		ret = append([]int{label}, ret...)
		label >>= 1
	}
	start := ret[len(ret)-1]

	for i := len(ret) - 2; i > 0; i-- {
		ret[i] = 3 * (2 << (i - 1)) - 1 -  start >> 1
		start = ret[i]
	}

	return ret
}

func distanceK(root, target *TreeNode, k int) (ans []int) {
	// 从 root 出发 DFS，记录每个结点的父结点
	parents := map[int]*TreeNode{}
	var findParents func(*TreeNode)
	findParents = func(node *TreeNode) {
		if node.Left != nil {
			parents[node.Left.Val] = node
			findParents(node.Left)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			findParents(node.Right)
		}
	}
	findParents(root)

	// 从 target 出发 DFS，寻找所有深度为 k 的结点
	var findAns func(*TreeNode, *TreeNode, int)
	findAns = func(node, from *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == k { // 将所有深度为 k 的结点的值计入结果
			ans = append(ans, node.Val)
			return
		}
		if node.Left != from {
			findAns(node.Left, node, depth+1)
		}
		if node.Right != from {
			findAns(node.Right, node, depth+1)
		}
		if parents[node.Val] != from {
			findAns(parents[node.Val], node, depth+1)
		}
	}
	findAns(target, nil, 0)
	return
}

//  二叉树中第二小的节点
var minVal, ans int
func findSecondMinimumValue(root *TreeNode) int {
	minVal = root.Val
	ans = -1
	rangeTree(root)
	return ans
}

func rangeTree(root *TreeNode)  {
	if root == nil || (ans != -1 && root.Val >= ans) {
		return
	}
	if root.Val > minVal {
		ans = root.Val
		return
	}
	rangeTree(root.Left)
	rangeTree(root.Right)
}

// 得到子序列的最少操作次数
func minOperations(target, arr []int) int {
	n := len(target)
	pos := make(map[int]int, n)
	for i, val := range target {
		pos[val] = i
	}
	var d []int
	for _, val := range arr {
		if idx, has := pos[val]; has {
			if p := sort.SearchInts(d, idx); p < len(d) {
				d[p] = idx
			} else {
				d = append(d, idx)
			}
		}
	}
	return n - len(d)
}


//  . 从相邻元素对还原数组
func restoreArray(adjacentPairs [][]int) []int {

	m := make(map[int][]int)
	for _, pair := range adjacentPairs {
		m[pair[0]] = append(m[pair[0]], pair[1])
		m[pair[1]] = append(m[pair[1]], pair[0])
	}
	s := make([]int, len(m))
	for i, v := range m {
		if len(v) == 1 {
			s[0] = i
		}
	}
	start := s[0]
	l := 1
	for l < len(m) {
		next := m[start]
		for _, v1 := range next {
			if (l == 1 && v1 != s[l-1]) ||  v1 != s[l-2]{
				s[l] = v1
				start = v1
				l++
				break
			}
		}
	}
	return s
}

//  检查是否所有字符出现次数相同
func areOccurrencesEqual(s string) bool {
	var h = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		h[s[i]]++
	}
	var pre int
	var flag bool
	fmt.Println(h)
	for _, i := range h {
		if !flag {
			pre = i
			flag = true
		} else if i != pre {
			return false
		}
	}
	return false
}

//  替换隐藏数字得到的最晚时间
func maximumTime(time string) string {
	var ret = strings.Split(time, "")
	for i := 0; i < len(time); i++ {
		if time[i] == '?' {
			if i == 0 {
				if ret[1] == "?" || ret[1] == "0" || ret[1] == "1" || ret[1] == "2" || ret[1] == "3"{
					ret[i] = "2"
				} else {
					ret[i] = "1"
				}
			} else if i == 1 {
				if ret[0] == "2" {
					ret[1] = "3"
				} else {
					ret[1] = "9"
				}
			} else if i == 3 {
				ret[3] = "5"
			} else {
				ret[4] = "9"
			}
		}
	}

	return strings.Join(ret, "")
}

// 检查是否区域内所有整数都被覆盖
func isCovered(ranges [][]int, left, right int) bool {
	diff := [52]int{} // 差分数组
	for _, r := range ranges {
		diff[r[0]]++
		diff[r[1]+1]--
	}
	cnt := 0 // 前缀和
	for i := 1; i <= 50; i++ {
		cnt += diff[i]
		if cnt <= 0 && left <= i && i <= right {
			return false
		}
	}
	return true
}

//  复制带随机指针的链表
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	var newNode = &Node{}
	t2 := newNode
	var m = make([]*Node, 0)
	var h1 = make(map[*Node]int)

	t1 := head
	for t1 != nil {
		t2.Next = &Node{Val:t1.Val,Random: t1.Random}
		m = append(m, t2.Next)
		h1[t1] = len(m) - 1
		t2 = t2.Next
		t1 = t1.Next
	}

	t3 := newNode.Next
	for t3 != nil {
		if t3.Random != nil {
			j := h1[t3.Random]
			t3.Random = m[j]
		}
		t3 = t3.Next
	}

	return newNode.Next
}


//  两个链表的第一个公共节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	aNode, bNode := headA, headB
	for aNode != bNode {
		if aNode == nil {
			aNode = headB
		} else {
			aNode = aNode.Next
		}
		if bNode == nil {
			bNode = headA
		} else {
			bNode = bNode.Next
		}
	}
	return aNode
}

//  最高频元素的频数
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	ans := 1
	for l, r, total := 0, 1, 0; r < len(nums); r++ {
		total += (nums[r] - nums[r-1]) * (r - l)
		for total > k {
			total -= nums[r] - nums[l]
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 变位词组
func groupAnagrams(strs []string) [][]string {
	var h = make(map[[26]int][]string)
	for _, str := range strs {
		t := [26]int{}
		for i := 0; i < len(str); i++ {
			t[str[i]-'a']++
		}
		h[t] = append(h[t], str)
	}
	var ret [][]string
	for _, s := range h {
		ret = append(ret, s)
	}
	return ret
}

//  有效的括号
func isValid(s string) bool {

	var h = map[byte]byte{'}': '{', ']': '[', ')': '('}
	var li list.List

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '('{
			li.PushBack(s[i])
		} else if li.Len() == 0{
			return false
		} else  {
			e := li.Back()
			if e.Value != h[s[i]] {
				return false
			}
			li.Remove(e)
		}
	}

	if li.Len() > 0 {
		return false
	}

	return true
}

// 最接近的三数之和
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var l = len(nums)
	var nearest = math.MaxInt32

	for fir := 0; fir < l; fir++ {
		if fir > 0 && nums[fir] == nums[fir-1] {
			continue
		}
		sec, thi := fir+1, l-1
		for sec < thi {
			sum := nums[fir] + nums[sec] + nums[thi]
			if sum == target {
				return sum
			}
			if abs(sum-target) < abs(nearest-target) {
				nearest = sum
			}
			if sum > target {
				t := thi - 1
				for sec < t && nums[t] == nums[thi] {
					t--
				}
				thi = t
			} else {
				t := sec + 1
				for t < thi && nums[t] == nums[sec] {
					t++
				}
				sec = t
			}
		}
	}

	return nearest
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	ret := make([][]int, 0)

	for fir := 0; fir < l; fir++ {
		if fir > 0 && nums[fir] == nums[fir-1] {
			continue
		}
		thi := l - 1
		for sec := fir + 1; sec < l; sec++ {
			if sec > fir+1 && nums[sec] == nums[sec-1] {
				continue
			}

			for sec < thi && nums[sec]+nums[thi] > -1*nums[fir] {
				thi--
			}

			if thi == sec {
				break
			}

			if nums[sec]+nums[thi] == -1*nums[fir] {
				ret = append(ret, []int{nums[sec], nums[thi], nums[fir]})
			}
		}
	}

	return ret
}

//  最长公共前缀
func longestCommonPrefix(strs []string) string {
	var ret string
	if len(strs) == 0 {
		return ret
	}

	fStr := strs[0]
	var max, end int = len(fStr), len(fStr)

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			return ret
		}
		for start := 0; start < len(strs[i]) && start < max; start++ {
			if strs[i][start] != fStr[start] {
				end = start
				break
			}
			end = start + 1
		}

		if end < max {
			max = end
		}

		if end == 0 {
			break
		}
	}

	return fStr[:end]
}

//  字符串转换整数 (atoi)
func myAtoi(s string) int {

	var symbol int = 1
	var str strings.Builder
	var flag bool

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && !flag {
			continue
		}
		if !flag && (s[i] == '-' || s[i] == '+') {
			if s[i] == '-' {
				symbol = -1
			}
			flag = true
		} else if s[i] >= '0' && s[i] <= '9' {
			str.WriteByte(s[i])
			flag = true
		} else {
			break
		}
	}

	i, _ := strconv.Atoi(str.String())

	ret := i * symbol

	if ret < math.MinInt32 {
		return math.MinInt64
	}

	if ret > math.MaxInt64 {
		return math.MaxInt64
	}
	return ret
}
