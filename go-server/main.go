package main

import{
	"fmt"
	"log"
	"net/http"
}

func main(){
	fileserver := http.fileserver(http.Dir("./static"))
	

}