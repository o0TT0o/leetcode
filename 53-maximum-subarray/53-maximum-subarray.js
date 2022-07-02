/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function(nums) {
    if ( nums.length <= 1){
        return nums[0]
    }
    let fullMax = null
    let index = 0
    let current = nums[index]
    let best = nums[index]
    index++
    while(index<nums.length){
        if ( current + nums[index] > 0 ){
            current  = current + nums[index] 
        }else{
            if ( best >= 0 ){
                current = 0
            }
        }
        if ( best >= 0 ){
            if ( current > best){
                best = current
            }
        }else{ 
            if ( nums[index] > best){
                best = nums[index]
            }
            if ( nums[index] < 0 ){
                current = 0
            }else{
                current = nums[index]
            }
        }
        //console.log("index: ",index)
        //console.log("nums[index]: ",nums[index])
        //console.log("current: ",current)
        //console.log("best: ",best)
        index++
    }
    return best
    
};