#Slices

An array has a fixed size. A slice, ion the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common that arrays.

Ty type `[]T` is a slice with elements of type `T`.

This expression creates a slice of the first five elements of the array `a`:

```go
a[0:5]
```
