class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        diff_map = {}
        result = []
        for i, n in enumerate(nums):
            diff = target - n
            print(diff)
            if diff in diff_map:
                result = [diff_map[diff], i]
                
            diff_map[n] = i
        return result
            