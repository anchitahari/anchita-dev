package rag

import (
	"math"
	"os"
	"path/filepath"
	"strings"
)

type Chunk struct {
	Text   string
	Source string
}

var chunks []Chunk

// load knowledge base
func LoadKnowledge(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".txt" {
			continue
		}
		path := filepath.Join(dir, f.Name())
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		// split into chunks of ~500 chars with overlap
		text := string(content)
		parts := chunkText(text, 500, 100)
		for _, part := range parts {
			chunks = append(chunks, Chunk{
				Text:   part,
				Source: f.Name(),
			})
		}
	}
	return nil
}

func chunkText(text string, size, overlap int) []string {
	words := strings.Fields(text)
	var chunks []string
	for i := 0; i < len(words); i += size - overlap {
		end := i + size
		if end > len(words) {
			end = len(words)
		}
		chunk := strings.Join(words[i:end], " ")
		chunks = append(chunks, chunk)
		if end == len(words) {
			break
		}
	}
	return chunks
}

// simple keyword-based retrieval
// returns the top n most relevant chunks for a query
func Retrieve(query string, topN int) []Chunk {
	query = strings.ToLower(query)
	queryWords := strings.Fields(query)

	type scored struct {
		chunk Chunk
		score float64
	}

	var scored_chunks []scored
	for _, chunk := range chunks {
		chunkLower := strings.ToLower(chunk.Text)
		score := 0.0
		for _, word := range queryWords {
			if len(word) < 3 {
				continue
			}
			if strings.Contains(chunkLower, word) {
				score += 1.0
			}
			// boost score for exact phrase match
			if strings.Contains(chunkLower, query) {
				score += 2.0
			}
		}
		// normalize by chunk length to avoid bias toward longer chunks
		words := float64(len(strings.Fields(chunk.Text)))
		if words > 0 {
			score = score / math.Sqrt(words)
		}
		scored_chunks = append(scored_chunks, struct {
			chunk Chunk
			score float64
		}{chunk, score})
	}

	// sort by score descending
	for i := 0; i < len(scored_chunks)-1; i++ {
		for j := i + 1; j < len(scored_chunks); j++ {
			if scored_chunks[j].score > scored_chunks[i].score {
				scored_chunks[i], scored_chunks[j] = scored_chunks[j], scored_chunks[i]
			}
		}
	}

	var result []Chunk
	for i := 0; i < topN && i < len(scored_chunks); i++ {
		if scored_chunks[i].score > 0 {
			result = append(result, scored_chunks[i].chunk)
		}
	}

	// if no relevant chunks found, return first 3 as fallback
	if len(result) == 0 && len(chunks) > 0 {
		end := 3
		if end > len(chunks) {
			end = len(chunks)
		}
		return chunks[:end]
	}

	return result
}