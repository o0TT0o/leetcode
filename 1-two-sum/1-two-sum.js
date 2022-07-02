/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    if (nums.length < 0){
        return []
    }
    
    let index1=0
    let index2=0
    while ( nums[index1] !== null){
        index2 = index1 +1 
        while ( nums[index2] != null){
            if ( nums[index1] + nums[index2] === target){
                return [index1, index2]
            }
            index2++
        }
        index1++
    }
    return []
};