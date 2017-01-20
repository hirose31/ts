# ts

`ts` adds a timestamp to the beginning of each line of input like the `ts` command of `moreutils`.

## SYNOPSIS

```
ts [-f format] [-V]
ts --help
```

```
$ printf "foo\nbar\nbaz\n" | ts
2017-01-20 21:28:14 foo
2017-01-20 21:28:14 bar
2017-01-20 21:28:14 baz

$ printf "foo\nbar\nbaz\n" | ts -f '%D %r'
01/20/17 9:29:50 PM foo
01/20/17 9:29:50 PM bar
01/20/17 9:29:50 PM baz
```



## OPTIONS

* `-f format`

    default is `"%Y-%m-%d %H:%M:%S"`

* `-V`

    output version information and exit

* `--help`

    display this help and exit

