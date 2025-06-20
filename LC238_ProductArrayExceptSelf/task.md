# Description
```
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.
```

# Approach
```
Considering a four element array, the result array can be defined as:
r[0] = n[1] * n[2] * n[3]
r[1] = n[0] * n[2] * n[3]
r[2] = n[0] * n[1] * n[3]
r[3] = n[0] * n[1] * n[2]


We can define the prefix as the product of all previous elements and the suffix as the product of all next elements.
Thus:
r[k] == Π(i=0, i < k: n[i]) * Π(i=k+1, i<m: n[i])

Using a for loop to compute the prefix and other to compute sufixes, we can have the solution for this problem.
```

## Time complexity
O(2*n) => O(n)

## Space complexity
O(1) for computing, O(n) for result
