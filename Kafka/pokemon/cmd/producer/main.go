package main

import (
	"log"
	"pokemon/internal/kafka"
	"pokemon/internal/model"
)

func main() {

	brokers := []string{"localhost:9092"}
	topic := "pokemon_topic"

	p, err := kafka.NewProducer(brokers, topic)
	if err != nil {
		log.Fatalf("producer init: %v", err)
	}

	defer p.Close()

	pokemons := []model.Pokemon{
		{ID: 1, Name: "Bulbasaur", Type: "Grass"},
		{ID: 4, Name: "Charmanader", Type: "Fire"},
		{ID: 8, Name:"Wartortl", Type: "Water"},
		{ID: 25, Name:"Pikachu", Type: "Electric"},
	}


	for _, pk := range pokemons {
		part, off, err := p.SendPokemon(pk)
		if err != nil {
			log.Printf("send failed: %v", err)
		}else {
			log.Printf("send %s -> partition=%d offset=%d", pk.Name, part, off)
		}


	}

}