package main

import (
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/n-arms/searching-for-dunsinane"
)

type wikiRenderer struct {
	parser   *parser.Parser
	renderer *html.Renderer
}

func makeWikiRenderer() wikiRenderer {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return wikiRenderer{
		parser:   p,
		renderer: renderer,
	}
}

func (w wikiRenderer) mdToHtml(md []byte) []byte {
	doc := w.parser.Parse(md)
	return markdown.Render(doc, w.renderer)
}

// clone the given wiki and return all the file paths as well as the temp directory to remove when you are finished with them
func cloneWiki(url string) ([]string, string) {
	target, err := os.MkdirTemp("", "wiki-processor-temp-*")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("git", "clone", url, target)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	markdownFiles, err := os.ReadDir(target)
	if err != nil {
		log.Fatal(err)
	}
	paths := []string{}
	for _, file := range markdownFiles {
		if !file.IsDir() {
			path := target + string(os.PathSeparator) + file.Name()
			paths = append(paths, path)
		}
	}
	return paths, target
}

func (w wikiRenderer) tokenizeMdFile(filePath string) []dunsinane.Token {
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	html := string(w.mdToHtml(fileContents))
	return dunsinane.TokenizeHtml(html)
}

func ProcessWiki(url string) (dunsinane.Index, []dunsinane.Token) {
	result := make(chan []dunsinane.Token)
	var wg sync.WaitGroup

	paths, tempDir := cloneWiki(url)
	defer os.RemoveAll(tempDir)

	for _, path := range paths {
		wg.Add(1)
		go func(path string) {
			w := makeWikiRenderer()
			result <- w.tokenizeMdFile(path)
			wg.Done()
		}(path)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	allTokens := []dunsinane.Token{}

	for tokens := range result {
		allTokens = append(allTokens, tokens...)
	}

	return dunsinane.MakeIndex(allTokens), allTokens
}
