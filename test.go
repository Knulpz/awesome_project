package main

type TreeNode struct {
	Val int
	    Left *TreeNode
	    Right *TreeNode
}

func main(){

}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil{
		return res
	}
	//借助队列实现层序上的一个接一个
	queue := make([]*TreeNode, 0)
	//记录一层的数据
	tmp_arr := make([]int, 0)
	queue = append(queue, root)
	//外层循环代表每层
	for len(queue) > 0{
		//内层循环遍历该层的每个节点
		for i := 0; i < len(queue); i++{
			//每次循环出队首元素（但是去世之前留下了自己的孩子）
			root = queue[0]
			queue = queue[1:]
			//先进左孩子
			if root.Left != nil{
				queue = append(queue,root.Left)
			}
			//再进右孩子
			if root.Right != nil{
				queue = append(queue,root.Right)
			}
			tmp_arr = append(tmp_arr,root.Val)
		}
		res = append(res, tmp_arr)
		tmp_arr = make([]int, 0)
	}
	return res
}




