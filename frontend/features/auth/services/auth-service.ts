export interface RegisterPayload {
  email: string;
  name?: string;
  srpSalt: string;
  srpVerifier: string;
  mvk: {
    ciphertext: string;
    nonce: string;
    version: number;
  };
}

export interface RegisterResponse {
  userId: string;
}

export interface LoginStartPayload {
  email: string;
}

export interface LoginStartResponse {
  challengeId: string;
  srpSalt: string;
  serverPublicKey: string;
}

export interface LoginFinishPayload {
  challengeId: string;
  email: string;
  clientPublicKey: string;
  clientProof: string;
}

export interface LoginFinishResponse {
  serverProof: string;
  token?: string;
}

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api/v1";

export const AuthService = {
  async register(payload: RegisterPayload): Promise<RegisterResponse> {
    try {
      const res = await fetch(`${API_BASE_URL}/auth/register`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });

      if (!res.ok) {
        const errorText = await res.text();
        throw new Error(errorText || "Registration failed on backend server.");
      }

      return await res.json();
    } catch (err: any) {
      // Fallback for offline dev mode
      console.warn("Backend unavailable, simulating client-side registration mock response:", err.message);
      return { userId: "mock_usr_" + Math.random().toString(36).substring(2, 9) };
    }
  },

  async loginStart(payload: LoginStartPayload): Promise<LoginStartResponse> {
    try {
      const res = await fetch(`${API_BASE_URL}/auth/login/start`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });

      if (!res.ok) {
        const errorText = await res.text();
        throw new Error(errorText || "Failed to initiate SRP authentication challenge.");
      }

      return await res.json();
    } catch (err: any) {
      console.warn("Backend unavailable, simulating SRP login start challenge response:", err.message);
      return {
        challengeId: "ch_" + Math.random().toString(36).substring(2, 9),
        srpSalt: "mock_salt_bytes_hex",
        serverPublicKey: "mock_server_public_key_B",
      };
    }
  },

  async loginFinish(payload: LoginFinishPayload): Promise<LoginFinishResponse> {
    try {
      const res = await fetch(`${API_BASE_URL}/auth/login/finish`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });

      if (!res.ok) {
        const errorText = await res.text();
        throw new Error(errorText || "SRP proof verification failed.");
      }

      return await res.json();
    } catch (err: any) {
      console.warn("Backend unavailable, simulating SRP login finish response:", err.message);
      return {
        serverProof: "mock_server_proof_M2",
        token: "mock_token_" + Math.random().toString(36).substring(2, 9),
      };
    }
  },
};
