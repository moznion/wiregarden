package net.moznion.wiregarden;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.38.0)",
    comments = "Source: protos/peers.proto")
public final class PeersGrpc {

  private PeersGrpc() {}

  public static final String SERVICE_NAME = "Peers";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<net.moznion.wiregarden.PeersProto.GetPeersRequest,
      net.moznion.wiregarden.PeersProto.GetPeersResponse> getGetPeersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPeers",
      requestType = net.moznion.wiregarden.PeersProto.GetPeersRequest.class,
      responseType = net.moznion.wiregarden.PeersProto.GetPeersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<net.moznion.wiregarden.PeersProto.GetPeersRequest,
      net.moznion.wiregarden.PeersProto.GetPeersResponse> getGetPeersMethod() {
    io.grpc.MethodDescriptor<net.moznion.wiregarden.PeersProto.GetPeersRequest, net.moznion.wiregarden.PeersProto.GetPeersResponse> getGetPeersMethod;
    if ((getGetPeersMethod = PeersGrpc.getGetPeersMethod) == null) {
      synchronized (PeersGrpc.class) {
        if ((getGetPeersMethod = PeersGrpc.getGetPeersMethod) == null) {
          PeersGrpc.getGetPeersMethod = getGetPeersMethod =
              io.grpc.MethodDescriptor.<net.moznion.wiregarden.PeersProto.GetPeersRequest, net.moznion.wiregarden.PeersProto.GetPeersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPeers"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  net.moznion.wiregarden.PeersProto.GetPeersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  net.moznion.wiregarden.PeersProto.GetPeersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PeersMethodDescriptorSupplier("GetPeers"))
              .build();
        }
      }
    }
    return getGetPeersMethod;
  }

  private static volatile io.grpc.MethodDescriptor<net.moznion.wiregarden.PeersProto.RegisterPeersRequest,
      net.moznion.wiregarden.PeersProto.RegisterPeersResponse> getRegisterPeersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RegisterPeers",
      requestType = net.moznion.wiregarden.PeersProto.RegisterPeersRequest.class,
      responseType = net.moznion.wiregarden.PeersProto.RegisterPeersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<net.moznion.wiregarden.PeersProto.RegisterPeersRequest,
      net.moznion.wiregarden.PeersProto.RegisterPeersResponse> getRegisterPeersMethod() {
    io.grpc.MethodDescriptor<net.moznion.wiregarden.PeersProto.RegisterPeersRequest, net.moznion.wiregarden.PeersProto.RegisterPeersResponse> getRegisterPeersMethod;
    if ((getRegisterPeersMethod = PeersGrpc.getRegisterPeersMethod) == null) {
      synchronized (PeersGrpc.class) {
        if ((getRegisterPeersMethod = PeersGrpc.getRegisterPeersMethod) == null) {
          PeersGrpc.getRegisterPeersMethod = getRegisterPeersMethod =
              io.grpc.MethodDescriptor.<net.moznion.wiregarden.PeersProto.RegisterPeersRequest, net.moznion.wiregarden.PeersProto.RegisterPeersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RegisterPeers"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  net.moznion.wiregarden.PeersProto.RegisterPeersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  net.moznion.wiregarden.PeersProto.RegisterPeersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PeersMethodDescriptorSupplier("RegisterPeers"))
              .build();
        }
      }
    }
    return getRegisterPeersMethod;
  }

  private static volatile io.grpc.MethodDescriptor<net.moznion.wiregarden.PeersProto.DeletePeersRequest,
      net.moznion.wiregarden.PeersProto.DeletePeersResponse> getDeletePeersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DeletePeers",
      requestType = net.moznion.wiregarden.PeersProto.DeletePeersRequest.class,
      responseType = net.moznion.wiregarden.PeersProto.DeletePeersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<net.moznion.wiregarden.PeersProto.DeletePeersRequest,
      net.moznion.wiregarden.PeersProto.DeletePeersResponse> getDeletePeersMethod() {
    io.grpc.MethodDescriptor<net.moznion.wiregarden.PeersProto.DeletePeersRequest, net.moznion.wiregarden.PeersProto.DeletePeersResponse> getDeletePeersMethod;
    if ((getDeletePeersMethod = PeersGrpc.getDeletePeersMethod) == null) {
      synchronized (PeersGrpc.class) {
        if ((getDeletePeersMethod = PeersGrpc.getDeletePeersMethod) == null) {
          PeersGrpc.getDeletePeersMethod = getDeletePeersMethod =
              io.grpc.MethodDescriptor.<net.moznion.wiregarden.PeersProto.DeletePeersRequest, net.moznion.wiregarden.PeersProto.DeletePeersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DeletePeers"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  net.moznion.wiregarden.PeersProto.DeletePeersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  net.moznion.wiregarden.PeersProto.DeletePeersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PeersMethodDescriptorSupplier("DeletePeers"))
              .build();
        }
      }
    }
    return getDeletePeersMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static PeersStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PeersStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PeersStub>() {
        @java.lang.Override
        public PeersStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PeersStub(channel, callOptions);
        }
      };
    return PeersStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static PeersBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PeersBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PeersBlockingStub>() {
        @java.lang.Override
        public PeersBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PeersBlockingStub(channel, callOptions);
        }
      };
    return PeersBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static PeersFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PeersFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PeersFutureStub>() {
        @java.lang.Override
        public PeersFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PeersFutureStub(channel, callOptions);
        }
      };
    return PeersFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class PeersImplBase implements io.grpc.BindableService {

    /**
     */
    public void getPeers(net.moznion.wiregarden.PeersProto.GetPeersRequest request,
        io.grpc.stub.StreamObserver<net.moznion.wiregarden.PeersProto.GetPeersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPeersMethod(), responseObserver);
    }

    /**
     */
    public void registerPeers(net.moznion.wiregarden.PeersProto.RegisterPeersRequest request,
        io.grpc.stub.StreamObserver<net.moznion.wiregarden.PeersProto.RegisterPeersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRegisterPeersMethod(), responseObserver);
    }

    /**
     */
    public void deletePeers(net.moznion.wiregarden.PeersProto.DeletePeersRequest request,
        io.grpc.stub.StreamObserver<net.moznion.wiregarden.PeersProto.DeletePeersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeletePeersMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetPeersMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                net.moznion.wiregarden.PeersProto.GetPeersRequest,
                net.moznion.wiregarden.PeersProto.GetPeersResponse>(
                  this, METHODID_GET_PEERS)))
          .addMethod(
            getRegisterPeersMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                net.moznion.wiregarden.PeersProto.RegisterPeersRequest,
                net.moznion.wiregarden.PeersProto.RegisterPeersResponse>(
                  this, METHODID_REGISTER_PEERS)))
          .addMethod(
            getDeletePeersMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                net.moznion.wiregarden.PeersProto.DeletePeersRequest,
                net.moznion.wiregarden.PeersProto.DeletePeersResponse>(
                  this, METHODID_DELETE_PEERS)))
          .build();
    }
  }

  /**
   */
  public static final class PeersStub extends io.grpc.stub.AbstractAsyncStub<PeersStub> {
    private PeersStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PeersStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PeersStub(channel, callOptions);
    }

    /**
     */
    public void getPeers(net.moznion.wiregarden.PeersProto.GetPeersRequest request,
        io.grpc.stub.StreamObserver<net.moznion.wiregarden.PeersProto.GetPeersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPeersMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void registerPeers(net.moznion.wiregarden.PeersProto.RegisterPeersRequest request,
        io.grpc.stub.StreamObserver<net.moznion.wiregarden.PeersProto.RegisterPeersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRegisterPeersMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deletePeers(net.moznion.wiregarden.PeersProto.DeletePeersRequest request,
        io.grpc.stub.StreamObserver<net.moznion.wiregarden.PeersProto.DeletePeersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeletePeersMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class PeersBlockingStub extends io.grpc.stub.AbstractBlockingStub<PeersBlockingStub> {
    private PeersBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PeersBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PeersBlockingStub(channel, callOptions);
    }

    /**
     */
    public net.moznion.wiregarden.PeersProto.GetPeersResponse getPeers(net.moznion.wiregarden.PeersProto.GetPeersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPeersMethod(), getCallOptions(), request);
    }

    /**
     */
    public net.moznion.wiregarden.PeersProto.RegisterPeersResponse registerPeers(net.moznion.wiregarden.PeersProto.RegisterPeersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRegisterPeersMethod(), getCallOptions(), request);
    }

    /**
     */
    public net.moznion.wiregarden.PeersProto.DeletePeersResponse deletePeers(net.moznion.wiregarden.PeersProto.DeletePeersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeletePeersMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class PeersFutureStub extends io.grpc.stub.AbstractFutureStub<PeersFutureStub> {
    private PeersFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PeersFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PeersFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<net.moznion.wiregarden.PeersProto.GetPeersResponse> getPeers(
        net.moznion.wiregarden.PeersProto.GetPeersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPeersMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<net.moznion.wiregarden.PeersProto.RegisterPeersResponse> registerPeers(
        net.moznion.wiregarden.PeersProto.RegisterPeersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRegisterPeersMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<net.moznion.wiregarden.PeersProto.DeletePeersResponse> deletePeers(
        net.moznion.wiregarden.PeersProto.DeletePeersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeletePeersMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_PEERS = 0;
  private static final int METHODID_REGISTER_PEERS = 1;
  private static final int METHODID_DELETE_PEERS = 2;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final PeersImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(PeersImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_PEERS:
          serviceImpl.getPeers((net.moznion.wiregarden.PeersProto.GetPeersRequest) request,
              (io.grpc.stub.StreamObserver<net.moznion.wiregarden.PeersProto.GetPeersResponse>) responseObserver);
          break;
        case METHODID_REGISTER_PEERS:
          serviceImpl.registerPeers((net.moznion.wiregarden.PeersProto.RegisterPeersRequest) request,
              (io.grpc.stub.StreamObserver<net.moznion.wiregarden.PeersProto.RegisterPeersResponse>) responseObserver);
          break;
        case METHODID_DELETE_PEERS:
          serviceImpl.deletePeers((net.moznion.wiregarden.PeersProto.DeletePeersRequest) request,
              (io.grpc.stub.StreamObserver<net.moznion.wiregarden.PeersProto.DeletePeersResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class PeersBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    PeersBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return net.moznion.wiregarden.PeersProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Peers");
    }
  }

  private static final class PeersFileDescriptorSupplier
      extends PeersBaseDescriptorSupplier {
    PeersFileDescriptorSupplier() {}
  }

  private static final class PeersMethodDescriptorSupplier
      extends PeersBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    PeersMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (PeersGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new PeersFileDescriptorSupplier())
              .addMethod(getGetPeersMethod())
              .addMethod(getRegisterPeersMethod())
              .addMethod(getDeletePeersMethod())
              .build();
        }
      }
    }
    return result;
  }
}
