func worker(ch chan string) {
    ch <- "done"
}

func main() {
    ch := make(chan string)

    go worker(ch)

    msg := <-ch
    fmt.Println(msg)
}

