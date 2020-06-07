学习笔记


###第七课：泛型递归

递归三要素模板：

一、找终止条件

二、找返回值

三、找本级递归应该做什么

比如求二叉树的高度：
	root
left		 right
* 终止条件为当前树节点为nil，为空
* 返回值为当前深度+1（左右子树+root）
* 本级递归要做的是找出左右节点的最大值，并将其+1传给下一级


###第八课：分治、回溯

分解成子问题，然后下沉解决每个子问题，最后再合并处理答案。类似递归模板

分治模板：
```$xslt
    func divide_couquer(proble interface{}, param1 interface{}, param2 interface{}.) {
    	// recursion terminator
    	if problem == nil {
    		print_result
    		return
    	}
    	// prepare data
    	data := prepare_data(proble)
    	subproblems := split_proble(problem, data)
    	// conquer subproblems
    	subresult1 := divide_conquer(subproblems[0], p1)
    	subresult2 := divide_conquer(subproblems[1], p2)
    	subresult3 := divide_conquer(subproblems[2], p3)
    	// process and generate the final result
    	result := process_result(subresult1, subresult2, subresult3)
    	// revert the current level states
    
    }
```

回溯是采用试错的思想，它尝试分布去解决一个问题。当通过尝试发现现有分布答案不正确时，将取消上一步甚至上几步的计算，再通过其他的可能分步再次尝试寻找为的答案。

一般回溯的时间复杂度为N的指数级.

回溯模板：
```$xslt
    result = []
    def backtrack(路径, 选择列表):
	if 满足结束条件:
		result.add(路径)
	return

	for 选择 in 选择列表:
		做选择
		backtrack(路径, 选择列表)
		撤销选择
```