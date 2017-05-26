# Exercise: Loops and Functions

As a simple way to play with functions and loops, implement the square root function using Newton's method.

In this case, Newton's method is to approximate Sqrt(x) by picking a starting point z and then repeating:

![newton](https://tour.golang.org/content/img/newton.png "Newton's method")

To begin with, just repeat that calculation 10 times and see how close you get to the answer for various values (1, 2, 3, ...).

Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small delta). See if that's more or fewer iterations. How close are you to the [math.Sqrt](https://golang.org/pkg/math/#Sqrt)?

Hint: to declare and initialize a floating point value, give it floating point syntax or use a conversion:

```go
z := float64(1)
z := 1.0
```

_Note__: can't do this.
