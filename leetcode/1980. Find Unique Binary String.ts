function getRandomBinaryString(length: number): string {
    let result = ''

    while (length--) {
        if (Math.random() > 0.5) {
            result += '1'
        } else {
            result += '0'
        }
    }

    return result
}

function findDifferentBinaryString(nums: string[]): string {
    const s = new Set(nums)

    while (true) {
        const newNum = getRandomBinaryString(nums.length)

        if (!s.has(newNum)) {
            return newNum
        }
    }
}
