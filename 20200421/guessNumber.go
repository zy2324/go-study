package main

//我们正在玩一个猜数字游戏。 游戏规则如下：
//我从 1 到 n 选择一个数字。 你需要猜我选择了哪个数字。
//每次你猜错了，我会告诉你这个数字是大了还是小了。
//你调用一个预先定义好的接口 guess(int num)，它会返回 3 个可能的结果（-1，1 或 0）：
//
//-1 : 我的数字比较小
// 1 : 我的数字比较大
//  0 : 恭喜！你猜对了！
//  示例 :
//
//  输入: n = 10, pick = 6
//  输出: 6

// 个人思路
// 首先是二分法
// 其次题目没看懂

import "fmt"

/**
* Forward declaration of guess API.
 * @param  num   your guess
  * @return 	     -1 if num is lower than the guess number
   *			      1 if num is higher than the guess number
    *               otherwise return 0
     * func guess(num int) int;
*/

// by ly爱编程
func guessNumber(n int) int {
	a, b := 1, n
	for a < b {
		mid := a + (b-a)/2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == 1 {
			a = mid + 1
		} else {
			b = mid - 1
		}
	}
	return a
}

func main() {
	fmt.Println("vim-go")
}

// 收获：求mid值得方法再次见到 (a+b)/2 = a-a/2+b/2 = a + b/2 - a/2
