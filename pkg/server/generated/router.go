// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package generated

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /api/v1/applications)
	GetApiV1Applications(w http.ResponseWriter, r *http.Request)

	// (GET /api/v1/clusters)
	GetApiV1Clusters(w http.ResponseWriter, r *http.Request)

	// (GET /api/v1/controlplanes)
	GetApiV1Controlplanes(w http.ResponseWriter, r *http.Request)

	// (DELETE /api/v1/organization)
	DeleteApiV1Organization(w http.ResponseWriter, r *http.Request)

	// (POST /api/v1/organization)
	PostApiV1Organization(w http.ResponseWriter, r *http.Request)

	// (GET /api/v1/projects)
	GetApiV1Projects(w http.ResponseWriter, r *http.Request)

	// (POST /api/v1/projects)
	PostApiV1Projects(w http.ResponseWriter, r *http.Request)

	// (DELETE /api/v1/projects/{projectName})
	DeleteApiV1ProjectsProjectName(w http.ResponseWriter, r *http.Request, projectName ProjectNameParameter)

	// (POST /api/v1/projects/{projectName}/controlplanes)
	PostApiV1ProjectsProjectNameControlplanes(w http.ResponseWriter, r *http.Request, projectName ProjectNameParameter)

	// (DELETE /api/v1/projects/{projectName}/controlplanes/{controlPlaneName})
	DeleteApiV1ProjectsProjectNameControlplanesControlPlaneName(w http.ResponseWriter, r *http.Request, projectName ProjectNameParameter, controlPlaneName ControlPlaneNameParameter)

	// (PUT /api/v1/projects/{projectName}/controlplanes/{controlPlaneName})
	PutApiV1ProjectsProjectNameControlplanesControlPlaneName(w http.ResponseWriter, r *http.Request, projectName ProjectNameParameter, controlPlaneName ControlPlaneNameParameter)

	// (POST /api/v1/projects/{projectName}/controlplanes/{controlPlaneName}/clusters)
	PostApiV1ProjectsProjectNameControlplanesControlPlaneNameClusters(w http.ResponseWriter, r *http.Request, projectName ProjectNameParameter, controlPlaneName ControlPlaneNameParameter)

	// (DELETE /api/v1/projects/{projectName}/controlplanes/{controlPlaneName}/clusters/{clusterName})
	DeleteApiV1ProjectsProjectNameControlplanesControlPlaneNameClustersClusterName(w http.ResponseWriter, r *http.Request, projectName ProjectNameParameter, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter)

	// (PUT /api/v1/projects/{projectName}/controlplanes/{controlPlaneName}/clusters/{clusterName})
	PutApiV1ProjectsProjectNameControlplanesControlPlaneNameClustersClusterName(w http.ResponseWriter, r *http.Request, projectName ProjectNameParameter, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter)

	// (GET /api/v1/projects/{projectName}/controlplanes/{controlPlaneName}/clusters/{clusterName}/kubeconfig)
	GetApiV1ProjectsProjectNameControlplanesControlPlaneNameClustersClusterNameKubeconfig(w http.ResponseWriter, r *http.Request, projectName ProjectNameParameter, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter)

	// (GET /api/v1/regions)
	GetApiV1Regions(w http.ResponseWriter, r *http.Request)

	// (GET /api/v1/regions/{regionName}/flavors)
	GetApiV1RegionsRegionNameFlavors(w http.ResponseWriter, r *http.Request, regionName RegionNameParameter)

	// (GET /api/v1/regions/{regionName}/images)
	GetApiV1RegionsRegionNameImages(w http.ResponseWriter, r *http.Request, regionName RegionNameParameter)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetApiV1Applications operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1Applications(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1Applications(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1Clusters operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1Clusters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1Clusters(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1Controlplanes operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1Controlplanes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1Controlplanes(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteApiV1Organization operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiV1Organization(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteApiV1Organization(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostApiV1Organization operation middleware
func (siw *ServerInterfaceWrapper) PostApiV1Organization(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiV1Organization(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1Projects operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1Projects(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1Projects(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostApiV1Projects operation middleware
func (siw *ServerInterfaceWrapper) PostApiV1Projects(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiV1Projects(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteApiV1ProjectsProjectName operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiV1ProjectsProjectName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "projectName" -------------
	var projectName ProjectNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectName", runtime.ParamLocationPath, chi.URLParam(r, "projectName"), &projectName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "projectName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteApiV1ProjectsProjectName(w, r, projectName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostApiV1ProjectsProjectNameControlplanes operation middleware
func (siw *ServerInterfaceWrapper) PostApiV1ProjectsProjectNameControlplanes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "projectName" -------------
	var projectName ProjectNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectName", runtime.ParamLocationPath, chi.URLParam(r, "projectName"), &projectName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "projectName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiV1ProjectsProjectNameControlplanes(w, r, projectName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteApiV1ProjectsProjectNameControlplanesControlPlaneName operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiV1ProjectsProjectNameControlplanesControlPlaneName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "projectName" -------------
	var projectName ProjectNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectName", runtime.ParamLocationPath, chi.URLParam(r, "projectName"), &projectName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "projectName", Err: err})
		return
	}

	// ------------- Path parameter "controlPlaneName" -------------
	var controlPlaneName ControlPlaneNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, chi.URLParam(r, "controlPlaneName"), &controlPlaneName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "controlPlaneName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteApiV1ProjectsProjectNameControlplanesControlPlaneName(w, r, projectName, controlPlaneName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutApiV1ProjectsProjectNameControlplanesControlPlaneName operation middleware
func (siw *ServerInterfaceWrapper) PutApiV1ProjectsProjectNameControlplanesControlPlaneName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "projectName" -------------
	var projectName ProjectNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectName", runtime.ParamLocationPath, chi.URLParam(r, "projectName"), &projectName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "projectName", Err: err})
		return
	}

	// ------------- Path parameter "controlPlaneName" -------------
	var controlPlaneName ControlPlaneNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, chi.URLParam(r, "controlPlaneName"), &controlPlaneName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "controlPlaneName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutApiV1ProjectsProjectNameControlplanesControlPlaneName(w, r, projectName, controlPlaneName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostApiV1ProjectsProjectNameControlplanesControlPlaneNameClusters operation middleware
func (siw *ServerInterfaceWrapper) PostApiV1ProjectsProjectNameControlplanesControlPlaneNameClusters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "projectName" -------------
	var projectName ProjectNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectName", runtime.ParamLocationPath, chi.URLParam(r, "projectName"), &projectName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "projectName", Err: err})
		return
	}

	// ------------- Path parameter "controlPlaneName" -------------
	var controlPlaneName ControlPlaneNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, chi.URLParam(r, "controlPlaneName"), &controlPlaneName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "controlPlaneName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiV1ProjectsProjectNameControlplanesControlPlaneNameClusters(w, r, projectName, controlPlaneName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteApiV1ProjectsProjectNameControlplanesControlPlaneNameClustersClusterName operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiV1ProjectsProjectNameControlplanesControlPlaneNameClustersClusterName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "projectName" -------------
	var projectName ProjectNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectName", runtime.ParamLocationPath, chi.URLParam(r, "projectName"), &projectName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "projectName", Err: err})
		return
	}

	// ------------- Path parameter "controlPlaneName" -------------
	var controlPlaneName ControlPlaneNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, chi.URLParam(r, "controlPlaneName"), &controlPlaneName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "controlPlaneName", Err: err})
		return
	}

	// ------------- Path parameter "clusterName" -------------
	var clusterName ClusterNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, chi.URLParam(r, "clusterName"), &clusterName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "clusterName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteApiV1ProjectsProjectNameControlplanesControlPlaneNameClustersClusterName(w, r, projectName, controlPlaneName, clusterName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutApiV1ProjectsProjectNameControlplanesControlPlaneNameClustersClusterName operation middleware
func (siw *ServerInterfaceWrapper) PutApiV1ProjectsProjectNameControlplanesControlPlaneNameClustersClusterName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "projectName" -------------
	var projectName ProjectNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectName", runtime.ParamLocationPath, chi.URLParam(r, "projectName"), &projectName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "projectName", Err: err})
		return
	}

	// ------------- Path parameter "controlPlaneName" -------------
	var controlPlaneName ControlPlaneNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, chi.URLParam(r, "controlPlaneName"), &controlPlaneName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "controlPlaneName", Err: err})
		return
	}

	// ------------- Path parameter "clusterName" -------------
	var clusterName ClusterNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, chi.URLParam(r, "clusterName"), &clusterName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "clusterName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutApiV1ProjectsProjectNameControlplanesControlPlaneNameClustersClusterName(w, r, projectName, controlPlaneName, clusterName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1ProjectsProjectNameControlplanesControlPlaneNameClustersClusterNameKubeconfig operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1ProjectsProjectNameControlplanesControlPlaneNameClustersClusterNameKubeconfig(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "projectName" -------------
	var projectName ProjectNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectName", runtime.ParamLocationPath, chi.URLParam(r, "projectName"), &projectName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "projectName", Err: err})
		return
	}

	// ------------- Path parameter "controlPlaneName" -------------
	var controlPlaneName ControlPlaneNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, chi.URLParam(r, "controlPlaneName"), &controlPlaneName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "controlPlaneName", Err: err})
		return
	}

	// ------------- Path parameter "clusterName" -------------
	var clusterName ClusterNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, chi.URLParam(r, "clusterName"), &clusterName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "clusterName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1ProjectsProjectNameControlplanesControlPlaneNameClustersClusterNameKubeconfig(w, r, projectName, controlPlaneName, clusterName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1Regions operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1Regions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1Regions(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1RegionsRegionNameFlavors operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1RegionsRegionNameFlavors(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "regionName" -------------
	var regionName RegionNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "regionName", runtime.ParamLocationPath, chi.URLParam(r, "regionName"), &regionName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "regionName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1RegionsRegionNameFlavors(w, r, regionName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1RegionsRegionNameImages operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1RegionsRegionNameImages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "regionName" -------------
	var regionName RegionNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "regionName", runtime.ParamLocationPath, chi.URLParam(r, "regionName"), &regionName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "regionName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1RegionsRegionNameImages(w, r, regionName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/applications", wrapper.GetApiV1Applications)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/clusters", wrapper.GetApiV1Clusters)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/controlplanes", wrapper.GetApiV1Controlplanes)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/v1/organization", wrapper.DeleteApiV1Organization)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/v1/organization", wrapper.PostApiV1Organization)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/projects", wrapper.GetApiV1Projects)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/v1/projects", wrapper.PostApiV1Projects)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/v1/projects/{projectName}", wrapper.DeleteApiV1ProjectsProjectName)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/v1/projects/{projectName}/controlplanes", wrapper.PostApiV1ProjectsProjectNameControlplanes)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/v1/projects/{projectName}/controlplanes/{controlPlaneName}", wrapper.DeleteApiV1ProjectsProjectNameControlplanesControlPlaneName)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/api/v1/projects/{projectName}/controlplanes/{controlPlaneName}", wrapper.PutApiV1ProjectsProjectNameControlplanesControlPlaneName)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/v1/projects/{projectName}/controlplanes/{controlPlaneName}/clusters", wrapper.PostApiV1ProjectsProjectNameControlplanesControlPlaneNameClusters)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/v1/projects/{projectName}/controlplanes/{controlPlaneName}/clusters/{clusterName}", wrapper.DeleteApiV1ProjectsProjectNameControlplanesControlPlaneNameClustersClusterName)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/api/v1/projects/{projectName}/controlplanes/{controlPlaneName}/clusters/{clusterName}", wrapper.PutApiV1ProjectsProjectNameControlplanesControlPlaneNameClustersClusterName)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/projects/{projectName}/controlplanes/{controlPlaneName}/clusters/{clusterName}/kubeconfig", wrapper.GetApiV1ProjectsProjectNameControlplanesControlPlaneNameClustersClusterNameKubeconfig)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/regions", wrapper.GetApiV1Regions)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/regions/{regionName}/flavors", wrapper.GetApiV1RegionsRegionNameFlavors)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/regions/{regionName}/images", wrapper.GetApiV1RegionsRegionNameImages)
	})

	return r
}
