/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {boolean}
 */
//tough test cases:
//  [5,4,6,null,null,3,7]  is false
//      [5,4,6,null,null,null,7] is true
//      [4,3,5,null,null,6,7] is true
//      [5,4,6,null,null,5,7] is false
var isValidBST = function(root) {
    
    
    function valid(node, left, right){
        if ( node === null ){
            return true
        }
        
        if ( node.val <= left || node.val >= right ){
            return false
        }
        
        return ( valid(node.left, left,  node.val) && valid( node.right, node.val, right) ) 
    }
    
    return valid(root, -Infinity, Infinity)
    
};