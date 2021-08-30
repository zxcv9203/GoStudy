package main

func namedReturn(i int) (ret int) {
    ret = i
    i += 2
    return
}

func main() {
    println(namedReturn(5))
}