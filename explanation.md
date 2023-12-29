# Go `tree` implementation

## Symbols

| Letter | Unicode  | Comment              |
|--------|----------|----------------------|
| `I`    | &#x2502; | `vertical`           |
| `K`    | &#x251C; | `vertical and right` |
| `L`    | &#x2514; | `up and right`       |
| `0`    |          | Space                |
| `F`    |          | The last dir or file |
| `C`    |          | Not the last         |

## Scheme

```
(root)              Is final       Pseudographics
|-- .git            C              I
|-- build           C              K
|   |-- got         CC             IK
|   .-- exe         CF             IL
|-- main.go         C              K
.-- test            F              L
    |-- dir0        FC             0
    |-- dir1        FC             0K
    |   |-- file    FCC            0IK
    |   .-- file    FCF            0IL
    .-- dir2        FF             0L
```

## Rules

Children inherit FC-column from the parent directory.

IKL-column for nested dirs:
- the last `F` converts to `L`
- the last `C` converts to `K`
- not last `F` to `0`
- not last `C` to `I`
