# Problem Set 2: Valid Parentheses

## Problem Description
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is
valid. An input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.

## Solution Overview
The function, **isValid**, checks if an input string contains properly balanced parentheses and keep track of opening parentheses encountered in the string and checks if they are closed properly. First It initializes an empty stack called stack to store characters encountered in the string then iterates through each character in the string. Inside the loop, it checks if the stack is not empty and if the character c is a closing parenthesis that matches the most recent opening parenthesis in the stack. If they match, it means the parentheses are balanced, so it pops the corresponding opening parenthesis from the stack by removing the last element. After processing all characters in the string, it checks if the stack is empty. If it's empty, it means all opening parentheses were properly closed by their corresponding closing parentheses, so it returns true. Otherwise, it returns false. Time complexity is O(n) or linear.

## Instructions to Run the Code
To run: I use VSCode and Command Line with Windows my Machine: Just type or copy this in commandLine without "": "go run valid_parentheses.go" Note: I include a test area for testing with print statement pls delete if it bothers.