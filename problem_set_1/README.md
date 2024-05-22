# Problem Set 1: Palindrome Pairs

## Problem Description
Given a list of unique words, your task is to find all pairs of distinct indices (i, j) in the list so that
the concatenation of the two words, i.e., words[i] + words[j], forms a palindrome.

## Solution Overview
My solution for this problem for **isPalindrome** function is iterate the string using a loop that goes from the beginning of the string to the middle of the string. Then inside the loop i put an if s[i] and s[len(s)-1-i] is not a match, it means the string is not a palindrome, so it returns false, and if the loop doesnt find any mismatch then it should be true or a palindrome. For the **palindromePairs** function takes a slice of strings words and returns a slice of pairs of indices. It iterates through each word in the words slice using a nested loop. The outer loop (i) goes through each word in words, and the inner loop (j) also iterates through each word in words, inside the nested loop, it checks if i is not equal to j to avoid pairing a word with itself and if the concatenation of the words words[i] and words[j] forms a palindrome. It checks this condition using the isPalindrome function. If the concatenation forms a palindrome, it appends the pair of indices [i, j] to the result slice, and then it returns the result slice containing all the pairs of indices that form palindromes when concatenated.

## Instructions to Run the Code
To run: I use VSCode and Command Line with Windows my Machine: Just type or copy this in commandLine without "": "go run palindrome_pairs.go" Note: I include a test area for testing with print statement pls delete if it bothers.