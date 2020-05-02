sample cat program with golang

#### usage
```shell script
$ ./main foo.txt bar.txt
foo1
foo2
foo3
bar1
bar2
bar3
```

With `-n` flag, display text with line number.
```shell script
$ ./main -n foo.txt bar.txt
1: foo1
2: foo2
3: foo3
4: bar1
5: bar2
6: bar3
```
