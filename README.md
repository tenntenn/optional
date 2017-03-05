# Optional Values for Go

## Usage

```
type Person struct {
    Name string
    Age  int
}

type Query struct {
    Name *String
    Age  *Int
}

func search(q *Query, ps []*Person) []*Person {
    var result []*Person{}
    for _, p := range ps {
        if q.Name.Match(p.Name) && q.Age.Match(p.Age) {
            result = append(result, p)
        }
    }
    return result
}
```
