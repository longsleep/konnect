/*
 * Copyright 2017 Kopano and its licensors
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, version 3,
 * as published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package identifier

// A LogonRequest is the request data as sent to the logon endpoint
type LogonRequest struct {
	State     string `json:"state"`
	RawMaxAge string `json:"max_age"`

	Params []string `json:"params"`
}

// A LogonResponse holds a response as sent by the logon endpoint.
type LogonResponse struct {
	Success bool   `json:"success"`
	State   string `json:"state"`
}

// A HelloRequest is the request data as send to the hello endpoint.
type HelloRequest struct {
	State  string `json:"state"`
	Prompt bool   `json:"prompt"`
}

// A HelloResponse holds a response  as sent by the hello endpoint.
type HelloResponse struct {
	State       string `json:"state"`
	Success     bool   `json:"success"`
	Username    string `json:"username,omitempty"`
	DisplayName string `json:"displayname,omitempty"`
}