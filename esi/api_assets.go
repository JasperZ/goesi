/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.7.6
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

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"fmt"

	"github.com/antihax/goesi/optional"
)

// Linger please
var (
	_ context.Context
)

type AssetsApiService service

/*
AssetsApiService Get character assets
Return a list of the characters assets  ---  This route is cached for up to 3600 seconds
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param characterId An EVE character ID
 * @param optional nil or *GetCharactersCharacterIdAssetsOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "Page" (optional.Int32) -  Which page of results to return
     * @param "Token" (optional.String) -  Access token to use if unable to set a header
     * @param "UserAgent" (optional.String) -  Client identifier, takes precedence over headers
     * @param "XUserAgent" (optional.String) -  Client identifier, takes precedence over User-Agent

@return []GetCharactersCharacterIdAssets200Ok
*/

type GetCharactersCharacterIdAssetsOpts struct {
	Datasource optional.String
	Page       optional.Int32
	Token      optional.String
	UserAgent  optional.String
	XUserAgent optional.String
}

func (a *AssetsApiService) GetCharactersCharacterIdAssets(ctx context.Context, characterId int32, localVarOptionals *GetCharactersCharacterIdAssetsOpts) ([]GetCharactersCharacterIdAssets200Ok, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GetCharactersCharacterIdAssets200OkList
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v3/characters/{character_id}/assets/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if characterId < 1 {
		return localVarReturnValue, nil, reportError("characterId must be greater than 1")
	}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarQueryParams.Add("token", parameterToString(localVarOptionals.Token.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserAgent.IsSet() {
		localVarQueryParams.Add("user_agent", parameterToString(localVarOptionals.UserAgent.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XUserAgent.IsSet() {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarOptionals.XUserAgent.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v []GetCharactersCharacterIdAssets200Ok
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v Forbidden
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
AssetsApiService Get corporation assets
Return a list of the corporation assets  ---  This route is cached for up to 3600 seconds  --- Requires one of the following EVE corporation role(s): Director
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param corporationId An EVE corporation ID
 * @param optional nil or *GetCorporationsCorporationIdAssetsOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "Page" (optional.Int32) -  Which page of results to return
     * @param "Token" (optional.String) -  Access token to use if unable to set a header
     * @param "UserAgent" (optional.String) -  Client identifier, takes precedence over headers
     * @param "XUserAgent" (optional.String) -  Client identifier, takes precedence over User-Agent

@return []GetCorporationsCorporationIdAssets200Ok
*/

type GetCorporationsCorporationIdAssetsOpts struct {
	Datasource optional.String
	Page       optional.Int32
	Token      optional.String
	UserAgent  optional.String
	XUserAgent optional.String
}

func (a *AssetsApiService) GetCorporationsCorporationIdAssets(ctx context.Context, corporationId int32, localVarOptionals *GetCorporationsCorporationIdAssetsOpts) ([]GetCorporationsCorporationIdAssets200Ok, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GetCorporationsCorporationIdAssets200OkList
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v3/corporations/{corporation_id}/assets/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if corporationId < 1 {
		return localVarReturnValue, nil, reportError("corporationId must be greater than 1")
	}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarQueryParams.Add("token", parameterToString(localVarOptionals.Token.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserAgent.IsSet() {
		localVarQueryParams.Add("user_agent", parameterToString(localVarOptionals.UserAgent.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XUserAgent.IsSet() {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarOptionals.XUserAgent.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v []GetCorporationsCorporationIdAssets200Ok
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v Forbidden
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
AssetsApiService Get character asset locations
Return locations for a set of item ids, which you can get from character assets endpoint. Coordinates for items in hangars or stations are set to (0,0,0)  ---
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param characterId An EVE character ID
 * @param itemIds A list of item ids
 * @param optional nil or *PostCharactersCharacterIdAssetsLocationsOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "Token" (optional.String) -  Access token to use if unable to set a header
     * @param "UserAgent" (optional.String) -  Client identifier, takes precedence over headers
     * @param "XUserAgent" (optional.String) -  Client identifier, takes precedence over User-Agent

@return []PostCharactersCharacterIdAssetsLocations200Ok
*/

type PostCharactersCharacterIdAssetsLocationsOpts struct {
	Datasource optional.String
	Token      optional.String
	UserAgent  optional.String
	XUserAgent optional.String
}

func (a *AssetsApiService) PostCharactersCharacterIdAssetsLocations(ctx context.Context, characterId int32, itemIds []int64, localVarOptionals *PostCharactersCharacterIdAssetsLocationsOpts) ([]PostCharactersCharacterIdAssetsLocations200Ok, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue PostCharactersCharacterIdAssetsLocations200OkList
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v2/characters/{character_id}/assets/locations/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if characterId < 1 {
		return localVarReturnValue, nil, reportError("characterId must be greater than 1")
	}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarQueryParams.Add("token", parameterToString(localVarOptionals.Token.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserAgent.IsSet() {
		localVarQueryParams.Add("user_agent", parameterToString(localVarOptionals.UserAgent.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XUserAgent.IsSet() {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarOptionals.XUserAgent.Value(), "")
	}
	// body params
	localVarPostBody = &itemIds
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v []PostCharactersCharacterIdAssetsLocations200Ok
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v Forbidden
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
AssetsApiService Get character asset names
Return names for a set of item ids, which you can get from character assets endpoint. Typically used for items that can customize names, like containers or ships.  ---
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param characterId An EVE character ID
 * @param itemIds A list of item ids
 * @param optional nil or *PostCharactersCharacterIdAssetsNamesOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "Token" (optional.String) -  Access token to use if unable to set a header
     * @param "UserAgent" (optional.String) -  Client identifier, takes precedence over headers
     * @param "XUserAgent" (optional.String) -  Client identifier, takes precedence over User-Agent

@return []PostCharactersCharacterIdAssetsNames200Ok
*/

type PostCharactersCharacterIdAssetsNamesOpts struct {
	Datasource optional.String
	Token      optional.String
	UserAgent  optional.String
	XUserAgent optional.String
}

func (a *AssetsApiService) PostCharactersCharacterIdAssetsNames(ctx context.Context, characterId int32, itemIds []int64, localVarOptionals *PostCharactersCharacterIdAssetsNamesOpts) ([]PostCharactersCharacterIdAssetsNames200Ok, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue PostCharactersCharacterIdAssetsNames200OkList
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/characters/{character_id}/assets/names/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if characterId < 1 {
		return localVarReturnValue, nil, reportError("characterId must be greater than 1")
	}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarQueryParams.Add("token", parameterToString(localVarOptionals.Token.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserAgent.IsSet() {
		localVarQueryParams.Add("user_agent", parameterToString(localVarOptionals.UserAgent.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XUserAgent.IsSet() {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarOptionals.XUserAgent.Value(), "")
	}
	// body params
	localVarPostBody = &itemIds
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v []PostCharactersCharacterIdAssetsNames200Ok
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v Forbidden
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
AssetsApiService Get corporation asset locations
Return locations for a set of item ids, which you can get from corporation assets endpoint. Coordinates for items in hangars or stations are set to (0,0,0)  ---  Requires one of the following EVE corporation role(s): Director
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param corporationId An EVE corporation ID
 * @param itemIds A list of item ids
 * @param optional nil or *PostCorporationsCorporationIdAssetsLocationsOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "Token" (optional.String) -  Access token to use if unable to set a header
     * @param "UserAgent" (optional.String) -  Client identifier, takes precedence over headers
     * @param "XUserAgent" (optional.String) -  Client identifier, takes precedence over User-Agent

@return []PostCorporationsCorporationIdAssetsLocations200Ok
*/

type PostCorporationsCorporationIdAssetsLocationsOpts struct {
	Datasource optional.String
	Token      optional.String
	UserAgent  optional.String
	XUserAgent optional.String
}

func (a *AssetsApiService) PostCorporationsCorporationIdAssetsLocations(ctx context.Context, corporationId int32, itemIds []int64, localVarOptionals *PostCorporationsCorporationIdAssetsLocationsOpts) ([]PostCorporationsCorporationIdAssetsLocations200Ok, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue PostCorporationsCorporationIdAssetsLocations200OkList
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v2/corporations/{corporation_id}/assets/locations/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if corporationId < 1 {
		return localVarReturnValue, nil, reportError("corporationId must be greater than 1")
	}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarQueryParams.Add("token", parameterToString(localVarOptionals.Token.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserAgent.IsSet() {
		localVarQueryParams.Add("user_agent", parameterToString(localVarOptionals.UserAgent.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XUserAgent.IsSet() {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarOptionals.XUserAgent.Value(), "")
	}
	// body params
	localVarPostBody = &itemIds
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v []PostCorporationsCorporationIdAssetsLocations200Ok
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v Forbidden
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
AssetsApiService Get coporation asset names
Return names for a set of item ids, which you can get from corporation assets endpoint. Only valid for items that can customize names, like containers or ships.  ---  Requires one of the following EVE corporation role(s): Director
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param corporationId An EVE corporation ID
 * @param itemIds A list of item ids
 * @param optional nil or *PostCorporationsCorporationIdAssetsNamesOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "Token" (optional.String) -  Access token to use if unable to set a header
     * @param "UserAgent" (optional.String) -  Client identifier, takes precedence over headers
     * @param "XUserAgent" (optional.String) -  Client identifier, takes precedence over User-Agent

@return []PostCorporationsCorporationIdAssetsNames200Ok
*/

type PostCorporationsCorporationIdAssetsNamesOpts struct {
	Datasource optional.String
	Token      optional.String
	UserAgent  optional.String
	XUserAgent optional.String
}

func (a *AssetsApiService) PostCorporationsCorporationIdAssetsNames(ctx context.Context, corporationId int32, itemIds []int64, localVarOptionals *PostCorporationsCorporationIdAssetsNamesOpts) ([]PostCorporationsCorporationIdAssetsNames200Ok, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue PostCorporationsCorporationIdAssetsNames200OkList
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/corporations/{corporation_id}/assets/names/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if corporationId < 1 {
		return localVarReturnValue, nil, reportError("corporationId must be greater than 1")
	}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarQueryParams.Add("token", parameterToString(localVarOptionals.Token.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserAgent.IsSet() {
		localVarQueryParams.Add("user_agent", parameterToString(localVarOptionals.UserAgent.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XUserAgent.IsSet() {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarOptionals.XUserAgent.Value(), "")
	}
	// body params
	localVarPostBody = &itemIds
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v []PostCorporationsCorporationIdAssetsNames200Ok
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v Forbidden
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
