package nasa

import "fmt"

type ApiResponse struct {
	Date        string `json:"date"`
	MediaType   string `json:"media_type"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	HdUrl       string `json:"hdurl"`
	Explanation string `json:"explanation"`
}

type ImagesResponse struct {
	Images []ApiResponse `json:"images"`
}

func (a ApiResponse) GetDate() string {
	return a.Date
}

func (a ApiResponse) GetTitle() string {
	return a.Title
}

func (a ApiResponse) GetUrl() string {
	return a.Url
}

func (a ApiResponse) GetHdUrl() string {
	return a.HdUrl
}

func (a ApiResponse) GetExplanation() string {
	return a.Explanation
}

func (a ApiResponse) GetMediaType() string {
	return a.MediaType
}

func (a ApiResponse) String() string {
	s := ``
	s += fmt.Sprintln("Title:", a.GetTitle())
	s += fmt.Sprintln("Date:", a.GetDate())
	s += fmt.Sprintln("Explanation:", a.GetExplanation())
	s += fmt.Sprintln("Media-Type:", a.GetMediaType())
	s += fmt.Sprintln("Url:", a.GetUrl())
	s += fmt.Sprintln("HdUrl:", a.GetHdUrl())
	return s
}

func (i ImagesResponse) String() string {
	s := ``
	for _, image := range i.Images {
		s += image.String()
	}
	return s
}
