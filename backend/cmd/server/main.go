package main

import (
	"log"
	"net/http"

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
	embeddingApp "github.com/diablovocado/declutr/modules/embedding/application"
	embeddingRepository "github.com/diablovocado/declutr/modules/embedding/repository"
	embeddingTransport "github.com/diablovocado/declutr/modules/embedding/transport"
	insightsApp "github.com/diablovocado/declutr/modules/insights/application"
	insightsRepository "github.com/diablovocado/declutr/modules/insights/repository"
	insightsTransport "github.com/diablovocado/declutr/modules/insights/transport"
	memoryApp "github.com/diablovocado/declutr/modules/memory/application"
	memoryRepository "github.com/diablovocado/declutr/modules/memory/repository"
	memoryTransport "github.com/diablovocado/declutr/modules/memory/transport"
	notificationApp "github.com/diablovocado/declutr/modules/notification/application"
	notificationRepository "github.com/diablovocado/declutr/modules/notification/repository"
	notificationTransport "github.com/diablovocado/declutr/modules/notification/transport"
	personaApp "github.com/diablovocado/declutr/modules/persona/application"
	personaRepository "github.com/diablovocado/declutr/modules/persona/repository"
	personaTransport "github.com/diablovocado/declutr/modules/persona/transport"
	searchApp "github.com/diablovocado/declutr/modules/search/application"
	searchRepository "github.com/diablovocado/declutr/modules/search/repository"
	searchTransport "github.com/diablovocado/declutr/modules/search/transport"
	versioningApp "github.com/diablovocado/declutr/modules/versioning/application"
	versioningRepository "github.com/diablovocado/declutr/modules/versioning/repository"
	versioningTransport "github.com/diablovocado/declutr/modules/versioning/transport"
	workflowApp "github.com/diablovocado/declutr/modules/workflow/application"
	workflowRepository "github.com/diablovocado/declutr/modules/workflow/repository"
	workflowTransport "github.com/diablovocado/declutr/modules/workflow/transport"
	"github.com/diablovocado/declutr/pkg/health"
	"github.com/diablovocado/declutr/shared/database"
	"github.com/diablovocado/declutr/shared/middleware"
)

func main() {
	db := database.Connect()

	userRepo := &repository.PostgresUserRepository{
		DB: db,
	}

	authService := &application.Service{
		UserRepo:   userRepo,
		Challenges: application.NewChallengeStore(),
		SRP:        application.NewEngine(),
	}

	http.HandleFunc("/health", health.Handler)

	http.HandleFunc(
		"/api/v1/auth/register",
		transport.RegisterHandler(authService),
	)

	http.HandleFunc(
		"/api/v1/auth/login/start",
		transport.LoginStartHandler(authService),
	)
	http.Handle(
		"/api/v1/me",
		middleware.Auth(userRepo)(transport.MeHandler()),
	)

	http.HandleFunc(
		"/api/v1/auth/login/finish",
		transport.LoginFinishHandler(authService),
	)

	// Context & Intent Engine Module initialization
	contextRepo := contextRepository.NewInMemoryContextRepository()
	contextService := contextApp.NewContextService(contextRepo, nil)
	contextAPI := contextTransport.NewAPI(contextService)

	http.HandleFunc("/api/v1/context", contextAPI.GetContextsHandler)
	http.HandleFunc("/api/v1/context/details", contextAPI.GetContextDetailHandler)
	http.HandleFunc("/api/v1/context/refresh", contextAPI.RefreshContextHandler)
	http.HandleFunc("/api/v1/context/intent", contextAPI.GetIntentHandler)
	http.HandleFunc("/api/v1/context/stats", contextAPI.GetStatsHandler)

	// Reverse Persona Engine Module initialization
	personaRepo := personaRepository.NewInMemoryPersonaRepository()
	personaSvc := personaApp.NewPersonaService(personaRepo)
	personaEngine := personaApp.NewPersonaEngine(personaSvc)
	_ = personaEngine // available for worker dispatch
	personaAPI := personaTransport.NewPersonaAPI(personaSvc)

	http.HandleFunc("/api/v1/persona", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			personaAPI.DeletePersona(w, r)
		} else {
			personaAPI.GetPersona(w, r)
		}
	})
	http.HandleFunc("/api/v1/persona/recommendations", personaAPI.GetRecommendations)
	http.HandleFunc("/api/v1/persona/settings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			personaAPI.UpdateSettings(w, r)
		} else {
			personaAPI.GetSettings(w, r)
		}
	})
	http.HandleFunc("/api/v1/persona/reset", personaAPI.ResetPersona)
	http.HandleFunc("/api/v1/persona/export", personaAPI.ExportPersona)
	http.HandleFunc("/api/v1/persona/signal", personaAPI.RecordSignal)
	http.HandleFunc("/api/v1/persona/history", personaAPI.GetHistory)

	// Memory Engine Module initialization
	// Pipeline: Context → Persona → Memory Formation → Knowledge Memory
	memoryRepo := memoryRepository.NewInMemoryMemoryRepository()
	memorySvc := memoryApp.NewMemoryService(memoryRepo)
	memoryEngine := memoryApp.NewMemoryEngine(memorySvc)
	_ = memoryEngine // available for worker dispatch
	memoryAPI := memoryTransport.NewMemoryAPI(memorySvc)

	http.HandleFunc("/api/v1/memory", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			memoryAPI.DeleteMemory(w, r)
		} else {
			memoryAPI.GetMemories(w, r)
		}
	})
	http.HandleFunc("/api/v1/memory/timeline", memoryAPI.GetTimeline)
	http.HandleFunc("/api/v1/memory/detail", memoryAPI.GetMemoryDetail)
	http.HandleFunc("/api/v1/memory/refresh", memoryAPI.RefreshMemory)
	http.HandleFunc("/api/v1/memory/pin", memoryAPI.PinMemory)
	http.HandleFunc("/api/v1/memory/archive", memoryAPI.ArchiveMemory)
	http.HandleFunc("/api/v1/memory/stats", memoryAPI.GetStats)
	http.HandleFunc("/api/v1/memory/reset", memoryAPI.ResetMemory)
	http.HandleFunc("/api/v1/memory/settings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			memoryAPI.UpdateSettings(w, r)
		} else {
			memoryAPI.GetSettings(w, r)
		}
	})

	// Embedding Engine Module initialization
	// Pipeline: Memory Engine → Embedding Engine → Vector Storage
	embeddingRepo := embeddingRepository.NewInMemoryVectorRepository()
	embeddingSvc := embeddingApp.NewEmbeddingService(embeddingRepo)
	embeddingEngine := embeddingApp.NewEmbeddingEngine(embeddingSvc)
	_ = embeddingEngine // available for worker dispatch
	embeddingAPI := embeddingTransport.NewEmbeddingAPI(embeddingSvc)

	http.HandleFunc("/api/v1/embedding/generate", embeddingAPI.GenerateEmbedding)
	http.HandleFunc("/api/v1/embedding/refresh", embeddingAPI.RefreshEmbeddings)
	http.HandleFunc("/api/v1/embedding/status", embeddingAPI.GetStatus)
	http.HandleFunc("/api/v1/embedding/stats", embeddingAPI.GetStats)
	http.HandleFunc("/api/v1/embedding/history", embeddingAPI.GetHistory)
	http.HandleFunc("/api/v1/embedding/provider", embeddingAPI.UpdateProvider)
	http.HandleFunc("/api/v1/embedding/rebuild", embeddingAPI.RebuildVersion)

	// Hybrid Knowledge Search Engine Module initialization
	searchRepo := searchRepository.NewInMemorySearchRepository()
	searchSvc := searchApp.NewSearchService(searchRepo, embeddingSvc)
	searchEngine := searchApp.NewHybridSearchEngine(searchSvc)
	_ = searchEngine // available for internal retrieval calls
	searchAPI := searchTransport.NewSearchAPI(searchSvc)

	http.HandleFunc("/api/v1/search/query", searchAPI.ExecuteSearch)
	http.HandleFunc("/api/v1/search/saved", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			searchAPI.SaveSearch(w, r)
		case http.MethodDelete:
			searchAPI.DeleteSavedSearch(w, r)
		default:
			searchAPI.GetSavedSearches(w, r)
		}
	})
	http.HandleFunc("/api/v1/search/history", searchAPI.GetHistory)
	http.HandleFunc("/api/v1/search/suggestions", searchAPI.GetSuggestions)
	http.HandleFunc("/api/v1/search/stats", searchAPI.GetStats)
	http.HandleFunc("/api/v1/search/preferences", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			searchAPI.UpdatePreferences(w, r)
		} else {
			searchAPI.GetPreferences(w, r)
		}
	})

	// Knowledge Insights & Timeline Engine Module initialization
	// Pipeline: Assets → Metadata → Entities → Relationships → Contexts → Memory → Search → Knowledge Insights
	insightsRepo := insightsRepository.NewInMemoryInsightsRepository()
	insightsSvc := insightsApp.NewInsightsService(insightsRepo)
	insightsEngine := insightsApp.NewKnowledgeInsightsEngine(insightsSvc)
	_ = insightsEngine // available for background processing
	insightsAPI := insightsTransport.NewInsightsAPI(insightsSvc)

	http.HandleFunc("/api/v1/insights/timeline", insightsAPI.GetTimeline)
	http.HandleFunc("/api/v1/insights", insightsAPI.GetInsights)
	http.HandleFunc("/api/v1/insights/milestones", insightsAPI.GetMilestones)
	http.HandleFunc("/api/v1/insights/dismiss", insightsAPI.DismissInsight)
	http.HandleFunc("/api/v1/insights/refresh", insightsAPI.RefreshInsights)
	http.HandleFunc("/api/v1/insights/stats", insightsAPI.GetStats)
	http.HandleFunc("/api/v1/insights/preferences", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			insightsAPI.UpdatePreferences(w, r)
		} else {
			insightsAPI.GetPreferences(w, r)
		}
	})

	// Declutr AI Copilot RAG Module initialization
	// Pipeline: User Question → Intent Parser → Hybrid Search → Context Builder → Prompt Builder → Grounded RAG Answer → Citations
	copilotRepo := copilotRepository.NewInMemoryCopilotRepository()
	copilotSvc := copilotApp.NewCopilotService(copilotRepo, searchSvc)
	copilotEngine := copilotApp.NewGroundedRAGEngine(copilotSvc)
	_ = copilotEngine // available for direct RAG execution
	copilotAPI := copilotTransport.NewCopilotAPI(copilotSvc)

	http.HandleFunc("/api/v1/copilot/conversations", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			copilotAPI.StartConversation(w, r)
		case http.MethodDelete:
			copilotAPI.DeleteConversation(w, r)
		default:
			copilotAPI.ListConversations(w, r)
		}
	})
	http.HandleFunc("/api/v1/copilot/messages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			copilotAPI.SendMessage(w, r)
		} else {
			copilotAPI.GetMessages(w, r)
		}
	})
	http.HandleFunc("/api/v1/copilot/feedback", copilotAPI.SaveFeedback)
	http.HandleFunc("/api/v1/copilot/messages/stream", copilotAPI.StreamMessage)

	// Workflow Automation Engine Module initialization
	// Pipeline: Event → Trigger Engine → Condition Engine → Rule Engine → Action Engine → Execution → History
	workflowRepo := workflowRepository.NewInMemoryWorkflowRepository()
	workflowSvc := workflowApp.NewWorkflowService(workflowRepo)
	workflowEngine := workflowApp.NewWorkflowExecutionEngine(workflowSvc)
	_ = workflowEngine // available for background event dispatching
	workflowAPI := workflowTransport.NewWorkflowAPI(workflowSvc)

	http.HandleFunc("/api/v1/workflows", func(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/api/v1/workflows/toggle", workflowAPI.ToggleWorkflow)
	http.HandleFunc("/api/v1/workflows/run", workflowAPI.RunWorkflow)
	http.HandleFunc("/api/v1/workflows/history", workflowAPI.GetHistory)
	http.HandleFunc("/api/v1/workflows/stats", workflowAPI.GetStats)

	// Notification Center & Proactive Intelligence Module initialization
	// Pipeline: Domain Event → Notification Rules → Priority Engine → Delivery Scheduler → Notification Center → User Action
	notificationRepo := notificationRepository.NewInMemoryNotificationRepository()
	notificationSvc := notificationApp.NewNotificationService(notificationRepo)
	notificationEngine := notificationApp.NewNotificationEventEngine(notificationSvc)
	_ = notificationEngine // available for domain event subscription
	notificationAPI := notificationTransport.NewNotificationAPI(notificationSvc)

	http.HandleFunc("/api/v1/notifications", notificationAPI.ListNotifications)
	http.HandleFunc("/api/v1/notifications/read", notificationAPI.MarkRead)
	http.HandleFunc("/api/v1/notifications/dismiss", notificationAPI.DismissNotification)
	http.HandleFunc("/api/v1/notifications/action", notificationAPI.ExecuteAction)
	http.HandleFunc("/api/v1/notifications/digests", notificationAPI.GetDigests)
	http.HandleFunc("/api/v1/notifications/preferences", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			notificationAPI.UpdatePreferences(w, r)
		} else {
			notificationAPI.GetPreferences(w, r)
		}
	})
	http.HandleFunc("/api/v1/notifications/stats", notificationAPI.GetStats)

	// Secure Sharing & Collaboration Platform Module initialization
	// Pipeline: User → Permission Engine → Share Manager → Access Validation → Audit System → Shared Resource
	collaborationRepo := collaborationRepository.NewInMemoryCollaborationRepository()
	collaborationSvc := collaborationApp.NewCollaborationService(collaborationRepo)
	permissionValidationEngine := collaborationApp.NewPermissionValidationEngine(collaborationSvc)
	_ = permissionValidationEngine // available for access validation
	collaborationAPI := collaborationTransport.NewCollaborationAPI(collaborationSvc)

	http.HandleFunc("/api/v1/shares", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			collaborationAPI.CreateShare(w, r)
		case http.MethodDelete:
			collaborationAPI.DeleteShare(w, r)
		default:
			collaborationAPI.ListShares(w, r)
		}
	})
	http.HandleFunc("/api/v1/shares/invite", collaborationAPI.InviteUser)
	http.HandleFunc("/api/v1/shares/invite/accept", collaborationAPI.AcceptInvitation)
	http.HandleFunc("/api/v1/shares/links", collaborationAPI.CreateLink)
	http.HandleFunc("/api/v1/shares/links/revoke", collaborationAPI.RevokeLink)
	http.HandleFunc("/api/v1/shares/comments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			collaborationAPI.AddComment(w, r)
		} else {
			collaborationAPI.ListComments(w, r)
		}
	})
	http.HandleFunc("/api/v1/shares/activity", collaborationAPI.GetActivity)
	http.HandleFunc("/api/v1/shares/stats", collaborationAPI.GetStats)

	// Version History, Recovery & Time Machine Module initialization
	// Pipeline: Asset Change → Version Manager → Snapshot Generator → Version Store → Recovery Engine → Restore
	versioningRepo := versioningRepository.NewInMemoryVersioningRepository()
	versioningSvc := versioningApp.NewVersioningService(versioningRepo)
	timeMachineEngine := versioningApp.NewTimeMachineRecoveryEngine(versioningSvc)
	_ = timeMachineEngine // available for auto snapshot capture
	versioningAPI := versioningTransport.NewVersioningAPI(versioningSvc)

	http.HandleFunc("/api/v1/versions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			// delete version endpoint
		} else {
			versioningAPI.ListVersions(w, r)
		}
	})
	http.HandleFunc("/api/v1/versions/snapshot", versioningAPI.CreateSnapshot)
	http.HandleFunc("/api/v1/versions/compare", versioningAPI.CompareVersions)
	http.HandleFunc("/api/v1/versions/restore", versioningAPI.RestoreVersion)
	http.HandleFunc("/api/v1/recyclebin", versioningAPI.ListRecycleBin)
	http.HandleFunc("/api/v1/recyclebin/restore", versioningAPI.RestoreRecycleItem)
	http.HandleFunc("/api/v1/recyclebin/purge", versioningAPI.PurgeRecycleItem)
	http.HandleFunc("/api/v1/versions/stats", versioningAPI.GetStats)

	// Backup, Disaster Recovery & Business Continuity Module initialization
	// Pipeline: Vault → Backup Scheduler → Snapshot Engine → Backup Storage → Integrity Verification → Recovery Manager → Restore
	backupRepo := backupRepository.NewInMemoryBackupRepository()
	backupSvc := backupApp.NewBackupService(backupRepo)
	disasterRecoveryEngine := backupApp.NewDisasterRecoveryEngine(backupSvc)
	_ = disasterRecoveryEngine // available for scheduled automated backups
	backupAPI := backupTransport.NewBackupAPI(backupSvc)

	http.HandleFunc("/api/v1/backups", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			backupAPI.CreateBackup(w, r)
		} else {
			backupAPI.ListBackups(w, r)
		}
	})
	http.HandleFunc("/api/v1/backups/detail", backupAPI.GetBackupDetails)
	http.HandleFunc("/api/v1/backups/schedule", backupAPI.ConfigureSchedule)
	http.HandleFunc("/api/v1/backups/restore", backupAPI.RestoreBackup)
	http.HandleFunc("/api/v1/backups/verify", backupAPI.VerifyIntegrity)
	http.HandleFunc("/api/v1/backups/cancel", backupAPI.CancelJob)
	http.HandleFunc("/api/v1/backups/stats", backupAPI.GetStats)

	log.Println("Declutr Backend Running on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
