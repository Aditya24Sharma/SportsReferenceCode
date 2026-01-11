# SportsReferenceCode
Code for Sports Reference 2026 Summer Engineering intern

# Head-to-Head Matrix (Go)

This project demonstrates how to build a matrix-style table in Go from a nested
JSON-like data structure representing head-to-head win/loss records.

## Problem Description

Each team has win (`W`) and loss (`L`) counts against other teams.
The goal is to create a table where:

- Rows represent teams
- Columns represent opponents
- Each cell shows wins against that opponent
- Teams do not play themselves (`--`)

## Approach

1. **Data Modeling**
   - Use Go maps to represent nested JSON
   - Define small structs for clarity

2. **Team Ordering**
   - Extract team keys
   - Sort alphabetically for deterministic output

3. **Nested Iteration**
   - Outer loop prints rows
   - Inner loop prints opponent cells

## Running the Code
go run matrix.go
