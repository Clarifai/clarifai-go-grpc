package clarifai_grpc

import (
	"clarifai_grpc/proto/clarifai/api"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"os"
	"testing"
)

func TestGetModel(t *testing.T) {
	GENERAL_MODEL_ID := "aaa03c23b3724a16a56b629203edc62c"

	client := makeClient()
	ctx := makeContext()

	getModelResponse, err := client.GetModel(ctx, &api.GetModelRequest{ModelId: GENERAL_MODEL_ID})
	if err != nil {
		panic(err)
	}
	if getModelResponse.Model.Name != "general" {
		t.Errorf("Expected model name `general`, got `%s`\n", getModelResponse.Model.Name)
	}
}

func TestListModelsWithPagination(t *testing.T) {
	client := makeClient()
	ctx := makeContext()

	listModelsResponse, err := client.ListModels(ctx, &api.ListModelsRequest{PerPage: 2})
	if err != nil {
		panic(err)
	}
	if len(listModelsResponse.Models) != 2 {
		t.Errorf("Expected 2 models, got %d\n", len(listModelsResponse.Models))
	}
}

func makeClient() api.V2Client {
	baseGrpcUrl := os.Getenv("CLARIFAI_BASE_GRPC")
	port := "443"

	conn, err := grpc.Dial(baseGrpcUrl + ":" + port, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		panic(err)
	}
	return api.NewV2Client(conn)
}

func makeContext() context.Context {
	apiKey := os.Getenv("CLARIFAI_API_KEY")
	return metadata.AppendToOutgoingContext(context.Background(), "Authorization", "Key " + apiKey)
}
