/**
 * @param {number[]} nums
 * @return {number}
 */

// [2,7,9,3,1]

// [2,9,10,3,4,5]

// [0,1,0,1,9,10,11]

// [1,2]
var rob = function(nums) {
    let rob1 = 0
    let rob2 = 0
    
    for(let i =0 ;i< nums.length; i++){
        let sum = Math.max( nums[i] + rob1 , rob2)
        rob1 = rob2
        rob2 = sum
    }
    
    return rob2
};