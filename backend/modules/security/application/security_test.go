package application_test

import (
	"context"
	"testing"

	"github.com/diablovocado/declutr/modules/security/application"
	"github.com/diablovocado/declutr/modules/security/domain"
	"github.com/diablovocado/declutr/modules/security/repository"
)

const testVaultID = "vault-test-001"

func setupService() (*application.SecurityCenterService, *application.RiskEngine) {
	repo := repository.NewInMemorySecurityRepository()
	svc := application.NewSecurityCenterService(repo)
	riskEng := application.NewRiskEngine(svc)
	return svc, riskEng
}

// TestSecurityDashboard validates security dashboard payload generation
func TestSecurityDashboard(t *testing.T) {
	svc, _ := setupService()

	dash, err := svc.GetDashboard(testVaultID)
	if err != nil {
		t.Fatalf("get dashboard failed: %v", err)
	}
	if dash.Score == nil || dash.Score.Score != 92 {
		t.Errorf("expected security score 92, got %v", dash.Score)
	}
	if len(dash.ActiveSessions) == 0 {
		t.Error("expected active sessions in dashboard")
	}

	t.Logf("PASS: Security Dashboard — Score: %d (Grade: %s), Sessions: %d",
		dash.Score.Score, dash.Score.Grade, len(dash.ActiveSessions))
}

// TestAuditLogging validates asynchronous audit event recording and category filtering
func TestAuditLogging(t *testing.T) {
	svc, _ := setupService()
	ctx := context.Background()

	err := svc.RecordAuditEvent(ctx, &domain.RecordAuditEventRequest{
		VaultID:   testVaultID,
		Category:  domain.AuditAuth,
		Action:    "PASSWORD_CHANGED",
		ActorID:   "usr-owner",
		ActorName: "Vault Owner",
		IPAddress: "192.168.1.100",
	})
	if err != nil {
		t.Fatalf("record audit event failed: %v", err)
	}

	events, err := svc.ListAuditEvents(testVaultID, domain.AuditAuth, 10)
	if err != nil {
		t.Fatalf("list audit events failed: %v", err)
	}
	if len(events) == 0 {
		t.Error("expected audit event in list")
	}

	t.Logf("PASS: Audit Logging — Recorded audit event %s (%s)", events[0].Action, events[0].Category)
}

// TestSessionTermination validates listing and terminating active user sessions
func TestSessionTermination(t *testing.T) {
	svc, _ := setupService()

	sessions, err := svc.ListSessions(testVaultID)
	if err != nil {
		t.Fatalf("list sessions failed: %v", err)
	}
	if len(sessions) == 0 {
		t.Error("expected active sessions")
	}

	targetSessID := sessions[0].SessionID
	err = svc.TerminateSession(&domain.TerminateSessionRequest{
		VaultID:   testVaultID,
		SessionID: targetSessID,
	})
	if err != nil {
		t.Fatalf("terminate session failed: %v", err)
	}

	t.Logf("PASS: Session Termination — Terminated session %s", targetSessID)
}

// TestDeviceTrustToggle validates device trust status updating
func TestDeviceTrustToggle(t *testing.T) {
	svc, _ := setupService()

	devices, err := svc.ListDevices(testVaultID)
	if err != nil {
		t.Fatalf("list devices failed: %v", err)
	}
	if len(devices) == 0 {
		t.Error("expected devices in list")
	}

	targetDevID := devices[0].DeviceID
	err = svc.SetDeviceTrust(&domain.TrustDeviceRequest{
		VaultID:  testVaultID,
		DeviceID: targetDevID,
		Trust:    false,
	})
	if err != nil {
		t.Fatalf("set device trust failed: %v", err)
	}

	t.Logf("PASS: Device Trust Toggle — Updated device %s trust state to false", targetDevID)
}

// TestRiskEngineScoring validates dynamic risk score computation and signal detection
func TestRiskEngineScoring(t *testing.T) {
	svc, riskEng := setupService()
	ctx := context.Background()

	risk, err := riskEng.AssessVaultRisk(ctx, testVaultID)
	if err != nil {
		t.Fatalf("assess vault risk failed: %v", err)
	}
	if risk.RiskScore != 15 || risk.RiskLevel != domain.RiskLow {
		t.Errorf("expected score 15 LOW risk, got %d %s", risk.RiskScore, risk.RiskLevel)
	}

	fetchedRisk, err := svc.GetRiskAssessment(testVaultID)
	if err != nil || fetchedRisk == nil {
		t.Fatalf("get risk assessment failed: %v", err)
	}

	t.Logf("PASS: Risk Engine Scoring — Assessed risk score %d (%s) with %d signals",
		risk.RiskScore, risk.RiskLevel, len(risk.Signals))
}

// TestSecurityRecommendations validates actionable security recommendation generation
func TestSecurityRecommendations(t *testing.T) {
	svc, _ := setupService()

	recs, err := svc.GetRecommendations(testVaultID)
	if err != nil {
		t.Fatalf("get recommendations failed: %v", err)
	}
	if len(recs) == 0 {
		t.Error("expected recommendations")
	}

	t.Logf("PASS: Security Recommendations — Retrieved %d actionable recommendations (e.g. %s)",
		len(recs), recs[0].Title)
}
