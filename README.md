BETTER.ls
=========

Better ls is just a tiny helper for your daily work inside the terminal. I was annoyed by typing something like:

```
$ ls -l | awk '!/^d/{print }' | wc -l
```

(or to create aliases on each system I am at... or pull my beloved dotfiles)

`bls` is my simple solution, it is a ls "replacement":
```
$ bls
bls.go     1623  -rw-r--r--
README.md  21    -rw-r--r--
---------------------------
2
```

| name   | bytes | mode       |
| -------|:-----:| ----------:|
| bls.go | 1623  | -rw-r--r-- |

have fun with it (it's MIT)!