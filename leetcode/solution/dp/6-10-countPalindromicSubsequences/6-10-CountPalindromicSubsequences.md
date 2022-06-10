# 动态规划

由于长字符串会依赖短字符串的回文序列数量，所以我们可以采用动态规划来实现。
设`dp[i][j]`表示字符串从`i`到`j`的回文序列个数，我们可以将长字符串看作短字符串左右加上两个字符
于是我们有`s[i,j] = s[i] + s [i+1,j-1] + s[j]`，如："bccb" 可以看作 "cc"两边分别加上"b"，此时我们分情况进行讨论
(1)若`s[i] == s[j]`，相当于我们给`s[i+1,j-1]` 左右加上两个相同的字符，然后我们计算回文序列的个数

*   ①`s[i+1,j-1]`中没有字符和`s[i]`相等
    设有字符串"bcb"，则"bcb"的回文子序列是：b、c、bb、bcb
    若两边加上相同的字符，相当于给"bcb"的回文子序列左右个加一个相同字符，仍然构成回文子序列
    假设我们给"bcb"左右加一个字符"a"，则相当于给"bcb"的子序列都左右加一个字符可构成新的回文子序列
    ![image.png](https://pic.leetcode-cn.com/1654823325-bgUZna-image.png)
    在加上"a"(字符本身就是一个回文子序列)和"aa"(两个相同字符的回文子序列)
    所以此时**dp\[i\]\[j\] = 2dp\[i+1\]\[j-1\] + 2**（本身的4个+新生成的4个+2个单独生成的）
*   ②`s[i+1,j-1]`中有一个字符和s\[i\]相等
    假设有一个字符相等，则之前已经记录了此单字符的回文子序列(只能加上"aa"，不能加"a")
    所以此时**dp\[i\]\[j\] = 2dp\[i+1\]\[j-1\] + 1**（本身的4个+新生成的4个+1个单独生成的）
*   ③`s[i+1,j-1]`中有两个及以上字符和s\[i\]相等
    若有两个及以上的字符,则我们需要找到其位置，并删掉重复计算的回文子序列，并且两个单独的之前也已经计算。
    假设有字符串"dabcbad"，我们向两边加入字符"a"
    则此时的"a"字符会和中间的"bcb"组成重复的回文子序列，因为之前已经有"a"和"bcb"组成回文子序列
    ![image.png](https://pic.leetcode-cn.com/1654824331-GObSjO-image.png)
    ![image.png](https://pic.leetcode-cn.com/1654824427-OHwUlw-image.png)
    (2)若`s[i] != s[j]`，则我们给之前任何一个回文子序列左右加上s\[i\]和s\[j\]都不能组成回文子序列，只能单独计算
    ![image.png](https://pic.leetcode-cn.com/1654819613-movDBj-image.png)
    综上所述，状态转移方程为：
    ![image.png](https://pic.leetcode-cn.com/1654819758-yvOLas-image.png)
    **代码如下：**

 ```
public int countPalindromicSubsequences(String s) {
        int mod = 1000000007;
        int n = s.length();
        int[][] dp = new int[n][n];
        //一个单字符是一个回文子序列
        for (int i = 0; i < n; i++) {
            dp[i][i] = 1;
        }
        //从长度为2的子串开始计算
        for (int len = 2; len <= n; len++) {
            //挨个计算长度为len的子串的回文子序列个数
            for (int i = 0; i + len <= n; i++) {
                int j = i + len - 1;
                //情况(1) 相等
                if (s.charAt(i) == s.charAt(j)) {
                    int left = i + 1;
                    int right = j - 1;
                    //找到第一个和s[i]相同的字符
                    while (left <= right && s.charAt(left) != s.charAt(i)) {
                        left++;
                    }
                    //找到第一个和s[j]相同的字符
                    while (left <= right && s.charAt(right) != s.charAt(j)) {
                        right--;
                    }
                    if (left > right) {
                        //情况① 没有重复字符
                        dp[i][j] = 2 * dp[i + 1][j - 1] + 2;
                    } else if (left == right) {
                        //情况② 出现一个重复字符
                        dp[i][j] = 2 * dp[i + 1][j - 1] + 1;
                    } else {
                        //情况③ 有两个及两个以上
                        dp[i][j] = 2 * dp[i + 1][j - 1] - dp[left + 1][right - 1];
                    }
                } else {
                    //情况(2) 不相等
                    dp[i][j] = dp[i][j - 1] + dp[i + 1][j] - dp[i + 1][j - 1];
                }
                //处理超范围结果
                dp[i][j] = (dp[i][j] >= 0) ? dp[i][j] % mod : dp[i][j] + mod;
            }
        }
        return dp[0][n - 1];
    }
```

**复杂度分析：**
时间复杂度：O(n^2)
空间复杂度：O(n^2)
来源:https://leetcode.cn/problems/count-different-palindromic-subsequences/solution/tong-ji-butong-by-jiang-hui-4-q5xf/