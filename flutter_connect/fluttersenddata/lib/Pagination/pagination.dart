// ignore_for_file: prefer_const_constructors

import 'dart:developer';

import 'package:grpc/grpc.dart';
import '../gen/services.pbgrpc.dart';

class Pagination {
  Paginate paginate = Paginate();
  static BroadcastClient? client;

  Pagination() {
    client = BroadcastClient(ClientChannel("192.168.29.93",
        port: 8080,
        options: ChannelOptions(
          credentials: ChannelCredentials.insecure(),
          codecRegistry:
              CodecRegistry(codecs: const [GzipCodec(), IdentityCodec()]),
        )));
  }

  Future<List<Word>> recieveVocab() async {
    //// Paginate paginate = Paginate();
    //! Error expected: bcoz we are not waiting for the vocab async.
    Vocab vocab = await client!.getData(paginate);
    log(vocab.toString());      //? WORKING, data is comming
    return vocab.word.toList();
    // vocab;

    //   await for (var msg in client!.createStream(connect)) {
    //     // client!.createStream(connect) is the 2nd rpc function.
    //     // Takes connect, returns stream of messages
    //     yield msg;
    //   }
  }
}
