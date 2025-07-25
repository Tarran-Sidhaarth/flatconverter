// DO NOT EDIT!
// cpp generated by Buffman 💪
// Versions:
// 		Buffman: 1.0.0
// 		Flatc: v25.2.10


#ifndef FLATBUFFERS_GENERATED_STATUS_COMMON_H_
#define FLATBUFFERS_GENERATED_STATUS_COMMON_H_

#include "flatbuffers/flatbuffers.h"

// Ensure the included flatbuffers.h is the same version as when this file was
// generated, otherwise it may not be compatible.
static_assert(FLATBUFFERS_VERSION_MAJOR == 25 &&
              FLATBUFFERS_VERSION_MINOR == 2 &&
              FLATBUFFERS_VERSION_REVISION == 10,
             "Non-compatible flatbuffers version included");

namespace common {

enum Status : int32_t {
  Status_STATUS_UNKNOWN = 0,
  Status_STATUS_ACTIVE = 1,
  Status_STATUS_INACTIVE = 2,
  Status_STATUS_PENDING = 3,
  Status_STATUS_SUSPENDED = 4,
  Status_MIN = Status_STATUS_UNKNOWN,
  Status_MAX = Status_STATUS_SUSPENDED
};

inline const Status (&EnumValuesStatus())[5] {
  static const Status values[] = {
    Status_STATUS_UNKNOWN,
    Status_STATUS_ACTIVE,
    Status_STATUS_INACTIVE,
    Status_STATUS_PENDING,
    Status_STATUS_SUSPENDED
  };
  return values;
}

inline const char * const *EnumNamesStatus() {
  static const char * const names[6] = {
    "STATUS_UNKNOWN",
    "STATUS_ACTIVE",
    "STATUS_INACTIVE",
    "STATUS_PENDING",
    "STATUS_SUSPENDED",
    nullptr
  };
  return names;
}

inline const char *EnumNameStatus(Status e) {
  if (::flatbuffers::IsOutRange(e, Status_STATUS_UNKNOWN, Status_STATUS_SUSPENDED)) return "";
  const size_t index = static_cast<size_t>(e);
  return EnumNamesStatus()[index];
}

}  // namespace common

#endif  // FLATBUFFERS_GENERATED_STATUS_COMMON_H_
