// ignore_for_file: prefer_const_constructors

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
    Vocab vocab = await client!.getData(paginate); // await worked, not went ahead without implementing this
    return vocab.word.toList();
  }
}
