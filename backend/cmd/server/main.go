package main

import (
	"context"
	"log"
	"net/http"
	"time"

	adminTransport "github.com/diablovocado/declutr/modules/admin/transport"
	"github.com/diablovocado/declutr/modules/auth/application"
	"github.com/diablovocado/declutr/modules/auth/repository"
	"github.com/diablovocado/declutr/modules/auth/transport"
	backupApp "github.com/diablovocado/declutr/modules/backup/application"
	backupRepository "github.com/diablovocado/declutr/modules/backup/repository"
	backupTransport "github.com/diablovocado/declutr/modules/backup/transport"
	collaborationApp "github.com/diablovocado/declutr/modules/collaboration/application"
	collaborationRepository "github.com/diablovocado/declutr/modules/collaboration/repository"
	collaborationTransport "github.com/diablovocado/declutr/modules/collaboration/transport"
	contextApp "github.com/diablovocado/declutr/modules/context/application"
	contextRepository "github.com/diablovocado/declutr/modules/context/repository"
	contextTransport "github.com/diablovocado/declutr/modules/context/transport"
	copilotApp "github.com/diablovocado/declutr/modules/copilot/application"
	copilotRepository "github.com/diablovocado/declutr/modules/copilot/repository"
	copilotTransport "github.com/diablovocado/declutr/modules/copilot/transport"
	devApp "github.com/diablovocado/declutr/modules/developer/application"
	devRepository "github.com/diablovocado/declutr/modules/developer/repository"
	devTransport "github.com/diablovocado/declutr/modules/developer/transport"
	embeddingApp "github.com/diablovocado/declutr/modules/embedding/application"
	embeddingRepository "github.com/diablovocado/declutr/modules/embedding/repository"
	embeddingTransport "github.com/diablovocado/declutr/modules/embedding/transport"
	insightsApp "github.com/diablovocado/declutr/modules/insights/application"
	insightsRepository "github.com/diablovocado/declutr/modules/insights/repository"
	insightsTransport "github.com/diablovocado/declutr/modules/insights/transport"
	integrationsApp "github.com/diablovocado/declutr/modules/integrations/application"
	integrationsRepository "github.com/diablovocado/declutr/modules/integrations/repository"
	integrationsTransport "github.com/diablovocado/declutr/modules/integrations/transport"
	memoryApp "github.com/diablovocado/declutr/modules/memory/application"
	memoryRepository "github.com/diablovocado/declutr/modules/memory/repository"
	memoryTransport "github.com/diablovocado/declutr/modules/memory/transport"
	notificationApp "github.com/diablovocado/declutr/modules/notification/application"
	notificationRepository "github.com/diablovocado/declutr/modules/notification/repository"
	notificationTransport "github.com/diablovocado/declutr/modules/notification/transport"
	orgApp "github.com/diablovocado/declutr/modules/organization/application"
	orgRepository "github.com/diablovocado/declutr/modules/organization/repository"
	orgTransport "github.com/diablovocado/declutr/modules/organization/transport"
	personaApp "github.com/diablovocado/declutr/modules/persona/application"
	personaRepository "github.com/diablovocado/declutr/modules/persona/repository"
	personaTransport "github.com/diablovocado/declutr/modules/persona/transport"
	searchApp "github.com/diablovocado/declutr/modules/search/application"
	searchRepository "github.com/diablovocado/declutr/modules/search/repository"
	searchTransport "github.com/diablovocado/declutr/modules/search/transport"
	securityApp "github.com/diablovocado/declutr/modules/security/application"
	securityRepository "github.com/diablovocado/declutr/modules/security/repository"
	securityTransport "github.com/diablovocado/declutr/modules/security/transport"
	syncApp "github.com/diablovocado/declutr/modules/sync/application"
	syncRepository "github.com/diablovocado/declutr/modules/sync/repository"
	syncTransport "github.com/diablovocado/declutr/modules/sync/transport"
	versioningApp "github.com/diablovocado/declutr/modules/versioning/application"
	versioningRepository "github.com/diablovocado/declutr/modules/versioning/repository"
	versioningTransport "github.com/diablovocado/declutr/modules/versioning/transport"
	workflowApp "github.com/diablovocado/declutr/modules/workflow/application"
	workflowRepository "github.com/diablovocado/declutr/modules/workflow/repository"
	workflowTransport "github.com/diablovocado/declutr/modules/workflow/transport"
	"github.com/diablovocado/declutr/pkg/health"
	"github.com/diablovocado/declutr/shared/config"
	"github.com/diablovocado/declutr/shared/database"
	"github.com/diablovocado/declutr/shared/middleware"
	"github.com/diablovocado/declutr/shared/observability"
	"github.com/diablovocado/declutr/shared/ratelimit"
	"github.com/diablovocado/declutr/shared/supervisor"
)

func main() {
	cfg := config.LoadConfig()
	logger := observability.InitLogger("declutr-backend", nil)
	logger.Info(context.Background(), "Starting Declutr Developer Platform Backend Engine", map[string]interface{}{
		"env":  cfg.Env,
		"port": cfg.Port,
	})

	db := database.Connect()

	userRepo := &repository.PostgresUserRepository{
		DB: db,
	}

	authService := &application.Service{
		UserRepo:   userRepo,
		Challenges: application.NewChallengeStore(),
		SRP:        application.NewEngine(),
	}

	mux := http.NewServeMux()

	// Developer Platform Engine Initialization
	devRepo := devRepository.NewInMemoryDeveloperRepository()
	devSvc := devApp.NewDeveloperService(devRepo)
	devAPI := devTransport.NewDeveloperAPI(devSvc)

	mux.HandleFunc("/api/v1/developer/keys", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			devAPI.CreateAPIKey(w, r)
		case http.MethodDelete:
			devAPI.RevokeAPIKey(w, r)
		default:
			devAPI.ListAPIKeys(w, r)
		}
	})
	mux.HandleFunc("/api/v1/developer/webhooks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			devAPI.RegisterWebhook(w, r)
		} else {
			devAPI.ListWebhooks(w, r)
		}
	})
	mux.HandleFunc("/api/v1/developer/webhooks/deliveries", devAPI.GetWebhookDeliveries)
	mux.HandleFunc("/api/v1/developer/oauth/apps", devAPI.CreateOAuthApp)
	mux.HandleFunc("/api/v1/developer/oauth/token", devAPI.ExchangeOAuthToken)
	mux.HandleFunc("/api/v1/developer/openapi", devAPI.ServeOpenAPISpec)

	// Enterprise Organizations & Multi-Tenancy Module Initialization
	orgRepo := orgRepository.NewInMemoryOrganizationRepository()
	orgSvc := orgApp.NewOrganizationService(orgRepo)
	orgAPI := orgTransport.NewOrganizationAPI(orgSvc)

	mux.HandleFunc("/api/v1/organizations", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			orgAPI.CreateOrganization(w, r)
		} else {
			orgAPI.ListOrganizations(w, r)
		}
	})
	mux.HandleFunc("/api/v1/organizations/settings", orgAPI.ManageSettings)
	mux.HandleFunc("/api/v1/organizations/invite", orgAPI.InviteMember)
	mux.HandleFunc("/api/v1/organizations/members", orgAPI.ListMembers)
	mux.HandleFunc("/api/v1/organizations/members/status", orgAPI.UpdateMemberStatus)
	mux.HandleFunc("/api/v1/organizations/ownership/transfer", orgAPI.TransferOwnership)
	mux.HandleFunc("/api/v1/organizations/roles", orgAPI.ManageRoles)
	mux.HandleFunc("/api/v1/organizations/groups", orgAPI.ManageGroups)
	mux.HandleFunc("/api/v1/organizations/workspaces", orgAPI.ManageWorkspaces)
	mux.HandleFunc("/api/v1/organizations/policies", orgAPI.ManagePolicies)
	mux.HandleFunc("/api/v1/organizations/sso/config", orgAPI.ManageSSO)
	mux.HandleFunc("/api/v1/organizations/directory", orgAPI.GetDirectory)

	// Diagnostic Health, Metrics, Readiness, Liveness, and Version Endpoints
	mux.HandleFunc("/health", health.Handler)
	mux.HandleFunc("/ready", health.ReadinessHandler)
	mux.HandleFunc("/live", health.LivenessHandler)
	mux.HandleFunc("/version", health.VersionHandler)
	mux.HandleFunc("/metrics", health.MetricsHandler)

	mux.HandleFunc("/api/v1/health", health.Handler)
	mux.HandleFunc("/api/v1/ready", health.ReadinessHandler)
	mux.HandleFunc("/api/v1/live", health.LivenessHandler)
	mux.HandleFunc("/api/v1/version", health.VersionHandler)
	mux.HandleFunc("/api/v1/metrics", health.MetricsHandler)

	// Admin & Internal System Observability Endpoints
	adminAPI := adminTransport.NewAdminAPI()
	mux.HandleFunc("/api/v1/admin/status", adminAPI.GetSystemStatus)
	mux.HandleFunc("/api/v1/admin/metrics", adminAPI.GetMetrics)
	mux.HandleFunc("/api/v1/admin/queues", adminAPI.GetQueuesAndWorkers)
	mux.HandleFunc("/api/v1/admin/workers", adminAPI.GetQueuesAndWorkers)
	mux.HandleFunc("/api/v1/admin/cache", adminAPI.GetCacheStats)
	mux.HandleFunc("/api/v1/admin/traces", adminAPI.GetTraces)

	// Auth API Endpoints
	mux.HandleFunc(
		"/api/v1/auth/register",
		transport.RegisterHandler(authService),
	)

	mux.HandleFunc(
		"/api/v1/auth/login/start",
		transport.LoginStartHandler(authService),
	)
	mux.Handle(
		"/api/v1/me",
		middleware.Auth(userRepo)(transport.MeHandler()),
	)

	mux.HandleFunc(
		"/api/v1/auth/login/finish",
		transport.LoginFinishHandler(authService),
	)

	// Context & Intent Engine Module initialization
	contextRepo := contextRepository.NewInMemoryContextRepository()
	contextService := contextApp.NewContextService(contextRepo, nil)
	contextAPI := contextTransport.NewAPI(contextService)

	mux.HandleFunc("/api/v1/context", contextAPI.GetContextsHandler)
	mux.HandleFunc("/api/v1/context/details", contextAPI.GetContextDetailHandler)
	mux.HandleFunc("/api/v1/context/refresh", contextAPI.RefreshContextHandler)
	mux.HandleFunc("/api/v1/context/intent", contextAPI.GetIntentHandler)
	mux.HandleFunc("/api/v1/context/stats", contextAPI.GetStatsHandler)

	// Reverse Persona Engine Module initialization
	personaRepo := personaRepository.NewInMemoryPersonaRepository()
	personaSvc := personaApp.NewPersonaService(personaRepo)
	personaEngine := personaApp.NewPersonaEngine(personaSvc)
	_ = personaEngine
	personaAPI := personaTransport.NewPersonaAPI(personaSvc)

	mux.HandleFunc("/api/v1/persona", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			personaAPI.DeletePersona(w, r)
		} else {
			personaAPI.GetPersona(w, r)
		}
	})
	mux.HandleFunc("/api/v1/persona/recommendations", personaAPI.GetRecommendations)
	mux.HandleFunc("/api/v1/persona/settings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			personaAPI.UpdateSettings(w, r)
		} else {
			personaAPI.GetSettings(w, r)
		}
	})
	mux.HandleFunc("/api/v1/persona/reset", personaAPI.ResetPersona)
	mux.HandleFunc("/api/v1/persona/export", personaAPI.ExportPersona)
	mux.HandleFunc("/api/v1/persona/signal", personaAPI.RecordSignal)
	mux.HandleFunc("/api/v1/persona/history", personaAPI.GetHistory)

	// Memory Engine Module initialization
	memoryRepo := memoryRepository.NewInMemoryMemoryRepository()
	memorySvc := memoryApp.NewMemoryService(memoryRepo)
	memoryEngine := memoryApp.NewMemoryEngine(memorySvc)
	_ = memoryEngine
	memoryAPI := memoryTransport.NewMemoryAPI(memorySvc)

	mux.HandleFunc("/api/v1/memory", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			memoryAPI.DeleteMemory(w, r)
		} else {
			memoryAPI.GetMemories(w, r)
		}
	})
	mux.HandleFunc("/api/v1/memory/timeline", memoryAPI.GetTimeline)
	mux.HandleFunc("/api/v1/memory/detail", memoryAPI.GetMemoryDetail)
	mux.HandleFunc("/api/v1/memory/refresh", memoryAPI.RefreshMemory)
	mux.HandleFunc("/api/v1/memory/pin", memoryAPI.PinMemory)
	mux.HandleFunc("/api/v1/memory/archive", memoryAPI.ArchiveMemory)
	mux.HandleFunc("/api/v1/memory/stats", memoryAPI.GetStats)
	mux.HandleFunc("/api/v1/memory/reset", memoryAPI.ResetMemory)
	mux.HandleFunc("/api/v1/memory/settings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			memoryAPI.UpdateSettings(w, r)
		} else {
			memoryAPI.GetSettings(w, r)
		}
	})

	// Embedding Engine Module initialization
	embeddingRepo := embeddingRepository.NewInMemoryVectorRepository()
	embeddingSvc := embeddingApp.NewEmbeddingService(embeddingRepo)
	embeddingEngine := embeddingApp.NewEmbeddingEngine(embeddingSvc)
	_ = embeddingEngine
	embeddingAPI := embeddingTransport.NewEmbeddingAPI(embeddingSvc)

	mux.HandleFunc("/api/v1/embedding/generate", embeddingAPI.GenerateEmbedding)
	mux.HandleFunc("/api/v1/embedding/refresh", embeddingAPI.RefreshEmbeddings)
	mux.HandleFunc("/api/v1/embedding/status", embeddingAPI.GetStatus)
	mux.HandleFunc("/api/v1/embedding/stats", embeddingAPI.GetStats)
	mux.HandleFunc("/api/v1/embedding/history", embeddingAPI.GetHistory)
	mux.HandleFunc("/api/v1/embedding/provider", embeddingAPI.UpdateProvider)
	mux.HandleFunc("/api/v1/embedding/rebuild", embeddingAPI.RebuildVersion)

	// Hybrid Knowledge Search Engine Module initialization
	searchRepo := searchRepository.NewInMemorySearchRepository()
	searchSvc := searchApp.NewSearchService(searchRepo, embeddingSvc)
	searchEngine := searchApp.NewHybridSearchEngine(searchSvc)
	_ = searchEngine
	searchAPI := searchTransport.NewSearchAPI(searchSvc)

	mux.HandleFunc("/api/v1/search/query", searchAPI.ExecuteSearch)
	mux.HandleFunc("/api/v1/search/saved", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			searchAPI.SaveSearch(w, r)
		case http.MethodDelete:
			searchAPI.DeleteSavedSearch(w, r)
		default:
			searchAPI.GetSavedSearches(w, r)
		}
	})
	mux.HandleFunc("/api/v1/search/history", searchAPI.GetHistory)
	mux.HandleFunc("/api/v1/search/suggestions", searchAPI.GetSuggestions)
	mux.HandleFunc("/api/v1/search/stats", searchAPI.GetStats)
	mux.HandleFunc("/api/v1/search/preferences", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			searchAPI.UpdatePreferences(w, r)
		} else {
			searchAPI.GetPreferences(w, r)
		}
	})

	// Knowledge Insights & Timeline Engine Module initialization
	insightsRepo := insightsRepository.NewInMemoryInsightsRepository()
	insightsSvc := insightsApp.NewInsightsService(insightsRepo)
	insightsEngine := insightsApp.NewKnowledgeInsightsEngine(insightsSvc)
	_ = insightsEngine
	insightsAPI := insightsTransport.NewInsightsAPI(insightsSvc)

	mux.HandleFunc("/api/v1/insights/timeline", insightsAPI.GetTimeline)
	mux.HandleFunc("/api/v1/insights", insightsAPI.GetInsights)
	mux.HandleFunc("/api/v1/insights/milestones", insightsAPI.GetMilestones)
	mux.HandleFunc("/api/v1/insights/dismiss", insightsAPI.DismissInsight)
	mux.HandleFunc("/api/v1/insights/refresh", insightsAPI.RefreshInsights)
	mux.HandleFunc("/api/v1/insights/stats", insightsAPI.GetStats)
	mux.HandleFunc("/api/v1/insights/preferences", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			insightsAPI.UpdatePreferences(w, r)
		} else {
			insightsAPI.GetPreferences(w, r)
		}
	})

	// Declutr AI Copilot RAG Module initialization
	copilotRepo := copilotRepository.NewInMemoryCopilotRepository()
	copilotSvc := copilotApp.NewCopilotService(copilotRepo, searchSvc)
	copilotEngine := copilotApp.NewGroundedRAGEngine(copilotSvc)
	_ = copilotEngine
	copilotAPI := copilotTransport.NewCopilotAPI(copilotSvc)

	mux.HandleFunc("/api/v1/copilot/conversations", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			copilotAPI.StartConversation(w, r)
		case http.MethodDelete:
			copilotAPI.DeleteConversation(w, r)
		default:
			copilotAPI.ListConversations(w, r)
		}
	})
	mux.HandleFunc("/api/v1/copilot/messages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			copilotAPI.SendMessage(w, r)
		} else {
			copilotAPI.GetMessages(w, r)
		}
	})
	mux.HandleFunc("/api/v1/copilot/feedback", copilotAPI.SaveFeedback)
	mux.HandleFunc("/api/v1/copilot/messages/stream", copilotAPI.StreamMessage)

	// Workflow Automation Engine Module initialization
	workflowRepo := workflowRepository.NewInMemoryWorkflowRepository()
	workflowSvc := workflowApp.NewWorkflowService(workflowRepo)
	workflowEngine := workflowApp.NewWorkflowExecutionEngine(workflowSvc)
	_ = workflowEngine
	workflowAPI := workflowTransport.NewWorkflowAPI(workflowSvc)

	mux.HandleFunc("/api/v1/workflows", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			workflowAPI.CreateWorkflow(w, r)
		case http.MethodPut:
			workflowAPI.UpdateWorkflow(w, r)
		case http.MethodDelete:
			workflowAPI.DeleteWorkflow(w, r)
		default:
			workflowAPI.ListWorkflows(w, r)
		}
	})
	mux.HandleFunc("/api/v1/workflows/toggle", workflowAPI.ToggleWorkflow)
	mux.HandleFunc("/api/v1/workflows/run", workflowAPI.RunWorkflow)
	mux.HandleFunc("/api/v1/workflows/history", workflowAPI.GetHistory)
	mux.HandleFunc("/api/v1/workflows/stats", workflowAPI.GetStats)

	// Notification Center & Proactive Intelligence Module initialization
	notificationRepo := notificationRepository.NewInMemoryNotificationRepository()
	notificationSvc := notificationApp.NewNotificationService(notificationRepo)
	notificationEngine := notificationApp.NewNotificationEventEngine(notificationSvc)
	_ = notificationEngine
	notificationAPI := notificationTransport.NewNotificationAPI(notificationSvc)

	mux.HandleFunc("/api/v1/notifications", notificationAPI.ListNotifications)
	mux.HandleFunc("/api/v1/notifications/read", notificationAPI.MarkRead)
	mux.HandleFunc("/api/v1/notifications/dismiss", notificationAPI.DismissNotification)
	mux.HandleFunc("/api/v1/notifications/action", notificationAPI.ExecuteAction)
	mux.HandleFunc("/api/v1/notifications/digests", notificationAPI.GetDigests)
	mux.HandleFunc("/api/v1/notifications/preferences", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			notificationAPI.UpdatePreferences(w, r)
		} else {
			notificationAPI.GetPreferences(w, r)
		}
	})
	mux.HandleFunc("/api/v1/notifications/stats", notificationAPI.GetStats)

	// Secure Sharing & Collaboration Platform Module initialization
	collaborationRepo := collaborationRepository.NewInMemoryCollaborationRepository()
	collaborationSvc := collaborationApp.NewCollaborationService(collaborationRepo)
	permissionValidationEngine := collaborationApp.NewPermissionValidationEngine(collaborationSvc)
	_ = permissionValidationEngine
	collaborationAPI := collaborationTransport.NewCollaborationAPI(collaborationSvc)

	mux.HandleFunc("/api/v1/shares", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			collaborationAPI.CreateShare(w, r)
		case http.MethodDelete:
			collaborationAPI.DeleteShare(w, r)
		default:
			collaborationAPI.ListShares(w, r)
		}
	})
	mux.HandleFunc("/api/v1/shares/invite", collaborationAPI.InviteUser)
	mux.HandleFunc("/api/v1/shares/invite/accept", collaborationAPI.AcceptInvitation)
	mux.HandleFunc("/api/v1/shares/links", collaborationAPI.CreateLink)
	mux.HandleFunc("/api/v1/shares/links/revoke", collaborationAPI.RevokeLink)
	mux.HandleFunc("/api/v1/shares/comments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			collaborationAPI.AddComment(w, r)
		} else {
			collaborationAPI.ListComments(w, r)
		}
	})
	mux.HandleFunc("/api/v1/shares/activity", collaborationAPI.GetActivity)
	mux.HandleFunc("/api/v1/shares/stats", collaborationAPI.GetStats)

	// Version History, Recovery & Time Machine Module initialization
	versioningRepo := versioningRepository.NewInMemoryVersioningRepository()
	versioningSvc := versioningApp.NewVersioningService(versioningRepo)
	timeMachineEngine := versioningApp.NewTimeMachineRecoveryEngine(versioningSvc)
	_ = timeMachineEngine
	versioningAPI := versioningTransport.NewVersioningAPI(versioningSvc)

	mux.HandleFunc("/api/v1/versions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			// delete version endpoint
		} else {
			versioningAPI.ListVersions(w, r)
		}
	})
	mux.HandleFunc("/api/v1/versions/snapshot", versioningAPI.CreateSnapshot)
	mux.HandleFunc("/api/v1/versions/compare", versioningAPI.CompareVersions)
	mux.HandleFunc("/api/v1/versions/restore", versioningAPI.RestoreVersion)
	mux.HandleFunc("/api/v1/recyclebin", versioningAPI.ListRecycleBin)
	mux.HandleFunc("/api/v1/recyclebin/restore", versioningAPI.RestoreRecycleItem)
	mux.HandleFunc("/api/v1/recyclebin/purge", versioningAPI.PurgeRecycleItem)
	mux.HandleFunc("/api/v1/versions/stats", versioningAPI.GetStats)

	// Backup, Disaster Recovery & Business Continuity Module initialization
	backupRepo := backupRepository.NewInMemoryBackupRepository()
	backupSvc := backupApp.NewBackupService(backupRepo)
	disasterRecoveryEngine := backupApp.NewDisasterRecoveryEngine(backupSvc)
	_ = disasterRecoveryEngine
	backupAPI := backupTransport.NewBackupAPI(backupSvc)

	mux.HandleFunc("/api/v1/backups", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			backupAPI.CreateBackup(w, r)
		} else {
			backupAPI.ListBackups(w, r)
		}
	})
	mux.HandleFunc("/api/v1/backups/detail", backupAPI.GetBackupDetails)
	mux.HandleFunc("/api/v1/backups/schedule", backupAPI.ConfigureSchedule)
	mux.HandleFunc("/api/v1/backups/restore", backupAPI.RestoreBackup)
	mux.HandleFunc("/api/v1/backups/verify", backupAPI.VerifyIntegrity)
	mux.HandleFunc("/api/v1/backups/cancel", backupAPI.CancelJob)
	mux.HandleFunc("/api/v1/backups/stats", backupAPI.GetStats)

	// Security Center, Audit Hub & Trust Platform Module initialization
	securityRepo := securityRepository.NewInMemorySecurityRepository()
	securitySvc := securityApp.NewSecurityCenterService(securityRepo)
	riskEngine := securityApp.NewRiskEngine(securitySvc)
	_ = riskEngine
	securityAPI := securityTransport.NewSecurityAPI(securitySvc)

	mux.HandleFunc("/api/v1/security/dashboard", securityAPI.GetDashboard)
	mux.HandleFunc("/api/v1/security/audit", securityAPI.ListAuditEvents)
	mux.HandleFunc("/api/v1/security/sessions", securityAPI.ListSessions)
	mux.HandleFunc("/api/v1/security/sessions/terminate", securityAPI.TerminateSession)
	mux.HandleFunc("/api/v1/security/devices", securityAPI.ListDevices)
	mux.HandleFunc("/api/v1/security/devices/trust", securityAPI.SetDeviceTrust)
	mux.HandleFunc("/api/v1/security/risk", securityAPI.GetRiskAssessment)
	mux.HandleFunc("/api/v1/security/recommendations", securityAPI.GetRecommendations)

	// Offline-First Sync Engine & Conflict Resolution Module initialization
	syncRepo := syncRepository.NewInMemorySyncRepository()
	syncSvc := syncApp.NewSyncService(syncRepo)
	syncEngine := syncApp.NewSyncEngine(syncSvc)
	_ = syncEngine
	syncAPI := syncTransport.NewSyncAPI(syncSvc)

	mux.HandleFunc("/api/v1/sync/push", syncAPI.PushChanges)
	mux.HandleFunc("/api/v1/sync/pull", syncAPI.PullChanges)
	mux.HandleFunc("/api/v1/sync/status", syncAPI.GetStatus)
	mux.HandleFunc("/api/v1/sync/conflicts", syncAPI.ListConflicts)
	mux.HandleFunc("/api/v1/sync/resolve", syncAPI.ResolveConflict)
	mux.HandleFunc("/api/v1/sync/register-device", syncAPI.RegisterDevice)
	mux.HandleFunc("/api/v1/sync/stats", syncAPI.GetStats)

	// Integration Platform & Connector Framework Module initialization
	integrationRepo := integrationsRepository.NewInMemoryIntegrationRepository()
	integrationSvc := integrationsApp.NewIntegrationService(integrationRepo)
	connectorRuntime := integrationsApp.NewConnectorRuntime(integrationSvc)
	_ = connectorRuntime
	integrationAPI := integrationsTransport.NewIntegrationAPI(integrationSvc)

	mux.HandleFunc("/api/v1/integrations", integrationAPI.ListIntegrations)
	mux.HandleFunc("/api/v1/integrations/install", integrationAPI.InstallConnector)
	mux.HandleFunc("/api/v1/integrations/configure", integrationAPI.ConfigureConnector)
	mux.HandleFunc("/api/v1/integrations/enable", integrationAPI.ToggleConnector)
	mux.HandleFunc("/api/v1/integrations/sync", integrationAPI.TriggerSync)
	mux.HandleFunc("/api/v1/integrations/status", integrationAPI.GetStatus)
	mux.HandleFunc("/api/v1/integrations/logs", integrationAPI.GetLogs)
	mux.HandleFunc("/api/v1/integrations/webhooks", integrationAPI.ProcessWebhook)

	// Launch Background Worker Supervisors
	sup := supervisor.GetSupervisor()
	sup.RegisterWorker("w-queue-1", "Queue Ingestion Worker Pool", "QUEUE", func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(5 * time.Second):
				// Poll and process queue items
			}
		}
	})
	sup.RegisterWorker("w-workflow-1", "Workflow Execution Engine", "WORKFLOW", func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(10 * time.Second):
				// Process workflows
			}
		}
	})
	sup.RegisterWorker("w-sync-1", "Sync Queue Daemon", "SYNC", func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(15 * time.Second):
				// Flush sync queues
			}
		}
	})
	_ = sup.StartWorker(context.Background(), "w-queue-1")
	_ = sup.StartWorker(context.Background(), "w-workflow-1")
	_ = sup.StartWorker(context.Background(), "w-sync-1")

	// Apply Middlewares: Tenant Isolation → Security Headers → Rate Limiter → Request Observability
	handler := middleware.TenantMiddleware(
		middleware.SecurityHeaders(
			ratelimit.RateLimitMiddleware(ratelimit.GlobalPolicy, func(r *http.Request) string {
				return ratelimit.ExtractIP(r)
			})(
				middleware.RequestObservability(mux),
			),
		),
	)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Declutr Developer Platform Backend Running on :%s (Environment: %s)", cfg.Port, cfg.Env)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server startup failed: %v", err)
	}
}
