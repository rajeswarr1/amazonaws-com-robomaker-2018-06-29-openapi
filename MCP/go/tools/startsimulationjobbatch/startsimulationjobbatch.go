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

func StartsimulationjobbatchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/startSimulationJobBatch", cfg.BaseURL)
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
		var result models.StartSimulationJobBatchResponse
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

func CreateStartsimulationjobbatchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_startSimulationJobBatch",
		mcp.WithDescription("Starts a new simulation job batch. The batch is defined using one or more <code>SimulationJobRequest</code> objects. "),
		mcp.WithObject("batchPolicy", mcp.Description("Input parameter: Information about the batch policy.")),
		mcp.WithString("clientRequestToken", mcp.Description("Input parameter: Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")),
		mcp.WithArray("createSimulationJobRequests", mcp.Required(), mcp.Description("Input parameter: A list of simulation job requests to create in the batch.")),
		mcp.WithObject("tags", mcp.Description("Input parameter: A map that contains tag keys and tag values that are attached to the deployment job batch.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    StartsimulationjobbatchHandler(cfg),
	}
}
