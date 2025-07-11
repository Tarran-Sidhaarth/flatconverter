// DO NOT EDIT!
// rust generated by Buffman 💪
// Versions:
// 		Buffman: 1.0.0
// 		Flatc: v25.2.10


// @generated

use crate::status_generated::*;
use crate::timestamp_generated::*;
use crate::address_generated::*;
use core::mem;
use core::cmp::Ordering;

extern crate flatbuffers;
use self::flatbuffers::{EndianScalar, Follow};

#[allow(unused_imports, dead_code)]
pub mod services {

  use crate::status_generated::*;
  use crate::timestamp_generated::*;
  use crate::address_generated::*;
  use core::mem;
  use core::cmp::Ordering;

  extern crate flatbuffers;
  use self::flatbuffers::{EndianScalar, Follow};

pub enum UserOffset {}
#[derive(Copy, Clone, PartialEq)]

pub struct User<'a> {
  pub _tab: flatbuffers::Table<'a>,
}

impl<'a> flatbuffers::Follow<'a> for User<'a> {
  type Inner = User<'a>;
  #[inline]
  unsafe fn follow(buf: &'a [u8], loc: usize) -> Self::Inner {
    Self { _tab: flatbuffers::Table::new(buf, loc) }
  }
}

impl<'a> User<'a> {
  pub const VT_USER_ID: flatbuffers::VOffsetT = 4;
  pub const VT_EMAIL: flatbuffers::VOffsetT = 6;
  pub const VT_FIRST_NAME: flatbuffers::VOffsetT = 8;
  pub const VT_LAST_NAME: flatbuffers::VOffsetT = 10;
  pub const VT_ADDRESS: flatbuffers::VOffsetT = 12;
  pub const VT_STATUS: flatbuffers::VOffsetT = 14;
  pub const VT_CREATED_AT: flatbuffers::VOffsetT = 16;
  pub const VT_UPDATED_AT: flatbuffers::VOffsetT = 18;
  pub const VT_ROLES: flatbuffers::VOffsetT = 20;

  #[inline]
  pub unsafe fn init_from_table(table: flatbuffers::Table<'a>) -> Self {
    User { _tab: table }
  }
  #[allow(unused_mut)]
  pub fn create<'bldr: 'args, 'args: 'mut_bldr, 'mut_bldr, A: flatbuffers::Allocator + 'bldr>(
    _fbb: &'mut_bldr mut flatbuffers::FlatBufferBuilder<'bldr, A>,
    args: &'args UserArgs<'args>
  ) -> flatbuffers::WIPOffset<User<'bldr>> {
    let mut builder = UserBuilder::new(_fbb);
    if let Some(x) = args.roles { builder.add_roles(x); }
    if let Some(x) = args.updated_at { builder.add_updated_at(x); }
    if let Some(x) = args.created_at { builder.add_created_at(x); }
    builder.add_status(args.status);
    if let Some(x) = args.address { builder.add_address(x); }
    if let Some(x) = args.last_name { builder.add_last_name(x); }
    if let Some(x) = args.first_name { builder.add_first_name(x); }
    if let Some(x) = args.email { builder.add_email(x); }
    if let Some(x) = args.user_id { builder.add_user_id(x); }
    builder.finish()
  }


  #[inline]
  pub fn user_id(&self) -> Option<&'a str> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<&str>>(User::VT_USER_ID, None)}
  }
  #[inline]
  pub fn email(&self) -> Option<&'a str> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<&str>>(User::VT_EMAIL, None)}
  }
  #[inline]
  pub fn first_name(&self) -> Option<&'a str> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<&str>>(User::VT_FIRST_NAME, None)}
  }
  #[inline]
  pub fn last_name(&self) -> Option<&'a str> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<&str>>(User::VT_LAST_NAME, None)}
  }
  #[inline]
  pub fn address(&self) -> Option<super::common::Address<'a>> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<super::common::Address>>(User::VT_ADDRESS, None)}
  }
  #[inline]
  pub fn status(&self) -> super::common::Status {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<super::common::Status>(User::VT_STATUS, Some(super::common::Status::STATUS_UNKNOWN)).unwrap()}
  }
  #[inline]
  pub fn created_at(&self) -> Option<super::common::Timestamp<'a>> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<super::common::Timestamp>>(User::VT_CREATED_AT, None)}
  }
  #[inline]
  pub fn updated_at(&self) -> Option<super::common::Timestamp<'a>> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<super::common::Timestamp>>(User::VT_UPDATED_AT, None)}
  }
  #[inline]
  pub fn roles(&self) -> Option<flatbuffers::Vector<'a, flatbuffers::ForwardsUOffset<&'a str>>> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<flatbuffers::Vector<'a, flatbuffers::ForwardsUOffset<&'a str>>>>(User::VT_ROLES, None)}
  }
}

impl flatbuffers::Verifiable for User<'_> {
  #[inline]
  fn run_verifier(
    v: &mut flatbuffers::Verifier, pos: usize
  ) -> Result<(), flatbuffers::InvalidFlatbuffer> {
    use self::flatbuffers::Verifiable;
    v.visit_table(pos)?
     .visit_field::<flatbuffers::ForwardsUOffset<&str>>("user_id", Self::VT_USER_ID, false)?
     .visit_field::<flatbuffers::ForwardsUOffset<&str>>("email", Self::VT_EMAIL, false)?
     .visit_field::<flatbuffers::ForwardsUOffset<&str>>("first_name", Self::VT_FIRST_NAME, false)?
     .visit_field::<flatbuffers::ForwardsUOffset<&str>>("last_name", Self::VT_LAST_NAME, false)?
     .visit_field::<flatbuffers::ForwardsUOffset<super::common::Address>>("address", Self::VT_ADDRESS, false)?
     .visit_field::<super::common::Status>("status", Self::VT_STATUS, false)?
     .visit_field::<flatbuffers::ForwardsUOffset<super::common::Timestamp>>("created_at", Self::VT_CREATED_AT, false)?
     .visit_field::<flatbuffers::ForwardsUOffset<super::common::Timestamp>>("updated_at", Self::VT_UPDATED_AT, false)?
     .visit_field::<flatbuffers::ForwardsUOffset<flatbuffers::Vector<'_, flatbuffers::ForwardsUOffset<&'_ str>>>>("roles", Self::VT_ROLES, false)?
     .finish();
    Ok(())
  }
}
pub struct UserArgs<'a> {
    pub user_id: Option<flatbuffers::WIPOffset<&'a str>>,
    pub email: Option<flatbuffers::WIPOffset<&'a str>>,
    pub first_name: Option<flatbuffers::WIPOffset<&'a str>>,
    pub last_name: Option<flatbuffers::WIPOffset<&'a str>>,
    pub address: Option<flatbuffers::WIPOffset<super::common::Address<'a>>>,
    pub status: super::common::Status,
    pub created_at: Option<flatbuffers::WIPOffset<super::common::Timestamp<'a>>>,
    pub updated_at: Option<flatbuffers::WIPOffset<super::common::Timestamp<'a>>>,
    pub roles: Option<flatbuffers::WIPOffset<flatbuffers::Vector<'a, flatbuffers::ForwardsUOffset<&'a str>>>>,
}
impl<'a> Default for UserArgs<'a> {
  #[inline]
  fn default() -> Self {
    UserArgs {
      user_id: None,
      email: None,
      first_name: None,
      last_name: None,
      address: None,
      status: super::common::Status::STATUS_UNKNOWN,
      created_at: None,
      updated_at: None,
      roles: None,
    }
  }
}

pub struct UserBuilder<'a: 'b, 'b, A: flatbuffers::Allocator + 'a> {
  fbb_: &'b mut flatbuffers::FlatBufferBuilder<'a, A>,
  start_: flatbuffers::WIPOffset<flatbuffers::TableUnfinishedWIPOffset>,
}
impl<'a: 'b, 'b, A: flatbuffers::Allocator + 'a> UserBuilder<'a, 'b, A> {
  #[inline]
  pub fn add_user_id(&mut self, user_id: flatbuffers::WIPOffset<&'b  str>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<_>>(User::VT_USER_ID, user_id);
  }
  #[inline]
  pub fn add_email(&mut self, email: flatbuffers::WIPOffset<&'b  str>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<_>>(User::VT_EMAIL, email);
  }
  #[inline]
  pub fn add_first_name(&mut self, first_name: flatbuffers::WIPOffset<&'b  str>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<_>>(User::VT_FIRST_NAME, first_name);
  }
  #[inline]
  pub fn add_last_name(&mut self, last_name: flatbuffers::WIPOffset<&'b  str>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<_>>(User::VT_LAST_NAME, last_name);
  }
  #[inline]
  pub fn add_address(&mut self, address: flatbuffers::WIPOffset<super::common::Address<'b >>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<super::common::Address>>(User::VT_ADDRESS, address);
  }
  #[inline]
  pub fn add_status(&mut self, status: super::common::Status) {
    self.fbb_.push_slot::<super::common::Status>(User::VT_STATUS, status, super::common::Status::STATUS_UNKNOWN);
  }
  #[inline]
  pub fn add_created_at(&mut self, created_at: flatbuffers::WIPOffset<super::common::Timestamp<'b >>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<super::common::Timestamp>>(User::VT_CREATED_AT, created_at);
  }
  #[inline]
  pub fn add_updated_at(&mut self, updated_at: flatbuffers::WIPOffset<super::common::Timestamp<'b >>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<super::common::Timestamp>>(User::VT_UPDATED_AT, updated_at);
  }
  #[inline]
  pub fn add_roles(&mut self, roles: flatbuffers::WIPOffset<flatbuffers::Vector<'b , flatbuffers::ForwardsUOffset<&'b  str>>>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<_>>(User::VT_ROLES, roles);
  }
  #[inline]
  pub fn new(_fbb: &'b mut flatbuffers::FlatBufferBuilder<'a, A>) -> UserBuilder<'a, 'b, A> {
    let start = _fbb.start_table();
    UserBuilder {
      fbb_: _fbb,
      start_: start,
    }
  }
  #[inline]
  pub fn finish(self) -> flatbuffers::WIPOffset<User<'a>> {
    let o = self.fbb_.end_table(self.start_);
    flatbuffers::WIPOffset::new(o.value())
  }
}

impl core::fmt::Debug for User<'_> {
  fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
    let mut ds = f.debug_struct("User");
      ds.field("user_id", &self.user_id());
      ds.field("email", &self.email());
      ds.field("first_name", &self.first_name());
      ds.field("last_name", &self.last_name());
      ds.field("address", &self.address());
      ds.field("status", &self.status());
      ds.field("created_at", &self.created_at());
      ds.field("updated_at", &self.updated_at());
      ds.field("roles", &self.roles());
      ds.finish()
  }
}
pub enum CreateUserRequestOffset {}
#[derive(Copy, Clone, PartialEq)]

pub struct CreateUserRequest<'a> {
  pub _tab: flatbuffers::Table<'a>,
}

impl<'a> flatbuffers::Follow<'a> for CreateUserRequest<'a> {
  type Inner = CreateUserRequest<'a>;
  #[inline]
  unsafe fn follow(buf: &'a [u8], loc: usize) -> Self::Inner {
    Self { _tab: flatbuffers::Table::new(buf, loc) }
  }
}

impl<'a> CreateUserRequest<'a> {
  pub const VT_EMAIL: flatbuffers::VOffsetT = 4;
  pub const VT_FIRST_NAME: flatbuffers::VOffsetT = 6;
  pub const VT_LAST_NAME: flatbuffers::VOffsetT = 8;
  pub const VT_ADDRESS: flatbuffers::VOffsetT = 10;
  pub const VT_ROLES: flatbuffers::VOffsetT = 12;

  #[inline]
  pub unsafe fn init_from_table(table: flatbuffers::Table<'a>) -> Self {
    CreateUserRequest { _tab: table }
  }
  #[allow(unused_mut)]
  pub fn create<'bldr: 'args, 'args: 'mut_bldr, 'mut_bldr, A: flatbuffers::Allocator + 'bldr>(
    _fbb: &'mut_bldr mut flatbuffers::FlatBufferBuilder<'bldr, A>,
    args: &'args CreateUserRequestArgs<'args>
  ) -> flatbuffers::WIPOffset<CreateUserRequest<'bldr>> {
    let mut builder = CreateUserRequestBuilder::new(_fbb);
    if let Some(x) = args.roles { builder.add_roles(x); }
    if let Some(x) = args.address { builder.add_address(x); }
    if let Some(x) = args.last_name { builder.add_last_name(x); }
    if let Some(x) = args.first_name { builder.add_first_name(x); }
    if let Some(x) = args.email { builder.add_email(x); }
    builder.finish()
  }


  #[inline]
  pub fn email(&self) -> Option<&'a str> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<&str>>(CreateUserRequest::VT_EMAIL, None)}
  }
  #[inline]
  pub fn first_name(&self) -> Option<&'a str> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<&str>>(CreateUserRequest::VT_FIRST_NAME, None)}
  }
  #[inline]
  pub fn last_name(&self) -> Option<&'a str> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<&str>>(CreateUserRequest::VT_LAST_NAME, None)}
  }
  #[inline]
  pub fn address(&self) -> Option<super::common::Address<'a>> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<super::common::Address>>(CreateUserRequest::VT_ADDRESS, None)}
  }
  #[inline]
  pub fn roles(&self) -> Option<flatbuffers::Vector<'a, flatbuffers::ForwardsUOffset<&'a str>>> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<flatbuffers::Vector<'a, flatbuffers::ForwardsUOffset<&'a str>>>>(CreateUserRequest::VT_ROLES, None)}
  }
}

impl flatbuffers::Verifiable for CreateUserRequest<'_> {
  #[inline]
  fn run_verifier(
    v: &mut flatbuffers::Verifier, pos: usize
  ) -> Result<(), flatbuffers::InvalidFlatbuffer> {
    use self::flatbuffers::Verifiable;
    v.visit_table(pos)?
     .visit_field::<flatbuffers::ForwardsUOffset<&str>>("email", Self::VT_EMAIL, false)?
     .visit_field::<flatbuffers::ForwardsUOffset<&str>>("first_name", Self::VT_FIRST_NAME, false)?
     .visit_field::<flatbuffers::ForwardsUOffset<&str>>("last_name", Self::VT_LAST_NAME, false)?
     .visit_field::<flatbuffers::ForwardsUOffset<super::common::Address>>("address", Self::VT_ADDRESS, false)?
     .visit_field::<flatbuffers::ForwardsUOffset<flatbuffers::Vector<'_, flatbuffers::ForwardsUOffset<&'_ str>>>>("roles", Self::VT_ROLES, false)?
     .finish();
    Ok(())
  }
}
pub struct CreateUserRequestArgs<'a> {
    pub email: Option<flatbuffers::WIPOffset<&'a str>>,
    pub first_name: Option<flatbuffers::WIPOffset<&'a str>>,
    pub last_name: Option<flatbuffers::WIPOffset<&'a str>>,
    pub address: Option<flatbuffers::WIPOffset<super::common::Address<'a>>>,
    pub roles: Option<flatbuffers::WIPOffset<flatbuffers::Vector<'a, flatbuffers::ForwardsUOffset<&'a str>>>>,
}
impl<'a> Default for CreateUserRequestArgs<'a> {
  #[inline]
  fn default() -> Self {
    CreateUserRequestArgs {
      email: None,
      first_name: None,
      last_name: None,
      address: None,
      roles: None,
    }
  }
}

pub struct CreateUserRequestBuilder<'a: 'b, 'b, A: flatbuffers::Allocator + 'a> {
  fbb_: &'b mut flatbuffers::FlatBufferBuilder<'a, A>,
  start_: flatbuffers::WIPOffset<flatbuffers::TableUnfinishedWIPOffset>,
}
impl<'a: 'b, 'b, A: flatbuffers::Allocator + 'a> CreateUserRequestBuilder<'a, 'b, A> {
  #[inline]
  pub fn add_email(&mut self, email: flatbuffers::WIPOffset<&'b  str>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<_>>(CreateUserRequest::VT_EMAIL, email);
  }
  #[inline]
  pub fn add_first_name(&mut self, first_name: flatbuffers::WIPOffset<&'b  str>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<_>>(CreateUserRequest::VT_FIRST_NAME, first_name);
  }
  #[inline]
  pub fn add_last_name(&mut self, last_name: flatbuffers::WIPOffset<&'b  str>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<_>>(CreateUserRequest::VT_LAST_NAME, last_name);
  }
  #[inline]
  pub fn add_address(&mut self, address: flatbuffers::WIPOffset<super::common::Address<'b >>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<super::common::Address>>(CreateUserRequest::VT_ADDRESS, address);
  }
  #[inline]
  pub fn add_roles(&mut self, roles: flatbuffers::WIPOffset<flatbuffers::Vector<'b , flatbuffers::ForwardsUOffset<&'b  str>>>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<_>>(CreateUserRequest::VT_ROLES, roles);
  }
  #[inline]
  pub fn new(_fbb: &'b mut flatbuffers::FlatBufferBuilder<'a, A>) -> CreateUserRequestBuilder<'a, 'b, A> {
    let start = _fbb.start_table();
    CreateUserRequestBuilder {
      fbb_: _fbb,
      start_: start,
    }
  }
  #[inline]
  pub fn finish(self) -> flatbuffers::WIPOffset<CreateUserRequest<'a>> {
    let o = self.fbb_.end_table(self.start_);
    flatbuffers::WIPOffset::new(o.value())
  }
}

impl core::fmt::Debug for CreateUserRequest<'_> {
  fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
    let mut ds = f.debug_struct("CreateUserRequest");
      ds.field("email", &self.email());
      ds.field("first_name", &self.first_name());
      ds.field("last_name", &self.last_name());
      ds.field("address", &self.address());
      ds.field("roles", &self.roles());
      ds.finish()
  }
}
pub enum CreateUserResponseOffset {}
#[derive(Copy, Clone, PartialEq)]

pub struct CreateUserResponse<'a> {
  pub _tab: flatbuffers::Table<'a>,
}

impl<'a> flatbuffers::Follow<'a> for CreateUserResponse<'a> {
  type Inner = CreateUserResponse<'a>;
  #[inline]
  unsafe fn follow(buf: &'a [u8], loc: usize) -> Self::Inner {
    Self { _tab: flatbuffers::Table::new(buf, loc) }
  }
}

impl<'a> CreateUserResponse<'a> {
  pub const VT_USER: flatbuffers::VOffsetT = 4;
  pub const VT_SUCCESS: flatbuffers::VOffsetT = 6;
  pub const VT_MESSAGE: flatbuffers::VOffsetT = 8;

  #[inline]
  pub unsafe fn init_from_table(table: flatbuffers::Table<'a>) -> Self {
    CreateUserResponse { _tab: table }
  }
  #[allow(unused_mut)]
  pub fn create<'bldr: 'args, 'args: 'mut_bldr, 'mut_bldr, A: flatbuffers::Allocator + 'bldr>(
    _fbb: &'mut_bldr mut flatbuffers::FlatBufferBuilder<'bldr, A>,
    args: &'args CreateUserResponseArgs<'args>
  ) -> flatbuffers::WIPOffset<CreateUserResponse<'bldr>> {
    let mut builder = CreateUserResponseBuilder::new(_fbb);
    if let Some(x) = args.message { builder.add_message(x); }
    if let Some(x) = args.user { builder.add_user(x); }
    builder.add_success(args.success);
    builder.finish()
  }


  #[inline]
  pub fn user(&self) -> Option<User<'a>> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<User>>(CreateUserResponse::VT_USER, None)}
  }
  #[inline]
  pub fn success(&self) -> bool {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<bool>(CreateUserResponse::VT_SUCCESS, Some(false)).unwrap()}
  }
  #[inline]
  pub fn message(&self) -> Option<&'a str> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<&str>>(CreateUserResponse::VT_MESSAGE, None)}
  }
}

impl flatbuffers::Verifiable for CreateUserResponse<'_> {
  #[inline]
  fn run_verifier(
    v: &mut flatbuffers::Verifier, pos: usize
  ) -> Result<(), flatbuffers::InvalidFlatbuffer> {
    use self::flatbuffers::Verifiable;
    v.visit_table(pos)?
     .visit_field::<flatbuffers::ForwardsUOffset<User>>("user", Self::VT_USER, false)?
     .visit_field::<bool>("success", Self::VT_SUCCESS, false)?
     .visit_field::<flatbuffers::ForwardsUOffset<&str>>("message", Self::VT_MESSAGE, false)?
     .finish();
    Ok(())
  }
}
pub struct CreateUserResponseArgs<'a> {
    pub user: Option<flatbuffers::WIPOffset<User<'a>>>,
    pub success: bool,
    pub message: Option<flatbuffers::WIPOffset<&'a str>>,
}
impl<'a> Default for CreateUserResponseArgs<'a> {
  #[inline]
  fn default() -> Self {
    CreateUserResponseArgs {
      user: None,
      success: false,
      message: None,
    }
  }
}

pub struct CreateUserResponseBuilder<'a: 'b, 'b, A: flatbuffers::Allocator + 'a> {
  fbb_: &'b mut flatbuffers::FlatBufferBuilder<'a, A>,
  start_: flatbuffers::WIPOffset<flatbuffers::TableUnfinishedWIPOffset>,
}
impl<'a: 'b, 'b, A: flatbuffers::Allocator + 'a> CreateUserResponseBuilder<'a, 'b, A> {
  #[inline]
  pub fn add_user(&mut self, user: flatbuffers::WIPOffset<User<'b >>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<User>>(CreateUserResponse::VT_USER, user);
  }
  #[inline]
  pub fn add_success(&mut self, success: bool) {
    self.fbb_.push_slot::<bool>(CreateUserResponse::VT_SUCCESS, success, false);
  }
  #[inline]
  pub fn add_message(&mut self, message: flatbuffers::WIPOffset<&'b  str>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<_>>(CreateUserResponse::VT_MESSAGE, message);
  }
  #[inline]
  pub fn new(_fbb: &'b mut flatbuffers::FlatBufferBuilder<'a, A>) -> CreateUserResponseBuilder<'a, 'b, A> {
    let start = _fbb.start_table();
    CreateUserResponseBuilder {
      fbb_: _fbb,
      start_: start,
    }
  }
  #[inline]
  pub fn finish(self) -> flatbuffers::WIPOffset<CreateUserResponse<'a>> {
    let o = self.fbb_.end_table(self.start_);
    flatbuffers::WIPOffset::new(o.value())
  }
}

impl core::fmt::Debug for CreateUserResponse<'_> {
  fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
    let mut ds = f.debug_struct("CreateUserResponse");
      ds.field("user", &self.user());
      ds.field("success", &self.success());
      ds.field("message", &self.message());
      ds.finish()
  }
}
pub enum GetUserRequestOffset {}
#[derive(Copy, Clone, PartialEq)]

pub struct GetUserRequest<'a> {
  pub _tab: flatbuffers::Table<'a>,
}

impl<'a> flatbuffers::Follow<'a> for GetUserRequest<'a> {
  type Inner = GetUserRequest<'a>;
  #[inline]
  unsafe fn follow(buf: &'a [u8], loc: usize) -> Self::Inner {
    Self { _tab: flatbuffers::Table::new(buf, loc) }
  }
}

impl<'a> GetUserRequest<'a> {
  pub const VT_USER_ID: flatbuffers::VOffsetT = 4;

  #[inline]
  pub unsafe fn init_from_table(table: flatbuffers::Table<'a>) -> Self {
    GetUserRequest { _tab: table }
  }
  #[allow(unused_mut)]
  pub fn create<'bldr: 'args, 'args: 'mut_bldr, 'mut_bldr, A: flatbuffers::Allocator + 'bldr>(
    _fbb: &'mut_bldr mut flatbuffers::FlatBufferBuilder<'bldr, A>,
    args: &'args GetUserRequestArgs<'args>
  ) -> flatbuffers::WIPOffset<GetUserRequest<'bldr>> {
    let mut builder = GetUserRequestBuilder::new(_fbb);
    if let Some(x) = args.user_id { builder.add_user_id(x); }
    builder.finish()
  }


  #[inline]
  pub fn user_id(&self) -> Option<&'a str> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<&str>>(GetUserRequest::VT_USER_ID, None)}
  }
}

impl flatbuffers::Verifiable for GetUserRequest<'_> {
  #[inline]
  fn run_verifier(
    v: &mut flatbuffers::Verifier, pos: usize
  ) -> Result<(), flatbuffers::InvalidFlatbuffer> {
    use self::flatbuffers::Verifiable;
    v.visit_table(pos)?
     .visit_field::<flatbuffers::ForwardsUOffset<&str>>("user_id", Self::VT_USER_ID, false)?
     .finish();
    Ok(())
  }
}
pub struct GetUserRequestArgs<'a> {
    pub user_id: Option<flatbuffers::WIPOffset<&'a str>>,
}
impl<'a> Default for GetUserRequestArgs<'a> {
  #[inline]
  fn default() -> Self {
    GetUserRequestArgs {
      user_id: None,
    }
  }
}

pub struct GetUserRequestBuilder<'a: 'b, 'b, A: flatbuffers::Allocator + 'a> {
  fbb_: &'b mut flatbuffers::FlatBufferBuilder<'a, A>,
  start_: flatbuffers::WIPOffset<flatbuffers::TableUnfinishedWIPOffset>,
}
impl<'a: 'b, 'b, A: flatbuffers::Allocator + 'a> GetUserRequestBuilder<'a, 'b, A> {
  #[inline]
  pub fn add_user_id(&mut self, user_id: flatbuffers::WIPOffset<&'b  str>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<_>>(GetUserRequest::VT_USER_ID, user_id);
  }
  #[inline]
  pub fn new(_fbb: &'b mut flatbuffers::FlatBufferBuilder<'a, A>) -> GetUserRequestBuilder<'a, 'b, A> {
    let start = _fbb.start_table();
    GetUserRequestBuilder {
      fbb_: _fbb,
      start_: start,
    }
  }
  #[inline]
  pub fn finish(self) -> flatbuffers::WIPOffset<GetUserRequest<'a>> {
    let o = self.fbb_.end_table(self.start_);
    flatbuffers::WIPOffset::new(o.value())
  }
}

impl core::fmt::Debug for GetUserRequest<'_> {
  fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
    let mut ds = f.debug_struct("GetUserRequest");
      ds.field("user_id", &self.user_id());
      ds.finish()
  }
}
pub enum GetUserResponseOffset {}
#[derive(Copy, Clone, PartialEq)]

pub struct GetUserResponse<'a> {
  pub _tab: flatbuffers::Table<'a>,
}

impl<'a> flatbuffers::Follow<'a> for GetUserResponse<'a> {
  type Inner = GetUserResponse<'a>;
  #[inline]
  unsafe fn follow(buf: &'a [u8], loc: usize) -> Self::Inner {
    Self { _tab: flatbuffers::Table::new(buf, loc) }
  }
}

impl<'a> GetUserResponse<'a> {
  pub const VT_USER: flatbuffers::VOffsetT = 4;
  pub const VT_FOUND: flatbuffers::VOffsetT = 6;

  #[inline]
  pub unsafe fn init_from_table(table: flatbuffers::Table<'a>) -> Self {
    GetUserResponse { _tab: table }
  }
  #[allow(unused_mut)]
  pub fn create<'bldr: 'args, 'args: 'mut_bldr, 'mut_bldr, A: flatbuffers::Allocator + 'bldr>(
    _fbb: &'mut_bldr mut flatbuffers::FlatBufferBuilder<'bldr, A>,
    args: &'args GetUserResponseArgs<'args>
  ) -> flatbuffers::WIPOffset<GetUserResponse<'bldr>> {
    let mut builder = GetUserResponseBuilder::new(_fbb);
    if let Some(x) = args.user { builder.add_user(x); }
    builder.add_found(args.found);
    builder.finish()
  }


  #[inline]
  pub fn user(&self) -> Option<User<'a>> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<User>>(GetUserResponse::VT_USER, None)}
  }
  #[inline]
  pub fn found(&self) -> bool {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<bool>(GetUserResponse::VT_FOUND, Some(false)).unwrap()}
  }
}

impl flatbuffers::Verifiable for GetUserResponse<'_> {
  #[inline]
  fn run_verifier(
    v: &mut flatbuffers::Verifier, pos: usize
  ) -> Result<(), flatbuffers::InvalidFlatbuffer> {
    use self::flatbuffers::Verifiable;
    v.visit_table(pos)?
     .visit_field::<flatbuffers::ForwardsUOffset<User>>("user", Self::VT_USER, false)?
     .visit_field::<bool>("found", Self::VT_FOUND, false)?
     .finish();
    Ok(())
  }
}
pub struct GetUserResponseArgs<'a> {
    pub user: Option<flatbuffers::WIPOffset<User<'a>>>,
    pub found: bool,
}
impl<'a> Default for GetUserResponseArgs<'a> {
  #[inline]
  fn default() -> Self {
    GetUserResponseArgs {
      user: None,
      found: false,
    }
  }
}

pub struct GetUserResponseBuilder<'a: 'b, 'b, A: flatbuffers::Allocator + 'a> {
  fbb_: &'b mut flatbuffers::FlatBufferBuilder<'a, A>,
  start_: flatbuffers::WIPOffset<flatbuffers::TableUnfinishedWIPOffset>,
}
impl<'a: 'b, 'b, A: flatbuffers::Allocator + 'a> GetUserResponseBuilder<'a, 'b, A> {
  #[inline]
  pub fn add_user(&mut self, user: flatbuffers::WIPOffset<User<'b >>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<User>>(GetUserResponse::VT_USER, user);
  }
  #[inline]
  pub fn add_found(&mut self, found: bool) {
    self.fbb_.push_slot::<bool>(GetUserResponse::VT_FOUND, found, false);
  }
  #[inline]
  pub fn new(_fbb: &'b mut flatbuffers::FlatBufferBuilder<'a, A>) -> GetUserResponseBuilder<'a, 'b, A> {
    let start = _fbb.start_table();
    GetUserResponseBuilder {
      fbb_: _fbb,
      start_: start,
    }
  }
  #[inline]
  pub fn finish(self) -> flatbuffers::WIPOffset<GetUserResponse<'a>> {
    let o = self.fbb_.end_table(self.start_);
    flatbuffers::WIPOffset::new(o.value())
  }
}

impl core::fmt::Debug for GetUserResponse<'_> {
  fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
    let mut ds = f.debug_struct("GetUserResponse");
      ds.field("user", &self.user());
      ds.field("found", &self.found());
      ds.finish()
  }
}
pub enum UpdateUserStatusRequestOffset {}
#[derive(Copy, Clone, PartialEq)]

pub struct UpdateUserStatusRequest<'a> {
  pub _tab: flatbuffers::Table<'a>,
}

impl<'a> flatbuffers::Follow<'a> for UpdateUserStatusRequest<'a> {
  type Inner = UpdateUserStatusRequest<'a>;
  #[inline]
  unsafe fn follow(buf: &'a [u8], loc: usize) -> Self::Inner {
    Self { _tab: flatbuffers::Table::new(buf, loc) }
  }
}

impl<'a> UpdateUserStatusRequest<'a> {
  pub const VT_USER_ID: flatbuffers::VOffsetT = 4;
  pub const VT_NEW_STATUS: flatbuffers::VOffsetT = 6;

  #[inline]
  pub unsafe fn init_from_table(table: flatbuffers::Table<'a>) -> Self {
    UpdateUserStatusRequest { _tab: table }
  }
  #[allow(unused_mut)]
  pub fn create<'bldr: 'args, 'args: 'mut_bldr, 'mut_bldr, A: flatbuffers::Allocator + 'bldr>(
    _fbb: &'mut_bldr mut flatbuffers::FlatBufferBuilder<'bldr, A>,
    args: &'args UpdateUserStatusRequestArgs<'args>
  ) -> flatbuffers::WIPOffset<UpdateUserStatusRequest<'bldr>> {
    let mut builder = UpdateUserStatusRequestBuilder::new(_fbb);
    builder.add_new_status(args.new_status);
    if let Some(x) = args.user_id { builder.add_user_id(x); }
    builder.finish()
  }


  #[inline]
  pub fn user_id(&self) -> Option<&'a str> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<&str>>(UpdateUserStatusRequest::VT_USER_ID, None)}
  }
  #[inline]
  pub fn new_status(&self) -> super::common::Status {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<super::common::Status>(UpdateUserStatusRequest::VT_NEW_STATUS, Some(super::common::Status::STATUS_UNKNOWN)).unwrap()}
  }
}

impl flatbuffers::Verifiable for UpdateUserStatusRequest<'_> {
  #[inline]
  fn run_verifier(
    v: &mut flatbuffers::Verifier, pos: usize
  ) -> Result<(), flatbuffers::InvalidFlatbuffer> {
    use self::flatbuffers::Verifiable;
    v.visit_table(pos)?
     .visit_field::<flatbuffers::ForwardsUOffset<&str>>("user_id", Self::VT_USER_ID, false)?
     .visit_field::<super::common::Status>("new_status", Self::VT_NEW_STATUS, false)?
     .finish();
    Ok(())
  }
}
pub struct UpdateUserStatusRequestArgs<'a> {
    pub user_id: Option<flatbuffers::WIPOffset<&'a str>>,
    pub new_status: super::common::Status,
}
impl<'a> Default for UpdateUserStatusRequestArgs<'a> {
  #[inline]
  fn default() -> Self {
    UpdateUserStatusRequestArgs {
      user_id: None,
      new_status: super::common::Status::STATUS_UNKNOWN,
    }
  }
}

pub struct UpdateUserStatusRequestBuilder<'a: 'b, 'b, A: flatbuffers::Allocator + 'a> {
  fbb_: &'b mut flatbuffers::FlatBufferBuilder<'a, A>,
  start_: flatbuffers::WIPOffset<flatbuffers::TableUnfinishedWIPOffset>,
}
impl<'a: 'b, 'b, A: flatbuffers::Allocator + 'a> UpdateUserStatusRequestBuilder<'a, 'b, A> {
  #[inline]
  pub fn add_user_id(&mut self, user_id: flatbuffers::WIPOffset<&'b  str>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<_>>(UpdateUserStatusRequest::VT_USER_ID, user_id);
  }
  #[inline]
  pub fn add_new_status(&mut self, new_status: super::common::Status) {
    self.fbb_.push_slot::<super::common::Status>(UpdateUserStatusRequest::VT_NEW_STATUS, new_status, super::common::Status::STATUS_UNKNOWN);
  }
  #[inline]
  pub fn new(_fbb: &'b mut flatbuffers::FlatBufferBuilder<'a, A>) -> UpdateUserStatusRequestBuilder<'a, 'b, A> {
    let start = _fbb.start_table();
    UpdateUserStatusRequestBuilder {
      fbb_: _fbb,
      start_: start,
    }
  }
  #[inline]
  pub fn finish(self) -> flatbuffers::WIPOffset<UpdateUserStatusRequest<'a>> {
    let o = self.fbb_.end_table(self.start_);
    flatbuffers::WIPOffset::new(o.value())
  }
}

impl core::fmt::Debug for UpdateUserStatusRequest<'_> {
  fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
    let mut ds = f.debug_struct("UpdateUserStatusRequest");
      ds.field("user_id", &self.user_id());
      ds.field("new_status", &self.new_status());
      ds.finish()
  }
}
pub enum UpdateUserStatusResponseOffset {}
#[derive(Copy, Clone, PartialEq)]

pub struct UpdateUserStatusResponse<'a> {
  pub _tab: flatbuffers::Table<'a>,
}

impl<'a> flatbuffers::Follow<'a> for UpdateUserStatusResponse<'a> {
  type Inner = UpdateUserStatusResponse<'a>;
  #[inline]
  unsafe fn follow(buf: &'a [u8], loc: usize) -> Self::Inner {
    Self { _tab: flatbuffers::Table::new(buf, loc) }
  }
}

impl<'a> UpdateUserStatusResponse<'a> {
  pub const VT_SUCCESS: flatbuffers::VOffsetT = 4;
  pub const VT_MESSAGE: flatbuffers::VOffsetT = 6;

  #[inline]
  pub unsafe fn init_from_table(table: flatbuffers::Table<'a>) -> Self {
    UpdateUserStatusResponse { _tab: table }
  }
  #[allow(unused_mut)]
  pub fn create<'bldr: 'args, 'args: 'mut_bldr, 'mut_bldr, A: flatbuffers::Allocator + 'bldr>(
    _fbb: &'mut_bldr mut flatbuffers::FlatBufferBuilder<'bldr, A>,
    args: &'args UpdateUserStatusResponseArgs<'args>
  ) -> flatbuffers::WIPOffset<UpdateUserStatusResponse<'bldr>> {
    let mut builder = UpdateUserStatusResponseBuilder::new(_fbb);
    if let Some(x) = args.message { builder.add_message(x); }
    builder.add_success(args.success);
    builder.finish()
  }


  #[inline]
  pub fn success(&self) -> bool {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<bool>(UpdateUserStatusResponse::VT_SUCCESS, Some(false)).unwrap()}
  }
  #[inline]
  pub fn message(&self) -> Option<&'a str> {
    // Safety:
    // Created from valid Table for this object
    // which contains a valid value in this slot
    unsafe { self._tab.get::<flatbuffers::ForwardsUOffset<&str>>(UpdateUserStatusResponse::VT_MESSAGE, None)}
  }
}

impl flatbuffers::Verifiable for UpdateUserStatusResponse<'_> {
  #[inline]
  fn run_verifier(
    v: &mut flatbuffers::Verifier, pos: usize
  ) -> Result<(), flatbuffers::InvalidFlatbuffer> {
    use self::flatbuffers::Verifiable;
    v.visit_table(pos)?
     .visit_field::<bool>("success", Self::VT_SUCCESS, false)?
     .visit_field::<flatbuffers::ForwardsUOffset<&str>>("message", Self::VT_MESSAGE, false)?
     .finish();
    Ok(())
  }
}
pub struct UpdateUserStatusResponseArgs<'a> {
    pub success: bool,
    pub message: Option<flatbuffers::WIPOffset<&'a str>>,
}
impl<'a> Default for UpdateUserStatusResponseArgs<'a> {
  #[inline]
  fn default() -> Self {
    UpdateUserStatusResponseArgs {
      success: false,
      message: None,
    }
  }
}

pub struct UpdateUserStatusResponseBuilder<'a: 'b, 'b, A: flatbuffers::Allocator + 'a> {
  fbb_: &'b mut flatbuffers::FlatBufferBuilder<'a, A>,
  start_: flatbuffers::WIPOffset<flatbuffers::TableUnfinishedWIPOffset>,
}
impl<'a: 'b, 'b, A: flatbuffers::Allocator + 'a> UpdateUserStatusResponseBuilder<'a, 'b, A> {
  #[inline]
  pub fn add_success(&mut self, success: bool) {
    self.fbb_.push_slot::<bool>(UpdateUserStatusResponse::VT_SUCCESS, success, false);
  }
  #[inline]
  pub fn add_message(&mut self, message: flatbuffers::WIPOffset<&'b  str>) {
    self.fbb_.push_slot_always::<flatbuffers::WIPOffset<_>>(UpdateUserStatusResponse::VT_MESSAGE, message);
  }
  #[inline]
  pub fn new(_fbb: &'b mut flatbuffers::FlatBufferBuilder<'a, A>) -> UpdateUserStatusResponseBuilder<'a, 'b, A> {
    let start = _fbb.start_table();
    UpdateUserStatusResponseBuilder {
      fbb_: _fbb,
      start_: start,
    }
  }
  #[inline]
  pub fn finish(self) -> flatbuffers::WIPOffset<UpdateUserStatusResponse<'a>> {
    let o = self.fbb_.end_table(self.start_);
    flatbuffers::WIPOffset::new(o.value())
  }
}

impl core::fmt::Debug for UpdateUserStatusResponse<'_> {
  fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
    let mut ds = f.debug_struct("UpdateUserStatusResponse");
      ds.field("success", &self.success());
      ds.field("message", &self.message());
      ds.finish()
  }
}
}  // pub mod services

