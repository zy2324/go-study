package main

//给定一个整数，写一个函数来判断它是否是 3 的幂次方。
//
//示例 1:
//
//输入: 27
//输出: true
//示例 2:
//
//输入: 0
//输出: false
//示例 3:
//
//输入: 9
//输出: true
//示例 4:
//
//输入: 45
//输出: false
//进阶：
//你能不使用循环或者递归来完成本题吗？

// 个人思路
// 忘了之前做的2的幂次方怎么写来着，应该就是照搬吧

// by comment Ben
func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}

	nums := strconv.FormatInt(int64(n), 3)
	return nums[0:1] == "1" && strings.Count(nums, "0") == len(nums)-1
}

// 收获：strconv.FormatInt 转为几进制；3的幂次，对应的3进制的特征，判断是否第一位为1其他位为0，符合条件则为3的幂。
