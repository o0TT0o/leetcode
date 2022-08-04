/**
 * @param {number[]} prices
 * @return {number}
 */
//[7,1,5,3,6,4]
// start: 1
// 1 - 5
// 1 - 3
// 1 - 6

// tough caces:
// [2,1,2,1,0,1,2] - expected: 2
// [2,4,1] - expected: 2
var maxProfit = function(prices) {
	let minPrices = prices[0]
    let maxProfile = 0
    for (let i = 0; i< prices.length; i++){
        const currentPrice = prices[i]
        const currentProfile = prices[i] - minPrices
        
        if ( currentPrice < minPrices  ){
            //console.log("Price: ", currentPrice , " ---> ", minPrices)
            minPrices = currentPrice
        }
        
        if ( currentProfile > maxProfile  ){
            //console.log("Profile: ", currentProfile , " ---> ", maxProfile)
            maxProfile = currentProfile
        }
        
    }

	return maxProfile;
};