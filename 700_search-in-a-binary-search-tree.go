/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }

    if val == root.Val {
        return root
    }

    if val < root.Val {
        return searchBST(root.Left, val);
    }

    if val > root.Val {
        return searchBST(root.Right, val);
    }

    return nil;
}
