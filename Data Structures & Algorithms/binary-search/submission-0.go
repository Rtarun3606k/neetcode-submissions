func search(nums []int, target int) int {
r := len(nums)-1
l:= 0

for l<=r{
    mid := int(uint(l+r)>>1)
    if nums[mid]==target{
        return mid
    }else if nums[mid]>target{
        r = mid-1
    }else{
        l = mid+1
    }
}

return -1


}
