package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// CancelSimulationJobBatchResponse represents the CancelSimulationJobBatchResponse schema from the OpenAPI specification
type CancelSimulationJobBatchResponse struct {
}

// ListSimulationJobsResponse represents the ListSimulationJobsResponse schema from the OpenAPI specification
type ListSimulationJobsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Simulationjobsummaries interface{} `json:"simulationJobSummaries"`
}

// RobotApplicationConfig represents the RobotApplicationConfig schema from the OpenAPI specification
type RobotApplicationConfig struct {
	Applicationversion interface{} `json:"applicationVersion,omitempty"`
	Launchconfig interface{} `json:"launchConfig"`
	Tools interface{} `json:"tools,omitempty"`
	Uploadconfigurations interface{} `json:"uploadConfigurations,omitempty"`
	Usedefaulttools interface{} `json:"useDefaultTools,omitempty"`
	Usedefaultuploadconfigurations interface{} `json:"useDefaultUploadConfigurations,omitempty"`
	Application interface{} `json:"application"`
}

// DescribeDeploymentJobResponse represents the DescribeDeploymentJobResponse schema from the OpenAPI specification
type DescribeDeploymentJobResponse struct {
	Fleet interface{} `json:"fleet,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Deploymentconfig interface{} `json:"deploymentConfig,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Failurereason interface{} `json:"failureReason,omitempty"`
	Robotdeploymentsummary interface{} `json:"robotDeploymentSummary,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Deploymentapplicationconfigs interface{} `json:"deploymentApplicationConfigs,omitempty"`
}

// ListSimulationJobBatchesResponse represents the ListSimulationJobBatchesResponse schema from the OpenAPI specification
type ListSimulationJobBatchesResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Simulationjobbatchsummaries interface{} `json:"simulationJobBatchSummaries,omitempty"`
}

// ListWorldsRequest represents the ListWorldsRequest schema from the OpenAPI specification
type ListWorldsRequest struct {
	Filters interface{} `json:"filters,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListWorldsResponse represents the ListWorldsResponse schema from the OpenAPI specification
type ListWorldsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Worldsummaries interface{} `json:"worldSummaries,omitempty"`
}

// RobotApplicationSummary represents the RobotApplicationSummary schema from the OpenAPI specification
type RobotApplicationSummary struct {
	Version interface{} `json:"version,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Robotsoftwaresuite interface{} `json:"robotSoftwareSuite,omitempty"`
}

// DescribeSimulationJobBatchResponse represents the DescribeSimulationJobBatchResponse schema from the OpenAPI specification
type DescribeSimulationJobBatchResponse struct {
	Arn interface{} `json:"arn,omitempty"`
	Batchpolicy interface{} `json:"batchPolicy,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Failedrequests interface{} `json:"failedRequests,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
	Failurereason interface{} `json:"failureReason,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Createdrequests interface{} `json:"createdRequests,omitempty"`
	Pendingrequests interface{} `json:"pendingRequests,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
}

// SyncDeploymentJobRequest represents the SyncDeploymentJobRequest schema from the OpenAPI specification
type SyncDeploymentJobRequest struct {
	Clientrequesttoken interface{} `json:"clientRequestToken"`
	Fleet interface{} `json:"fleet"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// ListWorldGenerationJobsResponse represents the ListWorldGenerationJobsResponse schema from the OpenAPI specification
type ListWorldGenerationJobsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Worldgenerationjobsummaries interface{} `json:"worldGenerationJobSummaries"`
}

// BatchPolicy represents the BatchPolicy schema from the OpenAPI specification
type BatchPolicy struct {
	Timeoutinseconds interface{} `json:"timeoutInSeconds,omitempty"`
	Maxconcurrency interface{} `json:"maxConcurrency,omitempty"`
}

// Source represents the Source schema from the OpenAPI specification
type Source struct {
	S3key interface{} `json:"s3Key,omitempty"`
	Architecture interface{} `json:"architecture,omitempty"`
	Etag interface{} `json:"etag,omitempty"`
	S3bucket interface{} `json:"s3Bucket,omitempty"`
}

// VPCConfigResponse represents the VPCConfigResponse schema from the OpenAPI specification
type VPCConfigResponse struct {
	Subnets interface{} `json:"subnets,omitempty"`
	Vpcid interface{} `json:"vpcId,omitempty"`
	Assignpublicip interface{} `json:"assignPublicIp,omitempty"`
	Securitygroups interface{} `json:"securityGroups,omitempty"`
}

// CreateRobotApplicationResponse represents the CreateRobotApplicationResponse schema from the OpenAPI specification
type CreateRobotApplicationResponse struct {
	Arn interface{} `json:"arn,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Robotsoftwaresuite interface{} `json:"robotSoftwareSuite,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Environment interface{} `json:"environment,omitempty"`
	Revisionid interface{} `json:"revisionId,omitempty"`
	Sources interface{} `json:"sources,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// PortForwardingConfig represents the PortForwardingConfig schema from the OpenAPI specification
type PortForwardingConfig struct {
	Portmappings interface{} `json:"portMappings,omitempty"`
}

// OutputLocation represents the OutputLocation schema from the OpenAPI specification
type OutputLocation struct {
	S3bucket interface{} `json:"s3Bucket,omitempty"`
	S3prefix interface{} `json:"s3Prefix,omitempty"`
}

// BatchDeleteWorldsResponse represents the BatchDeleteWorldsResponse schema from the OpenAPI specification
type BatchDeleteWorldsResponse struct {
	Unprocessedworlds interface{} `json:"unprocessedWorlds,omitempty"`
}

// Filter represents the Filter schema from the OpenAPI specification
type Filter struct {
	Name interface{} `json:"name,omitempty"`
	Values interface{} `json:"values,omitempty"`
}

// SimulationSoftwareSuite represents the SimulationSoftwareSuite schema from the OpenAPI specification
type SimulationSoftwareSuite struct {
	Name interface{} `json:"name,omitempty"`
	Version interface{} `json:"version,omitempty"`
}

// Environment represents the Environment schema from the OpenAPI specification
type Environment struct {
	Uri interface{} `json:"uri,omitempty"`
}

// CreateDeploymentJobResponse represents the CreateDeploymentJobResponse schema from the OpenAPI specification
type CreateDeploymentJobResponse struct {
	Deploymentconfig interface{} `json:"deploymentConfig,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Deploymentapplicationconfigs interface{} `json:"deploymentApplicationConfigs,omitempty"`
	Failurereason interface{} `json:"failureReason,omitempty"`
	Fleet interface{} `json:"fleet,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
}

// DescribeSimulationApplicationResponse represents the DescribeSimulationApplicationResponse schema from the OpenAPI specification
type DescribeSimulationApplicationResponse struct {
	Sources interface{} `json:"sources,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Renderingengine interface{} `json:"renderingEngine,omitempty"`
	Robotsoftwaresuite interface{} `json:"robotSoftwareSuite,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Environment interface{} `json:"environment,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Revisionid interface{} `json:"revisionId,omitempty"`
	Simulationsoftwaresuite interface{} `json:"simulationSoftwareSuite,omitempty"`
	Imagedigest interface{} `json:"imageDigest,omitempty"`
}

// Robot represents the Robot schema from the OpenAPI specification
type Robot struct {
	Lastdeploymentjob interface{} `json:"lastDeploymentJob,omitempty"`
	Architecture interface{} `json:"architecture,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Greengrassgroupid interface{} `json:"greenGrassGroupId,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Fleetarn interface{} `json:"fleetArn,omitempty"`
	Lastdeploymenttime interface{} `json:"lastDeploymentTime,omitempty"`
}

// DescribeFleetRequest represents the DescribeFleetRequest schema from the OpenAPI specification
type DescribeFleetRequest struct {
	Fleet interface{} `json:"fleet"`
}

// BatchDescribeSimulationJobResponse represents the BatchDescribeSimulationJobResponse schema from the OpenAPI specification
type BatchDescribeSimulationJobResponse struct {
	Jobs interface{} `json:"jobs,omitempty"`
	Unprocessedjobs interface{} `json:"unprocessedJobs,omitempty"`
}

// CancelSimulationJobRequest represents the CancelSimulationJobRequest schema from the OpenAPI specification
type CancelSimulationJobRequest struct {
	Job interface{} `json:"job"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"tags"`
}

// DeleteRobotRequest represents the DeleteRobotRequest schema from the OpenAPI specification
type DeleteRobotRequest struct {
	Robot interface{} `json:"robot"`
}

// TemplateSummary represents the TemplateSummary schema from the OpenAPI specification
type TemplateSummary struct {
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Version interface{} `json:"version,omitempty"`
}

// Fleet represents the Fleet schema from the OpenAPI specification
type Fleet struct {
	Createdat interface{} `json:"createdAt,omitempty"`
	Lastdeploymentjob interface{} `json:"lastDeploymentJob,omitempty"`
	Lastdeploymentstatus interface{} `json:"lastDeploymentStatus,omitempty"`
	Lastdeploymenttime interface{} `json:"lastDeploymentTime,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
}

// WorldSummary represents the WorldSummary schema from the OpenAPI specification
type WorldSummary struct {
	Template interface{} `json:"template,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Generationjob interface{} `json:"generationJob,omitempty"`
}

// DescribeSimulationJobResponse represents the DescribeSimulationJobResponse schema from the OpenAPI specification
type DescribeSimulationJobResponse struct {
	Simulationapplications interface{} `json:"simulationApplications,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
	Laststartedat interface{} `json:"lastStartedAt,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Robotapplications interface{} `json:"robotApplications,omitempty"`
	Maxjobdurationinseconds interface{} `json:"maxJobDurationInSeconds,omitempty"`
	Outputlocation interface{} `json:"outputLocation,omitempty"`
	Compute interface{} `json:"compute,omitempty"`
	Datasources interface{} `json:"dataSources,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Loggingconfig interface{} `json:"loggingConfig,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Failurereason interface{} `json:"failureReason,omitempty"`
	Vpcconfig interface{} `json:"vpcConfig,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Simulationtimemillis interface{} `json:"simulationTimeMillis,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Failurebehavior interface{} `json:"failureBehavior,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Networkinterface interface{} `json:"networkInterface,omitempty"`
	Iamrole interface{} `json:"iamRole,omitempty"`
}

// DeploymentApplicationConfig represents the DeploymentApplicationConfig schema from the OpenAPI specification
type DeploymentApplicationConfig struct {
	Application interface{} `json:"application"`
	Applicationversion interface{} `json:"applicationVersion"`
	Launchconfig interface{} `json:"launchConfig"`
}

// FailedCreateSimulationJobRequest represents the FailedCreateSimulationJobRequest schema from the OpenAPI specification
type FailedCreateSimulationJobRequest struct {
	Failedat interface{} `json:"failedAt,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
	Failurereason interface{} `json:"failureReason,omitempty"`
	Request interface{} `json:"request,omitempty"`
}

// ListSimulationApplicationsResponse represents the ListSimulationApplicationsResponse schema from the OpenAPI specification
type ListSimulationApplicationsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Simulationapplicationsummaries interface{} `json:"simulationApplicationSummaries,omitempty"`
}

// CreateRobotApplicationVersionResponse represents the CreateRobotApplicationVersionResponse schema from the OpenAPI specification
type CreateRobotApplicationVersionResponse struct {
	Environment interface{} `json:"environment,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Revisionid interface{} `json:"revisionId,omitempty"`
	Robotsoftwaresuite interface{} `json:"robotSoftwareSuite,omitempty"`
	Sources interface{} `json:"sources,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
}

// CancelDeploymentJobRequest represents the CancelDeploymentJobRequest schema from the OpenAPI specification
type CancelDeploymentJobRequest struct {
	Job interface{} `json:"job"`
}

// SimulationJob represents the SimulationJob schema from the OpenAPI specification
type SimulationJob struct {
	Arn interface{} `json:"arn,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Maxjobdurationinseconds interface{} `json:"maxJobDurationInSeconds,omitempty"`
	Robotapplications interface{} `json:"robotApplications,omitempty"`
	Simulationapplications interface{} `json:"simulationApplications,omitempty"`
	Compute interface{} `json:"compute,omitempty"`
	Datasources interface{} `json:"dataSources,omitempty"`
	Failurereason interface{} `json:"failureReason,omitempty"`
	Iamrole interface{} `json:"iamRole,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Failurebehavior interface{} `json:"failureBehavior,omitempty"`
	Loggingconfig interface{} `json:"loggingConfig,omitempty"`
	Vpcconfig interface{} `json:"vpcConfig,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
	Simulationtimemillis interface{} `json:"simulationTimeMillis,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Networkinterface interface{} `json:"networkInterface,omitempty"`
	Laststartedat interface{} `json:"lastStartedAt,omitempty"`
	Outputlocation interface{} `json:"outputLocation,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// UpdateRobotApplicationResponse represents the UpdateRobotApplicationResponse schema from the OpenAPI specification
type UpdateRobotApplicationResponse struct {
	Revisionid interface{} `json:"revisionId,omitempty"`
	Robotsoftwaresuite interface{} `json:"robotSoftwareSuite,omitempty"`
	Sources interface{} `json:"sources,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Environment interface{} `json:"environment,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// WorldConfig represents the WorldConfig schema from the OpenAPI specification
type WorldConfig struct {
	World interface{} `json:"world,omitempty"`
}

// RobotSoftwareSuite represents the RobotSoftwareSuite schema from the OpenAPI specification
type RobotSoftwareSuite struct {
	Name interface{} `json:"name,omitempty"`
	Version interface{} `json:"version,omitempty"`
}

// LoggingConfig represents the LoggingConfig schema from the OpenAPI specification
type LoggingConfig struct {
	Recordallrostopics interface{} `json:"recordAllRosTopics,omitempty"`
}

// DescribeWorldExportJobResponse represents the DescribeWorldExportJobResponse schema from the OpenAPI specification
type DescribeWorldExportJobResponse struct {
	Failurecode interface{} `json:"failureCode,omitempty"`
	Failurereason interface{} `json:"failureReason,omitempty"`
	Worlds interface{} `json:"worlds,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Iamrole interface{} `json:"iamRole,omitempty"`
	Outputlocation OutputLocation `json:"outputLocation,omitempty"` // The output location.
	Tags interface{} `json:"tags,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
}

// RegisterRobotResponse represents the RegisterRobotResponse schema from the OpenAPI specification
type RegisterRobotResponse struct {
	Robot interface{} `json:"robot,omitempty"`
	Fleet interface{} `json:"fleet,omitempty"`
}

// ListSimulationJobBatchesRequest represents the ListSimulationJobBatchesRequest schema from the OpenAPI specification
type ListSimulationJobBatchesRequest struct {
	Filters interface{} `json:"filters,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListSimulationApplicationsRequest represents the ListSimulationApplicationsRequest schema from the OpenAPI specification
type ListSimulationApplicationsRequest struct {
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Versionqualifier interface{} `json:"versionQualifier,omitempty"`
	Filters interface{} `json:"filters,omitempty"`
}

// CreateRobotApplicationVersionRequest represents the CreateRobotApplicationVersionRequest schema from the OpenAPI specification
type CreateRobotApplicationVersionRequest struct {
	Application interface{} `json:"application"`
	Currentrevisionid interface{} `json:"currentRevisionId,omitempty"`
	Imagedigest interface{} `json:"imageDigest,omitempty"`
	S3etags interface{} `json:"s3Etags,omitempty"`
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// UpdateWorldTemplateResponse represents the UpdateWorldTemplateResponse schema from the OpenAPI specification
type UpdateWorldTemplateResponse struct {
	Name interface{} `json:"name,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
}

// CancelSimulationJobResponse represents the CancelSimulationJobResponse schema from the OpenAPI specification
type CancelSimulationJobResponse struct {
}

// CreateWorldGenerationJobResponse represents the CreateWorldGenerationJobResponse schema from the OpenAPI specification
type CreateWorldGenerationJobResponse struct {
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
	Worldcount interface{} `json:"worldCount,omitempty"`
	Worldtags interface{} `json:"worldTags,omitempty"`
	Template interface{} `json:"template,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// S3Object represents the S3Object schema from the OpenAPI specification
type S3Object struct {
	Bucket interface{} `json:"bucket"`
	Etag interface{} `json:"etag,omitempty"`
	Key interface{} `json:"key"`
}

// S3KeyOutput represents the S3KeyOutput schema from the OpenAPI specification
type S3KeyOutput struct {
	Etag interface{} `json:"etag,omitempty"`
	S3key interface{} `json:"s3Key,omitempty"`
}

// DeregisterRobotRequest represents the DeregisterRobotRequest schema from the OpenAPI specification
type DeregisterRobotRequest struct {
	Fleet interface{} `json:"fleet"`
	Robot interface{} `json:"robot"`
}

// DescribeWorldExportJobRequest represents the DescribeWorldExportJobRequest schema from the OpenAPI specification
type DescribeWorldExportJobRequest struct {
	Job interface{} `json:"job"`
}

// DeleteSimulationApplicationResponse represents the DeleteSimulationApplicationResponse schema from the OpenAPI specification
type DeleteSimulationApplicationResponse struct {
}

// CancelWorldExportJobResponse represents the CancelWorldExportJobResponse schema from the OpenAPI specification
type CancelWorldExportJobResponse struct {
}

// DescribeWorldTemplateResponse represents the DescribeWorldTemplateResponse schema from the OpenAPI specification
type DescribeWorldTemplateResponse struct {
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
}

// ListWorldExportJobsRequest represents the ListWorldExportJobsRequest schema from the OpenAPI specification
type ListWorldExportJobsRequest struct {
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Filters interface{} `json:"filters,omitempty"`
}

// NetworkInterface represents the NetworkInterface schema from the OpenAPI specification
type NetworkInterface struct {
	Privateipaddress interface{} `json:"privateIpAddress,omitempty"`
	Publicipaddress interface{} `json:"publicIpAddress,omitempty"`
	Networkinterfaceid interface{} `json:"networkInterfaceId,omitempty"`
}

// CancelWorldGenerationJobResponse represents the CancelWorldGenerationJobResponse schema from the OpenAPI specification
type CancelWorldGenerationJobResponse struct {
}

// ListRobotsRequest represents the ListRobotsRequest schema from the OpenAPI specification
type ListRobotsRequest struct {
	Filters interface{} `json:"filters,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ComputeResponse represents the ComputeResponse schema from the OpenAPI specification
type ComputeResponse struct {
	Computetype interface{} `json:"computeType,omitempty"`
	Gpuunitlimit interface{} `json:"gpuUnitLimit,omitempty"`
	Simulationunitlimit interface{} `json:"simulationUnitLimit,omitempty"`
}

// ListWorldExportJobsResponse represents the ListWorldExportJobsResponse schema from the OpenAPI specification
type ListWorldExportJobsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Worldexportjobsummaries interface{} `json:"worldExportJobSummaries"`
}

// DeleteRobotApplicationResponse represents the DeleteRobotApplicationResponse schema from the OpenAPI specification
type DeleteRobotApplicationResponse struct {
}

// DataSource represents the DataSource schema from the OpenAPI specification
type DataSource struct {
	Destination interface{} `json:"destination,omitempty"`
	Name interface{} `json:"name,omitempty"`
	S3bucket interface{} `json:"s3Bucket,omitempty"`
	S3keys interface{} `json:"s3Keys,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
}

// RenderingEngine represents the RenderingEngine schema from the OpenAPI specification
type RenderingEngine struct {
	Version interface{} `json:"version,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// WorldGenerationJobSummary represents the WorldGenerationJobSummary schema from the OpenAPI specification
type WorldGenerationJobSummary struct {
	Succeededworldcount interface{} `json:"succeededWorldCount,omitempty"`
	Template interface{} `json:"template,omitempty"`
	Worldcount interface{} `json:"worldCount,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Failedworldcount interface{} `json:"failedWorldCount,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// UploadConfiguration represents the UploadConfiguration schema from the OpenAPI specification
type UploadConfiguration struct {
	Uploadbehavior interface{} `json:"uploadBehavior"`
	Name interface{} `json:"name"`
	Path interface{} `json:"path"`
}

// CreateFleetRequest represents the CreateFleetRequest schema from the OpenAPI specification
type CreateFleetRequest struct {
	Name interface{} `json:"name"`
	Tags interface{} `json:"tags,omitempty"`
}

// DeploymentJob represents the DeploymentJob schema from the OpenAPI specification
type DeploymentJob struct {
	Createdat interface{} `json:"createdAt,omitempty"`
	Deploymentapplicationconfigs interface{} `json:"deploymentApplicationConfigs,omitempty"`
	Deploymentconfig interface{} `json:"deploymentConfig,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
	Failurereason interface{} `json:"failureReason,omitempty"`
	Fleet interface{} `json:"fleet,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
}

// CancelWorldGenerationJobRequest represents the CancelWorldGenerationJobRequest schema from the OpenAPI specification
type CancelWorldGenerationJobRequest struct {
	Job interface{} `json:"job"`
}

// CreateWorldExportJobRequest represents the CreateWorldExportJobRequest schema from the OpenAPI specification
type CreateWorldExportJobRequest struct {
	Iamrole interface{} `json:"iamRole"`
	Outputlocation OutputLocation `json:"outputLocation"` // The output location.
	Tags interface{} `json:"tags,omitempty"`
	Worlds interface{} `json:"worlds"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
}

// DescribeWorldTemplateRequest represents the DescribeWorldTemplateRequest schema from the OpenAPI specification
type DescribeWorldTemplateRequest struct {
	Template interface{} `json:"template"`
}

// CreateSimulationApplicationResponse represents the CreateSimulationApplicationResponse schema from the OpenAPI specification
type CreateSimulationApplicationResponse struct {
	Renderingengine interface{} `json:"renderingEngine,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Environment interface{} `json:"environment,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Simulationsoftwaresuite interface{} `json:"simulationSoftwareSuite,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Revisionid interface{} `json:"revisionId,omitempty"`
	Robotsoftwaresuite interface{} `json:"robotSoftwareSuite,omitempty"`
	Sources interface{} `json:"sources,omitempty"`
}

// CreateRobotApplicationRequest represents the CreateRobotApplicationRequest schema from the OpenAPI specification
type CreateRobotApplicationRequest struct {
	Sources interface{} `json:"sources,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Environment interface{} `json:"environment,omitempty"`
	Name interface{} `json:"name"`
	Robotsoftwaresuite interface{} `json:"robotSoftwareSuite"`
}

// DescribeDeploymentJobRequest represents the DescribeDeploymentJobRequest schema from the OpenAPI specification
type DescribeDeploymentJobRequest struct {
	Job interface{} `json:"job"`
}

// RestartSimulationJobRequest represents the RestartSimulationJobRequest schema from the OpenAPI specification
type RestartSimulationJobRequest struct {
	Job interface{} `json:"job"`
}

// DeregisterRobotResponse represents the DeregisterRobotResponse schema from the OpenAPI specification
type DeregisterRobotResponse struct {
	Fleet interface{} `json:"fleet,omitempty"`
	Robot interface{} `json:"robot,omitempty"`
}

// FailureSummary represents the FailureSummary schema from the OpenAPI specification
type FailureSummary struct {
	Failures interface{} `json:"failures,omitempty"`
	Totalfailurecount interface{} `json:"totalFailureCount,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// DescribeSimulationJobBatchRequest represents the DescribeSimulationJobBatchRequest schema from the OpenAPI specification
type DescribeSimulationJobBatchRequest struct {
	Batch interface{} `json:"batch"`
}

// CreateSimulationApplicationVersionRequest represents the CreateSimulationApplicationVersionRequest schema from the OpenAPI specification
type CreateSimulationApplicationVersionRequest struct {
	Currentrevisionid interface{} `json:"currentRevisionId,omitempty"`
	Imagedigest interface{} `json:"imageDigest,omitempty"`
	S3etags interface{} `json:"s3Etags,omitempty"`
	Application interface{} `json:"application"`
}

// WorldFailure represents the WorldFailure schema from the OpenAPI specification
type WorldFailure struct {
	Failurecode interface{} `json:"failureCode,omitempty"`
	Failurecount interface{} `json:"failureCount,omitempty"`
	Samplefailurereason interface{} `json:"sampleFailureReason,omitempty"`
}

// Tool represents the Tool schema from the OpenAPI specification
type Tool struct {
	Command interface{} `json:"command"`
	Exitbehavior interface{} `json:"exitBehavior,omitempty"`
	Name interface{} `json:"name"`
	Streamoutputtocloudwatch interface{} `json:"streamOutputToCloudWatch,omitempty"`
	Streamui interface{} `json:"streamUI,omitempty"`
}

// RobotDeployment represents the RobotDeployment schema from the OpenAPI specification
type RobotDeployment struct {
	Failurereason interface{} `json:"failureReason,omitempty"`
	Progressdetail interface{} `json:"progressDetail,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Deploymentfinishtime interface{} `json:"deploymentFinishTime,omitempty"`
	Deploymentstarttime interface{} `json:"deploymentStartTime,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
}

// DescribeRobotResponse represents the DescribeRobotResponse schema from the OpenAPI specification
type DescribeRobotResponse struct {
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Greengrassgroupid interface{} `json:"greengrassGroupId,omitempty"`
	Fleetarn interface{} `json:"fleetArn,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Lastdeploymentjob interface{} `json:"lastDeploymentJob,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Architecture interface{} `json:"architecture,omitempty"`
	Lastdeploymenttime interface{} `json:"lastDeploymentTime,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// StartSimulationJobBatchResponse represents the StartSimulationJobBatchResponse schema from the OpenAPI specification
type StartSimulationJobBatchResponse struct {
	Batchpolicy interface{} `json:"batchPolicy,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Createdrequests interface{} `json:"createdRequests,omitempty"`
	Failedrequests interface{} `json:"failedRequests,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Failurereason interface{} `json:"failureReason,omitempty"`
	Pendingrequests interface{} `json:"pendingRequests,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
}

// SourceConfig represents the SourceConfig schema from the OpenAPI specification
type SourceConfig struct {
	Architecture interface{} `json:"architecture,omitempty"`
	S3bucket interface{} `json:"s3Bucket,omitempty"`
	S3key interface{} `json:"s3Key,omitempty"`
}

// ListFleetsRequest represents the ListFleetsRequest schema from the OpenAPI specification
type ListFleetsRequest struct {
	Filters interface{} `json:"filters,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// DeleteFleetResponse represents the DeleteFleetResponse schema from the OpenAPI specification
type DeleteFleetResponse struct {
}

// ListWorldGenerationJobsRequest represents the ListWorldGenerationJobsRequest schema from the OpenAPI specification
type ListWorldGenerationJobsRequest struct {
	Filters interface{} `json:"filters,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// CancelWorldExportJobRequest represents the CancelWorldExportJobRequest schema from the OpenAPI specification
type CancelWorldExportJobRequest struct {
	Job interface{} `json:"job"`
}

// DeleteWorldTemplateResponse represents the DeleteWorldTemplateResponse schema from the OpenAPI specification
type DeleteWorldTemplateResponse struct {
}

// CreateSimulationJobRequest represents the CreateSimulationJobRequest schema from the OpenAPI specification
type CreateSimulationJobRequest struct {
	Vpcconfig interface{} `json:"vpcConfig,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Failurebehavior interface{} `json:"failureBehavior,omitempty"`
	Robotapplications interface{} `json:"robotApplications,omitempty"`
	Compute interface{} `json:"compute,omitempty"`
	Iamrole interface{} `json:"iamRole"`
	Maxjobdurationinseconds interface{} `json:"maxJobDurationInSeconds"`
	Datasources interface{} `json:"dataSources,omitempty"`
	Loggingconfig interface{} `json:"loggingConfig,omitempty"`
	Outputlocation interface{} `json:"outputLocation,omitempty"`
	Simulationapplications interface{} `json:"simulationApplications,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// DescribeFleetResponse represents the DescribeFleetResponse schema from the OpenAPI specification
type DescribeFleetResponse struct {
	Lastdeploymentjob interface{} `json:"lastDeploymentJob,omitempty"`
	Lastdeploymentstatus interface{} `json:"lastDeploymentStatus,omitempty"`
	Lastdeploymenttime interface{} `json:"lastDeploymentTime,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Robots interface{} `json:"robots,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
}

// DescribeSimulationApplicationRequest represents the DescribeSimulationApplicationRequest schema from the OpenAPI specification
type DescribeSimulationApplicationRequest struct {
	Applicationversion interface{} `json:"applicationVersion,omitempty"`
	Application interface{} `json:"application"`
}

// SimulationJobSummary represents the SimulationJobSummary schema from the OpenAPI specification
type SimulationJobSummary struct {
	Arn interface{} `json:"arn,omitempty"`
	Computetype interface{} `json:"computeType,omitempty"`
	Datasourcenames interface{} `json:"dataSourceNames,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Robotapplicationnames interface{} `json:"robotApplicationNames,omitempty"`
	Simulationapplicationnames interface{} `json:"simulationApplicationNames,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// Compute represents the Compute schema from the OpenAPI specification
type Compute struct {
	Simulationunitlimit interface{} `json:"simulationUnitLimit,omitempty"`
	Computetype interface{} `json:"computeType,omitempty"`
	Gpuunitlimit interface{} `json:"gpuUnitLimit,omitempty"`
}

// DeleteSimulationApplicationRequest represents the DeleteSimulationApplicationRequest schema from the OpenAPI specification
type DeleteSimulationApplicationRequest struct {
	Application interface{} `json:"application"`
	Applicationversion interface{} `json:"applicationVersion,omitempty"`
}

// UpdateWorldTemplateRequest represents the UpdateWorldTemplateRequest schema from the OpenAPI specification
type UpdateWorldTemplateRequest struct {
	Template interface{} `json:"template"`
	Templatebody interface{} `json:"templateBody,omitempty"`
	Templatelocation interface{} `json:"templateLocation,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// SimulationApplicationConfig represents the SimulationApplicationConfig schema from the OpenAPI specification
type SimulationApplicationConfig struct {
	Launchconfig interface{} `json:"launchConfig"`
	Tools interface{} `json:"tools,omitempty"`
	Uploadconfigurations interface{} `json:"uploadConfigurations,omitempty"`
	Usedefaulttools interface{} `json:"useDefaultTools,omitempty"`
	Usedefaultuploadconfigurations interface{} `json:"useDefaultUploadConfigurations,omitempty"`
	Worldconfigs interface{} `json:"worldConfigs,omitempty"`
	Application interface{} `json:"application"`
	Applicationversion interface{} `json:"applicationVersion,omitempty"`
}

// VPCConfig represents the VPCConfig schema from the OpenAPI specification
type VPCConfig struct {
	Securitygroups interface{} `json:"securityGroups,omitempty"`
	Subnets interface{} `json:"subnets"`
	Assignpublicip interface{} `json:"assignPublicIp,omitempty"`
}

// DeleteRobotResponse represents the DeleteRobotResponse schema from the OpenAPI specification
type DeleteRobotResponse struct {
}

// RestartSimulationJobResponse represents the RestartSimulationJobResponse schema from the OpenAPI specification
type RestartSimulationJobResponse struct {
}

// GetWorldTemplateBodyRequest represents the GetWorldTemplateBodyRequest schema from the OpenAPI specification
type GetWorldTemplateBodyRequest struct {
	Template interface{} `json:"template,omitempty"`
	Generationjob interface{} `json:"generationJob,omitempty"`
}

// UpdateSimulationApplicationRequest represents the UpdateSimulationApplicationRequest schema from the OpenAPI specification
type UpdateSimulationApplicationRequest struct {
	Currentrevisionid interface{} `json:"currentRevisionId,omitempty"`
	Environment interface{} `json:"environment,omitempty"`
	Renderingengine interface{} `json:"renderingEngine,omitempty"`
	Robotsoftwaresuite interface{} `json:"robotSoftwareSuite"`
	Simulationsoftwaresuite interface{} `json:"simulationSoftwareSuite"`
	Sources interface{} `json:"sources,omitempty"`
	Application interface{} `json:"application"`
}

// CancelSimulationJobBatchRequest represents the CancelSimulationJobBatchRequest schema from the OpenAPI specification
type CancelSimulationJobBatchRequest struct {
	Batch interface{} `json:"batch"`
}

// SimulationApplicationSummary represents the SimulationApplicationSummary schema from the OpenAPI specification
type SimulationApplicationSummary struct {
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Robotsoftwaresuite interface{} `json:"robotSoftwareSuite,omitempty"`
	Simulationsoftwaresuite interface{} `json:"simulationSoftwareSuite,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
}

// ListDeploymentJobsResponse represents the ListDeploymentJobsResponse schema from the OpenAPI specification
type ListDeploymentJobsResponse struct {
	Deploymentjobs interface{} `json:"deploymentJobs,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// DescribeSimulationJobRequest represents the DescribeSimulationJobRequest schema from the OpenAPI specification
type DescribeSimulationJobRequest struct {
	Job interface{} `json:"job"`
}

// CreateWorldTemplateResponse represents the CreateWorldTemplateResponse schema from the OpenAPI specification
type CreateWorldTemplateResponse struct {
	Createdat interface{} `json:"createdAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
}

// DescribeRobotRequest represents the DescribeRobotRequest schema from the OpenAPI specification
type DescribeRobotRequest struct {
	Robot interface{} `json:"robot"`
}

// BatchDescribeSimulationJobRequest represents the BatchDescribeSimulationJobRequest schema from the OpenAPI specification
type BatchDescribeSimulationJobRequest struct {
	Jobs interface{} `json:"jobs"`
}

// ListDeploymentJobsRequest represents the ListDeploymentJobsRequest schema from the OpenAPI specification
type ListDeploymentJobsRequest struct {
	Filters interface{} `json:"filters,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// DeleteWorldTemplateRequest represents the DeleteWorldTemplateRequest schema from the OpenAPI specification
type DeleteWorldTemplateRequest struct {
	Template interface{} `json:"template"`
}

// CreateSimulationApplicationRequest represents the CreateSimulationApplicationRequest schema from the OpenAPI specification
type CreateSimulationApplicationRequest struct {
	Environment interface{} `json:"environment,omitempty"`
	Name interface{} `json:"name"`
	Renderingengine interface{} `json:"renderingEngine,omitempty"`
	Robotsoftwaresuite interface{} `json:"robotSoftwareSuite"`
	Simulationsoftwaresuite interface{} `json:"simulationSoftwareSuite"`
	Sources interface{} `json:"sources,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// BatchDeleteWorldsRequest represents the BatchDeleteWorldsRequest schema from the OpenAPI specification
type BatchDeleteWorldsRequest struct {
	Worlds interface{} `json:"worlds"`
}

// FinishedWorldsSummary represents the FinishedWorldsSummary schema from the OpenAPI specification
type FinishedWorldsSummary struct {
	Failuresummary interface{} `json:"failureSummary,omitempty"`
	Finishedcount interface{} `json:"finishedCount,omitempty"`
	Succeededworlds interface{} `json:"succeededWorlds,omitempty"`
}

// DescribeWorldGenerationJobRequest represents the DescribeWorldGenerationJobRequest schema from the OpenAPI specification
type DescribeWorldGenerationJobRequest struct {
	Job interface{} `json:"job"`
}

// CreateDeploymentJobRequest represents the CreateDeploymentJobRequest schema from the OpenAPI specification
type CreateDeploymentJobRequest struct {
	Fleet interface{} `json:"fleet"`
	Tags interface{} `json:"tags,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken"`
	Deploymentapplicationconfigs interface{} `json:"deploymentApplicationConfigs"`
	Deploymentconfig interface{} `json:"deploymentConfig,omitempty"`
}

// DescribeWorldGenerationJobResponse represents the DescribeWorldGenerationJobResponse schema from the OpenAPI specification
type DescribeWorldGenerationJobResponse struct {
	Failurereason interface{} `json:"failureReason,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Worldtags interface{} `json:"worldTags,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
	Template interface{} `json:"template,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Finishedworldssummary interface{} `json:"finishedWorldsSummary,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Worldcount interface{} `json:"worldCount,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// ListRobotApplicationsResponse represents the ListRobotApplicationsResponse schema from the OpenAPI specification
type ListRobotApplicationsResponse struct {
	Robotapplicationsummaries interface{} `json:"robotApplicationSummaries,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// CreateFleetResponse represents the CreateFleetResponse schema from the OpenAPI specification
type CreateFleetResponse struct {
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// PortMapping represents the PortMapping schema from the OpenAPI specification
type PortMapping struct {
	Jobport interface{} `json:"jobPort"`
	Applicationport interface{} `json:"applicationPort"`
	Enableonpublicip interface{} `json:"enableOnPublicIp,omitempty"`
}

// ListRobotsResponse represents the ListRobotsResponse schema from the OpenAPI specification
type ListRobotsResponse struct {
	Robots interface{} `json:"robots,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListFleetsResponse represents the ListFleetsResponse schema from the OpenAPI specification
type ListFleetsResponse struct {
	Fleetdetails interface{} `json:"fleetDetails,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// DeploymentLaunchConfig represents the DeploymentLaunchConfig schema from the OpenAPI specification
type DeploymentLaunchConfig struct {
	Environmentvariables interface{} `json:"environmentVariables,omitempty"`
	Launchfile interface{} `json:"launchFile"`
	Packagename interface{} `json:"packageName"`
	Postlaunchfile interface{} `json:"postLaunchFile,omitempty"`
	Prelaunchfile interface{} `json:"preLaunchFile,omitempty"`
}

// DescribeWorldResponse represents the DescribeWorldResponse schema from the OpenAPI specification
type DescribeWorldResponse struct {
	Tags interface{} `json:"tags,omitempty"`
	Template interface{} `json:"template,omitempty"`
	Worlddescriptionbody interface{} `json:"worldDescriptionBody,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Generationjob interface{} `json:"generationJob,omitempty"`
}

// UpdateRobotApplicationRequest represents the UpdateRobotApplicationRequest schema from the OpenAPI specification
type UpdateRobotApplicationRequest struct {
	Sources interface{} `json:"sources,omitempty"`
	Application interface{} `json:"application"`
	Currentrevisionid interface{} `json:"currentRevisionId,omitempty"`
	Environment interface{} `json:"environment,omitempty"`
	Robotsoftwaresuite interface{} `json:"robotSoftwareSuite"`
}

// ListWorldTemplatesResponse represents the ListWorldTemplatesResponse schema from the OpenAPI specification
type ListWorldTemplatesResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Templatesummaries interface{} `json:"templateSummaries,omitempty"`
}

// TemplateLocation represents the TemplateLocation schema from the OpenAPI specification
type TemplateLocation struct {
	S3bucket interface{} `json:"s3Bucket"`
	S3key interface{} `json:"s3Key"`
}

// CreateRobotRequest represents the CreateRobotRequest schema from the OpenAPI specification
type CreateRobotRequest struct {
	Architecture interface{} `json:"architecture"`
	Greengrassgroupid interface{} `json:"greengrassGroupId"`
	Name interface{} `json:"name"`
	Tags interface{} `json:"tags,omitempty"`
}

// LaunchConfig represents the LaunchConfig schema from the OpenAPI specification
type LaunchConfig struct {
	Packagename interface{} `json:"packageName,omitempty"`
	Portforwardingconfig interface{} `json:"portForwardingConfig,omitempty"`
	Streamui interface{} `json:"streamUI,omitempty"`
	Command interface{} `json:"command,omitempty"`
	Environmentvariables interface{} `json:"environmentVariables,omitempty"`
	Launchfile interface{} `json:"launchFile,omitempty"`
}

// CreateWorldGenerationJobRequest represents the CreateWorldGenerationJobRequest schema from the OpenAPI specification
type CreateWorldGenerationJobRequest struct {
	Template interface{} `json:"template"`
	Worldcount interface{} `json:"worldCount"`
	Worldtags interface{} `json:"worldTags,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// RegisterRobotRequest represents the RegisterRobotRequest schema from the OpenAPI specification
type RegisterRobotRequest struct {
	Robot interface{} `json:"robot"`
	Fleet interface{} `json:"fleet"`
}

// DescribeWorldRequest represents the DescribeWorldRequest schema from the OpenAPI specification
type DescribeWorldRequest struct {
	World interface{} `json:"world"`
}

// ListRobotApplicationsRequest represents the ListRobotApplicationsRequest schema from the OpenAPI specification
type ListRobotApplicationsRequest struct {
	Filters interface{} `json:"filters,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Versionqualifier interface{} `json:"versionQualifier,omitempty"`
}

// CreateSimulationApplicationVersionResponse represents the CreateSimulationApplicationVersionResponse schema from the OpenAPI specification
type CreateSimulationApplicationVersionResponse struct {
	Simulationsoftwaresuite interface{} `json:"simulationSoftwareSuite,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Environment interface{} `json:"environment,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Revisionid interface{} `json:"revisionId,omitempty"`
	Sources interface{} `json:"sources,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Renderingengine interface{} `json:"renderingEngine,omitempty"`
	Robotsoftwaresuite interface{} `json:"robotSoftwareSuite,omitempty"`
}

// WorldCount represents the WorldCount schema from the OpenAPI specification
type WorldCount struct {
	Floorplancount interface{} `json:"floorplanCount,omitempty"`
	Interiorcountperfloorplan interface{} `json:"interiorCountPerFloorplan,omitempty"`
}

// CreateSimulationJobResponse represents the CreateSimulationJobResponse schema from the OpenAPI specification
type CreateSimulationJobResponse struct {
	Laststartedat interface{} `json:"lastStartedAt,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Datasources interface{} `json:"dataSources,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
	Loggingconfig interface{} `json:"loggingConfig,omitempty"`
	Iamrole interface{} `json:"iamRole,omitempty"`
	Outputlocation interface{} `json:"outputLocation,omitempty"`
	Simulationapplications interface{} `json:"simulationApplications,omitempty"`
	Maxjobdurationinseconds interface{} `json:"maxJobDurationInSeconds,omitempty"`
	Vpcconfig interface{} `json:"vpcConfig,omitempty"`
	Simulationtimemillis interface{} `json:"simulationTimeMillis,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Compute interface{} `json:"compute,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Failurebehavior interface{} `json:"failureBehavior,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Robotapplications interface{} `json:"robotApplications,omitempty"`
}

// WorldExportJobSummary represents the WorldExportJobSummary schema from the OpenAPI specification
type WorldExportJobSummary struct {
	Createdat interface{} `json:"createdAt,omitempty"`
	Outputlocation OutputLocation `json:"outputLocation,omitempty"` // The output location.
	Status interface{} `json:"status,omitempty"`
	Worlds interface{} `json:"worlds,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
}

// DataSourceConfig represents the DataSourceConfig schema from the OpenAPI specification
type DataSourceConfig struct {
	Name interface{} `json:"name"`
	S3bucket interface{} `json:"s3Bucket"`
	S3keys interface{} `json:"s3Keys"`
	TypeField interface{} `json:"type,omitempty"`
	Destination interface{} `json:"destination,omitempty"`
}

// SimulationJobRequest represents the SimulationJobRequest schema from the OpenAPI specification
type SimulationJobRequest struct {
	Robotapplications interface{} `json:"robotApplications,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Failurebehavior interface{} `json:"failureBehavior,omitempty"`
	Maxjobdurationinseconds interface{} `json:"maxJobDurationInSeconds"`
	Simulationapplications interface{} `json:"simulationApplications,omitempty"`
	Usedefaultapplications interface{} `json:"useDefaultApplications,omitempty"`
	Vpcconfig VPCConfig `json:"vpcConfig,omitempty"` // If your simulation job accesses resources in a VPC, you provide this parameter identifying the list of security group IDs and subnet IDs. These must belong to the same VPC. You must provide at least one security group and two subnet IDs.
	Datasources interface{} `json:"dataSources,omitempty"`
	Iamrole interface{} `json:"iamRole,omitempty"`
	Loggingconfig LoggingConfig `json:"loggingConfig,omitempty"` // The logging configuration.
	Outputlocation OutputLocation `json:"outputLocation,omitempty"` // The output location.
	Compute interface{} `json:"compute,omitempty"`
}

// UpdateSimulationApplicationResponse represents the UpdateSimulationApplicationResponse schema from the OpenAPI specification
type UpdateSimulationApplicationResponse struct {
	Environment interface{} `json:"environment,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Revisionid interface{} `json:"revisionId,omitempty"`
	Simulationsoftwaresuite interface{} `json:"simulationSoftwareSuite,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Renderingengine interface{} `json:"renderingEngine,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Robotsoftwaresuite interface{} `json:"robotSoftwareSuite,omitempty"`
	Sources interface{} `json:"sources,omitempty"`
}

// DeploymentConfig represents the DeploymentConfig schema from the OpenAPI specification
type DeploymentConfig struct {
	Concurrentdeploymentpercentage interface{} `json:"concurrentDeploymentPercentage,omitempty"`
	Downloadconditionfile interface{} `json:"downloadConditionFile,omitempty"`
	Failurethresholdpercentage interface{} `json:"failureThresholdPercentage,omitempty"`
	Robotdeploymenttimeoutinseconds interface{} `json:"robotDeploymentTimeoutInSeconds,omitempty"`
}

// CreateWorldTemplateRequest represents the CreateWorldTemplateRequest schema from the OpenAPI specification
type CreateWorldTemplateRequest struct {
	Templatebody interface{} `json:"templateBody,omitempty"`
	Templatelocation interface{} `json:"templateLocation,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// GetWorldTemplateBodyResponse represents the GetWorldTemplateBodyResponse schema from the OpenAPI specification
type GetWorldTemplateBodyResponse struct {
	Templatebody interface{} `json:"templateBody,omitempty"`
}

// ListSimulationJobsRequest represents the ListSimulationJobsRequest schema from the OpenAPI specification
type ListSimulationJobsRequest struct {
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Filters interface{} `json:"filters,omitempty"`
}

// DescribeRobotApplicationRequest represents the DescribeRobotApplicationRequest schema from the OpenAPI specification
type DescribeRobotApplicationRequest struct {
	Applicationversion interface{} `json:"applicationVersion,omitempty"`
	Application interface{} `json:"application"`
}

// CreateWorldExportJobResponse represents the CreateWorldExportJobResponse schema from the OpenAPI specification
type CreateWorldExportJobResponse struct {
	Failurecode interface{} `json:"failureCode,omitempty"`
	Iamrole interface{} `json:"iamRole,omitempty"`
	Outputlocation OutputLocation `json:"outputLocation,omitempty"` // The output location.
	Status interface{} `json:"status,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
}

// CancelDeploymentJobResponse represents the CancelDeploymentJobResponse schema from the OpenAPI specification
type CancelDeploymentJobResponse struct {
}

// StartSimulationJobBatchRequest represents the StartSimulationJobBatchRequest schema from the OpenAPI specification
type StartSimulationJobBatchRequest struct {
	Createsimulationjobrequests interface{} `json:"createSimulationJobRequests"`
	Tags interface{} `json:"tags,omitempty"`
	Batchpolicy interface{} `json:"batchPolicy,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
}

// CreateRobotResponse represents the CreateRobotResponse schema from the OpenAPI specification
type CreateRobotResponse struct {
	Architecture interface{} `json:"architecture,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Greengrassgroupid interface{} `json:"greengrassGroupId,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// SimulationJobBatchSummary represents the SimulationJobBatchSummary schema from the OpenAPI specification
type SimulationJobBatchSummary struct {
	Pendingrequestcount interface{} `json:"pendingRequestCount,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Createdrequestcount interface{} `json:"createdRequestCount,omitempty"`
	Failedrequestcount interface{} `json:"failedRequestCount,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
}

// ListWorldTemplatesRequest represents the ListWorldTemplatesRequest schema from the OpenAPI specification
type ListWorldTemplatesRequest struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"tags,omitempty"`
}

// DescribeRobotApplicationResponse represents the DescribeRobotApplicationResponse schema from the OpenAPI specification
type DescribeRobotApplicationResponse struct {
	Tags interface{} `json:"tags,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Imagedigest interface{} `json:"imageDigest,omitempty"`
	Environment interface{} `json:"environment,omitempty"`
	Revisionid interface{} `json:"revisionId,omitempty"`
	Robotsoftwaresuite interface{} `json:"robotSoftwareSuite,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Sources interface{} `json:"sources,omitempty"`
}

// DeleteRobotApplicationRequest represents the DeleteRobotApplicationRequest schema from the OpenAPI specification
type DeleteRobotApplicationRequest struct {
	Application interface{} `json:"application"`
	Applicationversion interface{} `json:"applicationVersion,omitempty"`
}

// DeleteFleetRequest represents the DeleteFleetRequest schema from the OpenAPI specification
type DeleteFleetRequest struct {
	Fleet interface{} `json:"fleet"`
}

// SyncDeploymentJobResponse represents the SyncDeploymentJobResponse schema from the OpenAPI specification
type SyncDeploymentJobResponse struct {
	Failurereason interface{} `json:"failureReason,omitempty"`
	Fleet interface{} `json:"fleet,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Deploymentapplicationconfigs interface{} `json:"deploymentApplicationConfigs,omitempty"`
	Deploymentconfig interface{} `json:"deploymentConfig,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
}

// ProgressDetail represents the ProgressDetail schema from the OpenAPI specification
type ProgressDetail struct {
	Percentdone interface{} `json:"percentDone,omitempty"`
	Targetresource interface{} `json:"targetResource,omitempty"`
	Currentprogress interface{} `json:"currentProgress,omitempty"`
	Estimatedtimeremainingseconds interface{} `json:"estimatedTimeRemainingSeconds,omitempty"`
}

// EnvironmentVariableMap represents the EnvironmentVariableMap schema from the OpenAPI specification
type EnvironmentVariableMap struct {
}
