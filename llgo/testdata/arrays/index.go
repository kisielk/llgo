package main

func main() {
    var i [2]int
    j := &i
    i[0] = 123
    i[1] = 456
    println(i[0], i[1])
    println(j[0], j[1])
    i[0]++
    i[1]--
    println(i[0], i[1])
    println(j[0], j[1])
}

