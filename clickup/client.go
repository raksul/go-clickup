package clickup

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL      = "https://api.clickup.com/api/v2/"
	userAgent           = "go-clickup"
	headerRateRemaining = "X-RateLimit-Remaining"
	headerRateLimit     = "X-RateLimit-Limit"
	headerRateReset     = "X-RateLimit-Reset"
)

var errNonNilContext = errors.New("context must be non-nil")

// A Client manages communication with the Clickup API.
type Client struct {
	clientMu sync.Mutex   // clientMu protects the client during calls that modify the CheckRedirect func.
	client   *http.Client // HTTP client used to communicate with the API.
	APIKey   string

	BaseURL   *url.URL
	UserAgent string

	rateMu     sync.Mutex
	rateLimits Rate // Rate limits for the client as determined by the most recent API calls.

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Clickup API.
	Attachments     *AttachmentsService
	Authorization   *AuthorizationService
	Checklists      *ChecklistsService
	Comments        *CommentsService
	CustomFields    *CustomFieldsService
	Dependencies    *DependenciesService
	Goals           *GoalsService
	Tasks           *TasksService
	TaskTemplates   *TaskTemplatesService
	Teams           *TeamsService
	SharedHierarchy *SharedHierarchyService
	Spaces          *SpacesService
	Folders         *FoldersService
	Lists           *ListsService
	Members         *MembersService
	Tags            *TagsService
	Views           *ViewsService
	Webhooks        *WebhooksService
}

type service struct {
	client *Client
}

// Client returns the http.Client used by this Clickup client.
func (c *Client) Client() *http.Client {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	clientCopy := *c.client
	return &clientCopy
}

// RawType represents type of raw format of a request instead of JSON.
type RawType uint8

const (
	// Diff format.
	Diff RawType = 1 + iota
	// Patch format.
	Patch
)

// RawOptions specifies parameters when user wants to get raw format of
// a response instead of JSON.
type RawOptions struct {
	Type RawType
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewClient returns a new Clickup API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide a APIKey.
func NewClient(httpClient *http.Client, APIKey string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent, APIKey: APIKey}
	c.common.client = c
	c.Attachments = (*AttachmentsService)(&c.common)
	c.Authorization = (*AuthorizationService)(&c.common)
	c.Checklists = (*ChecklistsService)(&c.common)
	c.Comments = (*CommentsService)(&c.common)
	c.CustomFields = (*CustomFieldsService)(&c.common)
	c.Dependencies = (*DependenciesService)(&c.common)
	c.Goals = (*GoalsService)(&c.common)
	c.Tasks = (*TasksService)(&c.common)
	c.TaskTemplates = (*TaskTemplatesService)(&c.common)
	c.Teams = (*TeamsService)(&c.common)
	c.SharedHierarchy = (*SharedHierarchyService)(&c.common)
	c.Spaces = (*SpacesService)(&c.common)
	c.Folders = (*FoldersService)(&c.common)
	c.Lists = (*ListsService)(&c.common)
	c.Members = (*MembersService)(&c.common)
	c.Tags = (*TagsService)(&c.common)
	c.Views = (*ViewsService)(&c.common)
	c.Webhooks = (*WebhooksService)(&c.common)
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.APIKey)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

func (c *Client) NewMultiPartRequest(method, urlStr string, buf *bytes.Buffer) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.APIKey)

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Response is a Clickup API response. This wraps the standard http.Response
// returned from Clickup.
// Clickup does not currently support pagination, but if it did, we would provides
// convenient access to things like pagination links as if go-github.
type Response struct {
	*http.Response

	// Explicitly specify the Rate type so Rate's String() receiver doesn't
	// propagate to Response.
	Rate Rate
}

// newResponse creates a new Response for the provided http.Response.
// r must not be nil.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	response.Rate = parseRate(r)
	return response
}

// parseRate parses the rate related headers.
func parseRate(r *http.Response) Rate {
	var rate Rate
	if limit := r.Header.Get(headerRateLimit); limit != "" {
		rate.Limit, _ = strconv.Atoi(limit)
	}
	if remaining := r.Header.Get(headerRateRemaining); remaining != "" {
		rate.Remaining, _ = strconv.Atoi(remaining)
	}
	if reset := r.Header.Get(headerRateReset); reset != "" {
		if v, _ := strconv.ParseInt(reset, 10, 64); v != 0 {
			rate.Reset = Timestamp{time.Unix(v, 0)}
		}
	}
	return rate
}

type requestContext uint8

const (
	bypassRateLimitCheck requestContext = iota
)

// BareDo sends an API request and lets you handle the api response. If an error
// or API Error occurs, the error will contain more information. Otherwise you
// are supposed to read and close the response's Body. If rate limit is exceeded
// and reset time is in the future, BareDo returns *RateLimitError immediately
// without making a network API call.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it is
// canceled or times out, ctx.Err() will be returned.
func (c *Client) BareDo(ctx context.Context, req *http.Request) (*Response, error) {
	if ctx == nil {
		return nil, errNonNilContext
	}

	if bypass := ctx.Value(bypassRateLimitCheck); bypass == nil {
		// If we've hit rate limit, don't make further requests before Reset time.
		if err := c.checkRateLimitBeforeDo(req); err != nil {
			return &Response{
				Response: err.Response,
				Rate:     err.Rate,
			}, err
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = sanitizeURL(url).String()
				return nil, e
			}
		}

		return nil, err
	}

	response := newResponse(resp)

	c.rateMu.Lock()
	c.rateLimits = response.Rate
	c.rateMu.Unlock()

	err = CheckResponse(resp)
	if err != nil {
		defer resp.Body.Close()
	}
	return response, err
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer interface,
// the raw response body will be written to v, without attempting to first
// decode it. If v is nil, and no error happens, the response is returned as is.
// If rate limit is exceeded and reset time is in the future, Do returns
// *RateLimitError immediately without making a network API call.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it
// is canceled or times out, ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.BareDo(ctx, req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

// checkRateLimitBeforeDo does not make any network calls, but uses existing knowledge from
// current client state in order to quickly check if *RateLimitError can be immediately returned
// from Client.Do, and if so, returns it so that Client.Do can skip making a network API call unnecessarily.
// Otherwise it returns nil, and Client.Do should proceed normally.
func (c *Client) checkRateLimitBeforeDo(req *http.Request) *RateLimitError {
	c.rateMu.Lock()
	rate := c.rateLimits
	c.rateMu.Unlock()
	if !rate.Reset.Time.IsZero() && rate.Remaining == 0 && time.Now().Before(rate.Reset.Time) {
		// Create a fake response.
		resp := &http.Response{
			Status:     http.StatusText(http.StatusForbidden),
			StatusCode: http.StatusForbidden,
			Request:    req,
			Header:     make(http.Header),
			Body:       ioutil.NopCloser(strings.NewReader("")),
		}
		return &RateLimitError{
			Rate:     rate,
			Response: resp,
			Message:  fmt.Sprintf("API rate limit of %v still exceeded until %v, not making remote request.", rate.Limit, rate.Reset.Time),
		}
	}

	return nil
}

// compareHttpResponse returns whether two http.Response objects are equal or not.
// Currently, only StatusCode is checked. This function is used when implementing the
// Is(error) bool interface for the custom error types in this package.
func compareHttpResponse(r1, r2 *http.Response) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if r1 != nil && r2 != nil {
		return r1.StatusCode == r2.StatusCode
	}
	return false
}

/*
An ErrorResponse reports one or more errors caused by an API request.
*/
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
	Errors   []Error        `json:"errors"`  // more detail on individual errors
	// Block is only populated on certain types of errors such as code 451.
	Block *ErrorBlock `json:"block,omitempty"`
}

// ErrorBlock contains a further explanation for the reason of an error.
type ErrorBlock struct {
	Reason    string     `json:"reason,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.Message, r.Errors)
}

// Is returns whether the provided error equals this error.
func (r *ErrorResponse) Is(target error) bool {
	v, ok := target.(*ErrorResponse)
	if !ok {
		return false
	}

	if r.Message != v.Message || !compareHttpResponse(r.Response, v.Response) {
		return false
	}

	// Compare Errors.
	if len(r.Errors) != len(v.Errors) {
		return false
	}
	for idx := range r.Errors {
		if r.Errors[idx] != v.Errors[idx] {
			return false
		}
	}

	// Compare Block.
	if (r.Block != nil && v.Block == nil) || (r.Block == nil && v.Block != nil) {
		return false
	}
	if r.Block != nil && v.Block != nil {
		if r.Block.Reason != v.Block.Reason {
			return false
		}
		if (r.Block.CreatedAt != nil && v.Block.CreatedAt == nil) || (r.Block.CreatedAt ==
			nil && v.Block.CreatedAt != nil) {
			return false
		}
		if r.Block.CreatedAt != nil && v.Block.CreatedAt != nil {
			if *(r.Block.CreatedAt) != *(v.Block.CreatedAt) {
				return false
			}
		}
	}

	return true
}

// RateLimitError occurs when Clickup returns 429 Too Many Requests response with a rate limit
// remaining value of 0.
type RateLimitError struct {
	Rate     Rate           // Rate specifies last known rate limit for the client
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
}

func (r *RateLimitError) Error() string {
	return fmt.Sprintf("%v %v: %d %v %v",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.Message, formatRateReset(time.Until(r.Rate.Reset.Time)))
}

// Is returns whether the provided error equals this error.
func (r *RateLimitError) Is(target error) bool {
	v, ok := target.(*RateLimitError)
	if !ok {
		return false
	}

	return r.Rate == v.Rate &&
		r.Message == v.Message &&
		compareHttpResponse(r.Response, v.Response)
}

// AbuseRateLimitError occurs when Clickup returns 429 Too Many Requests response with the
type AbuseRateLimitError struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
}

func (r *AbuseRateLimitError) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.Message)
}

// Is returns whether the provided error equals this error.
func (r *AbuseRateLimitError) Is(target error) bool {
	v, ok := target.(*AbuseRateLimitError)
	if !ok {
		return false
	}

	return r.Message == v.Message && compareHttpResponse(r.Response, v.Response)
}

// sanitizeURL redacts the client_secret parameter from the URL which may be
// exposed to the user.
func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("client_secret")) > 0 {
		params.Set("client_secret", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}

/*
An Error reports more details on an individual error in an ErrorResponse.
These are the possible validation error codes:

	missing:
	    resource does not exist
	missing_field:
	    a required field on a resource has not been set
	invalid:
	    the formatting of a field is invalid
	already_exists:
	    another resource has the same valid as this field
	custom:
	    information is set in the Message field of the Error
*/
type Error struct {
	Resource string `json:"resource"` // resource on which the error occurred
	Field    string `json:"field"`    // field on which the error occurred
	Code     string `json:"code"`     // validation error code
	Message  string `json:"message"`  // Message describing the error. Errors with Code == "custom" will always have this set.
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v error caused by %v field on %v resource",
		e.Code, e.Field, e.Resource)
}

func (e *Error) UnmarshalJSON(data []byte) error {
	type aliasError Error // avoid infinite recursion by using type alias.
	if err := json.Unmarshal(data, (*aliasError)(e)); err != nil {
		return json.Unmarshal(data, &e.Message) // data can be json string.
	}
	return nil
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range or equal to 202 Accepted.
// API error responses are expected to have response
// body, and a JSON response body that maps to ErrorResponse.
//
// The error type will be *RateLimitError for rate limit exceeded errors,
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	switch {
	case r.StatusCode == http.StatusTooManyRequests && r.Header.Get(headerRateRemaining) == "0":
		return &RateLimitError{
			Rate:     parseRate(r),
			Response: errorResponse.Response,
			Message:  errorResponse.Message,
		}
	default:
		return errorResponse
	}
}

// Rate represents the rate limit for the current client.
type Rate struct {
	// The number of requests per hour the client is currently limited to.
	Limit int `json:"limit"`

	// The number of remaining requests the client can make this hour.
	Remaining int `json:"remaining"`

	// The time at which the current rate limit will reset.
	Reset Timestamp `json:"reset"`
}

func (r Rate) String() string {
	return Stringify(r)
}

// formatRateReset formats d to look like "[rate reset in 2s]" or
// "[rate reset in 87m02s]" for the positive durations. And like "[rate limit was reset 87m02s ago]"
// for the negative cases.
func formatRateReset(d time.Duration) string {
	isNegative := d < 0
	if isNegative {
		d *= -1
	}
	secondsTotal := int(0.5 + d.Seconds())
	minutes := secondsTotal / 60
	seconds := secondsTotal - minutes*60

	var timeString string
	if minutes > 0 {
		timeString = fmt.Sprintf("%dm%02ds", minutes, seconds)
	} else {
		timeString = fmt.Sprintf("%ds", seconds)
	}

	if isNegative {
		return fmt.Sprintf("[rate limit was reset %v ago]", timeString)
	}
	return fmt.Sprintf("[rate reset in %v]", timeString)
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
