package search

import (
	"WebpageArchiver/common"
	"github.com/meilisearch/meilisearch-go"
)

func AddDocFile(fileName string, originLink string) (err error) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   common.MEILIHOST,
		APIKey: common.MEILIAPIKey,
	})

	HTMLContent, err := common.GetHTMLFileContent(common.HTMLPath + fileName)
	if err != nil {
		return err
	}
	title, err := common.GetHTMLTitle(HTMLContent)
	if err != nil {
		return err
	}
	HTMLPuretext, err := common.ExtractHTMLText(HTMLContent)
	if err != nil {
		return err
	}

	documents := []map[string]interface{}{
		{
			"title":    title,
			"filename": fileName,
			"link":     originLink,
			"content":  HTMLPuretext,
		},
	}
	_, err = client.Index(common.MEILIBlogsIndex).AddDocuments(documents)
	return nil
}
