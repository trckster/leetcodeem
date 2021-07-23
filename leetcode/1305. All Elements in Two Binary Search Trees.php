<?php


class TreeNode
{
    public $val = null;
    public $left = null;
    public $right = null;

    function __construct($val = 0, $left = null, $right = null)
    {
        $this->val = $val;
        $this->left = $left;
        $this->right = $right;
    }
}

class Solution
{
    private $items = [];

    private function leftTreeTraversal($root)
    {
        if (is_null($root)) {
            return;
        }

        if ($root->left) {
            $this->leftTreeTraversal($root->left);
        }

        if ($root->right) {
            $this->leftTreeTraversal($root->right);
        }

        $this->items[] = $root->val;
    }


    /**
     * @param TreeNode $root1
     * @param TreeNode $root2
     * @return Integer[]
     */
    function getAllElements($root1, $root2)
    {
        $this->leftTreeTraversal($root1);
        $this->leftTreeTraversal($root2);

        sort($this->items);

        return $this->items;
    }
}