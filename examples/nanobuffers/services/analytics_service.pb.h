/* Automatically generated nanopb header */
/* Generated by nanopb-0.4.9.1 */

#ifndef PB_SERVICES_SERVICES_ANALYTICS_SERVICE_PB_H_INCLUDED
#define PB_SERVICES_SERVICES_ANALYTICS_SERVICE_PB_H_INCLUDED
#include <pb.h>
#include "common/timestamp.pb.h"
#include "services/user_service.pb.h"
#include "services/notification_service.pb.h"

#if PB_PROTO_HEADER_VERSION != 40
#error Regenerate this file with the current version of nanopb generator.
#endif

/* Struct definitions */
typedef struct _services_UserActivity {
    pb_callback_t activity_id;
    pb_callback_t user_id;
    pb_callback_t action;
    pb_callback_t resource;
    bool has_timestamp;
    common_Timestamp timestamp;
    pb_callback_t properties;
    pb_callback_t session_id;
} services_UserActivity;

typedef struct _services_UserActivity_PropertiesEntry {
    pb_callback_t key;
    pb_callback_t value;
} services_UserActivity_PropertiesEntry;

typedef struct _services_NotificationMetrics {
    pb_callback_t notification_id;
    services_NotificationType type;
    bool delivered;
    bool opened;
    bool clicked;
    bool has_delivered_at;
    common_Timestamp delivered_at;
    bool has_opened_at;
    common_Timestamp opened_at;
    bool has_clicked_at;
    common_Timestamp clicked_at;
} services_NotificationMetrics;

typedef struct _services_UserAnalyticsRequest {
    pb_callback_t user_id;
    bool has_start_date;
    common_Timestamp start_date;
    bool has_end_date;
    common_Timestamp end_date;
} services_UserAnalyticsRequest;

typedef struct _services_UserAnalyticsResponse {
    bool has_user;
    services_User user; /* Dependency on user_service.proto */
    pb_callback_t activities;
    pb_callback_t notification_metrics;
    int32_t total_activities;
    int32_t total_notifications_sent;
    int32_t total_notifications_opened;
} services_UserAnalyticsResponse;

typedef struct _services_SystemMetricsRequest {
    bool has_start_date;
    common_Timestamp start_date;
    bool has_end_date;
    common_Timestamp end_date;
} services_SystemMetricsRequest;

typedef struct _services_SystemMetricsResponse {
    int32_t total_users;
    int32_t active_users;
    int32_t total_notifications_sent;
    double notification_open_rate;
    pb_callback_t activity_counts;
    bool has_generated_at;
    common_Timestamp generated_at;
} services_SystemMetricsResponse;

typedef struct _services_SystemMetricsResponse_ActivityCountsEntry {
    pb_callback_t key;
    int32_t value;
} services_SystemMetricsResponse_ActivityCountsEntry;

typedef struct _services_TrackUserActivityResponse {
    bool success;
    pb_callback_t message;
} services_TrackUserActivityResponse;


#ifdef __cplusplus
extern "C" {
#endif

/* Initializer values for message structs */
#define services_UserActivity_init_default       {{{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}, false, common_Timestamp_init_default, {{NULL}, NULL}, {{NULL}, NULL}}
#define services_UserActivity_PropertiesEntry_init_default {{{NULL}, NULL}, {{NULL}, NULL}}
#define services_NotificationMetrics_init_default {{{NULL}, NULL}, _services_NotificationType_MIN, 0, 0, 0, false, common_Timestamp_init_default, false, common_Timestamp_init_default, false, common_Timestamp_init_default}
#define services_UserAnalyticsRequest_init_default {{{NULL}, NULL}, false, common_Timestamp_init_default, false, common_Timestamp_init_default}
#define services_UserAnalyticsResponse_init_default {false, services_User_init_default, {{NULL}, NULL}, {{NULL}, NULL}, 0, 0, 0}
#define services_SystemMetricsRequest_init_default {false, common_Timestamp_init_default, false, common_Timestamp_init_default}
#define services_SystemMetricsResponse_init_default {0, 0, 0, 0, {{NULL}, NULL}, false, common_Timestamp_init_default}
#define services_SystemMetricsResponse_ActivityCountsEntry_init_default {{{NULL}, NULL}, 0}
#define services_TrackUserActivityResponse_init_default {0, {{NULL}, NULL}}
#define services_UserActivity_init_zero          {{{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}, false, common_Timestamp_init_zero, {{NULL}, NULL}, {{NULL}, NULL}}
#define services_UserActivity_PropertiesEntry_init_zero {{{NULL}, NULL}, {{NULL}, NULL}}
#define services_NotificationMetrics_init_zero   {{{NULL}, NULL}, _services_NotificationType_MIN, 0, 0, 0, false, common_Timestamp_init_zero, false, common_Timestamp_init_zero, false, common_Timestamp_init_zero}
#define services_UserAnalyticsRequest_init_zero  {{{NULL}, NULL}, false, common_Timestamp_init_zero, false, common_Timestamp_init_zero}
#define services_UserAnalyticsResponse_init_zero {false, services_User_init_zero, {{NULL}, NULL}, {{NULL}, NULL}, 0, 0, 0}
#define services_SystemMetricsRequest_init_zero  {false, common_Timestamp_init_zero, false, common_Timestamp_init_zero}
#define services_SystemMetricsResponse_init_zero {0, 0, 0, 0, {{NULL}, NULL}, false, common_Timestamp_init_zero}
#define services_SystemMetricsResponse_ActivityCountsEntry_init_zero {{{NULL}, NULL}, 0}
#define services_TrackUserActivityResponse_init_zero {0, {{NULL}, NULL}}

/* Field tags (for use in manual encoding/decoding) */
#define services_UserActivity_activity_id_tag    1
#define services_UserActivity_user_id_tag        2
#define services_UserActivity_action_tag         3
#define services_UserActivity_resource_tag       4
#define services_UserActivity_timestamp_tag      5
#define services_UserActivity_properties_tag     6
#define services_UserActivity_session_id_tag     7
#define services_UserActivity_PropertiesEntry_key_tag 1
#define services_UserActivity_PropertiesEntry_value_tag 2
#define services_NotificationMetrics_notification_id_tag 1
#define services_NotificationMetrics_type_tag    2
#define services_NotificationMetrics_delivered_tag 3
#define services_NotificationMetrics_opened_tag  4
#define services_NotificationMetrics_clicked_tag 5
#define services_NotificationMetrics_delivered_at_tag 6
#define services_NotificationMetrics_opened_at_tag 7
#define services_NotificationMetrics_clicked_at_tag 8
#define services_UserAnalyticsRequest_user_id_tag 1
#define services_UserAnalyticsRequest_start_date_tag 2
#define services_UserAnalyticsRequest_end_date_tag 3
#define services_UserAnalyticsResponse_user_tag  1
#define services_UserAnalyticsResponse_activities_tag 2
#define services_UserAnalyticsResponse_notification_metrics_tag 3
#define services_UserAnalyticsResponse_total_activities_tag 4
#define services_UserAnalyticsResponse_total_notifications_sent_tag 5
#define services_UserAnalyticsResponse_total_notifications_opened_tag 6
#define services_SystemMetricsRequest_start_date_tag 1
#define services_SystemMetricsRequest_end_date_tag 2
#define services_SystemMetricsResponse_total_users_tag 1
#define services_SystemMetricsResponse_active_users_tag 2
#define services_SystemMetricsResponse_total_notifications_sent_tag 3
#define services_SystemMetricsResponse_notification_open_rate_tag 4
#define services_SystemMetricsResponse_activity_counts_tag 5
#define services_SystemMetricsResponse_generated_at_tag 6
#define services_SystemMetricsResponse_ActivityCountsEntry_key_tag 1
#define services_SystemMetricsResponse_ActivityCountsEntry_value_tag 2
#define services_TrackUserActivityResponse_success_tag 1
#define services_TrackUserActivityResponse_message_tag 2

/* Struct field encoding specification for nanopb */
#define services_UserActivity_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   activity_id,       1) \
X(a, CALLBACK, SINGULAR, STRING,   user_id,           2) \
X(a, CALLBACK, SINGULAR, STRING,   action,            3) \
X(a, CALLBACK, SINGULAR, STRING,   resource,          4) \
X(a, STATIC,   OPTIONAL, MESSAGE,  timestamp,         5) \
X(a, CALLBACK, REPEATED, MESSAGE,  properties,        6) \
X(a, CALLBACK, SINGULAR, STRING,   session_id,        7)
#define services_UserActivity_CALLBACK pb_default_field_callback
#define services_UserActivity_DEFAULT NULL
#define services_UserActivity_timestamp_MSGTYPE common_Timestamp
#define services_UserActivity_properties_MSGTYPE services_UserActivity_PropertiesEntry

#define services_UserActivity_PropertiesEntry_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   key,               1) \
X(a, CALLBACK, SINGULAR, STRING,   value,             2)
#define services_UserActivity_PropertiesEntry_CALLBACK pb_default_field_callback
#define services_UserActivity_PropertiesEntry_DEFAULT NULL

#define services_NotificationMetrics_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   notification_id,   1) \
X(a, STATIC,   SINGULAR, UENUM,    type,              2) \
X(a, STATIC,   SINGULAR, BOOL,     delivered,         3) \
X(a, STATIC,   SINGULAR, BOOL,     opened,            4) \
X(a, STATIC,   SINGULAR, BOOL,     clicked,           5) \
X(a, STATIC,   OPTIONAL, MESSAGE,  delivered_at,      6) \
X(a, STATIC,   OPTIONAL, MESSAGE,  opened_at,         7) \
X(a, STATIC,   OPTIONAL, MESSAGE,  clicked_at,        8)
#define services_NotificationMetrics_CALLBACK pb_default_field_callback
#define services_NotificationMetrics_DEFAULT NULL
#define services_NotificationMetrics_delivered_at_MSGTYPE common_Timestamp
#define services_NotificationMetrics_opened_at_MSGTYPE common_Timestamp
#define services_NotificationMetrics_clicked_at_MSGTYPE common_Timestamp

#define services_UserAnalyticsRequest_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   user_id,           1) \
X(a, STATIC,   OPTIONAL, MESSAGE,  start_date,        2) \
X(a, STATIC,   OPTIONAL, MESSAGE,  end_date,          3)
#define services_UserAnalyticsRequest_CALLBACK pb_default_field_callback
#define services_UserAnalyticsRequest_DEFAULT NULL
#define services_UserAnalyticsRequest_start_date_MSGTYPE common_Timestamp
#define services_UserAnalyticsRequest_end_date_MSGTYPE common_Timestamp

#define services_UserAnalyticsResponse_FIELDLIST(X, a) \
X(a, STATIC,   OPTIONAL, MESSAGE,  user,              1) \
X(a, CALLBACK, REPEATED, MESSAGE,  activities,        2) \
X(a, CALLBACK, REPEATED, MESSAGE,  notification_metrics,   3) \
X(a, STATIC,   SINGULAR, INT32,    total_activities,   4) \
X(a, STATIC,   SINGULAR, INT32,    total_notifications_sent,   5) \
X(a, STATIC,   SINGULAR, INT32,    total_notifications_opened,   6)
#define services_UserAnalyticsResponse_CALLBACK pb_default_field_callback
#define services_UserAnalyticsResponse_DEFAULT NULL
#define services_UserAnalyticsResponse_user_MSGTYPE services_User
#define services_UserAnalyticsResponse_activities_MSGTYPE services_UserActivity
#define services_UserAnalyticsResponse_notification_metrics_MSGTYPE services_NotificationMetrics

#define services_SystemMetricsRequest_FIELDLIST(X, a) \
X(a, STATIC,   OPTIONAL, MESSAGE,  start_date,        1) \
X(a, STATIC,   OPTIONAL, MESSAGE,  end_date,          2)
#define services_SystemMetricsRequest_CALLBACK NULL
#define services_SystemMetricsRequest_DEFAULT NULL
#define services_SystemMetricsRequest_start_date_MSGTYPE common_Timestamp
#define services_SystemMetricsRequest_end_date_MSGTYPE common_Timestamp

#define services_SystemMetricsResponse_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, INT32,    total_users,       1) \
X(a, STATIC,   SINGULAR, INT32,    active_users,      2) \
X(a, STATIC,   SINGULAR, INT32,    total_notifications_sent,   3) \
X(a, STATIC,   SINGULAR, DOUBLE,   notification_open_rate,   4) \
X(a, CALLBACK, REPEATED, MESSAGE,  activity_counts,   5) \
X(a, STATIC,   OPTIONAL, MESSAGE,  generated_at,      6)
#define services_SystemMetricsResponse_CALLBACK pb_default_field_callback
#define services_SystemMetricsResponse_DEFAULT NULL
#define services_SystemMetricsResponse_activity_counts_MSGTYPE services_SystemMetricsResponse_ActivityCountsEntry
#define services_SystemMetricsResponse_generated_at_MSGTYPE common_Timestamp

#define services_SystemMetricsResponse_ActivityCountsEntry_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   key,               1) \
X(a, STATIC,   SINGULAR, INT32,    value,             2)
#define services_SystemMetricsResponse_ActivityCountsEntry_CALLBACK pb_default_field_callback
#define services_SystemMetricsResponse_ActivityCountsEntry_DEFAULT NULL

#define services_TrackUserActivityResponse_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, BOOL,     success,           1) \
X(a, CALLBACK, SINGULAR, STRING,   message,           2)
#define services_TrackUserActivityResponse_CALLBACK pb_default_field_callback
#define services_TrackUserActivityResponse_DEFAULT NULL

extern const pb_msgdesc_t services_UserActivity_msg;
extern const pb_msgdesc_t services_UserActivity_PropertiesEntry_msg;
extern const pb_msgdesc_t services_NotificationMetrics_msg;
extern const pb_msgdesc_t services_UserAnalyticsRequest_msg;
extern const pb_msgdesc_t services_UserAnalyticsResponse_msg;
extern const pb_msgdesc_t services_SystemMetricsRequest_msg;
extern const pb_msgdesc_t services_SystemMetricsResponse_msg;
extern const pb_msgdesc_t services_SystemMetricsResponse_ActivityCountsEntry_msg;
extern const pb_msgdesc_t services_TrackUserActivityResponse_msg;

/* Defines for backwards compatibility with code written before nanopb-0.4.0 */
#define services_UserActivity_fields &services_UserActivity_msg
#define services_UserActivity_PropertiesEntry_fields &services_UserActivity_PropertiesEntry_msg
#define services_NotificationMetrics_fields &services_NotificationMetrics_msg
#define services_UserAnalyticsRequest_fields &services_UserAnalyticsRequest_msg
#define services_UserAnalyticsResponse_fields &services_UserAnalyticsResponse_msg
#define services_SystemMetricsRequest_fields &services_SystemMetricsRequest_msg
#define services_SystemMetricsResponse_fields &services_SystemMetricsResponse_msg
#define services_SystemMetricsResponse_ActivityCountsEntry_fields &services_SystemMetricsResponse_ActivityCountsEntry_msg
#define services_TrackUserActivityResponse_fields &services_TrackUserActivityResponse_msg

/* Maximum encoded size of messages (where known) */
/* services_UserActivity_size depends on runtime parameters */
/* services_UserActivity_PropertiesEntry_size depends on runtime parameters */
/* services_NotificationMetrics_size depends on runtime parameters */
/* services_UserAnalyticsRequest_size depends on runtime parameters */
/* services_UserAnalyticsResponse_size depends on runtime parameters */
/* services_SystemMetricsResponse_size depends on runtime parameters */
/* services_SystemMetricsResponse_ActivityCountsEntry_size depends on runtime parameters */
/* services_TrackUserActivityResponse_size depends on runtime parameters */
#define SERVICES_SERVICES_ANALYTICS_SERVICE_PB_H_MAX_SIZE services_SystemMetricsRequest_size
#define services_SystemMetricsRequest_size       48

#ifdef __cplusplus
} /* extern "C" */
#endif

#endif
