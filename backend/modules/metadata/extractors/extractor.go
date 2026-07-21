package extractors

import (
	"context"
	"io"
	"strings"
	"time"
	"crypto/sha256"
	"encoding/hex"

	"github.com/diablovocado/declutr/modules/metadata/domain"
)

type MetadataExtractor interface {
	Supports(mimeType string) bool
	Extract(ctx context.Context, assetID, vaultID, filename, mimeType string, size int64, reader io.Reader) (*domain.CompleteMetadata, error)
}

// Registry to hold and route to extractors
type ExtractorRegistry struct {
	extractors []MetadataExtractor
}

func NewExtractorRegistry() *ExtractorRegistry {
	return &ExtractorRegistry{
		extractors: []MetadataExtractor{
			&ImageExtractor{},
			&TextExtractor{},
			&MockComplexExtractor{}, // Catch-all for PDF, Video, Audio
		},
	}
}

func (r *ExtractorRegistry) GetExtractor(mimeType string) MetadataExtractor {
	for _, ext := range r.extractors {
		if ext.Supports(mimeType) {
			return ext
		}
	}
	return &BaseExtractor{} // Fallback
}

// BaseExtractor extracts just the common file stats
type BaseExtractor struct{}

func (e *BaseExtractor) Supports(mimeType string) bool {
	return true
}

func (e *BaseExtractor) Extract(ctx context.Context, assetID, vaultID, filename, mimeType string, size int64, reader io.Reader) (*domain.CompleteMetadata, error) {
	// Simple stream hash
	hasher := sha256.New()
	if reader != nil {
		io.Copy(hasher, reader)
	}
	hash := hex.EncodeToString(hasher.Sum(nil))

	now := time.Now()
	
	general := &domain.AssetMetadata{
		AssetID:         assetID,
		VaultID:         vaultID,
		Filename:        filename,
		Extension:       extractExtension(filename),
		MimeType:        mimeType,
		FileSize:        size,
		Hash:            hash,
		UploadDate:      now,
		LastExtractedAt: now,
	}

	return &domain.CompleteMetadata{General: general}, nil
}

// ImageExtractor (Basic dimensions stub)
type ImageExtractor struct {
	BaseExtractor
}

func (e *ImageExtractor) Supports(mimeType string) bool {
	return strings.HasPrefix(mimeType, "image/")
}

func (e *ImageExtractor) Extract(ctx context.Context, assetID, vaultID, filename, mimeType string, size int64, reader io.Reader) (*domain.CompleteMetadata, error) {
	meta, err := e.BaseExtractor.Extract(ctx, assetID, vaultID, filename, mimeType, size, reader)
	if err != nil {
		return nil, err
	}
	
	// Mock decoding dimensions
	meta.Properties = &domain.AssetProperties{
		AssetID: assetID,
		Properties: map[string]interface{}{
			"width":  1920,
			"height": 1080,
			"orientation": 1,
			"colorSpace": "sRGB",
		},
	}

	// Mock EXIF
	meta.Exif = &domain.AssetExif{
		AssetID: assetID,
		CameraMake: "Apple",
		CameraModel: "iPhone 15 Pro",
		ISO: func(i int) *int { return &i }(100),
	}

	return meta, nil
}

// TextExtractor (Basic char count stub)
type TextExtractor struct {
	BaseExtractor
}

func (e *TextExtractor) Supports(mimeType string) bool {
	return strings.HasPrefix(mimeType, "text/") || mimeType == "application/json"
}

func (e *TextExtractor) Extract(ctx context.Context, assetID, vaultID, filename, mimeType string, size int64, reader io.Reader) (*domain.CompleteMetadata, error) {
	meta, err := e.BaseExtractor.Extract(ctx, assetID, vaultID, filename, mimeType, size, reader)
	if err != nil {
		return nil, err
	}
	
	meta.Properties = &domain.AssetProperties{
		AssetID: assetID,
		Properties: map[string]interface{}{
			"encoding": "UTF-8",
			"characterCount": size,
			"lineCount": size / 80, // rough guess for stub
		},
	}

	return meta, nil
}

// MockComplexExtractor (For PDFs, Videos, Audio)
type MockComplexExtractor struct {
	BaseExtractor
}

func (e *MockComplexExtractor) Supports(mimeType string) bool {
	return strings.HasPrefix(mimeType, "video/") || strings.HasPrefix(mimeType, "audio/") || mimeType == "application/pdf"
}

func (e *MockComplexExtractor) Extract(ctx context.Context, assetID, vaultID, filename, mimeType string, size int64, reader io.Reader) (*domain.CompleteMetadata, error) {
	meta, err := e.BaseExtractor.Extract(ctx, assetID, vaultID, filename, mimeType, size, reader)
	if err != nil {
		return nil, err
	}
	
	props := map[string]interface{}{}
	
	if mimeType == "application/pdf" {
		props["pages"] = 12
		props["author"] = "Unknown Author"
		props["producer"] = "Ghostscript"
	} else if strings.HasPrefix(mimeType, "video/") {
		props["durationSeconds"] = 125
		props["resolution"] = "1920x1080"
		props["codec"] = "h264"
	} else if strings.HasPrefix(mimeType, "audio/") {
		props["durationSeconds"] = 180
		props["sampleRate"] = 44100
		props["channels"] = 2
	}

	meta.Properties = &domain.AssetProperties{
		AssetID: assetID,
		Properties: props,
	}

	return meta, nil
}


func extractExtension(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) > 1 {
		return "." + parts[len(parts)-1]
	}
	return ""
}
