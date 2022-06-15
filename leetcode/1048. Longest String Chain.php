<?php

class Word
{
    public string $word;
    public int $bestChain = 1;

    public function __construct(string $word)
    {
        $this->word = $word;
    }

    public function isValidPredecessorFor(Word $anotherWord): bool
    {
        $skipped = false;

        for ($i = 0; $i < mb_strlen($this->word); $i++) {
            $anotherWordI = $i + (int)$skipped;

            if ($anotherWord->word[$anotherWordI] !== $this->word[$i]) {
                if ($skipped) {
                    return false;
                } else {
                    $skipped = true;
                    $i--;
                }
            }
        }

        return true;
    }

    public function updateBestChain(Word $anotherWord): void
    {
        $this->bestChain = max($this->bestChain, $anotherWord->bestChain + 1);
    }
}

class Solution
{
    function longestStrChain(array $words): int
    {
        $levels = [];

        foreach ($words as $word) {
            $levels[mb_strlen($word)][] = new Word($word);
        }

        ksort($levels);
        $answer = 1;

        foreach ($levels as $length => $words) {
            $nextLevel = $levels[$length + 1] ?? false;

            if (!$nextLevel) {
                continue;
            }

            /** @var Word $predecessor */
            foreach ($words as $predecessor) {
                /** @var Word $word */
                foreach ($nextLevel as $word) {
                    if ($predecessor->isValidPredecessorFor($word)) {
                        $word->updateBestChain($predecessor);

                        if ($answer < $word->bestChain) {
                            $answer = $word->bestChain;
                        }
                    }
                }
            }
        }

        return $answer;
    }
}

$s = new Solution;
echo $s->longestStrChain(["qyssedya", "pabouk", "mjwdrbqwp", "vylodpmwp", "nfyqeowa", "pu", "paboukc", "qssedya", "lopmw", "nfyqowa", "vlodpmw", "mwdrqwp", "opmw", "qsda", "neo", "qyssedhyac", "pmw", "lodpmw", "mjwdrqwp", "eo", "nfqwa", "pabuk", "nfyqwa", "qssdya", "qsdya", "qyssedhya", "pabu", "nqwa", "pabqoukc", "pbu", "mw", "vlodpmwp", "x", "xr"]);
