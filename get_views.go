package telegraph

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// PageViews represents the number of page views for a Telegraph article.
type PageViews struct {
	// Number of page views for the target page.
	Views int `json:"views"`
}

// GetViews get the number of views for a Telegraph article. By default, the total number of page
// views will be returned. Returns a PageViews object on success.
func GetViews(path string, year, month, day, hour int) (*PageViews, error) {
	args := http.AcquireArgs()

	if hour > -1 {
		// If passed, the number of page views for the requested hour will be returned.
		args.Add("hour", strconv.Itoa(hour))
	}

	if day > 0 {
		// Required if hour is passed. If passed, the number of page views for the requested day will
		// be returned.
		args.Add("day", strconv.Itoa(day))
	}

	if month > 0 {
		// Required if day is passed. If passed, the number of page views for the requested month will
		// be returned.
		args.Add("month", strconv.Itoa(month))
	}

	if year > 0 {
		// Required if month is passed. If passed, the number of page views for the requested year
		// will be returned.
		args.Add("year", strconv.Itoa(year))
	}

	body, err := request("getViews", path, args)
	if err != nil {
		return nil, err
	}

	var resp PageViews
	err = json.Unmarshal(*body.Result, &resp)

	return &resp, err
}
