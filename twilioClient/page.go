package twilioClient

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/kevinburke/go-types"
	"golang.org/x/net/context"
)

type Page struct {
	FirstPageURI    string           `json:"first_page_uri"`
	Start           uint             `json:"start"`
	End             uint             `json:"end"`
	NumPages        uint             `json:"num_pages"`
	Total           uint             `json:"total"`
	NextPageURI     types.NullString `json:"next_page_uri"`
	PreviousPageURI types.NullString `json:"previous_page_uri"`
	PageSize        uint             `json:"page_size"`
}

var NoMoreResults = errors.New("twilio: No more results")

type PageIterator struct {
	twilioClient *TwilioClient
	nextPageURI  types.NullString
	data         url.Values
	count        uint
	pathPart     string
}

func (p *PageIterator) SetNextPageURI(npuri types.NullString) {
	if npuri.Valid == false {
		p.nextPageURI = npuri
		return
	}
	p.nextPageURI = npuri
}

func (p *PageIterator) Next(ctx context.Context, v interface{}) error {
	var err error
	switch {
	case p.nextPageURI.Valid:
		err = p.twilioClient.get(p.nextPageURI.String, url.Values{}, v)
	case p.count == 0:
		fmt.Println(p.pathPart)
		err = p.twilioClient.get(p.pathPart, p.data, v)
	default:
		return NoMoreResults
	}
	if err != nil {
		return err
	}
	p.count++
	return nil
}

func NewPageIterator(twilioClient *TwilioClient, data url.Values, pathPart string) *PageIterator {
	return &PageIterator{
		data:         data,
		twilioClient: twilioClient,
		count:        0,
		nextPageURI:  types.NullString{},
		pathPart:     pathPart,
	}
}

func NewNextPageIterator(twilioClient *TwilioClient, nextPageURI string) *PageIterator {
	if nextPageURI == "" {
		panic("nextpageuri is empty")
	}
	return &PageIterator{
		data:         url.Values{},
		twilioClient: twilioClient,
		nextPageURI:  types.NullString{Valid: true, String: nextPageURI},
		pathPart:     "",
		count:        0,
	}
}

func containsResultsInRange(start time.Time, end time.Time, results []time.Time) bool {
	for _, result := range results {
		if (result.Equal(start) || result.After(start)) && result.Before(end) {
			return true
		}
	}
	return false
}

func shouldContinuePaging(start time.Time, results []time.Time) bool {
	// the last result in results is the earliest. if the earliest result is
	// before the start, fetching more resources may return more results.
	if len(results) == 0 {
		panic("zero length result set")
	}
	last := results[len(results)-1]
	return last.After(start)
}
