// ignore_for_file: prefer_const_constructors

import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:fluttersenddata/Pagination/pagination.dart';
import 'package:fluttersenddata/gen/services.pb.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

// todo PAGINATION:
//todo  send limited amount from golang to flutter, lets say 4
//todo  then user presses a button, pagination function again gets called 
//todo  and the next 4 gets in

//todo So button push, 4, then button push, next 4 (and it push out 4, which then added to flutter frontend list). Sort via id?
//todo but will itbe useful. Bcoz we need pagination in DynmoDb. Not MySql at the very least (DONT DO FOR MYSQL)

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
  final wordsList = List<Word>;
  PaginationPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: FutureBuilder<List<Word>>(
          future: pagination.recieveVocab(),
          builder: (BuildContext context, AsyncSnapshot snapshot) {
            if (!snapshot.hasData) {
              return Center(
                child: CircularProgressIndicator(),
              );
            }

            log(snapshot.data.toString());
            // log(snapshot.data[1].term.toString());
            // wordsList.add(snapshot.data);
            //! This is causing problems bcoz we just have 1 vocab.
            //! We need toconvert it into list of words. And then you can access it via indexes
            //! But how? Bcoz it is just a big huge vocab
            // log(wordsList.toString());
            // log(snapshot.data.length.toString());

            return ListView.builder(
                itemCount: snapshot.data.length,
                itemBuilder: ((context, index) {
                  return ListTile(
                    title: Text(snapshot.data[index].term),
                    subtitle: Text(snapshot.data[index].defination),
                  );
                }));
          }),
    );
  }
}
