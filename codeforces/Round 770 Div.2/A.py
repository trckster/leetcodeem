def is_palindrome(s):
    for i in range(len(s) // 2):
        if s[i] != s[len(s) - i - 1]:
            return False

    return True


n = int(input())
for i in range(n):
    n, k = map(int, input().split())
    s = input()

    if k == 0 or is_palindrome(s):
        print(1)
    else:
        print(2)
