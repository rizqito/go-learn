package latihan

import "fmt"

func main(){
	i, j := 42, 2701
	p := &i //buat variabel p dengan pointer i
	fmt.Println(*p) //print i melalui pointer
	*p = 21; //pasang nilai i menjadi 21 melalui pointer
	fmt.Println(*p) //print i
	p = &j //nilai p pointer ke j
	j = (*p / 37) //bagi j dengan 7 melalui pointer
	fmt.Println(j) //print j
}