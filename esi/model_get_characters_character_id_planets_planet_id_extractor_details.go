/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.8.2
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

/* A list of GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails. */
//easyjson:json
type GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList []GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails

/* extractor_details object */
//easyjson:json
type GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails struct {
	CycleTime     int32                                         `json:"cycle_time,omitempty"`      /* in seconds */
	HeadRadius    float32                                       `json:"head_radius,omitempty"`     /* head_radius number */
	Heads         []GetCharactersCharacterIdPlanetsPlanetIdHead `json:"heads,omitempty"`           /* heads array */
	ProductTypeId int32                                         `json:"product_type_id,omitempty"` /* product_type_id integer */
	QtyPerCycle   int32                                         `json:"qty_per_cycle,omitempty"`   /* qty_per_cycle integer */
}
