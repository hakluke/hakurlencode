# hakurlencode
(en|de)code urls from the CLI

# Installation

```
go get github.com/hakluke/hakurlencode
```

# Usage

Pipe into the tool with no options to encode, for decoding, add `-d`.

```
~$ echo "test!" | hakurlencode
test%21
~$ echo "test%21" | hakurlencode -d
test!
```
