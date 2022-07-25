#!/usr/bin/env bash

easy_questions_counter=0
medium_questions_counter=0
hard_questions_counter=0

for entry in `pwd`/*/
do
  file=`ls $entry*.md`

  if grep -q 'Easy' "$file"; then
    ((easy_questions_counter = easy_questions_counter + 1))
  elif grep -q 'Medium' "$file"; then
     ((medium_questions_counter = medium_questions_counter + 1))
  elif grep -q 'Hard' "$file"; then
      ((hard_questions_counter = hard_questions_counter + 1))
  fi
done

cat > README.md  << EOM
# Leetcode

Repository containing *[Leetcode](https://www.leetcode.com)* questions solutions.

Each directory contains one *[Leetcode](https://www.leetcode.com)* question. Inside these directories are **three** files.

1. \`README.md\` - contains the description of the question.
2. \`solution.go\` - contains the functionality to solve the problem.
3. \`solution_test.go\` - contains unit tests to validate the code inside \`solution.go\`.

### Question Types:

    Easy: $easy_questions_counter
    Medium: $medium_questions_counter
    Hard: $hard_questions_counter

    Total: `expr $easy_questions_counter + $medium_questions_counter + $hard_questions_counter`
EOM
