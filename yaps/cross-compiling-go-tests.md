## Themes : 
- Cross-compiling Go
- QEMU

(go testing a Windows application on a github runner)
  
the runner configuration needed to be changed bc importing a go dependency that uses windows syscalls will result in `undefined: syscall.DLL`
go env shows the `GOOS` is still `linux`, this is misleading as the runner is labeled with `dt-ops::windows::docker`...
the setup env step confirms this:
```
  RUNNER_LABELS=dt-ops::windows::docker
  RUNNER_OS=Linux <--------------------
```
i guess it's because really `dt-ops::windows::docker` is a "unix shell with a remote docker daemon running on a windows platform", so 

i don't know of any instance of installing go for Windows and making `go test` work in the CI, so this may v well be a first "hack" of its kind. this runner is the only 1 i know of that is windows capable and i haven't found anything to hint otherwise...

so we will enable this "hackily"

# cross-compiling go tests

to cross compile on linux for windows: 
```
GOOS=windows GOARCH=amd64 go test ./...
```

of course, if you just add that, you won't be able to run `go test` 
the CI will give this error (but we have at least moved on from `undefined: syscall.DLL`):
```
fork/exec /tmp/go-build4120727679/b802/types.test.exe: exec format error
FAIL	github.tesla.com/design-technology/derived-output-event-system/internal/types
```
(* it is at some later point I find this line also mentioned in the QEMU blog post, see below,

...which means the rest of this section is sort of just a rabbit hole

one *can* try *compiling* the tests by 
```
GOOS=windows GOARCH=amd64 go test -c -o ./test_binaries ./...
```
("with multiple packages, -o must refer to a directory or /dev/null")

this also doesn't work: 
```
cannot write test binary worker.test for multiple packages:
github.tesla.com/design-technology/derived-output-event-system/app/converter/worker
```

so we try the single compile per package strategy:

```yaml
  mkdir -p test_binaries
  for pkg in $(go list -f '{{if .TestGoFiles}}{{.ImportPath}}{{end}}' ./...); do
    GOOS=windows GOARCH=amd64 go test -c -o test_binaries/$(basename $pkg).test $pkg
  done
```

fyi this will give you just packages with test files:
```
go list -f '{{if .TestGoFiles}}{{.ImportPath}}{{end}}' ./...
```

but alas, this just checks if the tests *compile*. just simply try, and :
```
go test -c -o conversion.test github.tesla.com/design-technology/derived-output-event-system/app/converter/conversion
```
and
```
> go test conversion.test
no required module provides package conversion.test; to add it:
        go get conversion.test
> ./conversion.test
(no output)
```
to see that testing doesn't work in the same way with a compiled binary as you would expect with `go test`.

it *is* what https://github.com/golang/go/issues/15513 is about. 

so at this point, I'm pretty sure reason(b) is why this would not work, and [this SO ans](https://stackoverflow.com/a/74389382) which leads to [this blog post about using QEMU](https://ctrl-c.us/posts/test-goarch)  
is the breakthrough for this all.

ofc we're not going down the QEMU rabbit hole, and this post is all just to say that we should just use the Windows Docker (as, in a sense, intended) to `go test` Windows-targeted apps.

# in conclusion: use QEMU or Docker for cross-platform go tests

*appendix:*
## reasons
(a) 
> "can't generate one binary for tests in multiple packages
because there is no guarantee that package A's test doesn't
interfere with package B's use of package A."
https://github.com/golang/go/issues/15513#issuecomment-216694348

(b)
> It's not possible to run go test for a target system that's different from the current system.
https://stackoverflow.com/questions/48895634/design-and-unit-test-cross-platform-application/48896657#48896657
https://ctrl-c.us/posts/test-goarch
