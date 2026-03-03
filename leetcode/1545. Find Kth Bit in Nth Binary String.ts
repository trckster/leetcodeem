function isKthBitOne(n: number, k: number): boolean {
    if (n == 1) {
        return false
    }

    const middle = 2 ** (n - 1)

    if (k == middle) {
        return true
    }

    if (k < middle) {
        return isKthBitOne(n - 1, k)
    }

    const reversedK = (2 ** n) - (k - middle)

    return !isKthBitOne(n - 1, reversedK - middle)
}

function findKthBit(n: number, k: number): string {
    return isKthBitOne(n, k) ? "1" : "0"
};

findKthBit(4, 11)
