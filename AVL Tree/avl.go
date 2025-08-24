package main

import "fmt"

//Treenode represents a node in the BST
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

//BST represents the binary search tree
type BST struct {
	Root *TreeNode
}

//Insert  adds a new value to BST
func (bst *BST) Insert(value int) {
	if bst.Root == nil {
		bst.Root = &TreeNode{Value: value}
	} else {
		bst.Root.insert(value)
	}
}

//insert is a helper method for inserting value(recursive)
func (node *TreeNode) insert(value int) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &TreeNode{Value: value}
		} else {
			node.Left.insert(value)
		}
	} else if value > node.Value {
		if node.Right == nil {
			node.Right = &TreeNode{Value: value}
		} else {
			node.Right.insert(value)
		}
	}
	// If value == node.value, we don't insert duplicate
}

//InOrderTraversal performs in-order traversal(left root right)
func (bst *BST) InOrderTraversal() []int {
	var result []int
	if bst.Root != nil {
		bst.Root.inOrder(&result)
	}
	return result
}

func (node *TreeNode) inOrder(result *[]int) {
	if node.Left != nil {
		node.Left.inOrder(result)
	}
	*result = append(*result, node.Value)
	if node.Right != nil {
		node.Right.inOrder(result)
	}
}

//PreOrderTraversal performs pre-order traversal (root, left, right)
func (bst *BST) PreOrderTraversal() []int {
	var result []int
	if bst.Root != nil {
		bst.Root.preOrder(&result)
	}
	return result
}

func (node *TreeNode) preOrder(result *[]int) {
	*result = append(*result, node.Value)
	if node.Left != nil {
		node.Left.preOrder(result)
	}
	if node.Right != nil {
		node.Right.preOrder(result)
	}
}

//PostOrderTraversal performs post-order traversal(left, Right, root)
func (bst *BST) PostOrderTraversal() []int {
	var result []int
	if bst.Root != nil {
		bst.Root.postOrder(&result)
	}
	return result
}
func (node *TreeNode) postOrder(result *[]int) {
	if node.Left != nil {
		node.Left.postOrder(result)
	}
	if node.Right != nil {
		node.Right.postOrder(result)
	}
	*result = append(*result, node.Value)
}

//Search looks for a value in the BST
func (bst *BST) Search(value int) bool {
	if bst.Root == nil {
		return false
	}
	return bst.Root.search(value)
}

func (node *TreeNode) search(value int) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	}
	if value < node.Value {
		return node.Left.search(value)
	}
	return node.Right.search(value)
}

//FindMin find the minimum value in the BST
func (bst *BST) FindMin() int {
	if bst.Root == nil {
		return -1 // or you could return an error
	}
	return bst.Root.findMin()
}
func (node *TreeNode) findMin() int {
	if node.Left == nil {
		return node.Value
	}
	return node.Left.findMin()
}

//FindMax finds the maximum value in the BST
func (bst *BST) FindMax() int {
	if bst.Root == nil {
		return -1 // or you could return an error
	}
	return bst.Root.findMax()
}

func (node *TreeNode) findMax() int {
	if node.Right == nil {
		return node.Value
	}
	return node.Right.findMax()
}

// Visualize prints a simple visual representation of the BST
func (bst *BST) Visualize() {
	if bst.Root == nil {
		fmt.Println("Tree is empty")
		return
	}
	fmt.Println("BST Structure:")
	bst.Root.visualize(0)
}

func (node *TreeNode) visualize(level int) {
	if node == nil {
		return
	}

	// Print right subtree first (so it appears on top in console)
	node.Right.visualize(level + 1)

	// Print current node
	for i := 0; i < level; i++ {
		fmt.Print("    ")
	}
	fmt.Printf("-> %d\n", node.Value)

	// Print left subtree
	node.Left.visualize(level + 1)
}
func main() {
	//create a new BST
	bst := &BST{}
	//Insert the given value
	values := []int{14, 12, 24, 10, 42, 32, 16, 43, 34}
	for _, value := range values {
		bst.Insert(value)
	}

	fmt.Println("Binary search tree operations: ")
	fmt.Println("===============================")

	//Display tree structure
	bst.Visualize()

	//Traversals
	fmt.Println("\nIn-Order traversal:", bst.InOrderTraversal())
	fmt.Println("\nPre-Order traversal:", bst.PreOrderTraversal())
	fmt.Println("\nPost-Order traversal:", bst.PostOrderTraversal())

	//Search operations
	fmt.Println("\nsearch operations")
	fmt.Println("====================")
	fmt.Printf("Is 32 in the tree? %t\n", bst.Search(32))
	fmt.Printf("Is 100 in the tree? %t\n", bst.Search(100))
	// Min/Max
	fmt.Printf("Minimum value: %d\n", bst.FindMin())
	fmt.Printf("Maximum value: %d\n", bst.FindMax())

	// Additional operations
	fmt.Println("\nAdditional operations:")
	fmt.Printf("Root value: %d\n", bst.Root.Value)
}
