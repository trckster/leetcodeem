<?php

class Solution
{
    private $letters = [];

    private function countLetters(string $s)
    {
        for ($i = 0; $i < strlen($s); $i++) {
            $current = $this->letters[$s[$i]] ?? 0;

            $this->letters[$s[$i]] = $current + 1;
        }
    }

    function checkInclusion(string $needle, $haystack): bool
    {
        if (strlen($needle) > strlen($haystack)) {
            return false;
        }

        $this->countLetters($needle);

        $runningLetters = [];
        for ($right = 0; $right < strlen($haystack); $right++) {
            $rightLetter = $haystack[$right];
            $current = $runningLetters[$rightLetter] ?? 0;

            $runningLetters[$rightLetter] = $current + 1;

            $left = $right - strlen($needle);
            if ($left >= 0) {
                $runningLetters[$haystack[$left]] -= 1;

                if (!$runningLetters[$haystack[$left]]) {
                    unset($runningLetters[$haystack[$left]]);
                }
            }

            $rightLetterCountInNeedle = $this->letters[$rightLetter] ?? 0;
            if ($runningLetters[$rightLetter] == $rightLetterCountInNeedle && $runningLetters == $this->letters) {
                return true;
            }
        }

        return false;
    }
}

$solution = new Solution;
var_dump($solution->checkInclusion('ab', 'eidboaoo'));