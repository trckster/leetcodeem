<?php

class Solution
{
    public function minOperations(array $numbers, int $target): int
    {
        $target = array_sum($numbers) - $target;

        if ($target < 0) {
            return -1;
        }

        $left = 0;
        $currentSum = 0;
        $biggestLength = -1;

        for ($right = 0; $right < count($numbers); $right++) {
            $currentSum += $numbers[$right];

            while ($left <= $right && $currentSum > $target) {
                $currentSum -= $numbers[$left];
                $left++;
            }

            if ($currentSum === $target) {
                $biggestLength = max($biggestLength, $right - $left + 1);
            }
        }

        if ($biggestLength === -1) {
            return -1;
        }

        return count($numbers) - $biggestLength;
    }
}

echo (new Solution)->minOperations([1, 1, 4, 2, 3], 5) . "\n";
echo (new Solution)->minOperations([5, 6, 7, 8, 9], 4) . "\n";
echo (new Solution)->minOperations([3, 2, 20, 1, 1, 3], 10) . "\n";
echo (new Solution)->minOperations([1, 1], 2) . "\n";

echo (new Solution)->minOperations([8828, 9581, 49, 9818, 9974, 9869, 9991, 10000, 10000, 10000, 9999, 9993, 9904, 8819, 1231, 6309], 134365) . "\n";