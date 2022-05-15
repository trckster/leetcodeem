<?php


class TreeNode
{
    public int $val;
    public ?TreeNode $left, $right;

    function __construct($val = 0, $left = null, $right = null)
    {
        $this->val = $val;
        $this->left = $left;
        $this->right = $right;
    }
}

class Solution
{
    function deepestLeavesSum(?TreeNode $root)
    {
        if (!$root) {
            return 0;
        }

        $queue = [$root];

        while (!empty($queue)) {
            $currentLevelSum = 0;

            $nodesInLevel = count($queue);
            while ($nodesInLevel--) {
                $node = array_shift($queue);

                $currentLevelSum += $node->val;

                if ($node->left) {
                    $queue[] = $node->left;
                }

                if ($node->right) {
                    $queue[] = $node->right;
                }
            }
        }

        return $currentLevelSum;
    }
}

$left = new TreeNode(5);
$right = new TreeNode(6);

$root = new TreeNode(1, $left, $right);

$s = new Solution;
echo $s->deepestLeavesSum($root);