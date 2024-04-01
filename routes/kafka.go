package routes

import (

	// "os"

	// "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
)



func Kafka_Route(app *fiber.App) {




	app.Get("/kafka", func(c *fiber.Ctx) error {

		// to produce messages:
		//topic := "my-topic2"
		topic := "test-topic"

		partition := 0



		conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
		if err != nil {
			log.Fatal("failed to dial leader:", err)
		}

		conn.SetWriteDeadline(time.Now().Add(10*time.Second))
		_, err = conn.WriteMessages(
			//kafka.Message{Value: []byte("kafka it´s fucking amazing!! its kafka, I just love it bro!!")},
			//kafka.Message{Value: []byte("let's fuckingggg GOooooooooooo!!!!")},
			kafka.Message{Value: []byte("We are communicating two microservices written in 2 diferent languages!!")},
			kafka.Message{Value: []byte("Come on bitchhhhh!!!!")},

		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		// if err := conn.Close(); err != nil {
		// 	log.Fatal("failed to close writer:", err)
		// }



		//-------------------------------------------------------------
        // esta es otra manera de Producir mensajes con Kafka:

		// writer := kafka.NewWriter(kafka.WriterConfig{
		// 	Brokers:  []string{"localhost:9092"},
		// 	Topic:    "test-topic",
		// 	Balancer: &kafka.LeastBytes{},
		// })

		// err := writer.WriteMessages(context.TODO(),
		// 	kafka.Message{
		// 		Key:   []byte("key"),
		// 		Value: []byte("Hello from Go!"),
		// 	},
		// )
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// fmt.Println("Message sent from Go!")

		// writer.Close()




        //----------------------------------------------
		//consume messages

		// to consume messages

		// conn.SetReadDeadline(time.Now().Add(10*time.Second))
		// batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

		// b := make([]byte, 10e3) // 10KB max per message
		// for {
		// 	n, err := batch.Read(b)
		// 	if err != nil {
		// 		break
		// 	}
		// 	fmt.Println(string(b[:n]))
		// }

		// // if err := batch.Close(); err != nil {
		// // 	log.Fatal("failed to close batch:", err)
		// // }

		// if err := conn.Close(); err != nil {
		// 	log.Fatal("failed to close connection:", err)
		// }



          //-----------------------------------------
		// config := kafka.ReaderConfig{

		// 	Brokers: []string {"localhost:9092"},
        //     Topic: topic,
		// 	GroupID: "g1",
		// 	MinBytes: 10,
		// }

		// reader := kafka.NewReader(config)


        // for {

        //    message, err := reader.ReadMessage(context.Background())
		//    if err != nil {
		// 	   fmt.Println("Some error ocurred", err)
		// 	   continue
		//    }

		//    fmt.Println("The message is... : ", string(message.Value))


		// }


        //-----------------------------------------
        //esto era para Confluent kafka go
		// p, err := kafka.NewProducer(&kafka.ConfigMap{
		// 	"bootstrap.servers": "localhost:9092",
		// 	"client.id": "something",
		// 	"acks": "all"})

		// if err != nil {
		// 	fmt.Printf("Failed to create producer: %s\n", err)
		// 	// os.Exit(1)
		// }

		// fmt.Printf("%v\n", p)

		return c.JSON("Trying kafka in Go")
	})
}

