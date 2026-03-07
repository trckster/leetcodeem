function countBestFlipping(s: string): [number, number] {
    let flipsNeededWhenStartingFromZero = 0
    let flipsNeededWhenStartingFromOne = 0

    for (let i = 0; i < s.length; i++) {
        const c = s[i]

        if (i % 2 === 0) {
            if (c === '0') {
                flipsNeededWhenStartingFromOne++
            } else {
                flipsNeededWhenStartingFromZero++
            }
        } else {
            if (c === '0') {
                flipsNeededWhenStartingFromZero++
            } else {
                flipsNeededWhenStartingFromOne++
            }
        }
    }

    return [flipsNeededWhenStartingFromZero, flipsNeededWhenStartingFromOne]
}

function minFlips(s: string): number {
    let [fromZero, fromOne] = countBestFlipping(s)
    let answer = Math.min(fromZero, fromOne)

    for (const c of s) {
        [fromZero, fromOne] = [fromOne, fromZero]

        if (c === '0') {
            fromZero--

            if (s.length % 2 == 0) {
                fromZero++
            } else {
                fromOne++
            }
        } else {
            fromOne--

            if (s.length % 2 == 0) {
                fromOne++
            } else {
                fromZero++
            }
        }

        console.log(fromZero, fromOne)
        answer = Math.min(answer, fromZero, fromOne)
    }

    return answer
}
