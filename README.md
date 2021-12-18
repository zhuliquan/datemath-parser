# datemath-parser

this package is pure go package, this package can parse date match expression, which used by ElasticSearch. 

## Date Math Definition

you can click [here](http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/mapping-date-format.html#date-math)  to see date math definition. The date type supports using date math expression when using it in a range date query. The expression starts with an "anchor" date, which can be either now or a date string (in the applicable format) ending with `||`. It can then follow by a math expression, supporting `+`, `-` and `/` (time rounding symbol). The support time units are:
| time unit symbol |  meaning |
| --- | ---     |
| y   | Years   |
| M   | Months  |
| w   | Weeks   |
| d   | Days    |
| h   | Hours   |
| H   | Hours   |
| m   | Minutes |
| s   | Seconds |

Here are some samples: 
`now+1h`, `now+1h+1m`, `now+1h/d`, `2012-01-01||+1M/d`.

Note, when doing range type searches, and the upper value is inclusive, the rounding will properly be rounded to the ceiling instead of flooring it.

## Usage

Returns an int64 representing timestamp in milliseconds
```golang
package main

import (
    "fmt"
    "github.com/zhuliquan/datemath-parser"
)

func main() {
    

    
}
```