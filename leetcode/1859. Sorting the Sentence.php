<?php

class Solution
{

    /**
     * @param String $s
     * @return String
     */
    function sortSentence(string $s)
    {
        $pieces = explode(' ', $s);

        $words = [];
        foreach ($pieces as $piece) {
            $number = (int) $piece[strlen($piece) - 1];
            $words[$number] = substr($piece, 0, strlen($piece) - 1);
        }

        ksort($words);

        return implode(' ', $words);
    }
}

$solution = new Solution;

echo $solution->sortSentence("is2 sentence4 This1 a3");