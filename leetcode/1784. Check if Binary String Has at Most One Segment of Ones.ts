function checkOnesSegment(s: string): boolean {
    let changes = 0
    for (let i = 1; i < s.length; i++) {
        if (s[i] !== s[i - 1]) {
            changes++
        }

        if (changes > 1) {
            return false
        }
    }

    return true
};
