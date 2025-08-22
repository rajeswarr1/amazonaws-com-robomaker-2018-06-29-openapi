package main

import (
	"github.com/aws-robomaker/mcp-server/config"
	"github.com/aws-robomaker/mcp-server/models"
	tools_createworldexportjob "github.com/aws-robomaker/mcp-server/tools/createworldexportjob"
	tools_createworldtemplate "github.com/aws-robomaker/mcp-server/tools/createworldtemplate"
	tools_describedeploymentjob "github.com/aws-robomaker/mcp-server/tools/describedeploymentjob"
	tools_startsimulationjobbatch "github.com/aws-robomaker/mcp-server/tools/startsimulationjobbatch"
	tools_updaterobotapplication "github.com/aws-robomaker/mcp-server/tools/updaterobotapplication"
	tools_listrobotapplications "github.com/aws-robomaker/mcp-server/tools/listrobotapplications"
	tools_describeworldgenerationjob "github.com/aws-robomaker/mcp-server/tools/describeworldgenerationjob"
	tools_deleterobot "github.com/aws-robomaker/mcp-server/tools/deleterobot"
	tools_getworldtemplatebody "github.com/aws-robomaker/mcp-server/tools/getworldtemplatebody"
	tools_createsimulationapplication "github.com/aws-robomaker/mcp-server/tools/createsimulationapplication"
	tools_deletesimulationapplication "github.com/aws-robomaker/mcp-server/tools/deletesimulationapplication"
	tools_tags "github.com/aws-robomaker/mcp-server/tools/tags"
	tools_batchdeleteworlds "github.com/aws-robomaker/mcp-server/tools/batchdeleteworlds"
	tools_describerobot "github.com/aws-robomaker/mcp-server/tools/describerobot"
	tools_createsimulationjob "github.com/aws-robomaker/mcp-server/tools/createsimulationjob"
	tools_updatesimulationapplication "github.com/aws-robomaker/mcp-server/tools/updatesimulationapplication"
	tools_updateworldtemplate "github.com/aws-robomaker/mcp-server/tools/updateworldtemplate"
	tools_createworldgenerationjob "github.com/aws-robomaker/mcp-server/tools/createworldgenerationjob"
	tools_listsimulationjobs "github.com/aws-robomaker/mcp-server/tools/listsimulationjobs"
	tools_cancelworldexportjob "github.com/aws-robomaker/mcp-server/tools/cancelworldexportjob"
	tools_createdeploymentjob "github.com/aws-robomaker/mcp-server/tools/createdeploymentjob"
	tools_createsimulationapplicationversion "github.com/aws-robomaker/mcp-server/tools/createsimulationapplicationversion"
	tools_deletefleet "github.com/aws-robomaker/mcp-server/tools/deletefleet"
	tools_listfleets "github.com/aws-robomaker/mcp-server/tools/listfleets"
	tools_syncdeploymentjob "github.com/aws-robomaker/mcp-server/tools/syncdeploymentjob"
	tools_deleterobotapplication "github.com/aws-robomaker/mcp-server/tools/deleterobotapplication"
	tools_describeworld "github.com/aws-robomaker/mcp-server/tools/describeworld"
	tools_cancelsimulationjobbatch "github.com/aws-robomaker/mcp-server/tools/cancelsimulationjobbatch"
	tools_describesimulationapplication "github.com/aws-robomaker/mcp-server/tools/describesimulationapplication"
	tools_listworlds "github.com/aws-robomaker/mcp-server/tools/listworlds"
	tools_cancelworldgenerationjob "github.com/aws-robomaker/mcp-server/tools/cancelworldgenerationjob"
	tools_createrobot "github.com/aws-robomaker/mcp-server/tools/createrobot"
	tools_describerobotapplication "github.com/aws-robomaker/mcp-server/tools/describerobotapplication"
	tools_listworldexportjobs "github.com/aws-robomaker/mcp-server/tools/listworldexportjobs"
	tools_listsimulationapplications "github.com/aws-robomaker/mcp-server/tools/listsimulationapplications"
	tools_deleteworldtemplate "github.com/aws-robomaker/mcp-server/tools/deleteworldtemplate"
	tools_describesimulationjobbatch "github.com/aws-robomaker/mcp-server/tools/describesimulationjobbatch"
	tools_listrobots "github.com/aws-robomaker/mcp-server/tools/listrobots"
	tools_describeworldexportjob "github.com/aws-robomaker/mcp-server/tools/describeworldexportjob"
	tools_describesimulationjob "github.com/aws-robomaker/mcp-server/tools/describesimulationjob"
	tools_cancelsimulationjob "github.com/aws-robomaker/mcp-server/tools/cancelsimulationjob"
	tools_batchdescribesimulationjob "github.com/aws-robomaker/mcp-server/tools/batchdescribesimulationjob"
	tools_describefleet "github.com/aws-robomaker/mcp-server/tools/describefleet"
	tools_registerrobot "github.com/aws-robomaker/mcp-server/tools/registerrobot"
	tools_restartsimulationjob "github.com/aws-robomaker/mcp-server/tools/restartsimulationjob"
	tools_deregisterrobot "github.com/aws-robomaker/mcp-server/tools/deregisterrobot"
	tools_listworldgenerationjobs "github.com/aws-robomaker/mcp-server/tools/listworldgenerationjobs"
	tools_listworldtemplates "github.com/aws-robomaker/mcp-server/tools/listworldtemplates"
	tools_createrobotapplication "github.com/aws-robomaker/mcp-server/tools/createrobotapplication"
	tools_createrobotapplicationversion "github.com/aws-robomaker/mcp-server/tools/createrobotapplicationversion"
	tools_createfleet "github.com/aws-robomaker/mcp-server/tools/createfleet"
	tools_listsimulationjobbatches "github.com/aws-robomaker/mcp-server/tools/listsimulationjobbatches"
	tools_canceldeploymentjob "github.com/aws-robomaker/mcp-server/tools/canceldeploymentjob"
	tools_describeworldtemplate "github.com/aws-robomaker/mcp-server/tools/describeworldtemplate"
	tools_listdeploymentjobs "github.com/aws-robomaker/mcp-server/tools/listdeploymentjobs"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_createworldexportjob.CreateCreateworldexportjobTool(cfg),
		tools_createworldtemplate.CreateCreateworldtemplateTool(cfg),
		tools_describedeploymentjob.CreateDescribedeploymentjobTool(cfg),
		tools_startsimulationjobbatch.CreateStartsimulationjobbatchTool(cfg),
		tools_updaterobotapplication.CreateUpdaterobotapplicationTool(cfg),
		tools_listrobotapplications.CreateListrobotapplicationsTool(cfg),
		tools_describeworldgenerationjob.CreateDescribeworldgenerationjobTool(cfg),
		tools_deleterobot.CreateDeleterobotTool(cfg),
		tools_getworldtemplatebody.CreateGetworldtemplatebodyTool(cfg),
		tools_createsimulationapplication.CreateCreatesimulationapplicationTool(cfg),
		tools_deletesimulationapplication.CreateDeletesimulationapplicationTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_batchdeleteworlds.CreateBatchdeleteworldsTool(cfg),
		tools_describerobot.CreateDescriberobotTool(cfg),
		tools_createsimulationjob.CreateCreatesimulationjobTool(cfg),
		tools_updatesimulationapplication.CreateUpdatesimulationapplicationTool(cfg),
		tools_updateworldtemplate.CreateUpdateworldtemplateTool(cfg),
		tools_createworldgenerationjob.CreateCreateworldgenerationjobTool(cfg),
		tools_listsimulationjobs.CreateListsimulationjobsTool(cfg),
		tools_cancelworldexportjob.CreateCancelworldexportjobTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_createdeploymentjob.CreateCreatedeploymentjobTool(cfg),
		tools_createsimulationapplicationversion.CreateCreatesimulationapplicationversionTool(cfg),
		tools_deletefleet.CreateDeletefleetTool(cfg),
		tools_listfleets.CreateListfleetsTool(cfg),
		tools_syncdeploymentjob.CreateSyncdeploymentjobTool(cfg),
		tools_deleterobotapplication.CreateDeleterobotapplicationTool(cfg),
		tools_describeworld.CreateDescribeworldTool(cfg),
		tools_cancelsimulationjobbatch.CreateCancelsimulationjobbatchTool(cfg),
		tools_describesimulationapplication.CreateDescribesimulationapplicationTool(cfg),
		tools_listworlds.CreateListworldsTool(cfg),
		tools_cancelworldgenerationjob.CreateCancelworldgenerationjobTool(cfg),
		tools_createrobot.CreateCreaterobotTool(cfg),
		tools_describerobotapplication.CreateDescriberobotapplicationTool(cfg),
		tools_listworldexportjobs.CreateListworldexportjobsTool(cfg),
		tools_listsimulationapplications.CreateListsimulationapplicationsTool(cfg),
		tools_deleteworldtemplate.CreateDeleteworldtemplateTool(cfg),
		tools_describesimulationjobbatch.CreateDescribesimulationjobbatchTool(cfg),
		tools_listrobots.CreateListrobotsTool(cfg),
		tools_describeworldexportjob.CreateDescribeworldexportjobTool(cfg),
		tools_describesimulationjob.CreateDescribesimulationjobTool(cfg),
		tools_cancelsimulationjob.CreateCancelsimulationjobTool(cfg),
		tools_batchdescribesimulationjob.CreateBatchdescribesimulationjobTool(cfg),
		tools_describefleet.CreateDescribefleetTool(cfg),
		tools_registerrobot.CreateRegisterrobotTool(cfg),
		tools_restartsimulationjob.CreateRestartsimulationjobTool(cfg),
		tools_deregisterrobot.CreateDeregisterrobotTool(cfg),
		tools_listworldgenerationjobs.CreateListworldgenerationjobsTool(cfg),
		tools_listworldtemplates.CreateListworldtemplatesTool(cfg),
		tools_createrobotapplication.CreateCreaterobotapplicationTool(cfg),
		tools_createrobotapplicationversion.CreateCreaterobotapplicationversionTool(cfg),
		tools_createfleet.CreateCreatefleetTool(cfg),
		tools_listsimulationjobbatches.CreateListsimulationjobbatchesTool(cfg),
		tools_canceldeploymentjob.CreateCanceldeploymentjobTool(cfg),
		tools_describeworldtemplate.CreateDescribeworldtemplateTool(cfg),
		tools_listdeploymentjobs.CreateListdeploymentjobsTool(cfg),
	}
}
