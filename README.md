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

Construct the `V2Client` object using which you access all the Clarifai API functionality, and a `Context` object which
is used for authentication - remember to insert your own Clarifai API key (or PAT).

```go
import (
	"context"
	"fmt"
	"github.com/Clarifai/clarifai-go-grpc/proto/clarifai/api"
	"github.com/Clarifai/clarifai-go-grpc/proto/clarifai/api/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

conn, err := grpc.Dial(
    "api.clarifai.com:443",
    grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
)
if err != nil {
    panic(err)
}
client := api.NewV2Client(conn)

ctx := metadata.AppendToOutgoingContext(
    context.Background(),
    "Authorization", "Key YOUR_CLARIFAI_API_KEY_OR_PAT",
)
```

Predict concepts in an image:

```go
// This is a publicly available model ID.
var GeneralModelId = "aaa03c23b3724a16a56b629203edc62c"
response, err := client.PostModelOutputs(
    ctx,
    &api.PostModelOutputsRequest{
        ModelId: GeneralModelId,
        Inputs: []*api.Input{
            {
                Data: &api.Data{
                    Image: &api.Image{
                        Url: "https://samples.clarifai.com/dog2.jpeg",
                    },
                },
            },
        },
    },
)
if err != nil {
    panic(err)
}
if response.Status.Code != status.StatusCode_SUCCESS {
    panic(fmt.Sprintf("Failed response: %s", response))
}

for _, concept := range response.Outputs[0].Data.Concepts {
    fmt.Printf("%s: %.2f\n", concept.Name, concept.Value)
}
```
