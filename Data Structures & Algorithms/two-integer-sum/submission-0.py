class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        answer = []
        for i in range(len(nums)):
            for j in range(len(nums)):
                if i != j:
                    if target == nums[i]+nums[j]:
                        answer.append(i)
                        answer.append(j)
                        break
        nums = set(answer)
        return list(nums)
        