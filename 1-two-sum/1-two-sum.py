class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        tmp = {}
        index = 0
        for i in nums:
            wanted = target-i
            if i in tmp:
                return [tmp[i], index]
            else:
                tmp[wanted] = index
            index=index+1
            #print(tmp)
        return []
            