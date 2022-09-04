class Solution:
    # test cases:
    #   "bvvf"
    #   "au"
    #   "abba"
    def lengthOfLongestSubstring(self, s: str) -> int:
        if len(s) == 0:
            return 0
        
        dic_substring = {}
        index_substring = 0
        index_start = 0
        longest = 1
        for char in s:
            if char in dic_substring:
                            #index_substring -1 - index_start +1
                distance = index_substring - index_start
                
                if distance > longest:
                    longest = distance
                    #print("start: {:n} , stop: {:n}, longest: {:n}".format(index_start, index_substring, longest))
                
                index_start = max(index_start, dic_substring[char] + 1)
                #print("index_start: ",index_start)
                del dic_substring[char]
                    
            dic_substring[char] = index_substring
            index_substring +=1
            #print(dic_substring)
            
        distance = index_substring - index_start
        if distance > longest:
            longest = distance
            #print("start: {:n} , stop: {:n}, longest: {:n}".format(index_start, index_substring, longest))
            
        return longest