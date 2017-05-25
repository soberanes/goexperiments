# Exported names

In Go, a name is exported if it begins with a capital letter. For example, ```goPizza``` is an exported name, as is ```goPi```, which is exported from the ```gomath``` package.

```gopizza``` and ```gopi``` do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

Run the code. Notice the error message.

To fix the error, rename ```gomath.pi``` to ```gomath.Pi``` and try it again.
