function minOperations(s: string): number {
    let startWithZero = 0, startWithOne = 0

    let expectedChar = '0'
    for (const c of s) {
        if (c === expectedChar) {
            startWithOne++
        } else {
            startWithZero++
        }

        expectedChar = expectedChar === '0' ? '1' : '0'
    }

    return Math.min(startWithZero, startWithOne)
};
