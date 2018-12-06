# gotypenames

Go tool to print all declared type names.

## Installation

`go get -u github.com/zoncoen/gotypenames`

## Usage

```
usage: gotypenames --filename=FILENAME [<flags>]

Go tool to print all declared type names.

Flags:
      --help                    Show context-sensitive help (also try --help-long and --help-man).
  -f, --filename=FILENAME       Target filename.
      --only-exported           Print only exported type name.
      --types=primitive... ...  Filter by type. (primitive, array, map, func, struct, interface, chan)
```
