/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function(n) {
    let cache = {}
    function climbing(n, current){
        //console.log("n : ",n , " current : ", current)
        if(current === n){
            return 1
        }
        if(current >= n){
            return 0
        }
        if( ! (current in cache) ){
            cache[current] = climbing(n, current+1) + climbing(n, current+2)
        }
        return cache[current]
    }
    
    return climbing(n, 0)
};