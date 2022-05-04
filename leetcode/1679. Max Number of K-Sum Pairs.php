<?php

class Solution
{
    function maxOperations(array $nums, int $sum)
    {
        $lookingForPair = [];
        $answer = 0;

        foreach ($nums as $num) {
            $lookingForSuchAmount = $lookingForPair[$sum - $num] ?? 0;

            if ($lookingForSuchAmount > 0) {
                $lookingForPair[$sum - $num]--;
                $answer++;
            } else {
                $numCount = $lookingForPair[$num] ?? 0;
                $lookingForPair[$num] = $numCount + 1;
            }
        }

        return $answer;
    }
}

$s = new Solution;
echo $s->maxOperations([2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2], 3);