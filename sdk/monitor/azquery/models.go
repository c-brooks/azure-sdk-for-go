//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azquery

import "time"

// BatchQueryRequest - An single request in a batch.
type BatchQueryRequest struct {
	// REQUIRED; The Analytics query. Learn more about the Analytics query syntax [https://azure.microsoft.com/documentation/articles/app-insights-analytics-reference/]
	Body *Body `json:"body,omitempty"`

	// REQUIRED; Unique ID corresponding to each request in the batch
	CorrelationID *string `json:"id,omitempty"`

	// REQUIRED; Primary Workspace ID of the query
	WorkspaceID *string `json:"workspace,omitempty"`

	// Optional. Headers of the request. Can use prefer header to set server timeout, query statistics and visualization information.
	// For more information, see
	// https://pkg.go.dev/github.com/c-brooks/azure-sdk-for-go/sdk/monitor/azquery#readme-increase-wait-time-include-statistics-include-render-visualization
	Headers map[string]*string `json:"headers,omitempty"`

	// The method of a single request in a batch, defaults to POST
	Method *BatchQueryRequestMethod `json:"method,omitempty"`

	// The query path of a single request in a batch, defaults to /query
	Path *BatchQueryRequestPath `json:"path,omitempty"`
}

// BatchQueryResponse - Contains the batch query response and the headers, id, and status of the request
type BatchQueryResponse struct {
	// Contains the tables, columns & rows resulting from a query.
	Body          *BatchQueryResults `json:"body,omitempty"`
	CorrelationID *string            `json:"id,omitempty"`

	// Dictionary of
	Headers map[string]*string `json:"headers,omitempty"`
	Status  *int32             `json:"status,omitempty"`
}

// BatchQueryResults - Contains the tables, columns & rows resulting from a query.
type BatchQueryResults struct {
	// The code and message for an error.
	Error *ErrorInfo `json:"error,omitempty"`

	// Statistics represented in JSON format.
	Statistics []byte `json:"statistics,omitempty"`

	// The list of tables, columns and rows.
	Tables []*Table `json:"tables,omitempty"`

	// Visualization data in JSON format.
	Visualization []byte `json:"render,omitempty"`
}

// BatchRequest - An array of requests.
type BatchRequest struct {
	// REQUIRED; An single request in a batch.
	Requests []*BatchQueryRequest `json:"requests,omitempty"`
}

// BatchResponse - Response to a batch query.
type BatchResponse struct {
	// An array of responses corresponding to each individual request in a batch.
	Responses []*BatchQueryResponse `json:"responses,omitempty"`
}

// Body - The Analytics query. Learn more about the Analytics query syntax [https://azure.microsoft.com/documentation/articles/app-insights-analytics-reference/]
type Body struct {
	// REQUIRED; The query to execute.
	Query *string `json:"query,omitempty"`

	// A list of workspaces to query in addition to the primary workspace.
	AdditionalWorkspaces []*string `json:"workspaces,omitempty"`

	// Optional. The timespan over which to query data. This is an ISO8601 time period value. This timespan is applied in addition
	// to any that are specified in the query expression.
	Timespan *TimeInterval `json:"timespan,omitempty"`
}

// Column - A column in a table.
type Column struct {
	// The name of this column.
	Name *string `json:"name,omitempty"`

	// The data type of this column.
	Type *LogsColumnType `json:"type,omitempty"`
}

// LocalizableString - The localizable string class.
type LocalizableString struct {
	// REQUIRED; the invariant value.
	Value *string `json:"value,omitempty"`

	// the locale specific value.
	LocalizedValue *string `json:"localizedValue,omitempty"`
}

// LogsClientQueryBatchOptions contains the optional parameters for the LogsClient.QueryBatch method.
type LogsClientQueryBatchOptions struct {
	// placeholder for future optional parameters
}

// LogsClientQueryWorkspaceOptions contains the optional parameters for the LogsClient.QueryWorkspace method.
type LogsClientQueryWorkspaceOptions struct {
	// Optional. The prefer header to set server timeout, query statistics and visualization information.
	Options *LogsQueryOptions
}

// MetadataValue - Represents a metric metadata value.
type MetadataValue struct {
	// the name of the metadata.
	Name *LocalizableString `json:"name,omitempty"`

	// the value of the metadata.
	Value *string `json:"value,omitempty"`
}

// Metric - The result data of a query.
type Metric struct {
	// REQUIRED; the metric Id.
	ID *string `json:"id,omitempty"`

	// REQUIRED; the name and the display name of the metric, i.e. it is localizable string.
	Name *LocalizableString `json:"name,omitempty"`

	// REQUIRED; the time series returned when a data query is performed.
	Timeseries []*TimeSeriesElement `json:"timeseries,omitempty"`

	// REQUIRED; the resource type of the metric resource.
	Type *string `json:"type,omitempty"`

	// REQUIRED; The unit of the metric.
	Unit *MetricUnit `json:"unit,omitempty"`

	// Detailed description of this metric.
	DisplayDescription *string `json:"displayDescription,omitempty"`

	// 'Success' or the error details on query failures for this metric.
	ErrorCode *string `json:"errorCode,omitempty"`

	// Error message encountered querying this specific metric.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// MetricAvailability - Metric availability specifies the time grain (aggregation interval or frequency) and the retention
// period for that time grain.
type MetricAvailability struct {
	// the retention period for the metric at the specified timegrain. Expressed as a duration 'PT1M', 'P1D', etc.
	Retention *string `json:"retention,omitempty"`

	// the time grain specifies the aggregation interval for the metric. Expressed as a duration 'PT1M', 'P1D', etc.
	TimeGrain *string `json:"timeGrain,omitempty"`
}

// MetricDefinition - Metric definition class specifies the metadata for a metric.
type MetricDefinition struct {
	// Custom category name for this metric.
	Category *string `json:"category,omitempty"`

	// the name and the display name of the dimension, i.e. it is a localizable string.
	Dimensions []*LocalizableString `json:"dimensions,omitempty"`

	// Detailed description of this metric.
	DisplayDescription *string `json:"displayDescription,omitempty"`

	// the resource identifier of the metric definition.
	ID *string `json:"id,omitempty"`

	// Flag to indicate whether the dimension is required.
	IsDimensionRequired *bool `json:"isDimensionRequired,omitempty"`

	// the collection of what aggregation intervals are available to be queried.
	MetricAvailabilities []*MetricAvailability `json:"metricAvailabilities,omitempty"`

	// The class of the metric.
	MetricClass *MetricClass `json:"metricClass,omitempty"`

	// the name and the display name of the metric, i.e. it is a localizable string.
	Name *LocalizableString `json:"name,omitempty"`

	// the namespace the metric belongs to.
	Namespace *string `json:"namespace,omitempty"`

	// the primary aggregation type value defining how to use the values for display.
	PrimaryAggregationType *AggregationType `json:"primaryAggregationType,omitempty"`

	// the resource identifier of the resource that emitted the metric.
	ResourceID *string `json:"resourceId,omitempty"`

	// the collection of what aggregation types are supported.
	SupportedAggregationTypes []*AggregationType `json:"supportedAggregationTypes,omitempty"`

	// The unit of the metric.
	Unit *MetricUnit `json:"unit,omitempty"`
}

// MetricDefinitionCollection - Represents collection of metric definitions.
type MetricDefinitionCollection struct {
	// REQUIRED; the values for the metric definitions.
	Value []*MetricDefinition `json:"value,omitempty"`
}

// MetricNamespace - Metric namespace class specifies the metadata for a metric namespace.
type MetricNamespace struct {
	// Kind of namespace
	Classification *NamespaceClassification `json:"classification,omitempty"`

	// The ID of the metric namespace.
	ID *string `json:"id,omitempty"`

	// The escaped name of the namespace.
	Name *string `json:"name,omitempty"`

	// Properties which include the fully qualified namespace name.
	Properties *MetricNamespaceName `json:"properties,omitempty"`

	// The type of the namespace.
	Type *string `json:"type,omitempty"`
}

// MetricNamespaceCollection - Represents collection of metric namespaces.
type MetricNamespaceCollection struct {
	// REQUIRED; The values for the metric namespaces.
	Value []*MetricNamespace `json:"value,omitempty"`
}

// MetricNamespaceName - The fully qualified metric namespace name.
type MetricNamespaceName struct {
	// The metric namespace name.
	MetricNamespaceName *string `json:"metricNamespaceName,omitempty"`
}

// MetricValue - Represents a metric value.
type MetricValue struct {
	// REQUIRED; the timestamp for the metric value in ISO 8601 format.
	TimeStamp *time.Time `json:"timeStamp,omitempty"`

	// the average value in the time range.
	Average *float64 `json:"average,omitempty"`

	// the number of samples in the time range. Can be used to determine the number of values that contributed to the average
	// value.
	Count *float64 `json:"count,omitempty"`

	// the greatest value in the time range.
	Maximum *float64 `json:"maximum,omitempty"`

	// the least value in the time range.
	Minimum *float64 `json:"minimum,omitempty"`

	// the sum of all of the values in the time range.
	Total *float64 `json:"total,omitempty"`
}

// MetricsClientListDefinitionsOptions contains the optional parameters for the MetricsClient.ListDefinitions method.
type MetricsClientListDefinitionsOptions struct {
	// Metric namespace to query metric definitions for.
	Metricnamespace *string
}

// MetricsClientListNamespacesOptions contains the optional parameters for the MetricsClient.ListNamespaces method.
type MetricsClientListNamespacesOptions struct {
	// The ISO 8601 conform Date start time from which to query for metric namespaces.
	StartTime *string
}

// MetricsClientQueryResourceOptions contains the optional parameters for the MetricsClient.QueryResource method.
type MetricsClientQueryResourceOptions struct {
	// The list of aggregation types (comma separated) to retrieve.
	Aggregation *string
	// The $filter is used to reduce the set of metric data returned. Example: Metric contains metadata A, B and C. - Return all
	// time series of C where A = a1 and B = b1 or b2 $filter=A eq 'a1' and B eq 'b1'
	// or B eq 'b2' and C eq '' - Invalid variant: $filter=A eq 'a1' and B eq 'b1' and C eq '' or B = 'b2' This is invalid because
	// the logical or operator cannot separate two different metadata names. -
	// Return all time series where A = a1, B = b1 and C = c1: $filter=A eq 'a1' and B eq 'b1' and C eq 'c1' - Return all time
	// series where A = a1 $filter=A eq 'a1' and B eq '' and C eq ''. Special case:
	// When dimension name or dimension value uses round brackets. Eg: When dimension name is dim (test) 1 Instead of using $filter=
	// "dim (test) 1 eq '' " use $filter= "dim %2528test%2529 1 eq '' " When
	// dimension name is dim (test) 3 and dimension value is dim3 (test) val Instead of using $filter= "dim (test) 3 eq 'dim3
	// (test) val' " use $filter= "dim %2528test%2529 3 eq 'dim3 %2528test%2529 val' "
	Filter *string
	// The interval (i.e. timegrain) of the query.
	Interval *string
	// The names of the metrics (comma separated) to retrieve. Special case: If a metricname itself has a comma in it then use
	// %2 to indicate it. Eg: 'Metric,Name1' should be 'Metric%2Name1'
	Metricnames *string
	// Metric namespace to query metric definitions for.
	Metricnamespace *string
	// The aggregation to use for sorting results and the direction of the sort. Only one order can be specified. Examples: sum
	// asc.
	Orderby *string
	// Reduces the set of data collected. The syntax allowed depends on the operation. See the operation's description for details.
	ResultType *ResultType
	// The timespan of the query. It is a string with the following format 'startDateTimeISO/endDateTimeISO'.
	Timespan *TimeInterval
	// The maximum number of records to retrieve. Valid only if $filter is specified. Defaults to 10.
	Top *int32
}

// Response - The response to a metrics query.
type Response struct {
	// REQUIRED; The timespan for which the data was retrieved. Its value consists of two datetimes concatenated, separated by
	// '/'. This may be adjusted in the future and returned back from what was originally
	// requested.
	Timespan *TimeInterval `json:"timespan,omitempty"`

	// REQUIRED; the value of the collection.
	Value []*Metric `json:"value,omitempty"`

	// The integer value representing the relative cost of the query.
	Cost *int32 `json:"cost,omitempty"`

	// The interval (window size) for which the metric data was returned in. This may be adjusted in the future and returned back
	// from what was originally requested. This is not present if a metadata request
	// was made.
	Interval *string `json:"interval,omitempty"`

	// The namespace of the metrics being queried
	Namespace *string `json:"namespace,omitempty"`

	// The region of the resource being queried for metrics.
	Resourceregion *string `json:"resourceregion,omitempty"`
}

// Results - Contains the tables, columns & rows resulting from a query.
type Results struct {
	// REQUIRED; The list of tables, columns and rows.
	Tables []*Table `json:"tables,omitempty"`

	// The code and message for an error.
	Error *ErrorInfo `json:"error,omitempty"`

	// Statistics represented in JSON format.
	Statistics []byte `json:"statistics,omitempty"`

	// Visualization data in JSON format.
	Visualization []byte `json:"render,omitempty"`
}

// Table - Contains the columns and rows for one table in a query response.
type Table struct {
	// REQUIRED; The list of columns in this table.
	Columns []*Column `json:"columns,omitempty"`

	// REQUIRED; The name of the table.
	Name *string `json:"name,omitempty"`

	// REQUIRED; The resulting rows from this query.
	Rows []Row `json:"rows,omitempty"`
}

// TimeSeriesElement - A time series result type. The discriminator value is always TimeSeries in this case.
type TimeSeriesElement struct {
	// An array of data points representing the metric values. This is only returned if a result type of data is specified.
	Data []*MetricValue `json:"data,omitempty"`

	// the metadata values returned if $filter was specified in the call.
	Metadatavalues []*MetadataValue `json:"metadatavalues,omitempty"`
}
