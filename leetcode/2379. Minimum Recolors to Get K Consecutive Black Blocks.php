<?php

class Solution
{
    function minimumRecolors($blocks, $k): int
    {
        $black = $white = 0;

        for ($i = 0; $i < $k; $i++) {
            if ($blocks[$i] === 'B') {
                $black++;
            } else {
                $white++;
            }
        }

        $answer = $white;

        for ($i = $k; $i < strlen($blocks); $i++) {
            if ($blocks[$i] === 'B') {
                $black++;
            } else {
                $white++;
            }

            if ($blocks[$i - $k] === 'B') {
                $black--;
            } else {
                $white--;
            }

            $answer = min($answer, $white);

            if (!$answer) {
                return 0;
            }
        }

        return $answer;
    }
}
