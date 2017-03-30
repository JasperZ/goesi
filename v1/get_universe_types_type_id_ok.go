/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.2.dev17
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

package goesiv1

/* 200 ok object */
type GetUniverseTypesTypeIdOk struct {

	/* category_id integer */
	CategoryId int32 `json:"category_id,omitempty"`

	/* graphic_id integer */
	GraphicId int32 `json:"graphic_id,omitempty"`

	/* group_id integer */
	GroupId int32 `json:"group_id,omitempty"`

	/* icon_id integer */
	IconId int32 `json:"icon_id,omitempty"`

	/* type_description string */
	TypeDescription string `json:"type_description,omitempty"`

	/* type_name string */
	TypeName string `json:"type_name,omitempty"`
}
