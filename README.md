# cows

> 400+ ASCII 🐮s

Get a list of cows or a random cow. The `Go` version. Original [JavaScript Cows](https://github.com/sindresorhus/cows).

## Install

```sh
go get github.com/bojand/cows
```

## API

```go
func All() []string
```

Get all the cows.

```go
func Random() string
```

Get a random cow.

## Usage

```go
package main

import (
	"fmt"

	"github.com/bojand/cows"
)

func main() {
	fmt.Println(cows.Random())
}
```

## Examples

```
         (__)
         (oo)
  /-------\/
 / |     ||
+  ||----||
   ~~    ~~
     Cow


       \(:)/
       (o|o)
  /-----\_/
 /|      |
^ ||----||
  ^^    ^^
 Klingon Cow


                 ________________
         ^__^   /                \
         (oo)  ( Milk is logical. )
  /-------\/ --'\________________/
 / |     ||
*  ||W---||
   ^^    ^^
Mr Spock's cow


      (__)
    /   oo      ______
   |  /\_|     |      \
   |  |___     |       |
   |   ---@    |_______|
*  |  |   ----   |    |
 \ |  |_____
  \|________|
        CompuCow


        ___________________________
       | (__)  (__)  (__)   (__)  |
       | ( oo  ( oo  ( oo   ( oo  |
_______| /\_|  /\_|  /\_|   /\_|  |________
|                                         |
|   _____                        _____    |
|___|   |________________________|   |____|
    |___|                        |___|
              Cow-pooling


  /--------------------/
 / |     ||           /          (__)
*  ||----||          /-----------(oo)
   ^^    ^^                       \/
        Network Virtual Cow
(with separate frontend and backend)


         (__)               (__)  |    |  (__)
         (--)               (--)  |    |  (--)
  /-------\/   /o    /-------\/   |    I   \/-------\
 / |  M  |----< o   / |  L  |----<T    I>----|  D  | \
*  ||----|   /  o  *  ||----|     I    I     |----||  *
   ^^    ^      |     ^^    ^          |     ^    ^^
                |                      |
              Teenage Mutant Ninja Cows

                                  @
               (__)    (__) _____/
            /| (oo) _  (oo)/----/_____    *
  _o\______/_|\_\/_/_|__\/|____|//////== *- *  * -
 /_________   \   00 |   00 |       /== -* * -
[_____/^^\_____\_____|_____/^^\_____]     *- * -
      \__/                 \__/
               Cow-mobile
```


## Related

- [JavaScript Cows](https://github.com/sindresorhus/cows) - Original Javascript ASCII cows package
- [vaca](https://github.com/sindresorhus/vaca) - Get a random ASCII cow 🐮
- [cows-docker](https://github.com/alexellis/cows-docker) - ASCII cows on Docker
- [Swift Cows](https://github.com/NozeIO/Noze.io/tree/develop/Sources/cows) - ASCII cows in Swift
- [Elixir Cows](https://github.com/sotojuan/excows) - ASCII cows in Elixir
- [Haskell Cows](https://github.com/passy/cows-hs) - ASCII cows in Haskell

## License

MIT © [Sindre Sorhus](http://sindresorhus.com)
MIT © [Bojan D](http://github.com/bojand)
