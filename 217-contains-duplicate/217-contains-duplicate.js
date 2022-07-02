/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
    let nonDuplicateArr = []
    for (let i = 0 ; i<nums.length; i++){
        let exist = false
        for (let j = 0 ; j<nonDuplicateArr.length; j++){
            if ( nums[i] === nonDuplicateArr[j] ){
                return true
            }
        }
        nonDuplicateArr.push(nums[i])
    }
    //console.log(nonDuplicateArr)
    return false
};