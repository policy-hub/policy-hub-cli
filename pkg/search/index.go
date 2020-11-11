package search

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/olekukonko/tablewriter"
	"github.com/policy-hub/policy-hub-cli/pkg/helpers"
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

func (e *Engine) ListResults(out io.Writer, res *bleve.SearchResult, metadata []metaconfig.Metadata) {
	table := tablewriter.NewWriter(out)
	table.SetHeader([]string{"name", "maintainers", "labels"})
	for _, hit := range res.Hits {
		var row []string
		row = append(row, hit.ID)
		for _, data := range metadata {
			if data.Name == hit.ID {
				row = append(row, strings.Join(data.Maintainers, ", "), strings.Join(data.Labels, ", "))
			}
		}

		table.Append(row)
	}

	if table.NumLines() > 0 {
		table.Render()
	} else {
		fmt.Fprintln(out, "No matches")
	}
}

func (e *Engine) Close() error {
	return e.index.Close()
}

// constructIndex builds a search index
func constructIndex() (bleve.Index, error) {
	cacheDir := helpers.IndexPath()
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		return setupIndexDirectory()
	}

	return nil, errors.New("called constructIndex while index already exists")
}

// loadIndex loads the index from disk
func loadIndex() (bleve.Index, error) {
	index, err := bleve.Open(indexDirectory())
	if err == bleve.ErrorIndexPathDoesNotExist {
		index, err := constructIndex()
		if err != nil {
			return nil, fmt.Errorf("construct index: %w", err)
		}

		return index, nil
	} else if err != nil {
		return nil, fmt.Errorf("open index: %w", err)
	}

	return index, nil
}

// setupIndexDirectory setups the index directory
func setupIndexDirectory() (bleve.Index, error) {
	cacheDir := helpers.IndexPath()
	if err := os.MkdirAll(cacheDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("make search dir: %w", err)
	}

	_, err := os.Create(filepath.Join(cacheDir, IndexVersion))
	if err != nil {
		return nil, fmt.Errorf("create version file: %w", err)
	}

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(indexDirectory(), mapping)
	if err != nil {
		return nil, fmt.Errorf("creating index: %w", err)
	}

	return index, nil
}

// indexDirectory returns the directory to store the search index in
func indexDirectory() string {
	return filepath.Join(helpers.IndexPath(), "index")
}
