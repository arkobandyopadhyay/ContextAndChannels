//Context Package Concept

//package main

//import (
//	"context"
//	"errors"
//	"fmt"
//	"time"
//)

//func api(ctx context.Context)(bool,error){
//	time.Sleep(400*time.Millisecond)
//
//	if ctx.Err() == context.DeadlineExceeded {
//		return false , errors.New("API Time Limit Exceeded")
//	}
//  
//	return true, nil
//}

//func main() {
//   ctx , cancel := context.WithTimeout(context.Background(),500*time.Millisecond)
//   defer cancel()
//
//   flag , error := api(ctx)
//
//   if(error!=nil){
//	  fmt.Println("API Time Limit Exceeded")
//   }
//   if(flag){
//	 fmt.Println("API is executing well")
//   }
//}

//Channels Concepts

package main

import {
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"sync"
}

const apiKey="6fad5235c96f6fc4443e48dd6a3a3c40"

func fetchWeather(city string,ch chan <- string , wq *sync.WaitGroup) interface(){
	var data struct{
		Main struct{
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	defer wq.Done()

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s" , city , apiKey)
	resp , err := http.Get(url)

	if err != nil{
		fmt.Printf("Error fetching weather for %s : %s\n" , city , err)
		return data
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body) , Decode(&data): err != nil {
		fmt.Printf("Error decoding weather data for",city, ":" , err)
		return data
	}

	ch <- fmt.Sprintf("This is the",city)

	return data
}

func main(){
	startNow := time.Now()
      
	cities := [] string {"Toronto ", "London" , "Paris" , "Tokyo"}
    
	ch = make(chan string)
	vars wq = sync.WaitGroup{}

	for _, city := range cities{
		wq.Add(1)
	    go fetchWeather(city,ch,wq)
	}

	go func(){
		wq.Wait()
		close(ch)
	}()

	for result := range ch{
		fmt.Println(result)
	}

	fmt.Println("This operation took:",time.Since(startNow))
}
