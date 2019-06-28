package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/iterator"
)

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		panic(fmt.Sprintf("%s environment variable not set.", k))
	}
	return v
}

func getSlugSnap(ctx context.Context, collection string, slug string) map[string]interface{} {
	var (
		collectionref = firestoreClient.Collection(collection)
		iter          = collectionref.Where("slug", "==", slug).Documents(ctx)
		result        map[string]interface{}
	)

	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logger.Errorf("unable to retrieve document for slug: %s error was: %s", slug, err.Error())
		}
		result = doc.Data()
	}
	return result
}
