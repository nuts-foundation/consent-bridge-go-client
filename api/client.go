/*
 * Nuts consent bridge api
 * Copyright (C) 2019. Nuts community
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nuts-foundation/consent-bridge-go-client/pkg"
	core "github.com/nuts-foundation/nuts-go-core"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

// HttpClient holds the server address and other basic settings for the http client
type HttpClient struct {
	ServerAddress string
	Timeout       time.Duration
	Logger 		  *logrus.Entry
	customClient  *http.Client
}

func (hb HttpClient) GetConsentRequestStateById(ctx context.Context, uuid string) (*FullConsentRequestState, error) {
	resp, err := hb.handleError("GetConsentRequestStateById", func() (*http.Response, error) {
		return hb.client().GetConsentRequestStateById(ctx, uuid)
	})

	if err == nil {
		defer resp.Body.Close()
	}

	// response is of type ConsentRequestState
	respObj := FullConsentRequestState{}

	err = json.NewDecoder(resp.Body).Decode(&respObj)
	if err != nil {
		err := fmt.Errorf("error in decoding ConsentRequestState: %w", err)
		hb.Logger.Error(err)
		return nil, err
	}

	return &respObj, nil
}

func (hb HttpClient) GetAttachmentBySecureHash(ctx context.Context, hash string) ([]byte, error) {
	resp, err := hb.handleError("GetAttachmentBySecureHash", func() (*http.Response, error) {
		return hb.client().GetAttachmentBySecureHash(ctx, hash)
	})

	if err == nil {
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	}

	return nil, err
}

func (hb HttpClient) handleError(name string, f func() (*http.Response, error) ) (*http.Response, error) {
	resp, err := f()

	if err != nil {
		err := fmt.Errorf("error in %s: %w", name, core.Wrap(err))
		hb.Logger.Error(err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		errorBody, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("error in %s, status: %d, reason: %v", name, resp.StatusCode, string(errorBody))
		hb.Logger.Error(err)
		return resp, err
	}

	return resp, nil
}

type BridgeClient interface {
	// GetConsentRequestStateById returns the consent request state metadata based on the uuid
	GetConsentRequestStateById(context.Context, string) (*FullConsentRequestState, error)
	// GetAttachmentBySecureHash retrieves an attachment by its hash
	GetAttachmentBySecureHash(context.Context, string) ([]byte, error)
}

// NewConsentBridgeClient returns a BridgeClient configured according to the current config
func NewConsentBridgeClient() BridgeClient {
	return HttpClient{
		ServerAddress: pkg.ConfigInstance().Address,
		Logger: logrus.WithField("component", "ConsentBridgeClient"),
		Timeout: time.Second * 60,
	}
}

func (hb HttpClient) client() *Client {
	if hb.customClient != nil {
		return &Client{
			Server: hb.ServerAddress,
			Client: hb.customClient,
		}
	}

	return &Client{
		Server: hb.ServerAddress,
	}
}
