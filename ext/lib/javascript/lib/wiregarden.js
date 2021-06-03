"use strict";

Object.defineProperty(exports, "__esModule", { value: true });

exports.DevicesPb = require('./protos/devices_pb');
exports.PeersPb = require('./protos/peers_pb');

const DevicesGrpcPb = require('./protos/devices_grpc_pb');
exports.DevicesService = DevicesGrpcPb.DevicesService;
exports.DevicesClient = DevicesGrpcPb.DevicesClient;

const PeersGrpcPb = require('./protos/peers_grpc_pb');
exports.PeersService = PeersGrpcPb.PeersService;
exports.PeersClient = PeersGrpcPb.PeersClient;

