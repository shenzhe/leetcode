package code

func Execuate_o33(postorder []int) bool {
	return verifyPostorder(postorder)
}

func verifyPostorder(postorder []int) bool {
	l := len(postorder)
	if l <= 2 {
		return true
	}

	rootVal := postorder[l-1]

	//遍历数组，找到第一个大于rootVal的值
	idx := 0
	for k, v := range postorder {
		if v >= rootVal { //跳出循环
			idx = k
			break
		}
	}

	//后面的数组应该都大于rootVal
	for i := idx; i < l-1; i++ {
		if postorder[i] < rootVal {
			return false
		}
	}

	//判断左数
	res := verifyPostorder(postorder[0:idx])

	//左树正确，判断右树
	if res {
		res = verifyPostorder(postorder[idx : l-1])
	}
	return res

}
