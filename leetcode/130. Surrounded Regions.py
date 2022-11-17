class Solution:
    def __init__(self):
        self.n = 0
        self.m = 0

    def process(self, board, i, j):
        if i < 0 or i + 1 > self.n or j < 0 or j + 1 > self.m or board[i][j] != '?':
            return

        board[i][j] = 'O'
        self.process(board, i + 1, j)
        self.process(board, i, j + 1)
        self.process(board, i - 1, j)
        self.process(board, i, j - 1)

    def solve(self, board: List[List[str]]):
        self.n = len(board)
        self.m = len(board[0])
        for i in range(self.n):
            for j in range(self.m):
                if board[i][j] == 'O':
                    board[i][j] = '?'

        for i in range(self.n):
            self.process(board, i, 0)
            self.process(board, i, self.m - 1)

        for i in range(self.m):
            self.process(board, 0, i)
            self.process(board, self.n - 1, i)

        for i in range(self.n):
            for j in range(self.m):
                if board[i][j] == '?':
                    board[i][j] = 'X'