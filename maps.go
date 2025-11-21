package main
import (
    "fmt"
    "maps"
)

func main() {

	// let's create a map
	// make(map[key-type]val-type)
	m := make(map[string]int)

    m["key1"] = 7
    m["key2"] = 13

    fmt.Println(m)
    fmt.Println(len(m))

    delete(m, "key2")
    fmt.Println(m)
    fmt.Println(len(m))

    clear(m)
    fmt.Println(m)
    fmt.Println(len(m))

    val, prs := m["key2"]
    fmt.Println("prs:", prs, val)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println(n)

    n2 := map[string]int{"foo": 1, "bar": 2}
    
    if maps.Equal(n, n2) {
    	fmt.Println("n == n2")
    }
}