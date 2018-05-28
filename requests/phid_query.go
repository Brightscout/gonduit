package requests

type PHIDQueryRequest struct {
	PHIDs map[string]string `json:"phids"`
	Request
}
