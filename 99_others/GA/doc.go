//Package GA 遗传算法的Golang实现
//遗传算法的计算流程 编码->获得初始群体->计算适应值->产生新的群体(选择、交叉、变异等)->满足要求->解码->获得近似解
//例子：求解函数 f(x) = x + 10*sin(5*x) + 7*cos(4*x) 在区间[0,9]的最大值。
//假设求解精度为3，那么需要求解的值可能为 0.000-9.000 可能性在2^13-2^14这个范围之内
//使用一个14为的二进制字符串即可表示，最终的求解域
//解码 f(x), x∈[lower_bound, upper_bound] x = lower_bound + decimal(chromosome)×(upper_bound-lower_bound)/( 2 ^chromosome_size- 1 )
package GA
