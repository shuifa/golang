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


func main() {
	fmt.Println(groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"}))
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
