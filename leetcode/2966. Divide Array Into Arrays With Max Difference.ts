function divideArray(nums: number[], k: number): number[][] {
    nums.sort((a, b) => a - b)
    const dividedNums = []
    for (let i = 0; i + 2 < nums.length; i += 3) {
        if (nums[i + 2] - nums[i] > k) {
            return []
        }

        dividedNums.push([nums[i], nums[i + 1], nums[i + 2]])
    }

    return dividedNums
}