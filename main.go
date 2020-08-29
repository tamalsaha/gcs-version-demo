package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"cloud.google.com/go/storage"
)

func main() {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Read the object1 from bucket.
	bucketName := "gcs-version-demo"
	filename := "license.txt"
	rc, err := client.Bucket(bucketName).Object(filename).NewReader(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()
	body, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))


	w := client.Bucket(bucketName).Object(filename).NewWriter(ctx)
	defer w.Close()
	w.Write()
	body, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

}
