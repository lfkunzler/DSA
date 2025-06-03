# Description
```
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
```

# Approach
```
We need to find the number that + the number in the current position = target.
Example: if the target is '10' and the current position 3 has the value of '7', we need to find the position of '3'.
Using a hash map, we can store the values and its position and to look at it while we iterate over the array.
```

## Time complexity
O(n) since we iterate over the full array once in the worst scenario.

## Space complexity
O(n) as we need to add every element of the array and its position in the worst scenario.

