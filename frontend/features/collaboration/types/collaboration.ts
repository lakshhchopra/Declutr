// Secure Sharing & Collaboration Platform TypeScript types

export type ResourceType =
  | 'ASSET'
  | 'FOLDER'
  | 'COLLECTION'
  | 'CONTEXT'
  | 'PROJECT'
  | 'TIMELINE_VIEW'
  | 'SEARCH_RESULT';

export type AccessType = 'PRIVATE' | 'INVITE_ONLY' | 'LINK_SHARING' | 'TEMPORARY_ACCESS';
export type MemberRole = 'READ_ONLY' | 'COMMENT_ONLY' | 'EDIT' | 'OWNER' | 'CO_OWNER';

export type ShareActionType =
  | 'VIEWED'
  | 'DOWNLOADED'
  | 'EDITED'
  | 'COMMENTED'
  | 'SHARED'
  | 'PERMISSION_CHANGED'
  | 'ACCESS_REVOKED'
  | 'INVITE_ACCEPTED';

export type InviteStatus = 'PENDING' | 'ACCEPTED' | 'REJECTED' | 'REVOKED';

export interface SharePermission {
  permissionId: string;
  shareId: string;
  role: MemberRole;
  canView: boolean;
  canDownload: boolean;
  canEdit: boolean;
  canDelete: boolean;
  canComment: boolean;
  canShare: boolean;
  canManageMembers: boolean;
  createdAt: string;
}

export interface ShareMember {
  memberId: string;
  shareId: string;
  userId: string;
  email: string;
  role: MemberRole;
  joinedAt: string;
  expiresAt?: string;
}

export interface ShareLink {
  linkId: string;
  shareId: string;
  linkToken: string;
  isPasswordProtected: boolean;
  disableDownload: boolean;
  disableReshare: boolean;
  viewCount: number;
  maxViews: number;
  downloadCount: number;
  maxDownloads: number;
  expiresAt?: string;
  createdAt: string;
}

export interface ShareComment {
  commentId: string;
  shareId: string;
  userId: string;
  userName: string;
  content: string;
  parentCommentId?: string;
  isResolved: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface ShareActivity {
  activityId: string;
  shareId: string;
  vaultId: string;
  actorId: string;
  actorName: string;
  actionType: ShareActionType;
  details?: Record<string, unknown>;
  createdAt: string;
}

export interface ShareInvitation {
  invitationId: string;
  shareId: string;
  inviterId: string;
  inviteeEmail: string;
  role: MemberRole;
  status: InviteStatus;
  token: string;
  createdAt: string;
  expiresAt?: string;
}

export interface Share {
  shareId: string;
  vaultId: string;
  resourceType: ResourceType;
  resourceId: string;
  title: string;
  accessType: AccessType;
  members?: ShareMember[];
  links?: ShareLink[];
  permissions?: SharePermission[];
  createdBy: string;
  createdAt: string;
  updatedAt: string;
}

export interface ShareStats {
  vaultId: string;
  totalShares: number;
  activeLinks: number;
  totalMembers: number;
  totalComments: number;
  auditActivityCount: number;
}
