package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// 剑指 offer 刷题

// 剑指 Offer 38. 字符串的排列
// 输入一个字符串，打印出该字符串中字符的所有排列。

func permutation(s string) []string {
	var res = make([]string, 0, 1000)
	bytes := []byte(s)
	var dfs func(x int)
	dfs = func(x int) {
		if x == len(bytes)-1 {
			res = append(res, string(bytes))
			return
		}
		dict := map[byte]bool{}
		for i := x; i < len(bytes); i++ {
			if dict[bytes[i]] {
				continue
			}
			bytes[x], bytes[i] = bytes[i], bytes[x]
			dict[bytes[x]] = true
			dfs(x + 1)
			bytes[x], bytes[i] = bytes[i], bytes[x]
		}
	}
	dfs(0)
	return res
}

// 剑指 Offer 59 - I. 滑动窗口的最大值
// 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nil
	}
	ans := make([]int, 1, len(nums)-k+1)
	c := Constructor()
	for i := 0; i < k; i++ {
		c.Push_back(nums[i])
	}
	ans[0] = c.Max_value()
	for n := k; n < len(nums); n++ {
		c.Pop_front()
		c.Push_back(nums[n])
		ans = append(ans, c.Max_value())
	}
	return ans
}

// 剑指 Offer 59 - II. 队列的最大值

// 请定义一个队列并实现函数 max_value 得到队列里的最大值，
// 要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
// 若队列为空，pop_front 和 max_value需要返回 -1

type MaxQueue struct {
	x []int
	y []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.y) == 0 {
		return -1
	}
	return this.y[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.x = append(this.x, value)
	l := len(this.y) - 1
	for ; l >= 0 && value > this.y[l]; l-- {
	}
	if l < len(this.y)-1 {
		this.y = this.y[:l+1]
	}
	this.y = append(this.y, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.x) == 0 {
		return -1
	}
	t := this.x[0]
	this.x = this.x[1:]
	if t == this.y[0] {
		this.y = this.y[1:]
	}

	return t
}

// 剑指 Offer 67. 把字符串转换成整数
// 写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数

func strToInt(str string) int {
	// 去除头部空格
	for ; len(str) > 0 && str[0] == ' '; str = str[1:] {
	}

	sb := strings.Builder{}

	for i := 0; i < len(str); i++ {
		if i == 0 {
			if str[i] != '+' && str[i] != '-' && (str[i] < '0' || str[i] > '9') {
				return 0
			}
			if (str[i] == '+' || str[i] == '-') && len(str) == i+1 {
				return 0
			}
		} else if str[i] < '0' || str[i] > '9' {
			break
		}
		sb.WriteByte(str[i])
	}

	s := sb.String()
	if len(s) == 0 {
		return 0
	}

	t := s[0]
	if t == '+' || t == '-' {
		s = s[1:]
	}
	var ans int
	for i := 0; i < len(s); i++ {
		ans = int(s[i]-'0') + ans*10
	}
	if t == '-' {
		ans = -ans
	}
	if ans > math.MaxInt32 {
		return math.MaxInt32
	}
	if ans < math.MinInt32 {
		return math.MinInt32
	}

	return ans
}

// 剑指 Offer 20. 表示数值的字符串
// 请实现一个函数用来判断字符串是否表示数值（包括整数和小数）

func isNumber(s string) bool {
	s = strings.ToUpper(s)

	// 先清除首尾空格

	for ; len(s) > 0 && s[len(s)-1] == ' '; s = s[:len(s)-1] {
	}
	if len(s) == 0 || s[len(s)-1] == 'E' || s[0] == 'E' {
		return false
	}

	t := strings.Split(s, "E")

	if len(t) > 2 {
		return false
	}

	isInt := func(s string) bool {
		for i := 0; i < len(s); i++ {
			if i == 0 {
				if s[i] != '+' && s[i] != '-' && (s[i] < '0' || s[i] > '9') {
					return false
				}
				if (s[i] == '+' || s[i] == '-') && i == len(s)-1 {
					return false
				}
			} else if s[i] < '0' || s[i] > '9' {
				return false
			}
		}
		return true
	}

	isFloat := func(s string) bool {
		point := true
		num := false
		for i := 0; i < len(s); i++ {
			if i == 0 {
				if s[i] != '+' && s[i] != '-' && (s[i] < '0' || s[i] > '9') && s[i] != '.' {
					return false
				}
				if s[i] == '.' {
					if i == len(s)-1 {
						return false
					}
					point = false
				}
				if s[i] >= '0' || s[i] > '9' {
					num = true
				}

			} else if s[i] == '.' {
				if !point {
					return false
				}
				point = false
			} else if s[i] < '0' || s[i] > '9' {
				return false
			} else {
				num = true
			}
		}

		fmt.Println(s, point, num)
		return !point && num
	}

	if !isFloat(t[0]) && !isInt(t[0]) {
		return false
	}

	if len(t) == 2 && !isInt(t[1]) {
		return false
	}

	return true
}

// 剑指 Offer 31. 栈的压入、弹出序列
// 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。
// 假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，
// 序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列

func validateStackSequences(pushed []int, popped []int) bool {

	var stack = make([]int, 0, len(popped))
	i := 0
	for _, v := range pushed {
		// 模拟入栈
		stack = append(stack, v)
		// 模拟出栈
		for ; len(stack) > 0 && stack[len(stack)-1] == popped[i]; stack = stack[:len(stack)-1] {
			i++
		}
	}
	return len(stack) == 0
}

// 剑指 Offer 29. 顺时针打印矩阵
// 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	n := len(matrix) * len(matrix[0])
	ans := make([]int, 0, n)

	colAdd := func(row, col, max int) {
		for col <= max {
			ans = append(ans, matrix[row][col])
			col++
		}
	}
	colDe := func(row, col, min int) {
		for col >= min {
			ans = append(ans, matrix[row][col])
			col--
		}
	}
	rowAdd := func(row, col, max int) {
		for row <= max {
			ans = append(ans, matrix[row][col])
			row++
		}
	}
	rowDe := func(row, col, min int) {
		for row >= min {
			ans = append(ans, matrix[row][col])
			row--
		}
	}
	minCol, maxCol := 0, len(matrix[0])-1
	minRow, maxRow := 0, len(matrix)-1
	row, col := 0, 0
	ans = append(ans, matrix[row][col])
	for len(ans) < n {
		colAdd(row, col+1, maxCol)
		if len(ans) == n {
			break
		}
		col = maxCol
		minRow++
		rowAdd(row+1, col, maxRow)
		if len(ans) == n {
			break
		}
		row = maxRow
		maxCol--
		colDe(row, col-1, minCol)
		if len(ans) == n {
			break
		}
		col = minCol
		maxRow--
		rowDe(row-1, col, minRow)
		if len(ans) == n {
			break
		}
		row = minRow
		minCol++
	}

	return ans
}

// 剑指 Offer 14- I. 剪绳子
// 给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），
// 每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？
// 例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18

func cuttingRope(n int) int {
	var ans = make([]int, n+1, n+1)
	ans[1] = 1
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 1; i <= n; i++ {
		tMax := 0
		for l, j := 1, i-1; l <= i/2; l++ {
			tMax = max(tMax, max(l, ans[l])*max(j, ans[j]))
			j--
		}
		ans[i] = tMax
		fmt.Println(ans)
	}
	return ans[n]
}

// 剑指 Offer 57 - II. 和为s的连续正数序列
// 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

func findContinuousSequence(target int) [][]int {
	var ans = make([][]int, 0)

	makeSlice := func(x, y int) []int {
		t := make([]int, y-x+1, y-x+1)
		for i := x; i <= y; i++ {
			t[i-x] = i
		}
		return t
	}

	left, right, sum := 1, 2, 3
	for left <= target/2 {
		if target > sum {
			right++
			sum += right
		} else {
			if target == sum {
				ans = append(ans, makeSlice(left, right))
			}
			sum -= left
			left++
		}
	}
	return ans
}

// 剑指 Offer 62. 圆圈中最后剩下的数字
// 0,1,···,n-1这n个数字排成一个圆圈，
// 从数字0开始，每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。
// 求出这个圆圈里剩下的最后一个数字

func lastRemaining(n int, m int) int {
	var ans int
	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}

// 剑指 Offer 66. 构建乘积数组
// 给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，
// 其中B[i] 的值是数组 A 中除了下标 i 以外的元素的积,
// 即B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法

func constructArr(a []int) []int {
	l := len(a)
	var ans = make([]int, l, l)
	ans[0] = 1
	for i := 1; i < l; i++ {
		ans[i] = ans[i-1] * a[i-1]
	}
	t := 1
	for i := l - 2; i >= 0; i-- {
		t *= a[i+1]
		ans[i] *= t
	}
	return ans
}

// 剑指 Offer 39. 数组中出现次数超过一半的数字
// 数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字

func majorityElement(nums []int) int {
	var ans, vote int
	for _, num := range nums {
		if vote == 0 {
			ans = num
		}
		if num == ans {
			vote += 1
		} else {
			vote -= 1
		}
	}
	return ans
}

// 剑指 Offer 56 - II. 数组中数字出现的次数 II
// 在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。
// 请找出那个只出现一次的数字。

func singleNumber(nums []int) int {

	var ans int
	for i := 0; i < 32; i++ {
		r := 0
		for _, num := range nums {
			r += (num >> i) & 1
		}
		ans += (r % 3) << i
	}
	return ans
}

func singleNumber2(nums []int) int {
	count := [32]int{}
	// 统计出现1的次数
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			if t := 1 << i; t&num == t {
				count[i]++
			}
		}
	}
	var ans int
	for i, c := range count {
		if c%3 != 0 {
			ans |= 1 << i
		}
	}
	return ans
}

// 剑指 Offer 56 - I. 数组中数字出现的次数
// 一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。
// 请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。

func singleNumbers(nums []int) []int {
	var n int
	// 找到单独两个数的异位值
	for _, num := range nums {
		n ^= num
	}
	// 找到这个异位值的任意一个位为1的地方
	t := 1
	for ; t&n == 0; t <<= 1 {
	}

	var x, y int
	// 进行分组比较，保证了相同的数在同一组，不同的两个数一定不再同一边
	for _, num := range nums {
		if num&t == 0 {
			x ^= num
		} else {
			y ^= num
		}
	}

	return []int{x, y}
}

// 剑指 Offer 65. 不用加减乘除做加法
// 写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。

func add(a int, b int) int {

	for b != 0 {
		c := a & b // 需要进位的地方
		a ^= b     // 本位相加
		b = c << 1 // 进行进位
	}
	return a
}

// 剑指 Offer 15. 二进制中1的个数
// 编写一个函数，输入是一个无符号整数（以二进制串的形式），
// 返回其二进制表达式中数字位数为 '1' 的个数

func hammingWeight(num uint32) int {
	var ans int
	for ; num > 0; num &= num - 1 {
		ans++
	}
	return ans
}

func hammingWeight2(num uint32) int {
	var ans int
	for ; num > 0; num >>= 1 {
		if num&1 == 1 {
			ans++
		}
	}
	return ans
}

// 剑指 Offer 33. 二叉搜索树的后序遍历序列
// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。
// 如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同

func verifyPostorder(postorder []int) bool {
	fmt.Println(postorder)
	if len(postorder) <= 2 {
		return true
	}
	// 在二叉搜索树中，左子树的元素是都小于根元素，右子树都大于根元素
	// 在后序遍历中，最后一个元素是根元素
	head := len(postorder) - 1
	for head != 0 {
		// popinter 统计符合二叉搜索树的后序遍历的节点数
		popinter := 0
		// 从前面开始遍历，小于的当前根元素的值是左子树的，当找到第一个大于当前根元素的值，可以确定后半段的元素都应是在当前节点的右子树
		for postorder[popinter] < postorder[head] {
			popinter++
		}
		// 如果后半段里面有小于根元素的值的元素，就说明这个不是二叉搜索树的后序遍历，跳出循环
		for postorder[popinter] > postorder[head] {
			popinter++
		}
		// popinter != head 或 popinter < head 说明该数组不是某二叉搜索树的后序遍历结果
		if popinter != head {
			return false
		}
		// 进入下一个节点继续验证
		head--
	}
	return true
}

func verifyPostorder2(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	root := postorder[len(postorder)-1]
	i := 0
	for ; i < len(postorder)-1; i++ {
		if postorder[i] > root {
			break
		}
	}
	left := postorder[:i]
	right := postorder[i : len(postorder)-1]
	for k := 0; k < len(right); k++ {
		if right[k] < root {
			return false
		}
	}
	lRet := verifyPostorder(left)
	rRet := verifyPostorder(right)

	return lRet && rRet
}

// 剑指 Offer 16. 数值的整数次方
// 实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	var ans float64 = 1

	if n < 0 {
		x, n = 1/x, -n
	}

	for n > 0 {
		if n&1 == 1 {
			ans *= x
		}
		x *= x
		n >>= 1
	}

	return ans
}

// 剑指 Offer 07. 重建二叉树
// 输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

// 剑指 Offer 68 - I. 二叉搜索树的最近公共祖先
// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	tMap := make(map[*TreeNode]*TreeNode)

	for len(queue) > 0 {
		t := queue
		queue = []*TreeNode{}
		for i := 0; i < len(t); i++ {
			if t[i].Left != nil {
				tMap[t[i].Left] = t[i]
				queue = append(queue, t[i].Left)
			}
			if t[i].Right != nil {
				tMap[t[i].Right] = t[i]
				queue = append(queue, t[i].Right)
			}
		}
	}

	pQueue := []*TreeNode{p}
	t1 := p

	for {
		v, ok := tMap[t1]
		if !ok {
			break
		}
		pQueue = append(pQueue, v)
		t1 = v
	}

	qMap := make(map[*TreeNode]bool)
	qMap[q] = true
	t2 := q

	for {
		v, ok := tMap[t2]
		if !ok {
			break
		}
		qMap[v] = true
		t2 = v
	}

	for k := 0; k < len(pQueue); k++ {
		if qMap[pQueue[k]] {
			return pQueue[k]
		}
	}

	return nil
}

// 剑指 Offer 64. 求1+2+…+n
// 求 1+2+...+n ，要求不能使用乘除法、
// for、while、if、else、switch、case等关键字及条件判断语句（A?B:C

func sumNums(n int) int {
	var ans int
	var sum func(n int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)

	return ans
}

// 剑指 Offer 55 - II. 平衡二叉树
// 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。
// 如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}

// 剑指 Offer 55 - I. 二叉树的深度
// 输入一棵二叉树的根节点，求该树的深度。
// 从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度

func maxDepth(root *TreeNode) int {

	var maxLevel int
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		level += 1
		if level > maxLevel {
			maxLevel = level
		}
		if root.Left != nil {
			dfs(root.Left, level)
		}
		if root.Right != nil {
			dfs(root.Right, level)
		}
	}
	dfs(root, 1)
	return maxLevel
}

// 剑指 Offer 41. 数据流中的中位数
// 如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，
// 那么中位数就是所有数值排序之后位于中间的数值。
// 如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

type maxHeap []int // 大顶堆
type minHeap []int // 小顶堆

// 每个堆都要heap.Interface的五个方法：Len, Less, Swap, Push, Pop
// 其实只有Less的区别。

// Len 返回堆的大小
func (m maxHeap) Len() int {
	return len(m)
}
func (m minHeap) Len() int {
	return len(m)
}

// Less 决定是大优先还是小优先
func (m maxHeap) Less(i, j int) bool { // 大根堆
	return m[i] > m[j]
}
func (m minHeap) Less(i, j int) bool { // 小根堆
	return m[i] < m[j]
}

// Swap 交换下标i, j元素的顺序
func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// Push 在堆的末尾添加一个元素，注意和heap.Push(heap.Interface, interface{})区分
func (m *maxHeap) Push(v interface{}) {
	*m = append(*m, v.(int))
}
func (m *minHeap) Push(v interface{}) {
	*m = append(*m, v.(int))
}

// Pop 删除堆尾的元素，注意和heap.Pop()区分
func (m *maxHeap) Pop() interface{} {
	old := *m
	n := len(old)
	v := old[n-1]
	*m = old[:n-1]
	return v
}
func (m *minHeap) Pop() interface{} {
	old := *m
	n := len(old)
	v := old[n-1]
	*m = old[:n-1]
	return v
}

// MedianFinder 维护两个堆，前一半是大顶堆，后一半是小顶堆，中位数由两个堆顶决定
type MedianFinder struct {
	maxH *maxHeap
	minH *minHeap
}

// Constructor 建两个空堆
func Constructor3() MedianFinder {
	return MedianFinder{
		new(maxHeap),
		new(minHeap),
	}
}

// AddNum 插入元素num
// 分两种情况插入：
// 1. 两个堆的大小相等，则小顶堆增加一个元素（增加的不一定是num）
// 2. 小顶堆比大顶堆多一个元素，大顶堆增加一个元素
// 这两种情况又分别对应两种情况：
// 1. num小于大顶堆的堆顶，则num插入大顶堆
// 2. num大于小顶堆的堆顶，则num插入小顶堆
// 插入完成后记得调整堆的大小使得两个堆的容量相等，或小顶堆大1

func (m *MedianFinder) AddNum(num int) {
	if m.maxH.Len() == m.minH.Len() {
		if m.minH.Len() == 0 || num >= (*m.minH)[0] {
			heap.Push(m.minH, num)
		} else {
			heap.Push(m.maxH, num)
			top := heap.Pop(m.maxH).(int)
			heap.Push(m.minH, top)
		}
	} else {
		if num > (*m.minH)[0] {
			heap.Push(m.minH, num)
			bottle := heap.Pop(m.minH).(int)
			heap.Push(m.maxH, bottle)
		} else {
			heap.Push(m.maxH, num)
		}
	}
}

// FindMedian FindMediam 输出中位数
func (m *MedianFinder) FindMedian() float64 {
	if m.minH.Len() == m.maxH.Len() {
		return float64((*m.maxH)[0])/2.0 + float64((*m.minH)[0])/2.0
	} else {
		return float64((*m.minH)[0])
	}
}

// 剑指 Offer 40. 最小的k个数
// 输入整数数组 arr ，找出其中最小的 k 个数。
// 例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

// 剑指 Offer 61. 扑克牌中的顺子
// 从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，
// 即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，
// Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。

func isStraight(nums []int) bool {

	var min, max = 14, 0
	var m = make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
		if m[nums[i]] {
			return false
		}
		m[nums[i]] = true
	}

	return max-min < 5
}

// 剑指 Offer 45. 把数组排成最小的数
// 输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，
// 打印能拼接出的所有数字中最小的一个

func minNumber(nums []int) string {

	var s = make([]string, len(nums))

	for i := 0; i < len(nums); i++ {
		s[i] = strconv.Itoa(nums[i])
	}

	sort.Slice(s, func(i, j int) bool {
		if s[i]+s[j] < s[j]+s[i] {
			return true
		}
		return false
	})

	return strings.Join(s, "")
}

// 剑指 Offer 54. 二叉搜索树的第k大节点
// 给定一棵二叉搜索树，请找出其中第 k 大的节点的值

func kthLargest(root *TreeNode, k int) int {

	var ans int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		k--
		if k == 0 {
			ans = root.Val
			return
		}
		dfs(root.Left)
	}
	dfs(root)
	return ans
}

// 剑指 Offer 36. 二叉搜索树与双向链表
// 输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。
// 要求不能创建任何新的节点，只能调整树中节点指针的指向

func treeToDoublyList(root *TreeNode) {
	var ans []*TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		ans = append(ans, root)
		dfs(root.Left)
	}
	dfs(root)
}

// 剑指 Offer 34. 二叉树中和为某一个值的路径
// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，
// 找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

func pathSum(root *TreeNode, target int) [][]int {

	var ans [][]int

	var sum func(root *TreeNode, nodes []int, sums int)
	sum = func(root *TreeNode, nodes []int, sums int) {
		if root == nil {
			return
		}

		nodes = append(nodes, root.Val)

		if root.Left == nil && root.Right == nil && sums == target {
			ans = append(ans, nodes)
			return
		}

		l, r := make([]int, len(nodes)), make([]int, len(nodes))
		copy(l, nodes)
		copy(r, nodes)

		if root.Left != nil {
			sum(root.Left, l, sums+root.Left.Val)
		}
		if root.Right != nil {
			sum(root.Right, r, sums+root.Right.Val)
		}

	}
	sum(root, []int{}, root.Val)
	return ans
}

// // 剑指 Offer 13. 机器人的运动范围
// 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
// 一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），
// 也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，
// 因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子

func movingCount(m int, n int, k int) int {
	var set = make(map[[2]int]bool)
	set[[2]int{0, -1}] = true

	for i := 0; i < m; i++ {
		for p := 0; p < n; p++ {
			if i/10+i%10+p/10+p%10 > k {
				continue
			}
			if set[[2]int{i - 1, p}] || set[[2]int{i + 1, p}] || set[[2]int{i, p - 1}] || set[[2]int{i, p + 1}] {
				set[[2]int{i, p}] = true
			}
		}
	}

	return len(set) - 1
}

// 剑指 Offer 12. 矩阵中的路径
// 给定一个m x n 二维字符网格board 和一个字符串单词word 。
// 如果word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
// 同一个单元格内的字母不允许被重复使用

func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])
	vectors := [4][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool {
		if board[r][c] != word[i] {
			return false
		}
		if i == len(word)-1 {
			return true
		}
		board[r][c] = '/'
		for v := 0; v < len(vectors); v++ {
			x := vectors[v][0] + r
			y := vectors[v][1] + c
			if x >= 0 && x < row && y >= 0 && y < col && board[x][y] != '/' {
				if dfs(x, y, i+1) {
					return true
				}
			}
		}
		board[r][c] = word[i]
		return false
	}
	for p := 0; p < row; p++ {
		for n := 0; n < col; n++ {
			if dfs(p, n, 0) {
				return true
			}
		}
	}
	return false
}

// 剑指 Offer 58 - I. 翻转单词顺序
// 输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
// 为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"

func reverseWords(s string) string {
	var ans []string
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if sb.Len() > 0 {
				ans = append([]string{sb.String()}, ans...)
			}
			sb.Reset()
			continue
		}
		sb.WriteByte(s[i])
		if i == len(s)-1 {
			ans = append([]string{sb.String()}, ans...)
		}
	}

	return strings.Join(ans, " ")
}

// 剑指 Offer 57. 和为s的两个数字
// 输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] == target {
			return []int{nums[left], nums[right]}
		} else if nums[left]+nums[right] > target {
			right--
		} else {
			left++
		}
	}
	return nil
}

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，
// 使得所有奇数在数组的前半部分，所有偶数在数组的后半部分

func exchange(nums []int) []int {

	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]&1 == 0 {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		} else {
			left--
		}
		if nums[right]&1 == 1 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		} else {
			right--
		}
	}
	return nums
}

// 剑指 Offer 52. 两个链表的第一个公共节点
// 输入两个链表，找出它们的第一个公共节点

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ta, tb := headA, headB

	for ta != tb {
		if ta == nil {
			ta = headB
		} else {
			ta = ta.Next
		}
		if tb == nil {
			tb = headA
		} else {
			tb = tb.Next
		}
	}
	return ta
}

// 剑指 Offer 25. 合并两个排序的链表
// 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	h := &ListNode{Next: nil}
	t := h

	for l1 != nil || l2 != nil {
		if l1 == nil {
			t.Next, l2 = l2, l2.Next
		} else if l2 == nil || l1.Val <= l2.Val {
			t.Next, l1 = l1, l1.Next
		} else {
			t.Next, l2 = l2, l2.Next
		}
		t = t.Next
	}
	return h.Next
}

// 剑指 Offer 22. 链表中倒数第k个节点
// 输入一个链表，输出该链表中倒数第k个节点。
// 为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点

func getKthFromEnd(head *ListNode, k int) *ListNode {
	left := head
	for head != nil {
		k--
		head = head.Next
		if k < 0 {
			left = left.Next
		}
	}
	return left
}

// 剑指 Offer 18. 删除链表的节点
// 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
// 返回删除后的链表的头节点

func deleteNode(head *ListNode, val int) *ListNode {
	tmp := &ListNode{Next: head}
	pre := tmp
	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
			break
		}
		pre = head
		head = head.Next
	}
	return tmp.Next
}

// 剑指 Offer 48. 最长不含重复字符的子字符串
// 请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度

func lengthOfLongestSubstring(s string) int {
	var m = make(map[byte]int, len(s))
	var maxLen, start int
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok && v >= start {
			start = v + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		m[s[i]] = i
	}
	return maxLen
}

// 剑指 Offer 46. 把数字翻译成字符串
// 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，
// 1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。
// 一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法

func translateNum(num int) int {
	src := strconv.Itoa(num)
	// 滚动数组法。。动态规则太难理解了
	second, first, result := 0, 1, 1
	for i := 1; i < len(src); i++ {
		second, first = first, result
		pre := src[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			result += second
		}
	}
	return result
}

// 剑指 Offer 47. 礼物的最大价值
// 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
// 你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
// 给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

func maxValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for n := 0; n < len(grid[i]); n++ {
			down, right := 0, 0
			if i > 0 {
				down = grid[i][n] + grid[i-1][n]
			}
			if n > 0 {
				right = grid[i][n] + grid[i][n-1]
			}
			if down > right {
				grid[i][n] = down
			} else if right > down {
				grid[i][n] = right
			} else if down != 0 {
				grid[i][n] = right
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

// 剑指 Offer 42. 连续子数组的最大和
// 输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值

func maxSubArray(nums []int) int {
	var maxSum = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i-1] + nums[i]
		}
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}
	return maxSum
}

// func maxSubArray(nums []int) int {
//
// 	var maxSum = math.MinInt64
// 	var sum, start int
//
// 	for i := 0; i < len(nums); i++ {
// 		start = i
// 		sum = 0
// 		for sum >= 0 && start < len(nums) {
// 			sum += nums[start]
// 			if sum > maxSum {
// 				maxSum = sum
// 			}
// 			start++
// 		}
// 	}
// 	return maxSum
// }

// 剑指 Offer 63. 股票的最大利润
// 假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？

func maxProfit(prices []int) int {
	var minPrice = math.MaxInt64
	var maxProfit int
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

// 剑指 Offer 10- I. 斐波那契数列
// 写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）

// 剑指 Offer 26. 树的子结构
// 输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}

	var ret bool

	// 当在 A 中找到 B 的根节点时，进入helper递归校验
	if A.Val == B.Val {
		ret = helper(A, B)
	}

	// ret == false，说明 B 的根节点不在当前 A 树顶中，进入 A 的左子树进行递归查找
	if !ret {
		ret = isSubStructure(A.Left, B)
	}

	// 当 B 的根节点不在当前 A 树顶和左子树中，进入 A 的右子树进行递归查找
	if !ret {
		ret = isSubStructure(A.Right, B)
	}
	return ret

	// 利用 || 的短路特性可写成
	// return helper(A,B) || isSubStructure(A.Left,B) || isSubStructure(A.Right,B)
}

// helper 校验 B 是否与 A 的一个子树拥有相同的结构和节点值
func helper(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	// a.Val == b.Val 递归校验 A B 左子树和右子树的结构和节点是否相同
	return helper(a.Left, b.Left) && helper(a.Right, b.Right)
}

// 剑指 Offer 28. 对称的二叉树
// 请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var recur func(root *TreeNode, right *TreeNode) bool
	recur = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		return recur(left.Left, right.Right) && recur(left.Right, right.Left)
	}

	return recur(root.Left, root.Right)
}

func isSymmetric2(root *TreeNode) bool {
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		t := queue
		queue = []*TreeNode{}
		allnil := false
		for i := 0; i < len(t); i++ {
			if i < len(t)/2 {
				if t[i] == nil && t[len(t)-i-1] != nil {
					return false
				}
				if t[i] != nil && t[len(t)-i-1] == nil {
					return false
				}
				if t[i] != nil && t[len(t)-i-1] != nil && t[i].Val != t[len(t)-i-1].Val {
					return false
				}
			}
			if t[i] == nil {
				queue = append(queue, nil, nil)
			} else {
				allnil = true
				queue = append(queue, t[i].Right)
			}
		}
		if !allnil {
			break
		}
	}
	return true

}

// 剑指 Offer 27. 二叉树的镜像
// 请完成一个函数，输入一个二叉树，该函数输出它的镜像

func mirrorTree(root *TreeNode) *TreeNode {
	var mirror func(root *TreeNode)
	mirror = func(root *TreeNode) {
		if root == nil {
			return
		}
		t := root.Left
		root.Left = root.Right
		root.Right = t
		mirror(root.Left)
		mirror(root.Right)
	}
	mirror(root)
	return root
}

// 剑指 Offer 32 - III. 从上到下打印二叉树 III
// 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，
// 第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}

	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		t := queue
		queue = []*TreeNode{}
		var level []int
		for i := 0; i < len(t); i++ {
			if t[i].Left != nil {
				queue = append(queue, t[i].Left)
			}
			if t[i].Right != nil {
				queue = append(queue, t[i].Right)
			}

			if len(ans)&1 == 0 {
				level = append(level, t[i].Val)
			} else {
				level = append([]int{t[i].Val}, level...)
			}
		}
		ans = append(ans, level)
	}
	return ans
}

// 剑指 Offer 32 - I. 从上到下打印二叉树
// 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

func levelOrder3(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}

	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		t := queue
		queue = []*TreeNode{}
		for i := 0; i < len(t); i++ {
			if t[i].Left != nil {
				queue = append(queue, t[i].Left)
			}
			if t[i].Right != nil {
				queue = append(queue, t[i].Right)
			}
			ans = append(ans, t[i].Val)
		}
	}
	return ans
}

// 剑指 Offer 32 - II. 从上到下打印二叉树 II
// 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

func levelOrder2(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}

	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		t := queue
		queue = []*TreeNode{}
		var level []int
		for i := 0; i < len(t); i++ {
			if t[i].Left != nil {
				queue = append(queue, t[i].Left)
			}
			if t[i].Right != nil {
				queue = append(queue, t[i].Right)
			}
			level = append(level, t[i].Val)
		}
		ans = append(ans, level)
	}

	return ans
}

func levelOrder1(root *TreeNode) [][]int {
	var ans [][]int
	var orderTree func(root *TreeNode, level int)
	orderTree = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(ans) > level {
			ans[level] = append(ans[level], root.Val)
		} else {
			ans = append(ans, []int{root.Val})
		}
		orderTree(root.Left, level+1)
		orderTree(root.Right, level+1)
	}
	orderTree(root, 0)
	return ans
}

// 剑指 Offer 04. 二维数组中的查找
// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，
// 每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，
// 输入这样的一个二维数组和一个整数，判断数组中是否含有该整数

func findNumberIn2DArray(matrix [][]int, target int) bool {
	var row, col = len(matrix) - 1, 0
	for row >= 0 && col <= len(matrix[0])-1 {
		if matrix[row][col] > target {
			row--
		} else if matrix[row][col] < target {
			col++
		} else {
			return true
		}
	}
	return false
}

// 剑指 Offer 11. 旋转数组的最小数字
// 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
// 给你一个可能存在重复元素值的数组numbers，
// 它原来是一个升序排列的数组，并按上述情形进行了一次旋转。请返回旋转数组的最小元素。
// 例如，数组[3,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，该数组的最小值为1

func minArray(numbers []int) int {
	target := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < target {
			return numbers[i]
		}
	}
	return target
}

// 剑指 Offer 50. 第一个只出现一次的字符
// 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

func firstUniqChar(s string) byte {
	var arr [26]int
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if arr[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

// 剑指 Offer 53 - II. 0～n-1中缺失的数字
// 一个长度为n-1的递增排序数组中的所有数字都是唯一的，
// 并且每个数字都在范围0～n-1之内。
// 在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
// 二分法永远滴神😘😘😘😘😘😘😘😘😘😘😘😘😘😘😘😘

func missingNumber(nums []int) int {
	var left, right = 0, len(nums) - 1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == middle {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

// 剑指 Offer 53 - I. 在排序数组中查找数字 I
// 统计一个数字在排序数组中出现的次数
// 二分法永远滴神😘😘😘😘😘😘😘😘😘😘😘😘😘😘😘😘

func search(nums []int, target int) int {
	var binarySearch = func(nums []int, target int) int {
		var l, r = 0, len(nums) - 1
		for l <= r {
			m := l + (r-l)/2
			if nums[m] <= target {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		return l
	}

	return binarySearch(nums, target) - binarySearch(nums, target-1)
}

// 剑指 Offer 03. 数组中重复的数字
// 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
// 数组中某些数字是重复的，但不知道有几个数字重复了，
// 也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i == nums[i] {
			continue
		}
		if nums[i] == nums[nums[i]] {
			return nums[i]
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		if i != nums[i] {
			i--
		}
	}

	return -1
}

// 剑指 Offer 58 - II. 左旋转字符串
// 字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
// 请定义一个函数实现字符串左旋转操作的功能。
// 比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

func reverseLeftWords(s string, n int) string {
	var right strings.Builder

	for i := n; i < len(s)+n; i++ {
		right.WriteByte(s[i%len(s)])
	}

	return right.String()
}

// 剑指 Offer 05. 替换空格
// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

func replaceSpace(s string) string {
	if s == "" {
		return s
	}
	var str strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			str.WriteString("%20")
		} else {
			str.WriteByte(s[i])
		}
	}
	return str.String()
}

// 剑指 Offer 35. 复杂链表的复制
// 请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，
// 每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 每个节点的下一个节点 就是 他的拷贝节点
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}
	// Next.Random = Random.Next
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	// 分离两个链表
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return headNew
}

// 自己想到的方法比较丑陋。。。🐶🐶🐶🐶🐶🐶🐶🐶🐶
// func copyRandomList(head *Node) *Node {
// 	var oldMap = make(map[*Node]*Node)
//
// 	var copyNode = &Node{
// 		Val:    0,
// 		Next:   nil,
// 		Random: nil,
// 	}
//
// 	var tmpCopy = copyNode
//
// 	for tmpHead := head; tmpHead != nil; tmpHead = tmpHead.Next {
// 		tmpCopy.Next = &Node{
// 			Val:    tmpHead.Val,
// 			Next:   nil,
// 			Random: nil,
// 		}
// 		oldMap[tmpHead] = tmpCopy.Next
// 		tmpCopy = tmpCopy.Next
// 		tmpHead = tmpHead.Next
// 	}
//
// 	// 处理random
// 	tmpCopy2 := copyNode.Next
// 	for head != nil {
// 		if head.Random == nil {
// 			tmpCopy2.Random = nil
// 		} else {
// 			tmpCopy2.Random = oldMap[head]
// 		}
// 		head = head.Next
// 		tmpCopy2 = tmpCopy2.Next
// 	}
//
// 	return copyNode.Next
// }

// 剑指 Offer 24. 反转链表
// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	for ; head != nil; head = next {
		next = head.Next
		head.Next = prev
		prev = head
	}
	return prev
}

// 剑指 Offer 06. 从尾到头打印链表
// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

func reversePrint(head *ListNode) []int {
	var ans []int
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	return ans
}

// 剑指 Offer 30. 包含min函数的栈
// 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的
// min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

type MinStack struct {
	stack    *list.List
	minStack *list.List // 辅助栈 在栈顶记录每次 push 的最小值
}

func Constructor2() MinStack {
	return MinStack{
		stack:    list.New(),
		minStack: list.New(),
	}
}

func (this *MinStack) Push(x int) {
	this.stack.PushBack(x)
	if this.minStack.Len() == 0 {
		this.minStack.PushBack(x)
	} else {
		minVal := this.minStack.Back()
		if minVal.Value.(int) > x {
			this.minStack.PushBack(x)
		} else {
			this.minStack.PushBack(minVal.Value.(int))
		}
	}

}

func (this *MinStack) Pop() {
	val := this.stack.Back()
	minVal := this.minStack.Back()
	this.minStack.Remove(minVal)
	this.stack.Remove(val)
}

func (this *MinStack) Top() int {
	if this.stack.Len() > 0 {
		return this.stack.Back().Value.(int)
	}
	return -1
}

func (this *MinStack) Min() int {
	if this.minStack.Len() > 0 {
		return this.minStack.Back().Value.(int)
	}
	return -1
}

//  写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
// F(0) = 0,   F(1) = 1
// F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
func fib(n int) int {
	if n < 2 {
		return n
	}
	var curr int

	p2, p1 := 0, 1

	for i := 2; i <= n; i++ {
		curr = p1 + p2
		p1, p2 = curr, p1
	}

	return curr
}

// CQueue
// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
// 分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead操作返回 -1 )
// type CQueue struct {
//	Value []int
// }
// func Constructor() CQueue {
//	return CQueue{Value: make([]int, 0)}
// }
// func (this *CQueue) AppendTail(value int) {
//	this.Value = append(this.Value, value)
// }
// func (this *CQueue) DeleteHead() int {
//	if len(this.Value) == 0 {
//		return -1
//	}
//	head := this.Value[0]
//	this.Value = this.Value[1:]
//	return head
// }

type CQueue struct {
	Stack1, Stack2 *list.List
}

//  func Constructor() CQueue {
// 	return CQueue{
// 		Stack1: list.New(),
// 		Stack2: list.New(),
// 	}
// }

func (this *CQueue) AppendTail(value int) {
	this.Stack1.PushBack(value)
}
func (this *CQueue) DeleteHead() int {
	if this.Stack2.Len() == 0 {
		for this.Stack1.Len() > 0 {
			this.Stack2.PushBack(this.Stack1.Remove(this.Stack1.Back()))
		}
	}
	if this.Stack2.Len() > 0 {
		e := this.Stack2.Back()
		this.Stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}
