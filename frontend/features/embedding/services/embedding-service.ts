import type {
  EmbeddingStats,
  EmbeddingStatusResponse,
  EmbeddingHistoryResponse,
  EmbeddingProviderConfig,
  StructuredRepresentationInput,
  Embedding,
  ProviderName,
} from '../types/embedding';

const BASE_URL = process.env.NEXT_PUBLIC_API_URL ?? 'http://localhost:8080/api/v1';

async function apiFetch<T>(url: string, options?: RequestInit): Promise<T> {
  const res = await fetch(url, {
    ...options,
    headers: { 'Content-Type': 'application/json', ...options?.headers },
  });
  if (!res.ok) throw new Error(`Embedding API error: ${res.status} ${res.statusText}`);
  return res.json();
}

const VAULT_ID = 'vault-demo';

// ─── Mock Data Fallback ────────────────────────────────────────────────────────

const MOCK_STATS: EmbeddingStats = {
  vaultId: VAULT_ID,
  totalEmbeddings: 142,
  totalChunks: 588,
  activeProvider: 'OPENAI',
  activeModel: 'text-embedding-3-small',
  dimensions: 1536,
  sourceTypeBreakdown: {
    DOCUMENT: 48,
    MEMORY: 32,
    CONTEXT: 24,
    ENTITY: 18,
    SUMMARY: 12,
    RELATIONSHIP: 8,
  },
  strategyBreakdown: {
    SEMANTIC: 240,
    HEADING_AWARE: 180,
    DOCUMENT_AWARE: 90,
    PAGE_AWARE: 48,
    HIERARCHICAL: 30,
  },
  latestVersionTag: 'v1.2.0',
};

const MOCK_STATUS: EmbeddingStatusResponse = {
  vaultId: VAULT_ID,
  activeProvider: 'OPENAI',
  activeModel: 'text-embedding-3-small',
  dimensions: 1536,
  activeJobs: 0,
  status: 'HEALTHY',
};

const MOCK_HISTORY: EmbeddingHistoryResponse = {
  vaultId: VAULT_ID,
  jobs: [
    { jobId: 'j-101', vaultId: VAULT_ID, targetType: 'DOCUMENT', targetId: 'doc-thesis', status: 'COMPLETED', processedChunks: 14, createdAt: new Date(Date.now() - 3600000).toISOString(), completedAt: new Date(Date.now() - 3500000).toISOString() },
    { jobId: 'j-102', vaultId: VAULT_ID, targetType: 'MEMORY', targetId: 'mem-japan', status: 'COMPLETED', processedChunks: 6, createdAt: new Date(Date.now() - 86400000).toISOString(), completedAt: new Date(Date.now() - 86300000).toISOString() },
    { jobId: 'j-103', vaultId: VAULT_ID, targetType: 'CONTEXT', targetId: 'ctx-medical', status: 'COMPLETED', processedChunks: 8, createdAt: new Date(Date.now() - 172800000).toISOString(), completedAt: new Date(Date.now() - 172700000).toISOString() },
  ],
  latestVersion: {
    versionId: 'v-12',
    vaultId: VAULT_ID,
    providerName: 'OPENAI',
    modelName: 'text-embedding-3-small',
    dimensions: 1536,
    versionTag: 'v1.2.0',
    isActive: true,
    totalEmbeddedItems: 142,
    upgradedAt: new Date(Date.now() - 3 * 86400000).toISOString(),
  },
};

export const EmbeddingService = {
  async getStats(vaultId: string = VAULT_ID): Promise<EmbeddingStats> {
    try {
      return await apiFetch<EmbeddingStats>(`${BASE_URL}/embedding/stats?vaultId=${vaultId}`);
    } catch {
      return MOCK_STATS;
    }
  },

  async getStatus(vaultId: string = VAULT_ID): Promise<EmbeddingStatusResponse> {
    try {
      return await apiFetch<EmbeddingStatusResponse>(`${BASE_URL}/embedding/status?vaultId=${vaultId}`);
    } catch {
      return MOCK_STATUS;
    }
  },

  async getHistory(vaultId: string = VAULT_ID): Promise<EmbeddingHistoryResponse> {
    try {
      return await apiFetch<EmbeddingHistoryResponse>(`${BASE_URL}/embedding/history?vaultId=${vaultId}`);
    } catch {
      return MOCK_HISTORY;
    }
  },

  async generateEmbedding(input: StructuredRepresentationInput): Promise<Embedding> {
    try {
      return await apiFetch<Embedding>(`${BASE_URL}/embedding/generate`, {
        method: 'POST',
        body: JSON.stringify({ input }),
      });
    } catch {
      return {
        embeddingId: 'emb-mock-1',
        vaultId: input.vaultId,
        sourceType: input.sourceType,
        sourceId: input.sourceId,
        providerName: 'OPENAI',
        modelName: 'text-embedding-3-small',
        modelVersion: 'v1',
        dimensions: 1536,
        representationText: input.title + '\n' + input.summary,
        contentHash: 'hash-mock',
        vectorData: new Array(1536).fill(0.01),
        isActive: true,
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString(),
      };
    }
  },

  async refreshEmbeddings(vaultId: string = VAULT_ID): Promise<void> {
    try {
      await apiFetch(`${BASE_URL}/embedding/refresh`, {
        method: 'POST',
        body: JSON.stringify({ vaultId }),
      });
    } catch { /* mock */ }
  },

  async updateProvider(vaultId: string, providerName: ProviderName, modelName: string, dimensions: number): Promise<void> {
    try {
      await apiFetch(`${BASE_URL}/embedding/provider`, {
        method: 'PUT',
        body: JSON.stringify({ vaultId, providerName, modelName, dimensions, batchSize: 32, isDefault: true }),
      });
    } catch { /* mock */ }
  },

  async rebuildVersion(vaultId: string, providerName: ProviderName, modelName: string, versionTag: string): Promise<void> {
    try {
      await apiFetch(`${BASE_URL}/embedding/rebuild`, {
        method: 'POST',
        body: JSON.stringify({ vaultId, providerName, modelName, versionTag }),
      });
    } catch { /* mock */ }
  },
};
