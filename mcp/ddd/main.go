// Command ddd-mcp serves the "DDD Vite Fait" book to MCP clients.
// Agents read only the section they need instead of the whole PDF.
package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type noInput struct{}

type getInput struct {
	ID string `json:"id" jsonschema:"section ID from list_sections"`
}

type searchInput struct {
	Query string `json:"query" jsonschema:"word or phrase to look for (book is in French)"`
	Limit int    `json:"limit,omitempty" jsonschema:"max results, default 5"`
}

func text(s string) *mcp.CallToolResult {
	return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: s}}}
}

func newServer(sections []Section) *mcp.Server {
	server := mcp.NewServer(&mcp.Implementation{Name: "ddd", Version: "1.0.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_sections",
		Description: "List the sections of the DDD book (ID and title). Cheap: call it first.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ noInput) (*mcp.CallToolResult, any, error) {
		var b strings.Builder
		for _, s := range sections {
			fmt.Fprintf(&b, "%s: %s\n", s.ID, s.Title)
		}
		return text(b.String()), nil, nil
	})

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_section",
		Description: "Get the full text of one section of the DDD book.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, in getInput) (*mcp.CallToolResult, any, error) {
		s, ok := find(sections, in.ID)
		if !ok {
			return nil, nil, fmt.Errorf("unknown section %q: use list_sections", in.ID)
		}
		return text(s.Body), nil, nil
	})

	mcp.AddTool(server, &mcp.Tool{
		Name:        "search",
		Description: "Find which sections mention a word or phrase. Returns section IDs with a short excerpt.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, in searchInput) (*mcp.CallToolResult, any, error) {
		limit := in.Limit
		if limit <= 0 {
			limit = 5
		}
		hits := search(sections, in.Query, limit)
		if len(hits) == 0 {
			return text("no match"), nil, nil
		}
		return text(strings.Join(hits, "\n")), nil, nil
	})

	return server
}

func main() {
	sections, err := loadSections()
	if err != nil {
		log.Fatal(err)
	}
	if err := newServer(sections).Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
