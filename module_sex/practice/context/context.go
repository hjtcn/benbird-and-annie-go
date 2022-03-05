package main

import (
	"context"
	"encoding/json"
)

func main() {

}

func withCancel() {
	context.TODO()
	context.Background()
	json.Marshal(struct {
		Name string
	}{})
}
