package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func filter(arr []string, pos int) ([]string, []string){
	var res1, res2 []string
	for _, str:=range arr {
		if str[pos]=='0' {
			res1=append(res1,str)
		}else {
			res2=append(res2,str)
		}
	}
	return res1, res2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//var a [12]int
	var b []string
	for scanner.Scan() {
		line:=scanner.Text()
		b=append(b,line)
		//for i:=0; i<12; i++{
		//	if line[i]=='1'{
		//		a[i]++
		//	}else{
		//		a[i]--
		//	}
		//}
	}
	// fmt.Println(a)
	defer file.Close()
	var r1, r2 []string
	for i:=0;i<12;i++ {
		if i==0 {
			l1,l2:=filter(b,i)
			if len(l1)>len(l2) {
				r1=l1
				r2=l2
			} else {
				r1=l2
				r2=l1
			}
		} else {
			if (len(r1)==1) && (len(r2)==1) {
				break
			}
			if len(r1)!=1 {
				l1,l2:=filter(r1,i)
				if len(l1)>len(l2) {
					r1=l1
				} else {
					r1=l2
				}
			}
			if len(r2)!=1 {
				l1,l2:=filter(r2,i)
				if len(l1)>len(l2) {
					r2=l2
				} else {
					r2=l1
				}
			}
		}
	}
	fmt.Println(r1,r2)
}