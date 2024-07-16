package main

import (
	"context"
	"encoding/json"
	"io"

	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

type Output struct {
	Resposta string `json:"Resposta"`
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	
	var msg Output

	msg.Resposta = "Rota1"

	json.NewEncoder(out).Encode(&msg)
}
