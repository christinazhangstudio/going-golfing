# go envs

## go mod tidy adds a toolchain directive and I don't want it there.
`GOTOOLCHAIN=auto`

https://github.com/golang/go/issues/65847#issuecomment-1957636062
which says:
```
> At work, we would definitely prefer to have the ability to disable the automatic toolchain upgrades 
(and possibly fail builds that would otherwise require a newer toolchain).

You do have that ability: you can set GOTOOLCHAIN=local in your process environment or GOENV file. 
(You can even edit $GOROOT/go.env to make that the default for your Go toolchain installation, if you are so inclined.)

You can also force go get to stick to a particular go version, the same way you would for any other dependency 
that you want to avoid upgrading, by passing that version as an explicit argument to go get: 
go get go@1.20 will downgrade to go 1.20, and go get knative.dev/pkg@9f033a7 go@1.20 will correctly error out 
(and report that those versions are not mutually compatible).

> Without the ability to opt out of automatic toolchain upgrades we are now forced to 
synchronize the Go version across CI and every developer machine.

That is exactly the problem that GOTOOLCHAIN=auto is supposed to mitigate: when everyone is on Go 1.21.0 or above, the developers'
machines should automatically download whatever toolchain is needed to work in the module.
```
