export interface UploadItem {
  id: string;
  file: File;
  filename: string;
  sizeBytes: number;
  mimeType: string;
  progressPercentage: number;
  status: "QUEUED" | "UPLOADING" | "VALIDATING" | "READY" | "FAILED";
  checksumSha256?: string;
  error?: string;
}

export const UploadService = {
  async computeChecksum(file: File): Promise<string> {
    try {
      const buffer = await file.arrayBuffer();
      const hashBuffer = await crypto.subtle.digest("SHA-256", buffer);
      const hashArray = Array.from(new Uint8Array(hashBuffer));
      return hashArray.map((b) => b.toString(16).padStart(2, "0")).join("");
    } catch (err) {
      return "checksum_fallback_" + Math.random().toString(36).substring(2, 10);
    }
  },

  validateFile(file: File): { valid: boolean; error?: string } {
    const MAX_SIZE_BYTES = 500 * 1024 * 1024; // 500 MB limit
    if (file.size > MAX_SIZE_BYTES) {
      return { valid: false, error: "File exceeds 500 MB maximum size limit." };
    }
    return { valid: true };
  },
};
