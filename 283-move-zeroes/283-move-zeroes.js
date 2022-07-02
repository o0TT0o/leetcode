/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var moveZeroes = function(nums) {
    let zeroCount = 0
    let nonZeroArr = []
    let index = 0
    let count = 0
    while(count < nums.length){
        if ( nums[index] === 0){
            nums.splice(index,1)
            nums.push(0)
            index--
        }
        //console.log(nums)
        index++
        count++
    }
};