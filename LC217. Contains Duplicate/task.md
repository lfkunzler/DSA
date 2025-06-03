# Description
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

# First Solution Idea
A map, where the key is the array elements and the value is true. If one array element is already mapped, return false.

## Time complexity
In the worst case, O(n), where we need to scan the whole array once.

## Space complexity
In the worst case, O(n), where we need to allocate a map element for every array element.
