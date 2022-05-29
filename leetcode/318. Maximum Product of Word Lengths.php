<?php

class Solution
{
    public function maxProduct(array $words): int
    {
        $bits = [];

        foreach ($words as $word) {
            $bits[] = 0;

            for ($i = 0; $i < strlen($word); $i++) {
                $bits[count($bits) - 1] |= 1 << (122 - ord($word[$i]));
            }
        }

        $answer = 0;

        for ($i = 0; $i < count($bits); $i++) {
            for ($j = $i + 1; $j < count($bits); $j++) {
                if (($bits[$i] & $bits[$j]) === 0) {
                    $answer = max($answer, strlen($words[$i]) * strlen($words[$j]));
                }
            }
        }

        return $answer;
    }
}

$s = new Solution;
echo $s->maxProduct(['abcw', 'baz', 'foo', 'bar', 'xtfn', 'abcdef']);

