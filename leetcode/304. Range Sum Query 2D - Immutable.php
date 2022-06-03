<?php

class NumMatrix
{
    private array $sums = [];

    function __construct(array $matrix)
    {
        foreach ($matrix as $i => $row) {
            foreach ($row as $j => $value) {
                $sumAbove = $this->sums[$i - 1][$j] ?? 0;
                $leftSum = $this->sums[$i][$j - 1] ?? 0;
                $surplus = $this->sums[$i - 1][$j - 1] ?? 0;

                $this->sums[$i][$j] = $sumAbove + $leftSum - $surplus + $value;
            }
        }
    }

    function sumRegion(int $row1, int $col1, int $row2, int $col2): int
    {
        $sumAbove = $this->sums[$row1 - 1][$col2] ?? 0;
        $leftSum = $this->sums[$row2][$col1 - 1] ?? 0;
        $surplus = $this->sums[$row1 - 1][$col1 - 1] ?? 0;

        return $this->sums[$row2][$col2] - $sumAbove - $leftSum + $surplus;
    }
}