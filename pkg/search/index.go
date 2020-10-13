package search

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/blevesearch/bleve"
	"github.com/policy-hub/policy-hub-cli/pkg/metaconfig"
)

// IndexVersion indicates the version of the search index.
// This is used to migrate between index versions.
const IndexVersion = "v1"


type Engine struct {
	index bleve.Index
}

// Load loads the engine and initializes the Index.
func Load() (*Engine, error) {
	e := &Engine{}
	if err := constructIndex(); err != nil {
		return nil, fmt.Errorf("construct index: %w", err)
	}

	index, err := loadIndex()
	if err != nil {
		return nil, fmt.Errorf("load index: %w", err)
	}
	e.index = index

	return e, nil
}

// Index indexes the metadata into the search engine.
// Metadata is indexed based on name.
func (e *Engine) Index(metadata []metaconfig.Metadata) error {
	for _, meta := range metadata {
		if err := e.index.Index(meta.Name, meta); err != nil {
			return fmt.Errorf("index metadata: %w", err)
		}
	}

	return nil
}

// Query queries the index and returns the SearchResult
func (e *Engine) Query(query string) (*bleve.SearchResult, error) {
	matchQ := bleve.NewMatchQuery(query)
	search := bleve.NewSearchRequest(matchQ)
	return e.index.Search(search)
} 

// constructIndex builds a search index
func constructIndex() error {
	cacheDir := cacheDirectory()
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		return setupIndexDirectory()
	}

	return nil
}

// loadIndex loads the index from disk
func loadIndex() (bleve.Index, error) {
	return bleve.Open(indexDirectory())
}

// setupIndexDirectory setups the index directory
func setupIndexDirectory() error {
	cacheDir := cacheDirectory()
	if err := os.MkdirAll(cacheDir, os.ModePerm); err != nil {
		return fmt.Errorf("make search dir: %w", err)
	}

	_, err := os.Create(filepath.Join(cacheDir, IndexVersion))
	if err != nil {
		return fmt.Errorf("create version file: %w", err)
	}
 
	mapping := bleve.NewIndexMapping()
	_, err = bleve.New(indexDirectory(), mapping)
	if err != nil {
		return fmt.Errorf("creating index: %w", err)
	}

	return nil
}

// cacheDirectory returns the directory to cache policy-cli configs
func cacheDirectory() string {
	const cacheDir = ".policy-hub"

	homeDir, _ := os.UserHomeDir()

	directory := filepath.Join(homeDir, cacheDir)
	directory = filepath.ToSlash(directory)

	return directory
}

// indexDirectory returns the directory to store the search index in
func indexDirectory() string {
	return filepath.Join(cacheDirectory(), "index")
}