# hyperkit-high-cpu
An example of an idle VM inducing high CPU in hyperkit

To build:
```
linuxkit build bad.yml
linuxkit build good.yml
```

To see a VM with low CPU:
```
linuxkit run good
```

To see a VM with high CPU:
```
linuxkit run bad
```

Note the only difference is the good VM uses Linux 4.19 while the bad VM uses Linux 5.4.
To try to make the bad VM better, all the bug mitigations have been disabled but this
doesn't seem to help.

The helper script `go run measure.go` will monitor the %cpu of the hyperkit process for
60 seconds and produce an average.
