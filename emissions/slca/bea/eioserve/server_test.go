/*
Copyright © 2018 the InMAP authors.
This file is part of InMAP.

InMAP is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

InMAP is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with InMAP.  If not, see <http://www.gnu.org/licenses/>.*/

package eioserve

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"testing"

	eioservepb "github.com/spatialmodel/inmap/emissions/slca/bea/eioserve/proto/eioservepb"
)

func TestServer_grpc(t *testing.T) {
	s, err := NewServer()
	if err != nil {
		t.Fatalf("failed to create server: %v", err)
	}

	go func() {
		http.ListenAndServeTLS(":10000", "test_data/cert.pem", "test_data/key.pem", s)
	}()

	t.Run("index", func(t *testing.T) {
		client := &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}

		res, err := client.Get("https://" + eioservepb.Address)
		if err != nil {
			t.Error(err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("Response code was %v; want 200", res.StatusCode)
		}

		expected := []byte("<!DOCTYPE html>")
		body := make([]byte, len(expected))
		_, err = res.Body.Read(body)
		if err != nil {
			t.Fatal(err)
		}

		if bytes.Compare(expected, body) != 0 {
			t.Errorf("Response body was '%s'; want '%s'", expected, body)
		}

	})

	/*c := eioclientpb.NewEIOServeClient("https://" + eioservepb.Address)

	ctx := context.Background()

	t.Run("DemandGroups", func(t *testing.T) {
		r, err := c.DemandGroups(ctx, &eioclientpb.Selection{
			DemandGroup:      eioservepb.All,
			DemandSector:     eioservepb.All,
			ProductionGroup:  eioservepb.All,
			ProductionSector: eioservepb.All,
			ImpactType:       "health_total",
			DemandType:       "All",
		})
		if err != nil {
			t.Error(err)
		}
		fmt.Println(r)
	})*/
}
