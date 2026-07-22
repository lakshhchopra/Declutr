package repository

import (
	"fmt"
	"sync"
	"time"

	"github.com/diablovocado/declutr/modules/collaboration/domain"
)

// CollaborationRepository defines persistence contract for shares, members, links, comments, and audit activity
type CollaborationRepository interface {
	CreateShare(s *domain.Share) error
	GetShare(shareID string) (*domain.Share, error)
	ListShares(vaultID string) ([]*domain.Share, error)
	DeleteShare(shareID string) error

	AddMember(member *domain.ShareMember) error
	RemoveMember(shareID string, memberID string) error
	ListMembers(shareID string) ([]*domain.ShareMember, error)

	CreateLink(link *domain.ShareLink) error
	RevokeLink(linkID string) error
	GetLinkByToken(token string) (*domain.ShareLink, error)

	AddComment(comment *domain.ShareComment) error
	ListComments(shareID string) ([]*domain.ShareComment, error)
	ResolveComment(commentID string) error

	AddActivity(act *domain.ShareActivity) error
	GetActivity(shareID string) ([]*domain.ShareActivity, error)
	ListAllActivity(vaultID string) ([]*domain.ShareActivity, error)

	CreateInvitation(inv *domain.ShareInvitation) error
	GetInvitationByToken(token string) (*domain.ShareInvitation, error)
	UpdateInvitationStatus(invID string, status domain.InviteStatus) error

	GetStats(vaultID string) (*domain.ShareStats, error)
	ClearAllData(vaultID string) error
}

// InMemoryCollaborationRepository is a thread-safe in-memory store
type InMemoryCollaborationRepository struct {
	mu          sync.RWMutex
	shares      map[string]*domain.Share           // shareID -> Share
	members     map[string][]*domain.ShareMember   // shareID -> Members
	links       map[string][]*domain.ShareLink     // shareID -> Links
	comments    map[string][]*domain.ShareComment  // shareID -> Comments
	activity    map[string][]*domain.ShareActivity // shareID -> Activity
	invitations map[string]*domain.ShareInvitation // token -> Invitation
}

// NewInMemoryCollaborationRepository creates a new in-memory collaboration repository
func NewInMemoryCollaborationRepository() *InMemoryCollaborationRepository {
	return &InMemoryCollaborationRepository{
		shares:      make(map[string]*domain.Share),
		members:     make(map[string][]*domain.ShareMember),
		links:       make(map[string][]*domain.ShareLink),
		comments:    make(map[string][]*domain.ShareComment),
		activity:    make(map[string][]*domain.ShareActivity),
		invitations: make(map[string]*domain.ShareInvitation),
	}
}

func (r *InMemoryCollaborationRepository) CreateShare(s *domain.Share) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	s.UpdatedAt = time.Now()
	r.shares[s.ShareID] = s
	return nil
}

func (r *InMemoryCollaborationRepository) GetShare(shareID string) (*domain.Share, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	s, ok := r.shares[shareID]
	if !ok {
		return nil, fmt.Errorf("share %s not found", shareID)
	}
	return s, nil
}

func (r *InMemoryCollaborationRepository) ListShares(vaultID string) ([]*domain.Share, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var list []*domain.Share
	for _, s := range r.shares {
		if s.VaultID == vaultID {
			list = append(list, s)
		}
	}
	if len(list) == 0 {
		list = defaultSampleShares(vaultID)
		for _, s := range list {
			r.shares[s.ShareID] = s
		}
	}
	return list, nil
}

func (r *InMemoryCollaborationRepository) DeleteShare(shareID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.shares, shareID)
	delete(r.members, shareID)
	delete(r.links, shareID)
	delete(r.comments, shareID)
	delete(r.activity, shareID)
	return nil
}

func (r *InMemoryCollaborationRepository) AddMember(member *domain.ShareMember) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.members[member.ShareID] = append(r.members[member.ShareID], member)
	return nil
}

func (r *InMemoryCollaborationRepository) RemoveMember(shareID string, memberID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	list := r.members[shareID]
	var updated []*domain.ShareMember
	for _, m := range list {
		if m.MemberID != memberID {
			updated = append(updated, m)
		}
	}
	r.members[shareID] = updated
	return nil
}

func (r *InMemoryCollaborationRepository) ListMembers(shareID string) ([]*domain.ShareMember, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.members[shareID], nil
}

func (r *InMemoryCollaborationRepository) CreateLink(link *domain.ShareLink) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.links[link.ShareID] = append(r.links[link.ShareID], link)
	return nil
}

func (r *InMemoryCollaborationRepository) RevokeLink(linkID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for sID, links := range r.links {
		var updated []*domain.ShareLink
		for _, l := range links {
			if l.LinkID != linkID {
				updated = append(updated, l)
			}
		}
		r.links[sID] = updated
	}
	return nil
}

func (r *InMemoryCollaborationRepository) GetLinkByToken(token string) (*domain.ShareLink, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, links := range r.links {
		for _, l := range links {
			if l.LinkToken == token {
				return l, nil
			}
		}
	}
	return nil, fmt.Errorf("link token %s not found", token)
}

func (r *InMemoryCollaborationRepository) AddComment(comment *domain.ShareComment) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.comments[comment.ShareID] = append(r.comments[comment.ShareID], comment)
	return nil
}

func (r *InMemoryCollaborationRepository) ListComments(shareID string) ([]*domain.ShareComment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.comments[shareID], nil
}

func (r *InMemoryCollaborationRepository) ResolveComment(commentID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, list := range r.comments {
		for _, c := range list {
			if c.CommentID == commentID {
				c.IsResolved = true
				c.UpdatedAt = time.Now()
				return nil
			}
		}
	}
	return fmt.Errorf("comment %s not found", commentID)
}

func (r *InMemoryCollaborationRepository) AddActivity(act *domain.ShareActivity) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.activity[act.ShareID] = append(r.activity[act.ShareID], act)
	return nil
}

func (r *InMemoryCollaborationRepository) GetActivity(shareID string) ([]*domain.ShareActivity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.activity[shareID], nil
}

func (r *InMemoryCollaborationRepository) ListAllActivity(vaultID string) ([]*domain.ShareActivity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var all []*domain.ShareActivity
	for _, acts := range r.activity {
		for _, act := range acts {
			if act.VaultID == vaultID {
				all = append(all, act)
			}
		}
	}
	if len(all) == 0 {
		return defaultSampleActivity(vaultID), nil
	}
	return all, nil
}

func (r *InMemoryCollaborationRepository) CreateInvitation(inv *domain.ShareInvitation) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.invitations[inv.Token] = inv
	return nil
}

func (r *InMemoryCollaborationRepository) GetInvitationByToken(token string) (*domain.ShareInvitation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	inv, ok := r.invitations[token]
	if !ok {
		return nil, fmt.Errorf("invitation token %s not found", token)
	}
	return inv, nil
}

func (r *InMemoryCollaborationRepository) UpdateInvitationStatus(invID string, status domain.InviteStatus) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, inv := range r.invitations {
		if inv.InvitationID == invID {
			inv.Status = status
			return nil
		}
	}
	return fmt.Errorf("invitation %s not found", invID)
}

func (r *InMemoryCollaborationRepository) GetStats(vaultID string) (*domain.ShareStats, error) {
	shares, _ := r.ListShares(vaultID)
	acts, _ := r.ListAllActivity(vaultID)

	linksCount := 0
	membersCount := 0
	commentsCount := 0

	for _, s := range shares {
		linksCount += len(s.Links)
		membersCount += len(s.Members)
		cList, _ := r.ListComments(s.ShareID)
		commentsCount += len(cList)
	}

	return &domain.ShareStats{
		VaultID:            vaultID,
		TotalShares:        len(shares),
		ActiveLinks:        linksCount,
		TotalMembers:       membersCount,
		TotalComments:      commentsCount,
		AuditActivityCount: len(acts),
	}, nil
}

func (r *InMemoryCollaborationRepository) ClearAllData(vaultID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for id, s := range r.shares {
		if s.VaultID == vaultID {
			delete(r.shares, id)
			delete(r.members, id)
			delete(r.links, id)
			delete(r.comments, id)
			delete(r.activity, id)
		}
	}
	return nil
}

// Sample Data Generators
func defaultSampleShares(vaultID string) []*domain.Share {
	now := time.Now()
	return []*domain.Share{
		{
			ShareID:      "share-japan-001",
			VaultID:      vaultID,
			ResourceType: domain.ResourceCollection,
			ResourceID:   "col-japan-vacation",
			Title:        "Japan Trip Photos & Itinerary",
			AccessType:   domain.AccessInviteOnly,
			Members: []domain.ShareMember{
				{MemberID: "mem-1", ShareID: "share-japan-001", UserID: "usr-owner", Email: "owner@declutr.local", Role: domain.RoleOwner, JoinedAt: now.Add(-7 * 24 * time.Hour)},
				{MemberID: "mem-2", ShareID: "share-japan-001", UserID: "usr-alex", Email: "alex@travel.org", Role: domain.RoleEdit, JoinedAt: now.Add(-3 * 24 * time.Hour)},
			},
			Links: []domain.ShareLink{
				{LinkID: "link-1", ShareID: "share-japan-001", LinkToken: "tok-japan-public-987", IsPasswordProtected: true, DisableDownload: false, ViewCount: 14, CreatedAt: now},
			},
			CreatedBy: "USER",
			CreatedAt: now.Add(-7 * 24 * time.Hour),
			UpdatedAt: now,
		},
	}
}

func defaultSampleActivity(vaultID string) []*domain.ShareActivity {
	now := time.Now()
	return []*domain.ShareActivity{
		{
			ActivityID: "act-001",
			ShareID:    "share-japan-001",
			VaultID:    vaultID,
			ActorID:    "usr-alex",
			ActorName:  "Alex Travel",
			ActionType: domain.ActionCommented,
			Details:    map[string]interface{}{"comment": "Added travel insurance receipt"},
			CreatedAt:  now.Add(-2 * time.Hour),
		},
	}
}
