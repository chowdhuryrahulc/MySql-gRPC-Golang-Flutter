///
//  Generated code. Do not modify.
//  source: services.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class Word extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Word', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'proto'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.O3)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'term')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'defination')
    ..aOB(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'favorite')
    ..a<$core.List<$core.int>>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'image', $pb.PbFieldType.OY)
    ..hasRequiredFields = false
  ;

  Word._() : super();
  factory Word({
    $core.int? id,
    $core.String? term,
    $core.String? defination,
    $core.bool? favorite,
    $core.List<$core.int>? image,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (term != null) {
      _result.term = term;
    }
    if (defination != null) {
      _result.defination = defination;
    }
    if (favorite != null) {
      _result.favorite = favorite;
    }
    if (image != null) {
      _result.image = image;
    }
    return _result;
  }
  factory Word.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Word.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Word clone() => Word()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Word copyWith(void Function(Word) updates) => super.copyWith((message) => updates(message as Word)) as Word; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Word create() => Word._();
  Word createEmptyInstance() => create();
  static $pb.PbList<Word> createRepeated() => $pb.PbList<Word>();
  @$core.pragma('dart2js:noInline')
  static Word getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Word>(create);
  static Word? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get term => $_getSZ(1);
  @$pb.TagNumber(2)
  set term($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTerm() => $_has(1);
  @$pb.TagNumber(2)
  void clearTerm() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get defination => $_getSZ(2);
  @$pb.TagNumber(3)
  set defination($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasDefination() => $_has(2);
  @$pb.TagNumber(3)
  void clearDefination() => clearField(3);

  @$pb.TagNumber(4)
  $core.bool get favorite => $_getBF(3);
  @$pb.TagNumber(4)
  set favorite($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasFavorite() => $_has(3);
  @$pb.TagNumber(4)
  void clearFavorite() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<$core.int> get image => $_getN(4);
  @$pb.TagNumber(5)
  set image($core.List<$core.int> v) { $_setBytes(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasImage() => $_has(4);
  @$pb.TagNumber(5)
  void clearImage() => clearField(5);
}

class Vocab extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Vocab', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'proto'), createEmptyInstance: create)
    ..pc<Word>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Word', $pb.PbFieldType.PM, protoName: 'Word', subBuilder: Word.create)
    ..hasRequiredFields = false
  ;

  Vocab._() : super();
  factory Vocab({
    $core.Iterable<Word>? word,
  }) {
    final _result = create();
    if (word != null) {
      _result.word.addAll(word);
    }
    return _result;
  }
  factory Vocab.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Vocab.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Vocab clone() => Vocab()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Vocab copyWith(void Function(Vocab) updates) => super.copyWith((message) => updates(message as Vocab)) as Vocab; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Vocab create() => Vocab._();
  Vocab createEmptyInstance() => create();
  static $pb.PbList<Vocab> createRepeated() => $pb.PbList<Vocab>();
  @$core.pragma('dart2js:noInline')
  static Vocab getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Vocab>(create);
  static Vocab? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Word> get word => $_getList(0);
}

class Paginate extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Paginate', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'proto'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  Paginate._() : super();
  factory Paginate() => create();
  factory Paginate.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Paginate.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Paginate clone() => Paginate()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Paginate copyWith(void Function(Paginate) updates) => super.copyWith((message) => updates(message as Paginate)) as Paginate; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Paginate create() => Paginate._();
  Paginate createEmptyInstance() => create();
  static $pb.PbList<Paginate> createRepeated() => $pb.PbList<Paginate>();
  @$core.pragma('dart2js:noInline')
  static Paginate getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Paginate>(create);
  static Paginate? _defaultInstance;
}

