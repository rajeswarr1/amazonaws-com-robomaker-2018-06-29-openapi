package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/aws-robomaker/mcp-server/config"
	"github.com/aws-robomaker/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreatesimulationjobHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/createSimulationJob", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.CreateSimulationJobResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCreatesimulationjobTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_createSimulationJob",
		mcp.WithDescription("<p>Creates a simulation job.</p> <note> <p>After 90 days, simulation jobs expire and will be deleted. They will no longer be accessible. </p> </note>"),
		mcp.WithString("clientRequestToken", mcp.Description("Input parameter: Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")),
		mcp.WithArray("dataSources", mcp.Description("Input parameter: <p>Specify data sources to mount read-only files from S3 into your simulation. These files are available under <code>/opt/robomaker/datasources/data_source_name</code>. </p> <note> <p>There is a limit of 100 files and a combined size of 25GB for all <code>DataSourceConfig</code> objects. </p> </note>")),
		mcp.WithObject("outputLocation", mcp.Description("Input parameter: The output location.")),
		mcp.WithArray("robotApplications", mcp.Description("Input parameter: The robot application to use in the simulation job.")),
		mcp.WithArray("simulationApplications", mcp.Description("Input parameter: The simulation application to use in the simulation job.")),
		mcp.WithObject("vpcConfig", mcp.Description("Input parameter: If your simulation job accesses resources in a VPC, you provide this parameter identifying the list of security group IDs and subnet IDs. These must belong to the same VPC. You must provide at least one security group and two subnet IDs.")),
		mcp.WithString("iamRole", mcp.Required(), mcp.Description("Input parameter: The IAM role name that allows the simulation instance to call the AWS APIs that are specified in its associated policies on your behalf. This is how credentials are passed in to your simulation job. ")),
		mcp.WithObject("loggingConfig", mcp.Description("Input parameter: The logging configuration.")),
		mcp.WithObject("compute", mcp.Description("Input parameter: Compute information for the simulation job.")),
		mcp.WithString("failureBehavior", mcp.Description("Input parameter: <p>The failure behavior the simulation job.</p> <dl> <dt>Continue</dt> <dd> <p>Leaves the instance running for its maximum timeout duration after a <code>4XX</code> error code.</p> </dd> <dt>Fail</dt> <dd> <p>Stop the simulation job and terminate the instance.</p> </dd> </dl>")),
		mcp.WithNumber("maxJobDurationInSeconds", mcp.Required(), mcp.Description("Input parameter: The maximum simulation job duration in seconds (up to 14 days or 1,209,600 seconds. When <code>maxJobDurationInSeconds</code> is reached, the simulation job will status will transition to <code>Completed</code>.")),
		mcp.WithObject("tags", mcp.Description("Input parameter: A map that contains tag keys and tag values that are attached to the simulation job.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreatesimulationjobHandler(cfg),
	}
}
