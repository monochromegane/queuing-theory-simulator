# Queuing Theory Simulator in Go

Run simple simulator for `M/M/S` model of queuing theory.

```sh
$ go run cmd/simulator/main.go --step 100 --server 2 --lambda 0.2 --mu 0.1
$ python plot.py
$ open out/plot.png
```

![plot](https://user-images.githubusercontent.com/1845486/65520008-e1663f00-df21-11e9-8c4b-a42f018c895a.png)

## License

[MIT](https://github.com/monochromegane/queuing-theory-simulator/blob/master/LICENSE)

## Author

[monochromegane](https://github.com/monochromegane)
