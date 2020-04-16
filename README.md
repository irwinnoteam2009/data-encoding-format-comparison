# data-encoding-format-comparison

Data encoding format comparison. Superhero object taken from https://github.com/nrwiersma/avro-benchmarks

## Avro

### Official Apache Avro repository

1. [c++](https://github.com/apache/avro/tree/master/lang/c++)
1. [c](https://github.com/apache/avro/tree/master/lang/c)
1. [C#](https://github.com/apache/avro/tree/master/lang/csharp)
1. [java](https://github.com/apache/avro/tree/master/lang/java)
1. [js](https://github.com/apache/avro/tree/master/lang/js)
1. [perl](https://github.com/apache/avro/tree/master/lang/perl)
1. [php](https://github.com/apache/avro/tree/master/lang/php)
1. [py](https://github.com/apache/avro/tree/master/lang/py)
1. [py3](https://github.com/apache/avro/tree/master/lang/py3)
1. [ruby](https://github.com/apache/avro/tree/master/lang/ruby)

### Unofficial Avro repositories

1. [go hamba/avro](https://github.com/hamba/avro)
1. [go go-avro/avro](https://github.com/go-avro/avro)
1. [go linkedin/goavro](https://github.com/linkedin/goavro)

## Msgpac

### Unofficial Msgpac

1. [go vmihailenco/msgpack](https://github.com/vmihailenco/msgpack)

## Generate bin files

```bash
go run src/tools/create-avro-go-arvo-superhero-bin/main.go  
go run src/tools/create-avro-hamba-superhero-bin/main.go
go run src/tools/create-avro-hamba-snappy-superhero-bin/main.go
go run src/tools/create-msgpac-vmihailenco-superhero-bin/main.go
```

## Bin size

| Name | Size | Note |
| ---- | ---- | ---- |
| avro-go-arvo-superhero.bin        | 100 bytes | Неясно почему на 2 байта меньше чем avro-hamba-superhero.bin |
| avro-hamba-superhero.bin          | 102 bytes | - |
| avro-hamba-snappy-superhero.bin   | 105 bytes | При сжатии маленьких данных объём становится больше, но сам формат отлично сжимается, в офф. доккументации Arvo рекомендуют сжимать snappy |
| msgpac-vmihailenco-superhero.bin  | 255 bytes | - |
| superhero.json                    | 317 bytes | - |

## Testing results

Go: 1.14 Machine: Macbook Pro 8-core, 2,2 GHz Intel Core i7 16 GB 1600 MHz DDR3

```bash
go test ./test -bench=. -benchmem -benchtime=10s
```

```bash
BenchmarkAvroGoAvroDecode-8              2707065              4450 ns/op             442 B/op         27 allocs/op
BenchmarkAvroGoAvroEncode-8              2220375              5301 ns/op             861 B/op         63 allocs/op
BenchmarkAvroHambaSnappyDecode-8        15746656               766 ns/op             176 B/op          5 allocs/op
BenchmarkAvroHambaSnappyEncode-8        14795838               810 ns/op             336 B/op          3 allocs/op
BenchmarkAvroHambaDecode-8              17261074               680 ns/op              64 B/op          4 allocs/op
BenchmarkAvroHambaEncode-8              20581195               580 ns/op             176 B/op          2 allocs/op
BenchmarkAvroLinkedinDecode-8            4923289              2440 ns/op            1776 B/op         40 allocs/op
BenchmarkAvroLinkedinEncode-8           12724077               952 ns/op             288 B/op         10 allocs/op
BenchmarkJSONIteratorGoDecode-8          4375777              2728 ns/op             192 B/op         25 allocs/op
BenchmarkJSONIteratorGoEncode-8          5101243              2357 ns/op             384 B/op          2 allocs/op
BenchmarkJSONPquernaFfjsonDecode-8       1573910              7690 ns/op             384 B/op         21 allocs/op
BenchmarkJSONPquernaFfjsonEncode-8       5141089              2323 ns/op             384 B/op          2 allocs/op
BenchmarkJSONDecode-8                    1564813              7711 ns/op             384 B/op         21 allocs/op
BenchmarkJSONEncode-8                    5335033              2255 ns/op             384 B/op          2 allocs/op
BenchmarkMsgpacVmihailencoDecode-8       3098145              3895 ns/op             224 B/op         26 allocs/op
BenchmarkMsgpacVmihailencoEncode-8       4936180              2412 ns/op             608 B/op          5 allocs/op
```

## Links

1. [Avro spec](https://avro.apache.org/docs/current/spec.html)
1. [Avro schema evolution](https://binx.io/blog/2018/12/09/apache-avro-handson/)
1. [Avro schema from JSON generator](https://toolslick.com/generation/metadata/avro-schema-from-json)
1. [Avro schema to JSON schema](https://json-schema-validator.herokuapp.com/avro.jsp)
1. [Сравнение форматов сериализации](https://habr.com/ru/post/458026/)
1. [The best serialization strategy for Event Sourcing](https://blog.softwaremill.com/the-best-serialization-strategy-for-event-sourcing-9321c299632b)

## TODO

1. Add more msgpac implementations
1. Add more other encoding format implementations like as protobuf, thrift(binary, compact), etc
