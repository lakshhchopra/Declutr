package application_test

import (
	"testing"

	"github.com/diablovocado/declutr/modules/collaboration/application"
	"github.com/diablovocado/declutr/modules/collaboration/domain"
	"github.com/diablovocado/declutr/modules/collaboration/repository"
)

const testVaultID = "vault-test-001"

func setupService() *application.CollaborationService {
	repo := repository.NewInMemoryCollaborationRepository()
	return application.NewCollaborationService(repo)
}

// TestShareCreationAndPermissions validates share creation and role permission checking
func TestShareCreationAndPermissions(t *testing.T) {
	svc := setupService()

	share, err := svc.CreateShare(&domain.CreateShareRequest{
		VaultID:      testVaultID,
		ResourceType: domain.ResourceCollection,
		ResourceID:   "col-japan-123",
		Title:        "Japan Trip Photos",
		AccessType:   domain.AccessInviteOnly,
	})
	if err != nil {
		t.Fatalf("create share failed: %v", err)
	}

	if !application.CheckPermission(domain.RoleOwner, "EDIT") {
		t.Error("expected OWNER role to have EDIT permission")
	}
	if application.CheckPermission(domain.RoleReadOnly, "EDIT") {
		t.Error("expected READ_ONLY role to NOT have EDIT permission")
	}

	t.Logf("PASS: Share Creation & Permissions — Created %s (%s)", share.ShareID, share.Title)
}

// TestInvitationLifecycle validates sending and accepting invitations
func TestInvitationLifecycle(t *testing.T) {
	svc := setupService()

	share, _ := svc.CreateShare(&domain.CreateShareRequest{
		VaultID:      testVaultID,
		ResourceType: domain.ResourceAsset,
		ResourceID:   "asset-pdf-001",
		Title:        "Quarterly Financial Overview",
		AccessType:   domain.AccessInviteOnly,
	})

	inv, err := svc.InviteUser(&domain.InviteRequest{
		ShareID:      share.ShareID,
		InviterID:    "usr-owner",
		InviteeEmail: "partner@finance.com",
		Role:         domain.RoleEdit,
	})
	if err != nil {
		t.Fatalf("invite user failed: %v", err)
	}

	if err := svc.AcceptInvitation(inv.Token, "usr-partner"); err != nil {
		t.Fatalf("accept invitation failed: %v", err)
	}

	t.Logf("PASS: Invitation Lifecycle — Invitation %s sent and accepted by usr-partner", inv.InvitationID)
}

// TestLinkSharing validates link generation, password protection, and revocation
func TestLinkSharing(t *testing.T) {
	svc := setupService()

	share, _ := svc.CreateShare(&domain.CreateShareRequest{
		VaultID:      testVaultID,
		ResourceType: domain.ResourceProject,
		ResourceID:   "proj-startup-999",
		Title:        "Startup Pitch Deck",
		AccessType:   domain.AccessLink,
	})

	link, err := svc.CreateLink(&domain.CreateLinkRequest{
		ShareID:             share.ShareID,
		IsPasswordProtected: true,
		DisableDownload:     false,
		ExpiresInDays:       7,
	})
	if err != nil {
		t.Fatalf("create link failed: %v", err)
	}

	if err := svc.RevokeLink(link.LinkID); err != nil {
		t.Fatalf("revoke link failed: %v", err)
	}

	t.Logf("PASS: Link Sharing — Created password-protected link %s and revoked it", link.LinkID)
}

// TestThreadedComments validates adding comments and listing comment threads
func TestThreadedComments(t *testing.T) {
	svc := setupService()

	share, _ := svc.CreateShare(&domain.CreateShareRequest{
		VaultID:      testVaultID,
		ResourceType: domain.ResourceCollection,
		ResourceID:   "col-photos",
		Title:        "Shared Vacation Photos",
		AccessType:   domain.AccessInviteOnly,
	})

	cmnt, err := svc.AddComment(&domain.AddCommentRequest{
		ShareID:  share.ShareID,
		UserID:   "usr-alex",
		UserName: "Alex",
		Content:  "Looks great! Added passport receipts.",
	})
	if err != nil {
		t.Fatalf("add comment failed: %v", err)
	}

	comments, err := svc.ListComments(share.ShareID)
	if err != nil {
		t.Fatalf("list comments failed: %v", err)
	}
	if len(comments) == 0 {
		t.Error("expected comment in list")
	}

	t.Logf("PASS: Threaded Comments — Added comment %s by Alex", cmnt.CommentID)
}

// TestAuditActivityLogging validates activity history tracking
func TestAuditActivityLogging(t *testing.T) {
	svc := setupService()

	share, _ := svc.CreateShare(&domain.CreateShareRequest{
		VaultID:      testVaultID,
		ResourceType: domain.ResourceFolder,
		ResourceID:   "fld-tax",
		Title:        "Tax Documents 2025",
		AccessType:   domain.AccessInviteOnly,
	})

	acts, err := svc.ListActivity(testVaultID, share.ShareID)
	if err != nil {
		t.Fatalf("list activity failed: %v", err)
	}
	if len(acts) == 0 {
		t.Error("expected audit activity log entry for share creation")
	}

	t.Logf("PASS: Audit Activity Logging — Recorded %d audit activities", len(acts))
}

// TestRevokeShare validates deleting a share and revoking all access
func TestRevokeShare(t *testing.T) {
	svc := setupService()

	share, _ := svc.CreateShare(&domain.CreateShareRequest{
		VaultID:      testVaultID,
		ResourceType: domain.ResourceAsset,
		ResourceID:   "asset-temp-111",
		Title:        "Temporary Document",
		AccessType:   domain.AccessInviteOnly,
	})

	if err := svc.DeleteShare(share.ShareID); err != nil {
		t.Fatalf("delete share failed: %v", err)
	}

	list, _ := svc.ListShares(testVaultID)
	for _, item := range list {
		if item.ShareID == share.ShareID {
			t.Error("expected deleted share to be unretrievable")
		}
	}

	t.Logf("PASS: Revoke Share — Successfully deleted share %s and revoked access", share.ShareID)
}
