// ignore_for_file: prefer_const_constructors

import 'package:flutter/material.dart';
import 'package:fluttersenddata/Pagination/pagination.dart';
import 'package:fluttersenddata/gen/services.pb.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: PaginationPage(),
    );
  }
}

class PaginationPage extends StatelessWidget {
  final pagination = Pagination();
  PaginationPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: FutureBuilder<Vocab>(
          future: pagination.recieveVocab(),
          builder: (context, snapshot) {
            if (!snapshot.hasData) {
              return Center(
                child: CircularProgressIndicator(),
              );
            }
            return ListView.builder(
                itemCount: 2,
                itemBuilder: ((context, index) {
                  return ListTile(
                    // title: Text(snapshot.data.word[index]),
                  );
                }));
          }),
    );
  }
}
