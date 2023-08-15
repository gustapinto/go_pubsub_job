package flag

import (
	"errors"
	"flag"
)

func PublisherCliFlags() (string, string, error) {
	projectName := flag.String("project", "", "The Google Cloud Project Id")
	topicName := flag.String("topic", "", "The Pub/Sub Topic Name")
	flag.Parse()

	if *projectName == "" {
		return "", "", errors.New("please specify the Google Cloud Project Id")
	}

	if *topicName == "" {
		return "", "", errors.New("please specify the Pub/Sub Topic Name")
	}

	return *projectName, *topicName, nil
}

func ConsumerCliFlags() (string, string, error) {
	projectName := flag.String("project", "", "The Google Cloud Project Id")
	subscriptionName := flag.String("subscription", "", "The Pub/Sub Topic Subscription Name")
	flag.Parse()

	if *projectName == "" {
		return "", "", errors.New("please specify the Google Cloud Project Id")
	}

	if *subscriptionName == "" {
		return "", "", errors.New("please specify the Pub/Sub Topic Subscription Name")
	}

	return *projectName, *subscriptionName, nil
}
