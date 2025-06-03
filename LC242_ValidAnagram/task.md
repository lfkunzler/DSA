# Description
Given two strings s and t, return true if t is an of s, and false otherwise.

# Approach
1. Two maps, where the key is the array elements and the value is the count of occurrences for each char.
2. We can adjust the implementation to use only one map.

## Time complexity
In the worst case, O(n), where we need to scan both arrays once.

## Space complexity
In the worst case, O(n), where we need to allocate a map element for both array elements.
