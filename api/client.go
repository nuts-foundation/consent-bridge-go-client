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

func (hb HttpClient) GetConsentRequestStateById(ctx context.Context, uuid string) (*ConsentRequestState, error) {
	resp, err := hb.client().GetConsentRequestStateById(ctx, uuid)

	if err != nil {
		err := fmt.Errorf("error in retrieving ConsentRequestState by id: %v", err)
		hb.Logger.Error(err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errorBody, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("error in retrieving ConsentRequestState, status: %d, reason: %v", resp.StatusCode, errorBody)
		hb.Logger.Error(err)
		return nil, err
	}

	// response is of type ConsentRequestState
	respObj := ConsentRequestState{}

	err = json.NewDecoder(resp.Body).Decode(&respObj)
	if err != nil {
		err := fmt.Errorf("error in decoding ConsentRequestState: %v", err)
		hb.Logger.Error(err)
		return nil, err
	}

	return &respObj, nil
}

func (hb HttpClient) AcceptConsentRequestState(ctx context.Context, uuid string, pas PartyAttachmentSignature) error {
	resp, err := hb.client().GetConsentRequestStateById(ctx, uuid)

	if err != nil {
		err := fmt.Errorf("error in retrieving ConsentRequestState by id: %v", err)
		hb.Logger.Error(err)
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		errorBody, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("error in retrieving ConsentRequestState, status: %d, reason: %v", resp.StatusCode, errorBody)
		hb.Logger.Error(err)
		return err
	}

	// todo do someting usefull with ConsentRequestJobState

	return nil
}

type BridgeClient interface {
	// GetConsentRequestStateById returns the consent request state metadata based on the uuid
	GetConsentRequestStateById(context.Context, string) (*ConsentRequestState, error)
	// AcceptConsentRequestState accept a consent request with a signature proof
	AcceptConsentRequestState(context.Context, string, PartyAttachmentSignature) error
}



func (hb HttpClient) client() *Client {
	if hb.customClient != nil {
		return &Client{
			Server: fmt.Sprintf("http://%v", hb.ServerAddress),
			Client: *hb.customClient,
		}
	}

	return &Client{
		Server: fmt.Sprintf("http://%v", hb.ServerAddress),
	}
}
