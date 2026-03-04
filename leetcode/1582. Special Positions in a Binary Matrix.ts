function fill(n: number): number[] {
    const arr = []
    while (n--) {
        arr.push(0)
    }
    return arr
}

function numSpecial(mat: number[][]): number {
    const rows = fill(mat.length), columns = fill(mat[0].length)

    for (let i = 0; i < mat.length; i++) {
        for (let j = 0; j < mat[i].length; j++) {
            if (mat[i][j] === 1) {
                rows[i]++
                columns[j]++
            }
        }
    }

    let answer = 0

    for (let i = 0; i < mat.length; i++) {
        for (let j = 0; j < mat[i].length; j++) {
            if (mat[i][j] === 1) {
                if (rows[i] === 1 && columns[j] === 1) {
                    answer++
                }
            }
        }
    }

    return answer
};
