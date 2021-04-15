# incr

```
$ incr ls
ls1
$ ls
README.md  go.mod     main.go
$ ls > $(incr ls)
$ ls
README.md  go.mod     ls1        main.go
$
$ ls > $(incr ls)
$ ls
README.md  go.mod     ls1        ls2        main.go
$ incr ls
ls3
```


