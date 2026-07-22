// Embedding Engine TypeScript types

export type SourceType =
  | 'DOCUMENT'
  | 'SUMMARY'
  | 'ENTITY'
  | 'CONTEXT'
  | 'RELATIONSHIP'
  | 'MEMORY'
  | 'COLLECTION'
  | 'NOTE'
  | 'CHAT';

export type ChunkStrategy =
  | 'SEMANTIC'
  | 'HIERARCHICAL'
  | 'DOCUMENT_AWARE'
  | 'PAGE_AWARE'
  | 'HEADING_AWARE';

export type ProviderName =
  | 'OPENAI'
  | 'GEMINI'
  | 'VOYAGE'
  | 'COHERE'
  | 'OLLAMA'
  | 'LOCAL';

export interface StructuredRepresentationInput {
  sourceType: SourceType;
  sourceId: string;
  vaultId: string;
  title: string;
  summary: string;
  content?: string;
  entities?: string[];
  relationships?: string[];
  contexts?: string[];
  intent?: string;
  memoryScore?: number;
  tags?: string[];
  classification?: string;
  metadata?: Record<string, string>;
}

export interface Embedding {
  embeddingId: string;
  vaultId: string;
  sourceType: SourceType;
  sourceId: string;
  providerName: string;
  modelName: string;
  modelVersion: string;
  dimensions: number;
  representationText: string;
  contentHash: string;
  vectorData: number[];
  isActive: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface EmbeddingChunk {
  chunkId: string;
  embeddingId: string;
  vaultId: string;
  chunkIndex: number;
  chunkStrategy: ChunkStrategy;
  chunkText: string;
  tokenCount: number;
  headingPath?: string;
  pageNumber?: number;
  vectorData: number[];
  createdAt: string;
}

export interface EmbeddingVersion {
  versionId: string;
  vaultId: string;
  providerName: string;
  modelName: string;
  dimensions: number;
  versionTag: string;
  isActive: boolean;
  totalEmbeddedItems: number;
  upgradedAt: string;
}

export interface EmbeddingJob {
  jobId: string;
  vaultId: string;
  targetType: SourceType;
  targetId: string;
  status: 'QUEUED' | 'PROCESSING' | 'COMPLETED' | 'FAILED';
  errorMessage?: string;
  processedChunks: number;
  createdAt: string;
  completedAt?: string;
}

export interface EmbeddingProviderConfig {
  providerId: string;
  vaultId: string;
  providerName: ProviderName;
  modelName: string;
  dimensions: number;
  batchSize: number;
  isDefault: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface EmbeddingStats {
  vaultId: string;
  totalEmbeddings: number;
  totalChunks: number;
  activeProvider: string;
  activeModel: string;
  dimensions: number;
  sourceTypeBreakdown: Record<string, number>;
  strategyBreakdown: Record<string, number>;
  latestVersionTag: string;
}

export interface EmbeddingStatusResponse {
  vaultId: string;
  activeProvider: ProviderName;
  activeModel: string;
  dimensions: number;
  activeJobs: number;
  status: 'HEALTHY' | 'DEGRADED' | 'PROCESSING';
}

export interface EmbeddingHistoryResponse {
  vaultId: string;
  jobs: EmbeddingJob[];
  latestVersion: EmbeddingVersion;
}
