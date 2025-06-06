 Description
```
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
```

# Approach
```
1. We store the values in a map as the key and the frequency as the value of this map. After that, we use a bucket sort to arrange the elements.
2. Using a bucket list, we use the position of the array to store the counter, ie, the position 9 has all the elements that have 9 occurrences.
```

## Time complexity
O(n) to fill the map, O(m) to fill the frequency array and O(k) to fill the answer array. O(n+m+k) => O(n)

## Space complexity
O(m) to fill the map, O(n) for the bucket array, O(k) to fill the answer array. O(n+m+k) => O(n)
