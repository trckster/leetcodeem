<?php

class Solution
{
    function sortedArrayToBST($nums): ?TreeNode
    {
        if (empty($nums)) {
            return null;
        }

        $middleIndex = (int)(count($nums) / 2);
        $middle = $nums[$middleIndex];

        $root = new TreeNode($middle);
        $root->left = $this->sortedArrayToBST(array_slice($nums, 0, $middleIndex));
        $root->right = $this->sortedArrayToBST(array_slice($nums, $middleIndex + 1));

        return $root;
    }
}