class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        seen = set()
        index = 0

        for i in range(0,len(nums)):
            if nums[i] not in seen:
                seen.add(nums[i])
                nums[index] = nums[i]
                index +=1
        return index
        