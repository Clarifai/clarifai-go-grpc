package tests

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/Clarifai/clarifai-go-grpc/proto/clarifai/api"
	"github.com/Clarifai/clarifai-go-grpc/proto/clarifai/api/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

var GeneralModelId = "aaa03c23b3724a16a56b629203edc62c"
var DogImageUrl = "https://samples.clarifai.com/dog2.jpeg"
var NonExistingUrl = "https://example.com/non-existing.jpg"

func TestGetModel(t *testing.T) {
	client := makeClient()
	ctx := makeContext()

	response, err := client.GetModel(ctx, &api.GetModelRequest{ModelId: GeneralModelId})
	check(err)
	assertSuccessResponse(t, response.Status)

	if response.Model.Name != "Image Recognition" {
		t.Errorf("Expected model name `Image Recognition`, got `%s`\n", response.Model.Name)
	}
}

func TestListModelsWithPagination(t *testing.T) {
	client := makeClient()
	ctx := makeContext()

	response, err := client.ListModels(ctx, &api.ListModelsRequest{PerPage: 2})
	check(err)
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
							Url: DogImageUrl,
						},
					},
				},
			},
		},
	)
	check(err)
	assertSuccessResponse(t, response.Status)

	if len(response.Outputs[0].Data.Concepts) == 0 {
		t.Errorf("Received no outputs")
	}
}

func TestPostModelOutputsWithFileBytes(t *testing.T) {
	client := makeClient()
	ctx := makeContext()

	fileBytes, err := ioutil.ReadFile("test_assets/red-truck.png")
	check(err)

	response, err := client.PostModelOutputs(
		ctx,
		&api.PostModelOutputsRequest{
			ModelId: GeneralModelId,
			Inputs: []*api.Input{
				{
					Data: &api.Data{
						Image: &api.Image{
							Base64: fileBytes,
						},
					},
				},
			},
		},
	)
	check(err)
	assertSuccessResponse(t, response.Status)

	if len(response.Outputs[0].Data.Concepts) == 0 {
		t.Errorf("Received no outputs")
	}
}

func TestFailedPostModelOutputs(t *testing.T) {
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
							Url: NonExistingUrl,
						},
					},
				},
			},
		},
	)
	check(err)
	if response.Status.Code != status.StatusCode_FAILURE {
		t.Errorf("Expected FAILURE response, got: %s\n", response.Status)
	}
	if response.Outputs[0].Status.Code != status.StatusCode_INPUT_DOWNLOAD_FAILED {
		t.Errorf("Expected INPUT_DOWNLOAD_FAILED response, got: %s\n", response.Status)
	}
}

func TestMixedStatusPostModelOutputs(t *testing.T) {
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
							Url: DogImageUrl,
						},
					},
				},
				{
					Data: &api.Data{
						Image: &api.Image{
							Url: NonExistingUrl,
						},
					},
				},
			},
		},
	)
	check(err)
	if response.Status.Code != status.StatusCode_MIXED_STATUS {
		t.Errorf("Expected MIXED_STATUS response, got: %s\n", response.Status)
	}
	if response.Outputs[0].Status.Code != status.StatusCode_SUCCESS {
		t.Errorf("Expected INPUT_DOWNLOAD_SUCCESS response, got: %s\n", response.Status)
	}
	if response.Outputs[1].Status.Code != status.StatusCode_INPUT_DOWNLOAD_FAILED {
		t.Errorf("Expected INPUT_DOWNLOAD_FAILED response, got: %s\n", response.Status)
	}
}

func TestPostPatchAndDeleteInput(t *testing.T) {
	client := makeClient()
	ctx := makeContext()

	postInputsResponse, err := client.PostInputs(
		ctx,
		&api.PostInputsRequest{
			Inputs: []*api.Input{
				{
					Data: &api.Data{
						Image: &api.Image{
							Url:               DogImageUrl,
							AllowDuplicateUrl: true,
						},
					},
				},
			},
		},
	)
	check(err)
	assertSuccessResponse(t, postInputsResponse.Status)

	inputId := postInputsResponse.Inputs[0].Id
	for true {
		getInputResponse, err := client.GetInput(ctx, &api.GetInputRequest{InputId: inputId})
		check(err)
		assertSuccessResponse(t, getInputResponse.Status)

		inputStatusCode := getInputResponse.Input.Status.Code
		if inputStatusCode == status.StatusCode_INPUT_DOWNLOAD_SUCCESS {
			break
		}
		if inputStatusCode != status.StatusCode_INPUT_DOWNLOAD_PENDING && inputStatusCode != status.StatusCode_INPUT_DOWNLOAD_IN_PROGRESS {
			t.Errorf("Waiting for input ID %s failed, status code is %s", inputId, inputStatusCode)
		}
		time.Sleep(1 * time.Second)
	}

	patchInputsResponse, err := client.PatchInputs(
		ctx,
		&api.PatchInputsRequest{
			Action: "overwrite",
			Inputs: []*api.Input{
				{
					Id: inputId,
					Data: &api.Data{
						Concepts: []*api.Concept{
							{
								Id: "very-red-truck",
							},
						},
					},
				},
			},
		},
	)
	check(err)
	assertSuccessResponse(t, patchInputsResponse.Status)

	deleteInputResponse, err := client.DeleteInput(ctx, &api.DeleteInputRequest{InputId: inputId})
	check(err)
	assertSuccessResponse(t, deleteInputResponse.Status)
}

func makeClient() api.V2Client {
	grpcBaseUrl := os.Getenv("CLARIFAI_GRPC_BASE")
	if len(grpcBaseUrl) == 0 {
		grpcBaseUrl = "api.clarifai.com"
	}
	port := "443"
	conn, err := grpc.Dial(grpcBaseUrl+":"+port, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	check(err)
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
