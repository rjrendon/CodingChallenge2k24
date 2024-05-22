
# Problem Set 3: Longest Increasing Subsequence

## Problem Description
Given an unsorted array of integers, find the length of the longest increasing subsequence.

## Solution Overview
My solution might be overkill for the given problem but it is more efficient interms of large set of arrays like billions of elements. I use **binary search algorithm**. to make the explanation short, it takes the input array nums, left and right indices and the target value, while the left index is less than the right index and it goes to middle index by dividing it by 2 and if the middle element is less than the target, it updates the left index to be one position after the middle. otherwise if the middle element is greater than or equal to the target, it updates the right index to be the middle so it would be O(log n). For **lengthOfLIS** function it initializes an empty slice piles to store the top elements of piles if the number is larger than all pile tops, it creates a new pile. else it replaces the top of the appropriate pile with the number. total time complexity for lengthOfLIS function is O(n log n) including the binarysearch n*times inside.

## Instructions to Run the Code
To run: I use VSCode and Command Line with my Windows Machine: Just type or copy this in commandLine without "": "go run longest_increasing_subsequence.go" Note: I include a test area for testing with print statement pls delete if it bothers.