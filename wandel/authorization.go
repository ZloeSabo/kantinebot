package wandel

import (
	"context"
	"fmt"
)

//NewIdentityResponse Response of NewIdentity request
type NewIdentityResponse struct {
	APIKey              string `json:"ApiKey"`
	FirebaseCustomToken string
	ProfileSecret       string
	ProfileToken        string
	ShortToken          string
	UserID              int32 `json:"UserId"`
}

type newIdentityResponseFull struct {
	Content NewIdentityResponse `json:"Content"`
}

//NewAuthorizationContext requests a new authorization token using the provided context
func (client *Client) NewAuthorizationContext(ctx context.Context) string {
	fullResponse := &newIdentityResponseFull{}

	err := getResource(ctx, client.httpclient, headersApp(client.app), client.endpoint+"NewIdentity", fullResponse, client)

	if err != nil {
		return ""
	}

	return fmt.Sprintf("0%v", fullResponse.Content.APIKey)
}

//NewAuthorization requests a new authorization token
func (client *Client) NewAuthorization() string {
	return client.NewAuthorizationContext(context.Background())
}
