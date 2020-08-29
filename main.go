package main

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/googleapi"
	sv1 "google.golang.org/api/storage/v1"
)

func main() {
	bucketName := "gcs-version-demo"
	filename := "license2.txt"

	// Use oauth2.NoContext if there isn't a good context to pass in.
	ctx := context.Background()
	gcs, err := sv1.NewService(ctx,  option.WithScopes(sv1.DevstorageReadWriteScope))
	if err != nil {
		log.Fatal(err)
	}
	obj, err := gcs.Objects.Get(bucketName, filename).Do()
	if e, ok := err.(*googleapi.Error); ok {
		if e.Code == 409 {
			log.Fatal(err)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(obj.Metadata)
}

func main2() {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Read the object1 from bucket.
	bucketName := "gcs-version-demo"
	filename := "license2.txt"
	rc, err := client.Bucket(bucketName).Object(filename).NewReader(ctx)
	if e, ok := err.(*googleapi.Error); ok {
		if e.Code == 409 {
			log.Fatal(err)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()
	body, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	//
	//w := client.Bucket(bucketName).Object(filename).NewWriter(ctx)
	//defer w.Close()
	//w.Write()
	//body, err := ioutil.ReadAll(rc)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(body))

}
