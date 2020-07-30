# sensu-go-basic-checks

This repo is for packaging sensu-go-basic-checks asset.

Sources:
- [go-check-plugins](https://github.com/mackerelio/go-check-plugins)
- [sensu-go-cpu-check](https://github.com/asachs01/sensu-go-cpu-check)
- [sensu-go-memory-check](https://github.com/asachs01/sensu-go-memory-check)
- [sensu-go-systemd-check](https://github.com/sardinasystems/sensu-go-systemd-check)
- [sensu-go-basic-metrics](https://github.com/fgouteroux/sensu-go-basic-metrics)

See individual plugin documentation in the upstream repository.

- [check-cpu](https://github.com/asachs01/sensu-go-cpu-check/blob/master/README.md)
- [check-memory](https://github.com/asachs01/sensu-go-memory-check/blob/master/README.md)
- [check-systemd](https://github.com/sardinasystems/sensu-go-systemd-check/blob/master/README.md)
- [check-cert-file](https://github.com/mackerelio/go-check-plugins/blob/master/check-cert-file/README.md)
- [check-disk](https://github.com/mackerelio/go-check-plugins/blob/master/check-disk/README.md)
- [check-file-age](https://github.com/mackerelio/go-check-plugins/blob/master/check-file-age/README.md)
- [check-file-size](https://github.com/mackerelio/go-check-plugins/blob/master/check-file-size/README.md)
- [check-http](https://github.com/mackerelio/go-check-plugins/blob/master/check-http/README.md)
- [check-load](https://github.com/mackerelio/go-check-plugins/blob/master/check-load/README.md)
- [check-log](https://github.com/mackerelio/go-check-plugins/blob/master/check-log/README.md)
- [check-ntpoffset](https://github.com/mackerelio/go-check-plugins/blob/master/check-ntpoffset/README.md)
- [check-ntservice](https://github.com/mackerelio/go-check-plugins/blob/master/check-ntservice/README.md)
- [check-ping](https://github.com/mackerelio/go-check-plugins/blob/master/check-ping/README.md)
- [check-procs](https://github.com/mackerelio/go-check-plugins/blob/master/check-procs/README.md)
- [check-smtp](https://github.com/mackerelio/go-check-plugins/blob/master/check-smtp/README.md)
- [check-ssl-cert](https://github.com/mackerelio/go-check-plugins/blob/master/check-ssl-cert/README.md)
- [check-tcp](https://github.com/mackerelio/go-check-plugins/blob/master/check-tcp/README.md)
- [check-uptime](https://github.com/mackerelio/go-check-plugins/blob/master/check-uptime/README.md)
- [check-windows-eventlog](https://github.com/mackerelio/go-check-plugins/blob/master/check-windows-eventlog/README.md)
- [metrics](https://github.com/fgouteroux/sensu-go-basic-metrics/blob/master/README.md)

## Release build

Some go sources are not both compatible linux/windows.

To build linux asset:
```sh
./linux_build.sh
```

To build windows asset:
```sh
./windows_build.sh
```
