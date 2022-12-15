///
//  Generated code. Do not modify.
//  source: services.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'services.pb.dart' as $0;
export 'services.pb.dart';

class BroadcastClient extends $grpc.Client {
  static final _$getData = $grpc.ClientMethod<$0.Paginate, $0.Vocab>(
      '/proto.Broadcast/GetData',
      ($0.Paginate value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Vocab.fromBuffer(value));

  BroadcastClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.Vocab> getData($0.Paginate request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getData, request, options: options);
  }
}

abstract class BroadcastServiceBase extends $grpc.Service {
  $core.String get $name => 'proto.Broadcast';

  BroadcastServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Paginate, $0.Vocab>(
        'GetData',
        getData_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Paginate.fromBuffer(value),
        ($0.Vocab value) => value.writeToBuffer()));
  }

  $async.Future<$0.Vocab> getData_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Paginate> request) async {
    return getData(call, await request);
  }

  $async.Future<$0.Vocab> getData($grpc.ServiceCall call, $0.Paginate request);
}
