<?php

class Solution
{
    const WALL = 1;
    const FREE = 0;

    private int $size;

    private function isBad(array &$grid): bool
    {
        return $grid[0][0] !== self::FREE || $grid[$this->size - 1][$this->size - 1] !== self::FREE;
    }

    public function runBfs(array &$grid): int
    {
        $answer = 1;

        $queue = [[0, 0]];
        $grid[0][0] = self::WALL;

        while (!empty($queue)) {
            $queueSize = count($queue);

            while ($queueSize--) {
                [$y, $x] = array_shift($queue);

                if ($x === $this->size - 1 && $y === $this->size - 1) {
                    return $answer;
                }

                if ($y - 1 >= 0 && $grid[$y - 1][$x] === self::FREE) {
                    $queue[] = [$y - 1, $x];
                    $grid[$y - 1][$x] = self::WALL;
                }

                if ($y - 1 >= 0 && $x + 1 < $this->size && $grid[$y - 1][$x + 1] === self::FREE) {
                    $queue[] = [$y - 1, $x + 1];
                    $grid[$y - 1][$x + 1] = self::WALL;
                }

                if ($x + 1 < $this->size && $grid[$y][$x + 1] === self::FREE) {
                    $queue[] = [$y, $x + 1];
                    $grid[$y][$x + 1] = self::WALL;
                }

                if ($y + 1 < $this->size && $x + 1 < $this->size && $grid[$y + 1][$x + 1] === self::FREE) {
                    $queue[] = [$y + 1, $x + 1];
                    $grid[$y + 1][$x + 1] = self::WALL;
                }

                if ($y + 1 < $this->size && $grid[$y + 1][$x] === self::FREE) {
                    $queue[] = [$y + 1, $x];
                    $grid[$y + 1][$x] = self::WALL;
                }

                if ($y + 1 < $this->size && $x - 1 >= 0 && $grid[$y + 1][$x - 1] === self::FREE) {
                    $queue[] = [$y + 1, $x - 1];
                    $grid[$y + 1][$x - 1] = self::WALL;
                }

                if ($x - 1 >= 0 && $grid[$y][$x - 1] === self::FREE) {
                    $queue[] = [$y, $x - 1];
                    $grid[$y][$x - 1] = self::WALL;
                }

                if ($y - 1 >= 0 && $x - 1 >= 0 && $grid[$y - 1][$x - 1] == self::FREE) {
                    $queue[] = [$y - 1, $x - 1];
                    $grid[$y - 1][$x - 1] = self::WALL;
                }
            }

            $answer++;
        }

        return -1;
    }

    function shortestPathBinaryMatrix(array $grid): int
    {
        $this->size = count($grid);

        if ($this->isBad($grid)) {
            return -1;
        }

        return $this->runBfs($grid);
    }
}

$s = new Solution;
echo $s->shortestPathBinaryMatrix([[0, 0, 0], [1, 1, 0], [1, 1, 0]]);