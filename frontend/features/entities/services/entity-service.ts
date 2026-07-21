export type EntityType = 'Person' | 'Organization' | 'Location' | 'Date' | 'Amount' | 'Product' | 'Identifier';

export interface EntityOccurrence {
  occurrenceId: string;
  entityId: string;
  assetId: string;
  analysisId: string;
  originalValue: string;
  confidenceScore: number;
  extractorVersion: string;
  createdAt: string;
}

export interface Entity {
  entityId: string;
  vaultId: string;
  type: EntityType;
  canonicalName: string;
  normalizedValue: string;
  description: string;
  aliases: string[];
  createdAt: string;
  updatedAt: string;
}

export interface AssetEntityResponse {
  entity: Entity;
  occurrence: EntityOccurrence;
}

export const EntityService = {
  async getEntitiesForAsset(assetId: string): Promise<AssetEntityResponse[]> {
    // Mock for UI dev
    return [
      {
        entity: {
          entityId: "ent_org_1",
          vaultId: "v_123",
          type: "Organization",
          canonicalName: "Google",
          normalizedValue: "Google",
          description: "A multinational technology company.",
          aliases: ["Google LLC", "Alphabet"],
          createdAt: new Date().toISOString(),
          updatedAt: new Date().toISOString(),
        },
        occurrence: {
          occurrenceId: "occ_1",
          entityId: "ent_org_1",
          assetId,
          analysisId: "ai_1",
          originalValue: "Google LLC",
          confidenceScore: 0.99,
          extractorVersion: "mock-v1",
          createdAt: new Date().toISOString(),
        }
      },
      {
        entity: {
          entityId: "ent_loc_1",
          vaultId: "v_123",
          type: "Location",
          canonicalName: "New York City",
          normalizedValue: "New York City",
          description: "City in New York State.",
          aliases: ["NYC", "New York"],
          createdAt: new Date().toISOString(),
          updatedAt: new Date().toISOString(),
        },
        occurrence: {
          occurrenceId: "occ_2",
          entityId: "ent_loc_1",
          assetId,
          analysisId: "ai_1",
          originalValue: "NYC",
          confidenceScore: 0.95,
          extractorVersion: "mock-v1",
          createdAt: new Date().toISOString(),
        }
      },
      {
        entity: {
          entityId: "ent_amt_1",
          vaultId: "v_123",
          type: "Amount",
          canonicalName: "1500.50 USD",
          normalizedValue: "1500.50 USD",
          description: "Monetary value in USD.",
          aliases: [],
          createdAt: new Date().toISOString(),
          updatedAt: new Date().toISOString(),
        },
        occurrence: {
          occurrenceId: "occ_3",
          entityId: "ent_amt_1",
          assetId,
          analysisId: "ai_1",
          originalValue: "$1,500.50",
          confidenceScore: 0.99,
          extractorVersion: "mock-v1",
          createdAt: new Date().toISOString(),
        }
      }
    ];
  },
};
