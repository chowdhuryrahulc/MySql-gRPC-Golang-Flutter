///
//  Generated code. Do not modify.
//  source: services.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use wordDescriptor instead')
const Word$json = const {
  '1': 'Word',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 5, '10': 'id'},
    const {'1': 'term', '3': 2, '4': 1, '5': 9, '10': 'term'},
    const {'1': 'defination', '3': 3, '4': 1, '5': 9, '10': 'defination'},
    const {'1': 'favorite', '3': 4, '4': 1, '5': 8, '10': 'favorite'},
    const {'1': 'image', '3': 5, '4': 1, '5': 12, '10': 'image'},
  ],
};

/// Descriptor for `Word`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List wordDescriptor = $convert.base64Decode('CgRXb3JkEg4KAmlkGAEgASgFUgJpZBISCgR0ZXJtGAIgASgJUgR0ZXJtEh4KCmRlZmluYXRpb24YAyABKAlSCmRlZmluYXRpb24SGgoIZmF2b3JpdGUYBCABKAhSCGZhdm9yaXRlEhQKBWltYWdlGAUgASgMUgVpbWFnZQ==');
@$core.Deprecated('Use vocabDescriptor instead')
const Vocab$json = const {
  '1': 'Vocab',
  '2': const [
    const {'1': 'Word', '3': 1, '4': 3, '5': 11, '6': '.proto.Word', '10': 'Word'},
  ],
};

/// Descriptor for `Vocab`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List vocabDescriptor = $convert.base64Decode('CgVWb2NhYhIfCgRXb3JkGAEgAygLMgsucHJvdG8uV29yZFIEV29yZA==');
@$core.Deprecated('Use paginateDescriptor instead')
const Paginate$json = const {
  '1': 'Paginate',
};

/// Descriptor for `Paginate`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List paginateDescriptor = $convert.base64Decode('CghQYWdpbmF0ZQ==');
