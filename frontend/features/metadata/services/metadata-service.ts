export interface AssetMetadata {
  assetId: string;
  vaultId: string;
  filename: string;
  extension?: string;
  mimeType?: string;
  fileSize: number;
  checksum?: string;
  hash?: string;
  encoding?: string;
  createdDate?: string;
  modifiedDate?: string;
  uploadDate: string;
  lastExtractedAt: string;
}

export interface AssetProperties {
  assetId: string;
  properties: Record<string, any>;
}

export interface AssetExif {
  assetId: string;
  cameraMake?: string;
  cameraModel?: string;
  lens?: string;
  gpsLat?: number;
  gpsLong?: number;
  iso?: number;
  exposure?: string;
  fStop?: number;
  focalLength?: number;
  dateTaken?: string;
  rawData?: Record<string, any>;
}

export interface MetadataVersion {
  versionId: string;
  assetId: string;
  source: string;
  extractorVersion: string;
  confidence?: number;
  snapshot: Record<string, any>;
  createdAt: string;
}

export interface CompleteMetadata {
  general: AssetMetadata;
  properties?: AssetProperties;
  exif?: AssetExif;
}

export const MetadataService = {
  async getMetadata(assetId: string): Promise<CompleteMetadata> {
    // Mock for UI dev
    return {
      general: {
        assetId,
        vaultId: "v_123",
        filename: "IMG_0492.HEIC",
        extension: ".HEIC",
        mimeType: "image/heic",
        fileSize: 4500123,
        hash: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
        uploadDate: new Date().toISOString(),
        lastExtractedAt: new Date().toISOString(),
      },
      properties: {
        assetId,
        properties: {
          width: 4032,
          height: 3024,
          orientation: 1,
          colorSpace: "Display P3",
        }
      },
      exif: {
        assetId,
        cameraMake: "Apple",
        cameraModel: "iPhone 15 Pro",
        lens: "Apple iPhone 15 Pro back camera 6.86mm f/1.78",
        gpsLat: 37.7749,
        gpsLong: -122.4194,
        iso: 80,
        fStop: 1.8,
        exposure: "1/120",
        dateTaken: new Date().toISOString(),
      }
    };
  },

  async getVersions(assetId: string): Promise<MetadataVersion[]> {
    return [
      {
        versionId: "ver_1",
        assetId,
        source: "SYSTEM_EXTRACTOR",
        extractorVersion: "1.0",
        snapshot: { status: "initial_extract" },
        createdAt: new Date().toISOString()
      }
    ];
  },

  async refreshMetadata(assetId: string): Promise<void> {
    console.log("Triggered metadata refresh for", assetId);
  }
};
