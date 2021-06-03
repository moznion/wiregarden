package net.moznion.wiregarden;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.38.0)",
    comments = "Source: protos/devices.proto")
public final class DevicesGrpc {

  private DevicesGrpc() {}

  public static final String SERVICE_NAME = "Devices";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<net.moznion.wiregarden.DevicesProto.GetDevicesRequest,
      net.moznion.wiregarden.DevicesProto.GetDevicesResponse> getGetDevicesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetDevices",
      requestType = net.moznion.wiregarden.DevicesProto.GetDevicesRequest.class,
      responseType = net.moznion.wiregarden.DevicesProto.GetDevicesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<net.moznion.wiregarden.DevicesProto.GetDevicesRequest,
      net.moznion.wiregarden.DevicesProto.GetDevicesResponse> getGetDevicesMethod() {
    io.grpc.MethodDescriptor<net.moznion.wiregarden.DevicesProto.GetDevicesRequest, net.moznion.wiregarden.DevicesProto.GetDevicesResponse> getGetDevicesMethod;
    if ((getGetDevicesMethod = DevicesGrpc.getGetDevicesMethod) == null) {
      synchronized (DevicesGrpc.class) {
        if ((getGetDevicesMethod = DevicesGrpc.getGetDevicesMethod) == null) {
          DevicesGrpc.getGetDevicesMethod = getGetDevicesMethod =
              io.grpc.MethodDescriptor.<net.moznion.wiregarden.DevicesProto.GetDevicesRequest, net.moznion.wiregarden.DevicesProto.GetDevicesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetDevices"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  net.moznion.wiregarden.DevicesProto.GetDevicesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  net.moznion.wiregarden.DevicesProto.GetDevicesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new DevicesMethodDescriptorSupplier("GetDevices"))
              .build();
        }
      }
    }
    return getGetDevicesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyRequest,
      net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyResponse> getUpdatePrivateKeyMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdatePrivateKey",
      requestType = net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyRequest.class,
      responseType = net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyRequest,
      net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyResponse> getUpdatePrivateKeyMethod() {
    io.grpc.MethodDescriptor<net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyRequest, net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyResponse> getUpdatePrivateKeyMethod;
    if ((getUpdatePrivateKeyMethod = DevicesGrpc.getUpdatePrivateKeyMethod) == null) {
      synchronized (DevicesGrpc.class) {
        if ((getUpdatePrivateKeyMethod = DevicesGrpc.getUpdatePrivateKeyMethod) == null) {
          DevicesGrpc.getUpdatePrivateKeyMethod = getUpdatePrivateKeyMethod =
              io.grpc.MethodDescriptor.<net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyRequest, net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdatePrivateKey"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyResponse.getDefaultInstance()))
              .setSchemaDescriptor(new DevicesMethodDescriptorSupplier("UpdatePrivateKey"))
              .build();
        }
      }
    }
    return getUpdatePrivateKeyMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static DevicesStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DevicesStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DevicesStub>() {
        @java.lang.Override
        public DevicesStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DevicesStub(channel, callOptions);
        }
      };
    return DevicesStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static DevicesBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DevicesBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DevicesBlockingStub>() {
        @java.lang.Override
        public DevicesBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DevicesBlockingStub(channel, callOptions);
        }
      };
    return DevicesBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static DevicesFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DevicesFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DevicesFutureStub>() {
        @java.lang.Override
        public DevicesFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DevicesFutureStub(channel, callOptions);
        }
      };
    return DevicesFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class DevicesImplBase implements io.grpc.BindableService {

    /**
     */
    public void getDevices(net.moznion.wiregarden.DevicesProto.GetDevicesRequest request,
        io.grpc.stub.StreamObserver<net.moznion.wiregarden.DevicesProto.GetDevicesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetDevicesMethod(), responseObserver);
    }

    /**
     */
    public void updatePrivateKey(net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyRequest request,
        io.grpc.stub.StreamObserver<net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdatePrivateKeyMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetDevicesMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                net.moznion.wiregarden.DevicesProto.GetDevicesRequest,
                net.moznion.wiregarden.DevicesProto.GetDevicesResponse>(
                  this, METHODID_GET_DEVICES)))
          .addMethod(
            getUpdatePrivateKeyMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyRequest,
                net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyResponse>(
                  this, METHODID_UPDATE_PRIVATE_KEY)))
          .build();
    }
  }

  /**
   */
  public static final class DevicesStub extends io.grpc.stub.AbstractAsyncStub<DevicesStub> {
    private DevicesStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DevicesStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DevicesStub(channel, callOptions);
    }

    /**
     */
    public void getDevices(net.moznion.wiregarden.DevicesProto.GetDevicesRequest request,
        io.grpc.stub.StreamObserver<net.moznion.wiregarden.DevicesProto.GetDevicesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetDevicesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updatePrivateKey(net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyRequest request,
        io.grpc.stub.StreamObserver<net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdatePrivateKeyMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class DevicesBlockingStub extends io.grpc.stub.AbstractBlockingStub<DevicesBlockingStub> {
    private DevicesBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DevicesBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DevicesBlockingStub(channel, callOptions);
    }

    /**
     */
    public net.moznion.wiregarden.DevicesProto.GetDevicesResponse getDevices(net.moznion.wiregarden.DevicesProto.GetDevicesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetDevicesMethod(), getCallOptions(), request);
    }

    /**
     */
    public net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyResponse updatePrivateKey(net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdatePrivateKeyMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class DevicesFutureStub extends io.grpc.stub.AbstractFutureStub<DevicesFutureStub> {
    private DevicesFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DevicesFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DevicesFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<net.moznion.wiregarden.DevicesProto.GetDevicesResponse> getDevices(
        net.moznion.wiregarden.DevicesProto.GetDevicesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetDevicesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyResponse> updatePrivateKey(
        net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdatePrivateKeyMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_DEVICES = 0;
  private static final int METHODID_UPDATE_PRIVATE_KEY = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final DevicesImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(DevicesImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_DEVICES:
          serviceImpl.getDevices((net.moznion.wiregarden.DevicesProto.GetDevicesRequest) request,
              (io.grpc.stub.StreamObserver<net.moznion.wiregarden.DevicesProto.GetDevicesResponse>) responseObserver);
          break;
        case METHODID_UPDATE_PRIVATE_KEY:
          serviceImpl.updatePrivateKey((net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyRequest) request,
              (io.grpc.stub.StreamObserver<net.moznion.wiregarden.DevicesProto.UpdatePrivateKeyResponse>) responseObserver);
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

  private static abstract class DevicesBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    DevicesBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return net.moznion.wiregarden.DevicesProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Devices");
    }
  }

  private static final class DevicesFileDescriptorSupplier
      extends DevicesBaseDescriptorSupplier {
    DevicesFileDescriptorSupplier() {}
  }

  private static final class DevicesMethodDescriptorSupplier
      extends DevicesBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    DevicesMethodDescriptorSupplier(String methodName) {
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
      synchronized (DevicesGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new DevicesFileDescriptorSupplier())
              .addMethod(getGetDevicesMethod())
              .addMethod(getUpdatePrivateKeyMethod())
              .build();
        }
      }
    }
    return result;
  }
}
