import math


class SqrtDecomposition:
    def __init__(self, arr):
        self.arr = arr
        self.block_length = int(math.sqrt(len(self.arr)))
        self.computed = self.build_decomposition()

    def xor_arr(self, arr):
        result = 0
        for i in arr:
            result ^= i
        return result

    def build_decomposition(self):
        result = []
        for i in range(0, len(self.arr), self.block_length):
            start = i
            end = min(len(self.arr), start + self.block_length)
            result.append(self.xor_arr(self.arr[start:end]))
        return result

    def calculate(self, left, right):
        left_block = (left - 1) // self.block_length + 1
        right_block = (right + 1) // self.block_length - 1
        if right == len(self.arr) - 1:
            right_block = len(self.computed) - 1

        if left_block > right_block:
            return self.xor_arr(self.arr[left: right + 1])

        middle_answer = self.xor_arr(self.computed[left_block:right_block + 1])

        left_until = left_block * self.block_length
        left_answer = self.xor_arr(self.arr[left:left_until])

        right_start = (right_block + 1) * self.block_length
        right_answer = self.xor_arr(self.arr[right_start:right + 1])

        return self.xor_arr([left_answer, middle_answer, right_answer])


class Solution:
    def xorQueries(self, arr, queries):
        decomposition = SqrtDecomposition(arr)

        answers = []
        for query in queries:
            answers.append(decomposition.calculate(query[0], query[1]))
        return answers


s = Solution()

print(s.xorQueries([1, 3, 4, 8], [[0, 1], [1, 2], [0, 3], [3, 3]]))
print(s.xorQueries([4, 8, 2, 10], [[2, 3], [1, 3], [0, 0], [0, 3]]))
print(s.xorQueries([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12], [[10, 10]]))
print(s.xorQueries([
    123590368, 887039577, 950326013, 848125275, 42974945, 201779811, 257304363, 843437225, 630064054, 452059902,
    173967478, 124079655, 484159445, 179726595, 731100115, 259223062, 170665637, 798870803, 907332142, 110417326,
    467188111, 977925843, 406172122, 865959215, 582961757, 976270471, 878696742, 315705417, 590782505, 272097063,
    764086436, 512185692, 337648577, 107512856, 222924498, 700133576, 340823202, 42541888, 29272941, 11280895,
    844884439, 993859379, 317344241, 780055968, 640562856, 343867186, 483016894, 420121952, 336382767, 427945350,
    67607931, 68925611, 980747211, 340751466, 645798692, 489473791, 119574376, 268502913, 231062011, 843032950,
    663365081, 835284838, 956928808, 582948604, 931593268, 738938178, 503499127, 710639308, 382066319, 278191472,
    196733926, 581541006, 223163180, 329983542, 213906448, 264549796, 387044831, 87369044, 880443655, 301492451,
    96008417, 808835417, 480931370, 97161221, 700090977, 616710306, 690916441, 363873691, 244197059, 419273145,
    329407130, 44079526, 351372795, 200588773
], [
    [40, 90], [74, 83], [31, 52], [12, 81], [78, 80], [31, 45], [2, 33], [51, 55], [34, 69], [9, 18], [2, 83],
    [1, 38],
    [45, 76], [60, 69], [12, 76], [41, 45], [65, 86], [22, 44], [19, 37], [40, 59], [13, 78], [77, 86], [16, 42],
    [18, 87], [92, 92], [40, 92], [79, 91], [86, 89], [22, 60], [55, 89], [20, 26], [91, 93], [31, 47], [8, 65],
    [55, 90], [32, 66], [56, 90], [58, 58], [50, 71], [21, 54], [62, 63], [82, 88], [73, 73], [7, 52], [74, 78],
    [75, 79], [17, 50], [35, 60], [72, 84], [22, 33], [29, 91], [0, 22], [67, 77], [64, 92], [83, 90], [87, 92],
    [93, 93], [30, 70], [63, 84], [61, 75], [91, 92], [43, 78], [78, 86], [82, 85], [6, 15], [65, 85], [47, 57],
    [65, 89], [26, 65], [38, 82], [38, 92], [70, 81], [21, 80], [76, 78], [15, 92], [65, 83], [48, 59], [19, 51],
    [54, 67], [72, 73], [63, 84], [50, 90], [44, 68], [65, 92], [21, 90], [93, 93], [67, 69], [32, 72], [12, 46],
    [10, 27], [78, 80], [56, 71], [48, 75], [50, 60], [41, 69], [16, 78], [27, 42], [55, 93], [68, 81], [15, 52],
    [35, 50], [48, 83], [0, 24], [67, 81], [74, 74], [3, 83], [77, 84], [33, 46], [22, 58], [18, 87], [25, 59],
    [39, 76], [32, 85], [87, 90], [21, 90], [45, 76], [53, 60], [26, 75], [26, 62], [13, 16], [15, 87], [1, 70],
    [37, 80], [92, 92], [9, 73], [47, 83], [39, 66], [64, 85], [45, 93], [67, 77], [0, 15], [56, 84], [44, 63],
    [69, 81], [43, 93], [93, 93], [14, 62], [48, 61], [71, 71], [35, 75], [76, 92], [25, 84], [76, 92], [52, 71],
    [89, 90], [57, 90], [25, 71], [67, 67], [86, 92], [74, 87], [51, 72], [79, 88], [93, 93], [63, 86], [31, 71],
    [83, 87], [80, 80], [52, 92], [19, 69], [34, 88], [22, 31], [77, 77], [44, 60], [90, 93], [87, 91], [38, 47],
    [59, 75], [62, 72], [59, 91], [5, 39], [65, 68], [75, 88], [8, 53], [8, 92], [56, 57], [21, 85], [90, 91],
    [88, 88], [51, 91], [88, 90], [77, 86], [26, 93], [26, 56], [42, 59], [8, 17], [89, 93], [84, 89], [59, 91],
    [71, 72], [21, 59], [83, 91], [34, 56], [78, 85], [50, 85], [51, 62], [61, 77], [78, 88], [91, 91], [33, 72],
    [90, 91], [84, 84], [79, 85], [40, 67], [31, 81], [34, 46], [9, 89], [93, 93], [74, 88], [74, 78], [77, 85],
    [58, 91], [20, 37], [17, 73], [46, 65], [51, 66], [14, 40], [91, 93], [39, 43], [13, 42], [50, 70], [63, 92],
    [12, 35], [5, 12], [76, 76], [27, 31], [63, 85], [67, 93], [92, 93], [43, 85], [35, 42], [78, 83], [12, 40],
    [51, 65], [63, 77], [48, 58], [29, 59], [36, 65], [70, 88], [49, 62], [57, 73], [42, 73], [75, 78], [27, 37],
    [5, 6], [0, 61], [40, 64], [74, 83], [25, 76], [20, 39], [3, 4], [49, 58], [85, 93], [7, 79], [48, 64],
    [16, 26],
    [59, 78], [1, 5], [68, 69], [67, 93], [16, 21], [35, 84], [15, 70], [11, 35], [3, 66], [81, 83], [35, 78],
    [24, 81], [49, 70], [80, 84], [33, 74], [81, 84], [31, 34], [75, 93], [22, 66], [54, 92], [89, 93], [81, 89],
    [7, 52], [70, 83], [68, 74], [91, 93], [54, 58], [91, 92], [78, 80], [32, 43], [12, 31], [7, 33], [54, 56],
    [6, 87], [11, 76], [60, 92], [47, 53], [40, 42], [16, 84], [4, 60], [85, 87], [50, 78], [3, 70], [34, 39],
    [32, 83], [41, 46], [38, 40], [49, 52], [93, 93], [40, 87], [16, 49], [48, 55], [86, 90], [12, 66], [31, 63],
    [71, 77], [42, 63], [65, 90], [50, 87], [61, 67], [16, 73], [67, 84], [92, 92], [37, 84], [20, 45], [47, 71],
    [66, 76], [12, 64], [44, 52], [73, 75], [5, 43], [83, 91], [40, 66], [38, 58], [45, 62], [41, 88], [66, 82],
    [1, 68], [15, 34], [40, 86], [41, 91], [41, 77], [8, 65], [35, 65], [58, 81], [48, 53], [74, 75], [17, 23],
    [67, 82], [73, 81], [31, 75], [73, 83], [46, 87], [47, 72], [39, 68], [76, 86], [68, 84], [21, 24], [18, 50],
    [87, 88], [72, 76], [14, 37], [52, 91], [6, 18], [69, 90], [34, 79], [13, 39], [33, 37], [80, 89], [67, 87],
    [10, 19], [27, 49], [65, 92], [55, 56], [75, 86], [62, 80], [28, 53], [76, 91], [30, 84], [57, 80], [69, 75],
    [61, 65], [32, 58], [25, 26], [68, 92], [48, 80], [62, 66], [51, 90], [65, 90], [74, 92], [54, 56], [45, 74],
    [0, 24], [38, 82], [88, 93], [0, 69], [15, 53], [65, 93], [40, 89], [69, 89], [73, 90], [36, 69], [52, 86],
    [66, 79], [77, 86], [57, 76], [16, 80], [56, 93], [17, 87], [20, 52], [81, 81], [54, 90], [4, 51], [53, 78],
    [36, 78], [85, 85], [11, 22], [0, 49], [34, 63], [34, 84], [47, 87], [61, 82], [49, 78], [14, 75], [45, 54],
    [53, 62], [2, 24], [33, 56], [16, 91]
]))