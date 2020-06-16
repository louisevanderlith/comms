import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/requester.dart';

import 'bodies/message.dart';

Future<HttpRequest> sendMessage(Message obj) async {
  var apiroute = getEndpoint("comms");
  var url = "${apiroute}/messages";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}
