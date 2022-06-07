<?php

class Solution
{
    function merge(array &$main, int $firstLength, array $second, int $secondLength): void
    {
        $first = array_slice($main, 0, $firstLength);

        for ($i = $firstIndex = $secondIndex = 0; $i < count($main); $i++) {
            if ($firstIndex >= count($first) || $secondIndex < count($second) && $first[$firstIndex] > $second[$secondIndex]) {
                $main[$i] = $second[$secondIndex++];
            } else {
                $main[$i] = $first[$firstIndex++];
            }
        }
    }
}
