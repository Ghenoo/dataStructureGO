package main

type TreeNode struct { // Árvore binária
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	treeNode := TreeNode{}
	treeNode.Data = "Raiz"
	treeNode.Left = &TreeNode{Data: "Esquerda"}
	treeNode.Right = &TreeNode{Data: "Direita"}
	treeNode.Left.Left = &TreeNode{Data: "Esquerda da esquerda"}
	treeNode.Left.Right = &TreeNode{Data: "Direita da esquerda"}
	treeNode.Right.Left = &TreeNode{Data: "Esquerda da direita"}
	treeNode.Right.Right = &TreeNode{Data: "Direita da direita"}

	// E assim infinitamente ou até serem nulos.
}