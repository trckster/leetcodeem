function createCounter(n: number): () => number {
    let x = n - 1

    return function () {
        x++
        return x
    }
}
