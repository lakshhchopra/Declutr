package chunking

import (
	"regexp"
	"strings"

	"github.com/diablovocado/declutr/modules/embedding/domain"
)

// Chunker defines the contract for document chunking algorithms
type Chunker interface {
	Chunk(text string) ([]domain.ChunkResult, error)
}

// GetChunker instantiates a chunker for the requested strategy
func GetChunker(strategy domain.ChunkStrategy) Chunker {
	switch strategy {
	case domain.StrategyHeading:
		return &HeadingAwareChunker{}
	case domain.StrategyPage:
		return &PageAwareChunker{}
	case domain.StrategyHierarchical:
		return &HierarchicalChunker{}
	case domain.StrategyDocument:
		return &DocumentAwareChunker{}
	case domain.StrategySemantic:
		fallthrough
	default:
		return &SemanticChunker{TargetChunkSize: 500}
	}
}

// ============================================================
// 1. Semantic Chunker — splits on paragraph / sentence boundaries
// ============================================================

type SemanticChunker struct {
	TargetChunkSize int // approx token target (words)
}

func (c *SemanticChunker) Chunk(text string) ([]domain.ChunkResult, error) {
	if strings.TrimSpace(text) == "" {
		return nil, nil
	}

	paragraphs := strings.Split(text, "\n\n")
	var chunks []domain.ChunkResult
	var currentChunk strings.Builder
	wordCount := 0
	chunkIdx := 0

	for _, para := range paragraphs {
		para = strings.TrimSpace(para)
		if para == "" {
			continue
		}

		words := len(strings.Fields(para))
		if wordCount+words > c.TargetChunkSize && currentChunk.Len() > 0 {
			chunks = append(chunks, domain.ChunkResult{
				Text:       currentChunk.String(),
				Index:      chunkIdx,
				Strategy:   domain.StrategySemantic,
				TokenCount: wordCount,
			})
			chunkIdx++
			currentChunk.Reset()
			wordCount = 0
		}

		if currentChunk.Len() > 0 {
			currentChunk.WriteString("\n\n")
		}
		currentChunk.WriteString(para)
		wordCount += words
	}

	if currentChunk.Len() > 0 {
		chunks = append(chunks, domain.ChunkResult{
			Text:       currentChunk.String(),
			Index:      chunkIdx,
			Strategy:   domain.StrategySemantic,
			TokenCount: wordCount,
		})
	}

	return chunks, nil
}

// ============================================================
// 2. Heading-Aware Chunker — splits on markdown headings (# Title)
// ============================================================

type HeadingAwareChunker struct{}

var headingRegex = regexp.MustCompile(`(?m)^(#{1,6})\s+(.+)$`)

func (c *HeadingAwareChunker) Chunk(text string) ([]domain.ChunkResult, error) {
	if strings.TrimSpace(text) == "" {
		return nil, nil
	}

	indices := headingRegex.FindAllStringSubmatchIndex(text, -1)
	if len(indices) == 0 {
		// Fallback to semantic if no headings found
		return (&SemanticChunker{TargetChunkSize: 500}).Chunk(text)
	}

	var chunks []domain.ChunkResult
	headingStack := make(map[int]string) // level -> title
	chunkIdx := 0

	lastEnd := 0
	for _, idx := range indices {
		if idx[0] > lastEnd {
			sectionText := strings.TrimSpace(text[lastEnd:idx[0]])
			if sectionText != "" {
				chunks = append(chunks, domain.ChunkResult{
					Text:        sectionText,
					Index:       chunkIdx,
					Strategy:    domain.StrategyHeading,
					TokenCount:  len(strings.Fields(sectionText)),
					HeadingPath: buildHeadingPath(headingStack),
				})
				chunkIdx++
			}
		}

		level := idx[3] - idx[2]
		title := text[idx[4]:idx[5]]
		headingStack[level] = title
		// Clear deeper heading levels
		for l := level + 1; l <= 6; l++ {
			delete(headingStack, l)
		}
		lastEnd = idx[1]
	}

	if lastEnd < len(text) {
		remaining := strings.TrimSpace(text[lastEnd:])
		if remaining != "" {
			chunks = append(chunks, domain.ChunkResult{
				Text:        remaining,
				Index:       chunkIdx,
				Strategy:    domain.StrategyHeading,
				TokenCount:  len(strings.Fields(remaining)),
				HeadingPath: buildHeadingPath(headingStack),
			})
		}
	}

	return chunks, nil
}

func buildHeadingPath(stack map[int]string) string {
	var parts []string
	for l := 1; l <= 6; l++ {
		if title, ok := stack[l]; ok {
			parts = append(parts, title)
		}
	}
	return strings.Join(parts, " > ")
}

// ============================================================
// 3. Page-Aware Chunker — splits on page breaks (--- or \f or Page N)
// ============================================================

type PageAwareChunker struct{}

var pageBreakRegex = regexp.MustCompile(`(?i)(?:\f|---|\bPage\s+\d+\b)`)

func (c *PageAwareChunker) Chunk(text string) ([]domain.ChunkResult, error) {
	if strings.TrimSpace(text) == "" {
		return nil, nil
	}

	pages := pageBreakRegex.Split(text, -1)
	var chunks []domain.ChunkResult
	chunkIdx := 0

	for pageNum, page := range pages {
		pageText := strings.TrimSpace(page)
		if pageText == "" {
			continue
		}
		chunks = append(chunks, domain.ChunkResult{
			Text:       pageText,
			Index:      chunkIdx,
			Strategy:   domain.StrategyPage,
			TokenCount: len(strings.Fields(pageText)),
			PageNumber: pageNum + 1,
		})
		chunkIdx++
	}

	if len(chunks) == 0 {
		return (&SemanticChunker{TargetChunkSize: 500}).Chunk(text)
	}

	return chunks, nil
}

// ============================================================
// 4. Hierarchical Chunker — parent-child chunks
// ============================================================

type HierarchicalChunker struct{}

func (c *HierarchicalChunker) Chunk(text string) ([]domain.ChunkResult, error) {
	// First split into large parent sections, then smaller child chunks
	parents, err := (&SemanticChunker{TargetChunkSize: 1000}).Chunk(text)
	if err != nil {
		return nil, err
	}

	var chunks []domain.ChunkResult
	chunkIdx := 0
	for _, parent := range parents {
		children, _ := (&SemanticChunker{TargetChunkSize: 250}).Chunk(parent.Text)
		for _, child := range children {
			chunks = append(chunks, domain.ChunkResult{
				Text:       child.Text,
				Index:      chunkIdx,
				Strategy:   domain.StrategyHierarchical,
				TokenCount: child.TokenCount,
			})
			chunkIdx++
		}
	}

	return chunks, nil
}

// ============================================================
// 5. Document-Aware Chunker — structured component chunking
// ============================================================

type DocumentAwareChunker struct{}

func (c *DocumentAwareChunker) Chunk(text string) ([]domain.ChunkResult, error) {
	sections := strings.Split(text, "\n---\n")
	var chunks []domain.ChunkResult
	for i, sec := range sections {
		secText := strings.TrimSpace(sec)
		if secText == "" {
			continue
		}
		chunks = append(chunks, domain.ChunkResult{
			Text:       secText,
			Index:      i,
			Strategy:   domain.StrategyDocument,
			TokenCount: len(strings.Fields(secText)),
		})
	}
	if len(chunks) == 0 {
		return (&SemanticChunker{TargetChunkSize: 500}).Chunk(text)
	}
	return chunks, nil
}
