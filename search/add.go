package search

import (
	"WebpageArchiver/common"
	"github.com/meilisearch/meilisearch-go"
)

func addDoc(fileName string) (status string) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   common.MEILIHOST,
		APIKey: common.MEILIAPIKey,
	})

	HTMLContent, err := common.GetHTMLFileContent(common.HTMLPath + fileName)
	if err != nil {
		return
	}
	title, err := common.GetHTMLTitle(HTMLContent)
	if err != nil {
		return
	}
	HTMLPuretext, err := common.ExtractHTMLText(HTMLContent)
	if err != nil {
		return
	}

	documents := []map[string]interface{}{
		{
			"title":    title,
			"filename": fileName,
			"link":     "",
			"content":  HTMLPuretext,
		},
	}
	_, err = client.Index(common.MEILIBlogsIndex).AddDocuments(documents)
	return "Done"
}
