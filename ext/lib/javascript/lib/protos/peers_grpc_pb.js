// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var protos_peers_pb = require('../protos/peers_pb.js');

function serialize_DeletePeersRequest(arg) {
  if (!(arg instanceof protos_peers_pb.DeletePeersRequest)) {
    throw new Error('Expected argument of type DeletePeersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_DeletePeersRequest(buffer_arg) {
  return protos_peers_pb.DeletePeersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_DeletePeersResponse(arg) {
  if (!(arg instanceof protos_peers_pb.DeletePeersResponse)) {
    throw new Error('Expected argument of type DeletePeersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_DeletePeersResponse(buffer_arg) {
  return protos_peers_pb.DeletePeersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GetPeersRequest(arg) {
  if (!(arg instanceof protos_peers_pb.GetPeersRequest)) {
    throw new Error('Expected argument of type GetPeersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_GetPeersRequest(buffer_arg) {
  return protos_peers_pb.GetPeersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GetPeersResponse(arg) {
  if (!(arg instanceof protos_peers_pb.GetPeersResponse)) {
    throw new Error('Expected argument of type GetPeersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_GetPeersResponse(buffer_arg) {
  return protos_peers_pb.GetPeersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_RegisterPeersRequest(arg) {
  if (!(arg instanceof protos_peers_pb.RegisterPeersRequest)) {
    throw new Error('Expected argument of type RegisterPeersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_RegisterPeersRequest(buffer_arg) {
  return protos_peers_pb.RegisterPeersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_RegisterPeersResponse(arg) {
  if (!(arg instanceof protos_peers_pb.RegisterPeersResponse)) {
    throw new Error('Expected argument of type RegisterPeersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_RegisterPeersResponse(buffer_arg) {
  return protos_peers_pb.RegisterPeersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var PeersService = exports.PeersService = {
  getPeers: {
    path: '/Peers/GetPeers',
    requestStream: false,
    responseStream: false,
    requestType: protos_peers_pb.GetPeersRequest,
    responseType: protos_peers_pb.GetPeersResponse,
    requestSerialize: serialize_GetPeersRequest,
    requestDeserialize: deserialize_GetPeersRequest,
    responseSerialize: serialize_GetPeersResponse,
    responseDeserialize: deserialize_GetPeersResponse,
  },
  registerPeers: {
    path: '/Peers/RegisterPeers',
    requestStream: false,
    responseStream: false,
    requestType: protos_peers_pb.RegisterPeersRequest,
    responseType: protos_peers_pb.RegisterPeersResponse,
    requestSerialize: serialize_RegisterPeersRequest,
    requestDeserialize: deserialize_RegisterPeersRequest,
    responseSerialize: serialize_RegisterPeersResponse,
    responseDeserialize: deserialize_RegisterPeersResponse,
  },
  deletePeers: {
    path: '/Peers/DeletePeers',
    requestStream: false,
    responseStream: false,
    requestType: protos_peers_pb.DeletePeersRequest,
    responseType: protos_peers_pb.DeletePeersResponse,
    requestSerialize: serialize_DeletePeersRequest,
    requestDeserialize: deserialize_DeletePeersRequest,
    responseSerialize: serialize_DeletePeersResponse,
    responseDeserialize: deserialize_DeletePeersResponse,
  },
};

exports.PeersClient = grpc.makeGenericClientConstructor(PeersService);
