package clarifai_grpc

import (
	"clarifai_grpc/proto/clarifai/api"
	"clarifai_grpc/proto/clarifai/api/status"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"os"
	"testing"
)

var GeneralModelId = "aaa03c23b3724a16a56b629203edc62c"
var DogImageUrl = "https://samples.clarifai.com/dog2.jpeg"

func TestGetModel(t *testing.T) {
	client := makeClient()
	ctx := makeContext()

	response, err := client.GetModel(ctx, &api.GetModelRequest{ModelId: GeneralModelId})
	if err != nil {
		panic(err)
	}
	assertSuccessResponse(t, response.Status)

	if response.Model.Name != "general" {
		t.Errorf("Expected model name `general`, got `%s`\n", response.Model.Name)
	}
}

func TestListModelsWithPagination(t *testing.T) {
	client := makeClient()
	ctx := makeContext()

	response, err := client.ListModels(ctx, &api.ListModelsRequest{PerPage: 2})
	if err != nil {
		panic(err)
	}
	assertSuccessResponse(t, response.Status)

	if len(response.Models) != 2 {
		t.Errorf("Expected 2 models, got %d\n", len(response.Models))
	}
}

func TestPostModelOutputsWithUrl(t *testing.T) {
	client := makeClient()
	ctx := makeContext()

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
	assertSuccessResponse(t, response.Status)

	if len(response.Outputs[0].Data.Concepts) == 0 {
		t.Errorf("Received no outputs")
	}
}

func TestPostModelOutputsWithFileBytes(t *testing.T) {
	client := makeClient()
	ctx := makeContext()

	fileBytes, err := ioutil.ReadFile("test_assets/red-truck.png")
	if err != nil {
		panic(err)
	}

	response, err := client.PostModelOutputs(
		ctx,
		&api.PostModelOutputsRequest{
			ModelId: GeneralModelId,
			Inputs: []*api.Input{
				{
					Data: &api.Data{
						Image: &api.Image{
							Base64: fileBytes}}}}})
	if err != nil {
		panic(err)
	}
	assertSuccessResponse(t, response.Status)

	if len(response.Outputs[0].Data.Concepts) == 0 {
		t.Errorf("Received no outputs")
	}
}

func makeClient() api.V2Client {
	baseGrpcUrl := os.Getenv("CLARIFAI_BASE_GRPC")
	port := "443"

	conn, err := grpc.Dial(baseGrpcUrl+":"+port, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		panic(err)
	}
	return api.NewV2Client(conn)
}

func makeContext() context.Context {
	apiKey := os.Getenv("CLARIFAI_API_KEY")
	return metadata.AppendToOutgoingContext(context.Background(), "Authorization", "Key "+apiKey)
}

func assertSuccessResponse(t *testing.T, statusObj *status.Status) {
	if statusObj.Code != status.StatusCode_SUCCESS {
		t.Errorf("Unexpected status: %s", statusObj)
	}
}
