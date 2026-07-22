/**
 * Official Declutr TypeScript SDK
 */

export interface DeclutrConfig {
  apiKey?: string;
  accessToken?: string;
  baseUrl?: string;
}

export class DeclutrClient {
  private apiKey?: string;
  private accessToken?: string;
  private baseUrl: string;

  constructor(config: DeclutrConfig) {
    this.apiKey = config.apiKey;
    this.accessToken = config.accessToken;
    this.baseUrl = config.baseUrl || "http://localhost:8080";
  }

  private async request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
    const headers: Record<string, string> = {
      "Content-Type": "application/json",
      ...(options.headers as Record<string, string>),
    };

    if (this.apiKey) {
      headers["Authorization"] = `Bearer ${this.apiKey}`;
    } else if (this.accessToken) {
      headers["Authorization"] = `Bearer ${this.accessToken}`;
    }

    const res = await fetch(`${this.baseUrl}${endpoint}`, {
      ...options,
      headers,
    });

    if (!res.ok) {
      throw new Error(`Declutr API Error [${res.status}]: ${await res.text()}`);
    }

    return res.json();
  }

  // Search API
  public async search(query: string, filters?: Record<string, any>): Promise<any> {
    return this.request("/api/v1/search/query", {
      method: "POST",
      body: JSON.stringify({ query, filters }),
    });
  }

  // AI Copilot API
  public async chat(conversationId: string, message: string): Promise<any> {
    return this.request("/api/v1/copilot/messages", {
      method: "POST",
      body: JSON.stringify({ conversation_id: conversationId, content: message }),
    });
  }

  // Workflows API
  public async runWorkflow(workflowId: string, inputData?: any): Promise<any> {
    return this.request("/api/v1/workflows/run", {
      method: "POST",
      body: JSON.stringify({ workflow_id: workflowId, input: inputData }),
    });
  }

  // Webhooks Management API
  public async registerWebhook(url: string, events: string[]): Promise<any> {
    return this.request("/api/v1/developer/webhooks", {
      method: "POST",
      body: JSON.stringify({ url, events }),
    });
  }
}
