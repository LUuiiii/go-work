package main
import "fmt"
func kaisamima(n int,ora string ) string{
	n = n
	var target string
	var ans string
	target ="abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(ora); i++ {
 
	   for j := 0; j < len(target); j++ {
		  if ora[i]==target[j]{
	   
 
			 out_of_bounds:=(j+n)/26
 
			 if out_of_bounds>0{
				ans+=string(target[j+n-out_of_bounds*26])
				break
			 }
			 ans+=string(target[j+n])
			 break
 
		  }
 
	   }
 
	}
	return ans
 }

 func main(){
	fmt.Println(kaisamima(1,"z"));
	
 }