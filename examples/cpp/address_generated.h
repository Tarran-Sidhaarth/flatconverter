// DO NOT EDIT!
// cpp generated by Buffman 💪
// Versions:
// 		Buffman: 1.0.0
// 		Flatc: v25.2.10


#ifndef FLATBUFFERS_GENERATED_ADDRESS_COMMON_H_
#define FLATBUFFERS_GENERATED_ADDRESS_COMMON_H_

#include "flatbuffers/flatbuffers.h"

// Ensure the included flatbuffers.h is the same version as when this file was
// generated, otherwise it may not be compatible.
static_assert(FLATBUFFERS_VERSION_MAJOR == 25 &&
              FLATBUFFERS_VERSION_MINOR == 2 &&
              FLATBUFFERS_VERSION_REVISION == 10,
             "Non-compatible flatbuffers version included");

namespace common {

struct Address;
struct AddressBuilder;

struct Address FLATBUFFERS_FINAL_CLASS : private ::flatbuffers::Table {
  typedef AddressBuilder Builder;
  enum FlatBuffersVTableOffset FLATBUFFERS_VTABLE_UNDERLYING_TYPE {
    VT_STREET = 4,
    VT_CITY = 6,
    VT_STATE = 8,
    VT_POSTAL_CODE = 10,
    VT_COUNTRY = 12
  };
  const ::flatbuffers::String *street() const {
    return GetPointer<const ::flatbuffers::String *>(VT_STREET);
  }
  const ::flatbuffers::String *city() const {
    return GetPointer<const ::flatbuffers::String *>(VT_CITY);
  }
  const ::flatbuffers::String *state() const {
    return GetPointer<const ::flatbuffers::String *>(VT_STATE);
  }
  const ::flatbuffers::String *postal_code() const {
    return GetPointer<const ::flatbuffers::String *>(VT_POSTAL_CODE);
  }
  const ::flatbuffers::String *country() const {
    return GetPointer<const ::flatbuffers::String *>(VT_COUNTRY);
  }
  bool Verify(::flatbuffers::Verifier &verifier) const {
    return VerifyTableStart(verifier) &&
           VerifyOffset(verifier, VT_STREET) &&
           verifier.VerifyString(street()) &&
           VerifyOffset(verifier, VT_CITY) &&
           verifier.VerifyString(city()) &&
           VerifyOffset(verifier, VT_STATE) &&
           verifier.VerifyString(state()) &&
           VerifyOffset(verifier, VT_POSTAL_CODE) &&
           verifier.VerifyString(postal_code()) &&
           VerifyOffset(verifier, VT_COUNTRY) &&
           verifier.VerifyString(country()) &&
           verifier.EndTable();
  }
};

struct AddressBuilder {
  typedef Address Table;
  ::flatbuffers::FlatBufferBuilder &fbb_;
  ::flatbuffers::uoffset_t start_;
  void add_street(::flatbuffers::Offset<::flatbuffers::String> street) {
    fbb_.AddOffset(Address::VT_STREET, street);
  }
  void add_city(::flatbuffers::Offset<::flatbuffers::String> city) {
    fbb_.AddOffset(Address::VT_CITY, city);
  }
  void add_state(::flatbuffers::Offset<::flatbuffers::String> state) {
    fbb_.AddOffset(Address::VT_STATE, state);
  }
  void add_postal_code(::flatbuffers::Offset<::flatbuffers::String> postal_code) {
    fbb_.AddOffset(Address::VT_POSTAL_CODE, postal_code);
  }
  void add_country(::flatbuffers::Offset<::flatbuffers::String> country) {
    fbb_.AddOffset(Address::VT_COUNTRY, country);
  }
  explicit AddressBuilder(::flatbuffers::FlatBufferBuilder &_fbb)
        : fbb_(_fbb) {
    start_ = fbb_.StartTable();
  }
  ::flatbuffers::Offset<Address> Finish() {
    const auto end = fbb_.EndTable(start_);
    auto o = ::flatbuffers::Offset<Address>(end);
    return o;
  }
};

inline ::flatbuffers::Offset<Address> CreateAddress(
    ::flatbuffers::FlatBufferBuilder &_fbb,
    ::flatbuffers::Offset<::flatbuffers::String> street = 0,
    ::flatbuffers::Offset<::flatbuffers::String> city = 0,
    ::flatbuffers::Offset<::flatbuffers::String> state = 0,
    ::flatbuffers::Offset<::flatbuffers::String> postal_code = 0,
    ::flatbuffers::Offset<::flatbuffers::String> country = 0) {
  AddressBuilder builder_(_fbb);
  builder_.add_country(country);
  builder_.add_postal_code(postal_code);
  builder_.add_state(state);
  builder_.add_city(city);
  builder_.add_street(street);
  return builder_.Finish();
}

inline ::flatbuffers::Offset<Address> CreateAddressDirect(
    ::flatbuffers::FlatBufferBuilder &_fbb,
    const char *street = nullptr,
    const char *city = nullptr,
    const char *state = nullptr,
    const char *postal_code = nullptr,
    const char *country = nullptr) {
  auto street__ = street ? _fbb.CreateString(street) : 0;
  auto city__ = city ? _fbb.CreateString(city) : 0;
  auto state__ = state ? _fbb.CreateString(state) : 0;
  auto postal_code__ = postal_code ? _fbb.CreateString(postal_code) : 0;
  auto country__ = country ? _fbb.CreateString(country) : 0;
  return common::CreateAddress(
      _fbb,
      street__,
      city__,
      state__,
      postal_code__,
      country__);
}

}  // namespace common

#endif  // FLATBUFFERS_GENERATED_ADDRESS_COMMON_H_
