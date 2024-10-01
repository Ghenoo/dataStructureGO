package main

import "fmt"


type Node struct {
    data string
    next *Node
}


func (n *Node) printNodes() {
    for n != nil {
        fmt.Println(n.data)
        n = n.next
    }
}

func main() {
    nodes := Node{
        data: "first",
        next: &Node{
            data: "second",
            next: &Node{
                data: "third",
                next: nil,
            }, 
        },
    } 
    nodes.printNodes()
}