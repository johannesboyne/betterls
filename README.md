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

## How to build it for your platform?

First of all, you can look inside the prebuild dir, maybe I've did the work for you. Otherwise here are the commands:
```
$ GOOS=windows GOARCH=386 go build -o bls.exe bls.go #windows
$ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bls.linux_amd64 bls.go #linux
```

## Protip

Move/link the build go binary to your `/usr/bin/` directory, thus you can call it from everywhere!
