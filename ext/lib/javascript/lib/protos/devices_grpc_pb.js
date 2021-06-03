// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var protos_devices_pb = require('../protos/devices_pb.js');
var protos_peers_pb = require('../protos/peers_pb.js');

function serialize_GetDevicesRequest(arg) {
  if (!(arg instanceof protos_devices_pb.GetDevicesRequest)) {
    throw new Error('Expected argument of type GetDevicesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_GetDevicesRequest(buffer_arg) {
  return protos_devices_pb.GetDevicesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GetDevicesResponse(arg) {
  if (!(arg instanceof protos_devices_pb.GetDevicesResponse)) {
    throw new Error('Expected argument of type GetDevicesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_GetDevicesResponse(buffer_arg) {
  return protos_devices_pb.GetDevicesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_UpdatePrivateKeyRequest(arg) {
  if (!(arg instanceof protos_devices_pb.UpdatePrivateKeyRequest)) {
    throw new Error('Expected argument of type UpdatePrivateKeyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_UpdatePrivateKeyRequest(buffer_arg) {
  return protos_devices_pb.UpdatePrivateKeyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_UpdatePrivateKeyResponse(arg) {
  if (!(arg instanceof protos_devices_pb.UpdatePrivateKeyResponse)) {
    throw new Error('Expected argument of type UpdatePrivateKeyResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_UpdatePrivateKeyResponse(buffer_arg) {
  return protos_devices_pb.UpdatePrivateKeyResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var DevicesService = exports.DevicesService = {
  getDevices: {
    path: '/Devices/GetDevices',
    requestStream: false,
    responseStream: false,
    requestType: protos_devices_pb.GetDevicesRequest,
    responseType: protos_devices_pb.GetDevicesResponse,
    requestSerialize: serialize_GetDevicesRequest,
    requestDeserialize: deserialize_GetDevicesRequest,
    responseSerialize: serialize_GetDevicesResponse,
    responseDeserialize: deserialize_GetDevicesResponse,
  },
  updatePrivateKey: {
    path: '/Devices/UpdatePrivateKey',
    requestStream: false,
    responseStream: false,
    requestType: protos_devices_pb.UpdatePrivateKeyRequest,
    responseType: protos_devices_pb.UpdatePrivateKeyResponse,
    requestSerialize: serialize_UpdatePrivateKeyRequest,
    requestDeserialize: deserialize_UpdatePrivateKeyRequest,
    responseSerialize: serialize_UpdatePrivateKeyResponse,
    responseDeserialize: deserialize_UpdatePrivateKeyResponse,
  },
};

exports.DevicesClient = grpc.makeGenericClientConstructor(DevicesService);
