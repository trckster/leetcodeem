<?php

class Solution
{
    private array $graph = [], $emailToGraphNodeId = [], $newAccounts = [];

    private ?string $currentName = null;
    private array $currentEmails = [];

    private function isEmail(string $email): bool
    {
        return substr_count($email, '@') > 0;
    }

    private function addAccountToGraph(array &$account)
    {
        $graphIds = [];

        for ($i = 1; $i < count($account); $i++) {
            $email = $account[$i];

            $graphId = $this->emailToGraphNodeId[$email] ?? null;

            if (is_null($graphId)) {
                $graphId = count($this->graph);

                $this->graph[] = [
                    'value' => $email,
                    'visited' => false,
                    'links' => []
                ];

                $this->emailToGraphNodeId[$email] = $graphId;
            }

            $graphIds[] = $graphId;
        }

        for ($i = 1; $i < count($graphIds); $i++) {
            $graphId1 = $graphIds[$i - 1];
            $graphId2 = $graphIds[$i];

            $this->graph[$graphId1]['links'][] = $graphId2;
            $this->graph[$graphId2]['links'][] = $graphId1;
        }

        $this->graph[] = [
            'value' => $account[0],
            'visited' => false,
            'links' => [$graphIds[0]]
        ];
        $this->graph[$graphIds[0]]['links'][] = count($this->graph) - 1;
    }

    private function constructGraph(array &$accounts)
    {
        foreach ($accounts as $account) {
            $this->addAccountToGraph($account);
        }
    }

    private function dfs(int $id)
    {
        $currentNode = $this->graph[$id];

        if (!$this->isEmail($currentNode['value'])) {
            $this->currentName = $currentNode['value'];
        } else {
            $this->currentEmails[] = $currentNode['value'];
        }

        $this->graph[$id]['visited'] = true;

        foreach ($currentNode['links'] as $nodeId) {
            if (!$this->graph[$nodeId]['visited']) {
                $this->dfs($nodeId);
            }
        }
    }

    private function graphToAccounts()
    {
        for ($i = 0; $i < count($this->graph); $i++) {
            $node = $this->graph[$i];

            if ($node['visited']) {
                continue;
            }

            $this->dfs($i);

            sort($this->currentEmails);

            array_unshift($this->currentEmails, $this->currentName);

            $this->newAccounts[] = $this->currentEmails;

            $this->currentName = null;
            $this->currentEmails = [];
        }
    }

    public function accountsMerge(array $accounts): array
    {
        $this->constructGraph($accounts);

        $this->graphToAccounts();

        return $this->newAccounts;
    }
}

$solution = new Solution;
$result = $solution->accountsMerge([
    ["Alex", "Alex2@m.co", "Alex6@m.co", "Alex1@m.co", "Alex5@m.co", "Alex9@m.co"],
    ["Alex", "Alex9@m.co", "Alex4@m.co", "Alex3@m.co", "Alex0@m.co", "Alex11@m.co"],
    ["Alex", "Alex9@m.co", "Alex5@m.co", "Alex6@m.co", "Alex7@m.co", "Alex12@m.co"],
    ["Alex", "Alex6@m.co", "Alex5@m.co", "Alex12@m.co", "Alex11@m.co", "Alex10@m.co"],
    ["Alex", "Alex1@m.co", "Alex5@m.co", "Alex10@m.co", "Alex9@m.co", "Alex11@m.co"],
    ["Alex", "Alex9@m.co", "Alex5@m.co", "Alex1@m.co", "Alex3@m.co", "Alex1@m.co"],
    ["Alex", "Alex10@m.co", "Alex0@m.co", "Alex3@m.co", "Alex0@m.co", "Alex0@m.co"],
    ["Alex", "Alex7@m.co", "Alex3@m.co", "Alex6@m.co", "Alex7@m.co", "Alex2@m.co"],
    ["Alex", "Alex8@m.co", "Alex6@m.co", "Alex4@m.co", "Alex6@m.co", "Alex2@m.co"],
    ["Alex", "Alex10@m.co", "Alex9@m.co", "Alex11@m.co", "Alex10@m.co", "Alex9@m.co"],
    ["Alex", "Alex0@m.co", "Alex2@m.co", "Alex9@m.co", "Alex1@m.co", "Alex6@m.co"],
    ["Alex", "Alex6@m.co", "Alex5@m.co", "Alex7@m.co", "Alex8@m.co", "Alex10@m.co"]
]);

print_r($result);