![Clarifai logo](docs/logo.png)

# Clarifai Go gRPC Client

This is the official Clarifai gRPC Go client for interacting with our powerful recognition
[API](https://docs.clarifai.com).
Clarifai provides a platform for data scientists, developers, researchers and enterprises to master the entire 
artificial intelligence lifecycle. Gather valuable business insights from images, video and text using computer vision 
and natural language processing.

* Try the Clarifai demo at: https://clarifai.com/demo
* Sign up for a free account at: https://portal.clarifai.com/signup
* Read the documentation at: https://docs.clarifai.com/

# Installation

```
go get github.com/Clarifai/clarifai-go-grpc@INSERT_VERSION_HERE
```

## Getting started


```go
import (
	"clarifai_grpc/proto/clarifai/api"
	"clarifai_grpc/proto/clarifai/api/status"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"os"
)

baseGrpcUrl := "api.clarifai.com"
port := "443"

conn, err := grpc.Dial(baseGrpcUrl+":"+port, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
if err != nil {
    panic(err)
}
client := api.NewV2Client(conn)

apiKey := "YOUR_CLARIFAI_API_KEY_OR_PAT"
ctx := metadata.AppendToOutgoingContext(context.Background(), "Authorization", "Key "+apiKey)
```

Predict concepts in an image:

```go
response, err := client.PostModelOutputs(
    ctx,
    &api.PostModelOutputsRequest{
        ModelId: GeneralModelId,
        Inputs: []*api.Input{
            {
                Data: &api.Data{
                    Image: &api.Image{
                        Url: DogImageUrl}}}}})
if err != nil {
    panic(err)
}
if statusObj.Code != status.StatusCode_SUCCESS {
    panic(fmt.Sprintf("Unexpected status: %s", statusObj))
}

if len(response.Outputs[0].Data.Concepts) == 0 {
    t.Errorf("Received no outputs")
}
```
