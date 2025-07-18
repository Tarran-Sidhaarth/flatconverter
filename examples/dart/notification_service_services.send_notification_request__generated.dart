// DO NOT EDIT!
// dart generated by Buffman 💪
// Versions:
// 		Buffman: 1.0.0
// 		Flatc: v25.2.10
// ignore_for_file: unused_import, unused_field, unused_element, unused_local_variable, constant_identifier_names

library services.send_notification_request_;

import 'dart:typed_data' show Uint8List;
import 'package:flat_buffers/flat_buffers.dart' as fb;

import './notification_service_services_generated.dart' as services;
import './notification_service_services.notification__generated.dart' as services_notification_;

import './address_common_generated.dart' as common;
import './status_common_generated.dart' as common;
import './timestamp_common_generated.dart' as common;
import './user_service_services_generated.dart' as services;

class MetadataEntry {
  MetadataEntry._(this._bc, this._bcOffset);
  factory MetadataEntry(List<int> bytes) {
    final rootRef = fb.BufferContext.fromBytes(bytes);
    return reader.read(rootRef, 0);
  }

  static const fb.Reader<MetadataEntry> reader = _MetadataEntryReader();

  final fb.BufferContext _bc;
  final int _bcOffset;

  String? get key => const fb.StringReader().vTableGetNullable(_bc, _bcOffset, 4);
  String? get value => const fb.StringReader().vTableGetNullable(_bc, _bcOffset, 6);

  @override
  String toString() {
    return 'MetadataEntry{key: ${key}, value: ${value}}';
  }
}

class _MetadataEntryReader extends fb.TableReader<MetadataEntry> {
  const _MetadataEntryReader();

  @override
  MetadataEntry createObject(fb.BufferContext bc, int offset) => 
    MetadataEntry._(bc, offset);
}

class MetadataEntryBuilder {
  MetadataEntryBuilder(this.fbBuilder);

  final fb.Builder fbBuilder;

  void begin() {
    fbBuilder.startTable(2);
  }

  int addKeyOffset(int? offset) {
    fbBuilder.addOffset(0, offset);
    return fbBuilder.offset;
  }
  int addValueOffset(int? offset) {
    fbBuilder.addOffset(1, offset);
    return fbBuilder.offset;
  }

  int finish() {
    return fbBuilder.endTable();
  }
}

class MetadataEntryObjectBuilder extends fb.ObjectBuilder {
  final String? _key;
  final String? _value;

  MetadataEntryObjectBuilder({
    String? key,
    String? value,
  })
      : _key = key,
        _value = value;

  /// Finish building, and store into the [fbBuilder].
  @override
  int finish(fb.Builder fbBuilder) {
    final int? keyOffset = _key == null ? null
        : fbBuilder.writeString(_key!);
    final int? valueOffset = _value == null ? null
        : fbBuilder.writeString(_value!);
    fbBuilder.startTable(2);
    fbBuilder.addOffset(0, keyOffset);
    fbBuilder.addOffset(1, valueOffset);
    return fbBuilder.endTable();
  }

  /// Convenience method to serialize to byte list.
  @override
  Uint8List toBytes([String? fileIdentifier]) {
    final fbBuilder = fb.Builder(deduplicateTables: false);
    fbBuilder.finish(finish(fbBuilder), fileIdentifier);
    return fbBuilder.buffer;
  }
}
