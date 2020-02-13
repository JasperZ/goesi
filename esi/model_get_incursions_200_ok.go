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

/* A list of GetIncursions200Ok. */
//easyjson:json
type GetIncursions200OkList []GetIncursions200Ok

/* 200 ok object */
//easyjson:json
type GetIncursions200Ok struct {
	ConstellationId      int32   `json:"constellation_id,omitempty"`        /* The constellation id in which this incursion takes place */
	FactionId            int32   `json:"faction_id,omitempty"`              /* The attacking faction's id */
	HasBoss              bool    `json:"has_boss,omitempty"`                /* Whether the final encounter has boss or not */
	InfestedSolarSystems []int32 `json:"infested_solar_systems,omitempty"`  /* A list of infested solar system ids that are a part of this incursion */
	Influence            float32 `json:"influence,omitempty"`               /* Influence of this incursion as a float from 0 to 1 */
	StagingSolarSystemId int32   `json:"staging_solar_system_id,omitempty"` /* Staging solar system for this incursion */
	State                string  `json:"state,omitempty"`                   /* The state of this incursion */
	Type_                string  `json:"type,omitempty"`                    /* The type of this incursion */
}
