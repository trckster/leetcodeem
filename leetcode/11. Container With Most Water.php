<?php

class Solution
{
    function maxArea($height)
    {
        arsort($height);

        $bestArea = 0;
        $left = PHP_INT_MAX;
        $right = -1;

        foreach ($height as $i => $size) {
            if ($i < $left) {
                $left = $i;
            }

            if ($i > $right) {
                $right = $i;
            }

            $areaHeight = min($height[$left], $height[$right]);
            $areaWidth = $right - $left;

            $area = $areaHeight * $areaWidth;

            $bestArea = max($area, $bestArea);
        }

        return $bestArea;
    }
}


$s = new Solution;
echo $s->maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]);
