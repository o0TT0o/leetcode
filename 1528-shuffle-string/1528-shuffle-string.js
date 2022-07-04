/**
 * @param {string} s
 * @param {number[]} indices
 * @return {string}
 */
var restoreString = function(s, indices) {
    let result = s.split('')
    for ( let i = 0; i<indices.length; i++){
        //console.log(indices[i],s[i])
        result[indices[i]] = s[i]
    }
    return result.join('')
};