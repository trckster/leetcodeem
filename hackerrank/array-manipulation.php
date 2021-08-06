<?php

function arrayManipulation($n, $queries)
{
    $arr = [];

    foreach ($queries as $query) {
        $start = $query[0] - 1;
        $end = $query[1];
        $value = $query[2];

        $arr[$start] = ($arr[$start] ?? 0) + $value;

        if ($end < $n) {
            $arr[$end] = ($arr[$end] ?? 0) - $value;
        }
    }

    ksort($arr);

    $answer = 0;
    $currentSum = 0;

    foreach ($arr as $change) {
        $currentSum += $change;
        $answer = max($answer, $currentSum);
    }

    return $answer;
}

$result = arrayManipulation(10, [[2, 6, 8], [3, 5, 7], [1, 8, 1], [5, 9, 15]]);

echo $result;
