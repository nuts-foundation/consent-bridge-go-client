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
	"bytes"
	"context"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"testing"
)

// RoundTripFunc
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newHttpClient(fn RoundTripFunc) HttpClient {
	return HttpClient{
		ServerAddress: "http://localhost:1323",
		customClient: &http.Client{
			Transport: RoundTripFunc(fn),
		},
		Logger: logrus.StandardLogger().WithField("component", "API-client"),
	}
}

func newHttpClientWithBody(status int, body []byte) HttpClient {
	return newHttpClient(func(req *http.Request) *http.Response {
		// Test request parameters
		return &http.Response{
			StatusCode: status,
			Body:       ioutil.NopCloser(bytes.NewReader(body)),
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		}
	})
}

func TestHttpClient_GetConsentRequestStateById(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		client := newHttpClientWithBody(200, []byte("{}"))

		_, err := client.GetConsentRequestStateById(context.TODO(), "uuid")

		if err != nil {
			t.Errorf("Expected no error, got [%v]", err)
		}
	})

	t.Run("404 returns error", func(t *testing.T) {
		client := newHttpClientWithBody(404, []byte{})

		_, err := client.GetConsentRequestStateById(context.TODO(), "uuid")

		if err == nil {
			t.Errorf("Expected error, got nothing")
		}
	})
}