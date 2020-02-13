/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.3.8
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

/* A list of GetCorporationsCorporationIdContacts200Ok. */
//easyjson:json
type GetCorporationsCorporationIdContacts200OkList []GetCorporationsCorporationIdContacts200Ok

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdContacts200Ok struct {
	ContactId   int32   `json:"contact_id,omitempty"`   /* contact_id integer */
	ContactType string  `json:"contact_type,omitempty"` /* contact_type string */
	IsWatched   bool    `json:"is_watched,omitempty"`   /* Whether this contact is being watched */
	LabelIds    []int64 `json:"label_ids,omitempty"`    /* label_ids array */
	Standing    float32 `json:"standing,omitempty"`     /* Standing of the contact */
}
