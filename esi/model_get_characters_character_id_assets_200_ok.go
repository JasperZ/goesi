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

/* A list of GetCharactersCharacterIdAssets200Ok. */
//easyjson:json
type GetCharactersCharacterIdAssets200OkList []GetCharactersCharacterIdAssets200Ok

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdAssets200Ok struct {
	IsBlueprintCopy bool   `json:"is_blueprint_copy,omitempty"` /* is_blueprint_copy boolean */
	IsSingleton     bool   `json:"is_singleton,omitempty"`      /* is_singleton boolean */
	ItemId          int64  `json:"item_id,omitempty"`           /* item_id integer */
	LocationFlag    string `json:"location_flag,omitempty"`     /* location_flag string */
	LocationId      int64  `json:"location_id,omitempty"`       /* location_id integer */
	LocationType    string `json:"location_type,omitempty"`     /* location_type string */
	Quantity        int32  `json:"quantity,omitempty"`          /* quantity integer */
	TypeId          int32  `json:"type_id,omitempty"`           /* type_id integer */
}
