package code

import (
	"strconv"
	"strings"
)

type Codec struct {
	strArr   []string
	nilStr   string
	splitStr string
}

func NewCodec() Codec {
	return Codec{
		nilStr:   "#",
		splitStr: "|",
	}
}

// Serializes a tree to a single string.
func (cc *Codec) Serialize(root *TreeNode) string {
	if root == nil {
		return cc.nilStr
	}
	str := []string{}
	cc.strArr = str
	cc.real_serialize(root)
	return strings.Join(cc.strArr, cc.splitStr)
}

func (cc *Codec) real_serialize(root *TreeNode) {
	if root == nil {
		cc.strArr = append(cc.strArr, cc.nilStr)
		return
	}
	// fmt.Println(root.Val)
	cc.strArr = append(cc.strArr, strconv.Itoa(root.Val))
	// fmt.Println(cc.str)
	cc.real_serialize(root.Left)
	cc.real_serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (cc *Codec) Deserialize(data string) *TreeNode {
	// fmt.Println(data)
	if data == "" || data == "#" {
		return nil
	}
	cc.strArr = strings.Split(data, cc.splitStr)
	// fmt.Println(cc.str)
	root := cc.real_deserialize()
	return root
}

func (cc *Codec) real_deserialize() *TreeNode {
	// fmt.Println(cc.str)
	if len(cc.strArr) == 0 {
		return nil
	}

	first := cc.strArr[0]
	// fmt.Println("first", first)
	cc.strArr = cc.strArr[1:]

	if first == cc.nilStr {
		return nil
	}

	v, _ := strconv.Atoi(first)
	root := &TreeNode{
		Val: v,
	}
	root.Left = cc.real_deserialize()
	root.Right = cc.real_deserialize()
	return root
}
