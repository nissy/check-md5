# check-md5
Check the MD5 checksum

### Example of mackerel-agent.conf

```
[plugin.checks.md5]
command = "check-md5 -c check-md5.conf"
```

### Example of check-md5.conf

```
[[Files]]
PATH = "/usr/bin/python"
HASH = "faf96ffcd1955149edc54cd5e4195a0d"

[[Files]]
PATH = "/etc/redhat-release"
HASH = "af2cb935515b9d48999ecb2b1f4122e6"
```

### Save hash config

```
$ cat check-md5.conf
[[Files]]
PATH = "/usr/bin/python"

$ check-md5 -s -c check-md5.conf

$ cat check-md5.conf
[[Files]]
PATH = "/usr/bin/python"
HASH = "faf96ffcd1955149edc54cd5e4195a0d"
```

### Help

```
Usage: check-md5 [options]
  -c string
        set cfgiguration file (default "check-md5.conf")
  -h    this help
  -s    save cfgiguration file md5 hash
  -v    show this build version
```