package main

import(
	"fmt"
	"encoding/json"
)

type ColorGroup struct {  
    ID     int  
    Name   string  
    Sc     float32 `json:"sc,float64"`
}  
func main(){
	group := ColorGroup{  
		ID:     1,  
		Name:   "Reds",  
		Sc:  34, 
	}  

	b, err := json.Marshal(group)  
	if err != nil {  
		fmt.Println("error:", err)  
	}  
	fmt.Printf("%#v\n",string(b))
}
