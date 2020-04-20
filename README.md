# Cosine similarity calculation for Golang

GoLang implementation of [sklearn.metrics.pairwise.cosine_similarity](https://scikit-learn.org/stable/modules/generated/sklearn.metrics.pairwise.cosine_similarity.html). It Computes cosine similarity between samples in X and Y. Cosine similarity, or the cosine kernel, computes similarity as the normalized dot product of X and Y:
```
K(X, Y) = <X, Y> / (||X||*||Y||)
```

## Installation

```bash
go get github.com/khaibin/go-cosinesimilarity
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/khaibin/go-cosinesimilarity"
)

func main() {
    x := [][]float64{{1, 2, 3}, {8, 10, 15}}
    y := [][]float64{{1, 3, 4}}
    result := cosinesimilarity.Compute(x, y)
    fmt.Println(result)
}
```

## License
[MIT](https://choosealicense.com/licenses/mit/)