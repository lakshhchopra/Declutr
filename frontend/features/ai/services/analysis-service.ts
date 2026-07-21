export interface Classification {
  documentCategory: string;
  documentType: string;
  isScanned: boolean;
  isCorrupted: boolean;
  isIncomplete: boolean;
  qualityScore: number;
  confidenceScore: number;
}

export interface Tag {
  name: string;
  confidenceScore: number;
}

export interface Topic {
  name: string;
  confidenceScore: number;
}

export interface AIAnalysis {
  analysisId: string;
  documentId: string;
  assetId: string;
  title: string;
  shortSummary: string;
  detailedSummary: string;
  language: string;
  writingStyle: string;
  sentiment: string;
  complexity: string;
  readingLevel: string;
  estimatedReadingTime: number;
  documentPurpose: string;
  confidenceScore: number;
  classification: Classification;
  tags: Tag[];
  topics: Topic[];
  createdAt: string;
  updatedAt: string;
}

export const AnalysisService = {
  async getAnalysis(assetId: string): Promise<AIAnalysis> {
    // Mock for UI dev
    return {
      analysisId: "ai_1234",
      documentId: "doc_1234",
      assetId,
      title: "Alpha Project Requirements",
      shortSummary: "A brief overview of the Alpha project constraints.",
      detailedSummary: "This document outlines the Alpha project requirements, including strict security and performance constraints required for phase one delivery.",
      language: "en",
      writingStyle: "Technical",
      sentiment: "Neutral",
      complexity: "Medium",
      readingLevel: "College",
      estimatedReadingTime: 60,
      documentPurpose: "Project Planning",
      confidenceScore: 0.95,
      classification: {
        documentCategory: "General Note",
        documentType: "Markdown Document",
        isScanned: false,
        isCorrupted: false,
        isIncomplete: false,
        qualityScore: 0.99,
        confidenceScore: 0.98,
      },
      tags: [
        { name: "Project", confidenceScore: 0.9 },
        { name: "Planning", confidenceScore: 0.85 },
      ],
      topics: [
        { name: "Software Engineering", confidenceScore: 0.99 },
      ],
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    };
  },

  async refreshAnalysis(assetId: string): Promise<void> {
    console.log("Triggered analysis refresh for", assetId);
  }
};
