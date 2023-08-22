package main

import (
	jobapp "go_pubsub_job/internal/app/job"
	"go_pubsub_job/internal/domain/job"
	"go_pubsub_job/internal/infrastructure/ctx"
	"go_pubsub_job/internal/infrastructure/flag"
	"log"

	"cloud.google.com/go/pubsub"
)

func main() {
	projectName, subscriptionName, err := flag.ConsumerCliFlags()
	if err != nil {
		log.Fatalf("Err: %+v\n", err)
	}

	_ctx, cancel := ctx.NewTimeoutContext()
	defer cancel()

	client, err := pubsub.NewClient(_ctx, projectName)
	if err != nil {
		log.Fatalf("Err: %+v", err)
	}

	consumer := jobapp.PubSubJobConsumer{
		Client:       *client,
		Subscription: *client.Subscription(subscriptionName),
	}

	// TODO - Abstrair lógica de consumo em um service, também separar a função
	// que lida com os resultados em uma interface do tipo ResultHandler.handle(Result),
	//
	// JobService.RunJobs(ResultHandler.Handle(), JobConsumer.Consume())
	consumer.Consume(func(r job.JobState) {
		// TODO - Publicar resultados em duas tabelaa do big query, uma genérica
		// para todos os resultados e outra específica, usando o tipo de job
		// para determinar qual tabela usar
		log.Printf("Result: %+v", r)
	})
}
