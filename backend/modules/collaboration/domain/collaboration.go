package domain

import "time"

// ResourceType specifies what resource is being shared
type ResourceType string

const (
	ResourceAsset        ResourceType = "ASSET"
	ResourceFolder       ResourceType = "FOLDER"
	ResourceCollection   ResourceType = "COLLECTION"
	ResourceContext      ResourceType = "CONTEXT"
	ResourceProject      ResourceType = "PROJECT"
	ResourceTimelineView ResourceType = "TIMELINE_VIEW"
	ResourceSearchResult ResourceType = "SEARCH_RESULT"
)

// AccessType defines sharing scope
type AccessType string

const (
	AccessPrivate    AccessType = "PRIVATE"
	AccessInviteOnly AccessType = "INVITE_ONLY"
	AccessLink       AccessType = "LINK_SHARING"
	AccessTemporary  AccessType = "TEMPORARY_ACCESS"
)

// MemberRole defines member access tier
type MemberRole string

const (
	RoleReadOnly    MemberRole = "READ_ONLY"
	RoleCommentOnly MemberRole = "COMMENT_ONLY"
	RoleEdit        MemberRole = "EDIT"
	RoleOwner       MemberRole = "OWNER"
	RoleCoOwner     MemberRole = "CO_OWNER"
)

// ShareActionType defines auditable activity actions
type ShareActionType string

const (
	ActionViewed           ShareActionType = "VIEWED"
	ActionDownloaded       ShareActionType = "DOWNLOADED"
	ActionEdited           ShareActionType = "EDITED"
	ActionCommented        ShareActionType = "COMMENTED"
	ActionShared           ShareActionType = "SHARED"
	ActionPermissionChanged ShareActionType = "PERMISSION_CHANGED"
	ActionAccessRevoked    ShareActionType = "ACCESS_REVOKED"
	ActionInviteAccepted   ShareActionType = "INVITE_ACCEPTED"
)

// InviteStatus represents status of an invitation
type InviteStatus string

const (
	InvitePending  InviteStatus = "PENDING"
	InviteAccepted InviteStatus = "ACCEPTED"
	InviteRejected InviteStatus = "REJECTED"
	InviteRevoked  InviteStatus = "REVOKED"
)

// SharePermission model
type SharePermission struct {
	PermissionID     string     `json:"permissionId"`
	ShareID          string     `json:"shareId"`
	Role             MemberRole `json:"role"`
	CanView          bool       `json:"canView"`
	CanDownload      bool       `json:"canDownload"`
	CanEdit          bool       `json:"canEdit"`
	CanDelete        bool       `json:"canDelete"`
	CanComment       bool       `json:"canComment"`
	CanShare         bool       `json:"canShare"`
	CanManageMembers bool       `json:"canManageMembers"`
	CreatedAt        time.Time  `json:"createdAt"`
}

// ShareMember model
type ShareMember struct {
	MemberID  string     `json:"memberId"`
	ShareID   string     `json:"shareId"`
	UserID    string     `json:"userId"`
	Email     string     `json:"email"`
	Role      MemberRole `json:"role"`
	JoinedAt  time.Time  `json:"joinedAt"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

// ShareLink model
type ShareLink struct {
	LinkID              string     `json:"linkId"`
	ShareID             string     `json:"shareId"`
	LinkToken           string     `json:"linkToken"`
	IsPasswordProtected bool       `json:"isPasswordProtected"`
	PasswordHash        string     `json:"-"`
	DisableDownload     bool       `json:"disableDownload"`
	DisableReshare      bool       `json:"disableReshare"`
	ViewCount           int        `json:"viewCount"`
	MaxViews            int        `json:"maxViews"`
	DownloadCount       int        `json:"downloadCount"`
	MaxDownloads        int        `json:"maxDownloads"`
	ExpiresAt           *time.Time `json:"expiresAt,omitempty"`
	CreatedAt           time.Time  `json:"createdAt"`
}

// ShareComment model
type ShareComment struct {
	CommentID       string    `json:"commentId"`
	ShareID         string    `json:"shareId"`
	UserID          string    `json:"userId"`
	UserName        string    `json:"userName"`
	Content         string    `json:"content"`
	ParentCommentID string    `json:"parentCommentId,omitempty"`
	IsResolved      bool      `json:"isResolved"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// ShareActivity audit log model
type ShareActivity struct {
	ActivityID string                 `json:"activityId"`
	ShareID    string                 `json:"shareId"`
	VaultID    string                 `json:"vaultId"`
	ActorID    string                 `json:"actorId"`
	ActorName  string                 `json:"actorName"`
	ActionType ShareActionType        `json:"actionType"`
	Details    map[string]interface{} `json:"details,omitempty"`
	CreatedAt  time.Time              `json:"createdAt"`
}

// ShareInvitation model
type ShareInvitation struct {
	InvitationID string       `json:"invitationId"`
	ShareID      string       `json:"shareId"`
	InviterID    string       `json:"inviterId"`
	InviteeEmail string       `json:"inviteeEmail"`
	Role         MemberRole   `json:"role"`
	Status       InviteStatus `json:"status"`
	Token        string       `json:"token"`
	CreatedAt    time.Time    `json:"createdAt"`
	ExpiresAt    *time.Time   `json:"expiresAt,omitempty"`
}

// Share model
type Share struct {
	ShareID      string            `json:"shareId"`
	VaultID      string            `json:"vaultId"`
	ResourceType ResourceType      `json:"resourceType"`
	ResourceID   string            `json:"resourceId"`
	Title        string            `json:"title"`
	AccessType   AccessType        `json:"accessType"`
	Members      []ShareMember     `json:"members,omitempty"`
	Links        []ShareLink       `json:"links,omitempty"`
	Permissions  []SharePermission `json:"permissions,omitempty"`
	CreatedBy    string            `json:"createdBy"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
}

// ShareStats metrics
type ShareStats struct {
	VaultID           string  `json:"vaultId"`
	TotalShares       int     `json:"totalShares"`
	ActiveLinks       int     `json:"activeLinks"`
	TotalMembers      int     `json:"totalMembers"`
	TotalComments     int     `json:"totalComments"`
	AuditActivityCount int    `json:"auditActivityCount"`
}

// CreateShareRequest payload
type CreateShareRequest struct {
	VaultID      string       `json:"vaultId"`
	ResourceType ResourceType `json:"resourceType"`
	ResourceID   string       `json:"resourceId"`
	Title        string       `json:"title"`
	AccessType   AccessType   `json:"accessType"`
}

// InviteRequest payload
type InviteRequest struct {
	ShareID      string     `json:"shareId"`
	InviterID    string     `json:"inviterId"`
	InviteeEmail string     `json:"inviteeEmail"`
	Role         MemberRole `json:"role"`
}

// CreateLinkRequest payload
type CreateLinkRequest struct {
	ShareID             string     `json:"shareId"`
	IsPasswordProtected bool       `json:"isPasswordProtected"`
	Password            string     `json:"password,omitempty"`
	DisableDownload     bool       `json:"disableDownload"`
	ExpiresInDays       int        `json:"expiresInDays,omitempty"`
}

// AddCommentRequest payload
type AddCommentRequest struct {
	ShareID         string `json:"shareId"`
	UserID          string `json:"userId"`
	UserName        string `json:"userName"`
	Content         string `json:"content"`
	ParentCommentID string `json:"parentCommentId,omitempty"`
}
