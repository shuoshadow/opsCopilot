package server

import (
	"github.com/gin-gonic/gin"
	_ "github.com/shaowenchen/ops/swagger"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(r *gin.Engine) {
	v1Hosts := r.Group("/api/v1/namespaces/:namespace/hosts").Use(AuthMiddleware())
	{
		v1Hosts.GET("", ListHosts)
	}
	v1Clusters := r.Group("/api/v1/namespaces/:namespace/clusters").Use(AuthMiddleware())
	{
		v1Clusters.GET("", ListClusters)
		v1Clusters.GET(":cluster", GetCluster)
		v1Clusters.GET(":cluster/nodes", GetClusterNodes)
	}
	v1Tasks := r.Group("/api/v1/namespaces/:namespace/tasks").Use(AuthMiddleware())
	{
		v1Tasks.GET("", ListTasks)
		v1Tasks.POST("", CreateTask)
		v1Tasks.GET("/:task", GetTask)
		v1Tasks.PUT("/:task", PutTask)
		v1Tasks.DELETE("/:task", DeleteTask)
	}
	v1Taskruns := r.Group("/api/v1/namespaces/:namespace/taskruns").Use(AuthMiddleware())
	{
		v1Taskruns.GET("", ListTaskRun)
		v1Taskruns.POST("", CreateTaskRun)
		v1Taskruns.POST("/sync", CreateTaskRunSync)
		v1Taskruns.GET("/:taskrun", GetTaskRun)
	}
	v1Pipelines := r.Group("/api/v1/namespaces/:namespace/pipelines").Use(AuthMiddleware())
	{
		v1Pipelines.GET("", ListPipelines)
		v1Pipelines.POST("", CreatePipeline)
		v1Pipelines.GET("/:pipeline", GetPipeline)
		v1Pipelines.PUT("/:pipeline", PutPipeline)
		v1Pipelines.DELETE("/:pipeline", DeletePipeline)
		v1Pipelines.GET("tools", ListPipelineTools)
	}
	v1Pipelineruns := r.Group("/api/v1/namespaces/:namespace/pipelineruns").Use(AuthMiddleware())
	{
		v1Pipelineruns.GET("", ListPipelineRuns)
		v1Pipelineruns.POST("", CreatePipelineRun)
		v1Pipelineruns.POST("/sync", CreatePipelineRunSync)
		v1Pipelineruns.GET("/:pipelinerun", GetPipelineRun)
	}
	v1Copilot := r.Group("/api/v1/copilot").Use(AuthMiddleware())
	{
		v1Copilot.POST("", PostCopilot)
		v1Copilot.POST("/plain", PostCopilotPlain)
	}
	v1Login := r.Group("/api/v1/login").Use(AuthMiddleware())
	{
		v1Login.GET("/check", LoginCheck)
	}
	v1Summary := r.Group("/api/v1/summary").Use(AuthMiddleware())
	{
		v1Summary.GET("", GetSummary)
	}
	v1Events := r.Group("/api/v1/events")
	{
		v1Events.GET("", ListEvents)
		v1Events.GET("/:event", GetEvents)
	}
}

func SetupRouteWithoutAuth(r *gin.Engine) {
	v1Events := r.Group("/api/v1/namespaces/:namespace/events")
	{
		v1Events.POST("/:event", CreateEvent)
	}
	r.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func SetHealthzRouter(r *gin.Engine) {
	root := r.Group("/")
	{
		root.GET("healthz", Healthz)
		root.GET("readyz", Healthz)
	}
}
