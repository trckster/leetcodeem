<?php

class Solution
{
    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function findDisappearedNumbers($nums)
    {
        for ($i = 0; $i < count($nums);) {
            $rightPosition = $nums[$i] - 1;

            if ($i + 1 !== $nums[$i] && $nums[$i] !== $nums[$rightPosition]) {
                [$nums[$i], $nums[$rightPosition]] = [$nums[$rightPosition], $nums[$i]];
            } else {
                $i++;
            }
        }

        $missing = [];
        for ($i = 0; $i < count($nums); $i++) {
            if ($i + 1 !== $nums[$i]) {
                $missing[] = $i + 1;
            }
        }

        return $missing;
    }
}

$solution = new Solution;
print_r($solution->findDisappearedNumbers([4, 3, 2, 7, 8, 2, 3, 1]));
