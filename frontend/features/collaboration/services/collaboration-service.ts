import type {
  Share,
  ShareComment,
  ShareActivity,
  ShareStats,
  ShareLink,
  MemberRole,
} from '../types/collaboration';

const BASE_URL = process.env.NEXT_PUBLIC_API_URL ?? 'http://localhost:8080/api/v1';

async function apiFetch<T>(url: string, options?: RequestInit): Promise<T> {
  const res = await fetch(url, {
    ...options,
    headers: { 'Content-Type': 'application/json', ...options?.headers },
  });
  if (!res.ok) throw new Error(`Collaboration API error: ${res.status} ${res.statusText}`);
  return res.json();
}

const VAULT_ID = 'vault-demo';

// ─── Mock Data Fallback ────────────────────────────────────────────────────────

const MOCK_SHARES: Share[] = [
  {
    shareId: 'share-japan-001',
    vaultId: VAULT_ID,
    resourceType: 'COLLECTION',
    resourceId: 'col-japan-vacation',
    title: 'Japan Trip Photos & Itinerary',
    accessType: 'INVITE_ONLY',
    members: [
      { memberId: 'm1', shareId: 'share-japan-001', userId: 'usr-owner', email: 'owner@declutr.local', role: 'OWNER', joinedAt: new Date(Date.now() - 7 * 86400000).toISOString() },
      { memberId: 'm2', shareId: 'share-japan-001', userId: 'usr-alex', email: 'alex@travel.org', role: 'EDIT', joinedAt: new Date(Date.now() - 3 * 86400000).toISOString() },
    ],
    links: [
      { linkId: 'l1', shareId: 'share-japan-001', linkToken: 'tok-japan-public-987', isPasswordProtected: true, disableDownload: false, disableReshare: true, viewCount: 14, maxViews: 0, downloadCount: 2, maxDownloads: 0, createdAt: new Date().toISOString() },
    ],
    createdBy: 'USER',
    createdAt: new Date(Date.now() - 7 * 86400000).toISOString(),
    updatedAt: new Date().toISOString(),
  },
];

const MOCK_COMMENTS: ShareComment[] = [
  {
    commentId: 'cmnt-1',
    shareId: 'share-japan-001',
    userId: 'usr-alex',
    userName: 'Alex Travel',
    content: 'Uploaded the Tokyo Metro map PDF to the collection!',
    isResolved: false,
    createdAt: new Date(Date.now() - 2 * 3600000).toISOString(),
    updatedAt: new Date(Date.now() - 2 * 3600000).toISOString(),
  },
];

const MOCK_ACTIVITIES: ShareActivity[] = [
  {
    activityId: 'act-1',
    shareId: 'share-japan-001',
    vaultId: VAULT_ID,
    actorId: 'usr-alex',
    actorName: 'Alex Travel',
    actionType: 'COMMENTED',
    details: { content: 'Uploaded Tokyo Metro map PDF' },
    createdAt: new Date(Date.now() - 2 * 3600000).toISOString(),
  },
  {
    activityId: 'act-2',
    shareId: 'share-japan-001',
    vaultId: VAULT_ID,
    actorId: 'usr-owner',
    actorName: 'Vault Owner',
    actionType: 'SHARED',
    details: { title: 'Japan Trip Photos & Itinerary' },
    createdAt: new Date(Date.now() - 7 * 86400000).toISOString(),
  },
];

export const CollaborationService = {
  async getShares(vaultId: string = VAULT_ID): Promise<Share[]> {
    try {
      const res = await apiFetch<{ shares: Share[] }>(`${BASE_URL}/shares?vaultId=${vaultId}`);
      return res.shares ?? [];
    } catch {
      return MOCK_SHARES;
    }
  },

  async createShare(title: string, resourceType: Share['resourceType'], resourceId: string, accessType: Share['accessType'], vaultId: string = VAULT_ID): Promise<Share> {
    try {
      return await apiFetch<Share>(`${BASE_URL}/shares`, {
        method: 'POST',
        body: JSON.stringify({ vaultId, title, resourceType, resourceId, accessType }),
      });
    } catch {
      return {
        shareId: `share-${Date.now()}`,
        vaultId,
        resourceType,
        resourceId,
        title,
        accessType,
        createdBy: 'USER',
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString(),
      };
    }
  },

  async deleteShare(shareId: string): Promise<void> {
    try {
      await apiFetch(`${BASE_URL}/shares?shareId=${shareId}`, { method: 'DELETE' });
    } catch { /* mock */ }
  },

  async inviteUser(shareId: string, inviteeEmail: string, role: MemberRole): Promise<void> {
    try {
      await apiFetch(`${BASE_URL}/shares/invite`, {
        method: 'POST',
        body: JSON.stringify({ shareId, inviterId: 'usr-owner', inviteeEmail, role }),
      });
    } catch { /* mock */ }
  },

  async createLink(shareId: string, isPasswordProtected: boolean, disableDownload: boolean): Promise<ShareLink> {
    try {
      return await apiFetch<ShareLink>(`${BASE_URL}/shares/links`, {
        method: 'POST',
        body: JSON.stringify({ shareId, isPasswordProtected, disableDownload }),
      });
    } catch {
      return {
        linkId: `link-${Date.now()}`,
        shareId,
        linkToken: `tok-${Date.now()}`,
        isPasswordProtected,
        disableDownload,
        disableReshare: true,
        viewCount: 0,
        maxViews: 0,
        downloadCount: 0,
        maxDownloads: 0,
        createdAt: new Date().toISOString(),
      };
    }
  },

  async addComment(shareId: string, content: string): Promise<ShareComment> {
    try {
      return await apiFetch<ShareComment>(`${BASE_URL}/shares/comments`, {
        method: 'POST',
        body: JSON.stringify({ shareId, userId: 'usr-owner', userName: 'Vault Owner', content }),
      });
    } catch {
      return {
        commentId: `cmnt-${Date.now()}`,
        shareId,
        userId: 'usr-owner',
        userName: 'Vault Owner',
        content,
        isResolved: false,
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString(),
      };
    }
  },

  async getComments(shareId: string): Promise<ShareComment[]> {
    try {
      const res = await apiFetch<{ comments: ShareComment[] }>(`${BASE_URL}/shares/comments?shareId=${shareId}`);
      return res.comments ?? [];
    } catch {
      return MOCK_COMMENTS;
    }
  },

  async getActivity(vaultId: string = VAULT_ID): Promise<ShareActivity[]> {
    try {
      const res = await apiFetch<{ activity: ShareActivity[] }>(`${BASE_URL}/shares/activity?vaultId=${vaultId}`);
      return res.activity ?? [];
    } catch {
      return MOCK_ACTIVITIES;
    }
  },

  async getStats(vaultId: string = VAULT_ID): Promise<ShareStats> {
    try {
      return await apiFetch<ShareStats>(`${BASE_URL}/shares/stats?vaultId=${vaultId}`);
    } catch {
      return {
        vaultId,
        totalShares: 1,
        activeLinks: 1,
        totalMembers: 2,
        totalComments: 1,
        auditActivityCount: 2,
      };
    }
  },
};
