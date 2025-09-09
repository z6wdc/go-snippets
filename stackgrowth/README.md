# Stack Growth

This program demonstrates how Go’s runtime manages stack growth during recursive calls and how the size of arrays affects stack memory allocation.

By default, Go starts with a small stack (about 2KB) and automatically expands it as needed.

When functions use large arrays, the stack must grow more frequently, which can impact performance in deeply recursive scenarios.

This example helps illustrate how function parameters influence stack usage and when stack growth occurs in Go programs.

## Concept Overview

The program explores the relationship between:
- **Recursive function calls** and stack frame allocation
- **Go's stack growth mechanism** when the stack needs to expand

## How It Works

### Key Components

1. **Recursive Function**: `stackCopy()` calls itself recursively up to 10 times
2. **Array Parameter**: A fixed-size array `[size]int{}` is passed by value to each recursive call
3. **Stack Observation**: Each call prints its recursion depth, string pointer address, and string value

### Stack Memory Behavior

- **Small Arrays** (`size = 10`): Minimal stack growth, efficient memory usage
- **Large Arrays** (`size = 1024`): Significant stack growth due to array copying

When `size` is set to a large value like 1024, each recursive call must allocate space for a 1024-integer array on the stack, causing more frequent stack expansions.

## Experiment Instructions

1. **Run with small array**:
   ```bash
   go run main.go
   ```
   
2. **Modify for large array**:
   - Change `const size = 10` to `const size = 1024`
   - Run again to observe the difference

## Expected Output

The program prints for each recursion level:
- **Recursion depth** (0-9)
- **String pointer address** (same across all calls)
- **String value** ("HELLO")
