# Linear Programming Demos

## What is linear programming?

Linear programming (LP) is a way to find the best choice when you have
limited resources. You describe what you can choose, what you want to
maximize or minimize, and the limits you must obey.

For example, a business might choose how many products to make to earn as
much money as possible without using more materials than it has.

## Demos

- [Pete's Pottery Paradise](PetesPotteryParadise/README.md) shows how to
  choose how many small and large vases to make with limited clay and glaze.

## How Google OR-Tools helps

Google OR-Tools is a software toolkit for solving optimization problems.
You give it the choices (variables), the goal (such as maximizing revenue),
and the limits (constraints). A solver in OR-Tools searches for a solution
that obeys every limit and gives the best possible result.

In Pete's Pottery Paradise, the model uses the amount of each vase to make
as its choices, vase revenue as its goal, and the clay and glaze inventories
as limits. Because the number of vases must be whole numbers, this demo uses
OR-Tools' CP-SAT solver rather than a continuous LP solver.
