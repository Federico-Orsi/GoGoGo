package routes

import (
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Canal que se activará después de la duración especificada
//timerChan := time.After(duration)
// Bloquear hasta que el canal se active
//<-timerChan

// var wg *sync.WaitGroup

func hi(n int){
	time.Sleep(time.Second * 1)
	fmt.Println(n)
}

func do_something(d time.Duration, wg *sync.WaitGroup){
	fmt.Println("Doing something")
	fmt.Println("Waiting for it...")
	time.Sleep(d)
	fmt.Println("Something was done")
	wg.Done()
}

func first_channel(d time.Duration, str string,str_chan chan string){
	fmt.Println("Doing something")
	fmt.Println("Waiting for it...")
	time.Sleep(d)
	fmt.Println("Something was done")
	str_chan <- str

}

func Go_routines_Route(app *fiber.App) {

    wg := sync.WaitGroup{}


	app.Get("/routines", func(c *fiber.Ctx) error {

		now := time.Now()
        // duration := 1 * time.Second


		for i := 1; i <= 1000; i++ {
		go hi(i)

		}

	  timeSince := time.Since(now)
      fmt.Printf("El programa duró: %v segundos\n", timeSince)


		return c.JSON("go routines kick off")
	})


	app.Get("/wg", func(c *fiber.Ctx) error {

		now := time.Now()

        wg.Add(2)
		go do_something(14 * time.Second, &wg)
		go do_something(15 * time.Second, &wg)
		wg.Wait()


	    timeSince := time.Since(now)
        fmt.Printf("wg have finished properly!! El programa duró: %v segundos\n", timeSince)


		return c.JSON("Wait Groups kick off")
	})

	app.Get("/channels", func(c *fiber.Ctx) error {


		now := time.Now()
        first_chan := make(chan string)


		go first_channel(10 * time.Second, "1st Channel", first_chan)
		go first_channel(5 * time.Second, "2nd Channel", first_chan)
		go first_channel(2 * time.Second, "3rd Channel", first_chan)




        timeSince := time.Since(now)
        fmt.Printf("El programa duró: %v segundos\n", timeSince)

		//Muy importante realizar el loop sobre el canal y el cierre del mismo dentro de una go routine para que funcione correctamente!!
        go func (){
			for value := range first_chan {
				fmt.Println("Valor recibido:", value)
			}
	       close(first_chan)

		}()

		// message1 := <- first_chan
		// fmt.Println(message1)

		// message2 := <- first_chan
		// fmt.Println(message2)



		return c.JSON("Gogogoo fucking Channels")
	})

}


