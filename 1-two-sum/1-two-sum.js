/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    let tmp = new Map()
    for (let i = 0; i< nums.length;i++) {
        const value = nums[i]
        const paired = target - value
        if (tmp.has(paired)){
            return [tmp.get(paired), i]
        }else{
            tmp.set(nums[i],i)
        }
        //console.log(tmp)
    }
    
    return []
};