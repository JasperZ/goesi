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

/* A list of GetUniversePlanetsPlanetIdOk. */
//easyjson:json
type GetUniversePlanetsPlanetIdOkList []GetUniversePlanetsPlanetIdOk

/* 200 ok object */
//easyjson:json
type GetUniversePlanetsPlanetIdOk struct {
	Name     string                             `json:"name,omitempty"`      /* name string */
	PlanetId int32                              `json:"planet_id,omitempty"` /* planet_id integer */
	Position GetUniversePlanetsPlanetIdPosition `json:"position,omitempty"`
	SystemId int32                              `json:"system_id,omitempty"` /* The solar system this planet is in */
	TypeId   int32                              `json:"type_id,omitempty"`   /* type_id integer */
}
