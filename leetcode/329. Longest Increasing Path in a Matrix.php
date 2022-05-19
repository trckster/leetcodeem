<?php

class Solution
{
    private array $cache = [];
    private array $matrix;

    public function init(array &$matrix)
    {
        $this->cache = [];
        $this->matrix = $matrix;
    }

    private function calculateIncreasingPathLength(int $i, int $j): void
    {
        $increasingPathLength = 1;

        if ($i - 1 >= 0 && $this->matrix[$i][$j] < $this->matrix[$i - 1][$j]) {
            if (!isset($this->cache[$i - 1][$j])) {
                $this->calculateIncreasingPathLength($i - 1, $j);
            }

            $increasingPathLength = $this->cache[$i - 1][$j] + 1;
        }

        if ($j + 1 < count($this->matrix[0]) && $this->matrix[$i][$j] < $this->matrix[$i][$j + 1]) {
            if (!isset($this->cache[$i][$j + 1])) {
                $this->calculateIncreasingPathLength($i, $j + 1);
            }

            $increasingPathLength = max($increasingPathLength, $this->cache[$i][$j + 1] + 1);
        }

        if ($i + 1 < count($this->matrix) && $this->matrix[$i][$j] < $this->matrix[$i + 1][$j]) {
            if (!isset($this->cache[$i + 1][$j])) {
                $this->calculateIncreasingPathLength($i + 1, $j);
            }

            $increasingPathLength = max($increasingPathLength, $this->cache[$i + 1][$j] + 1);
        }

        if ($j - 1 >= 0 && $this->matrix[$i][$j] < $this->matrix[$i][$j - 1]) {
            if (!isset($this->cache[$i][$j - 1])) {
                $this->calculateIncreasingPathLength($i, $j - 1);
            }

            $increasingPathLength = max($increasingPathLength, $this->cache[$i][$j - 1] + 1);
        }

        $this->cache[$i][$j] = $increasingPathLength;
    }

    public function longestIncreasingPath(array $matrix): int
    {
        $this->init($matrix);

        $answer = 0;

        for ($i = 0; $i < count($matrix); $i++) {
            for ($j = 0; $j < count($matrix[$i]); $j++) {
                if (!isset($this->cache[$i][$j])) {
                    $this->calculateIncreasingPathLength($i, $j);
                }

                $answer = max($answer, $this->cache[$i][$j]);
            }
        }

        return $answer;
    }
}

$s = new Solution;
echo $s->longestIncreasingPath([[9, 9, 4], [6, 6, 8], [2, 1, 1]]);
