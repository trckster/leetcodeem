<?php

class Node
{
    public $val;
    public ?Node $left, $right, $next;

    function __construct($val = 0)
    {
        $this->val = $val;
        $this->left = null;
        $this->right = null;
        $this->next = null;
    }
}

class Solution
{
    public function connect(?Node $root): ?Node
    {
        if (!$root) {
            return null;
        }

        $queue = [$root];

        while (!empty($queue)) {
            $currentLevelNodes = $queue;
            $queue = [];

            for ($i = 0; $i < count($currentLevelNodes); $i++) {
                if ($i + 1 < count($currentLevelNodes)) {
                    $currentLevelNodes[$i]->next = $currentLevelNodes[$i + 1];
                }

                if ($currentLevelNodes[$i]->left) {
                    $queue[] = $currentLevelNodes[$i]->left;
                }

                if ($currentLevelNodes[$i]->right) {
                    $queue[] = $currentLevelNodes[$i]->right;
                }
            }
        }

        return $root;
    }
}