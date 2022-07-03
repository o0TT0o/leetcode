/**
 * @param {number[]} nums
 * @param {number} k
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var rotate = function(nums, k) {
    if ( nums.length === 1){
        return
    }
    
    if ( nums.length === 2){
        //console.log(k%2)
        if ( k%2 === 0){
            return
        }else{
            nums.reverse()
            return
        }
    }
    
    let tmpArr = []
    
    if ( k > nums.length ){
        k = k-nums.length 
    }
    
    
    const count = k
    

    while(k>0){
        tmpArr.push(nums[nums.length-k]) //5,6,7
        k--
    }
    
    k = count
    let index = nums.length-1-k
    //console.log(index, nums)
    while(index >=0){
        nums[index + k] = nums[index]
        index--
    }
    
    
    for(let i = 0 ; i< tmpArr.length ; i++){
        nums[i] = tmpArr[i]
    }
    
};