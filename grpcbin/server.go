package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"strings"
	"time"

	pb "github.com/Anddd7/rubber-duck/grpcbin/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type ServeCmd struct {
}

func (cmd *ServeCmd) Run(globals *Globals) error {
	port := fmt.Sprintf(":%d", globals.Port)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		slog.Error("failed to listen", "err", err)
		return err
	}

	s := grpc.NewServer()
	pb.RegisterGrpcbinServiceServer(s, &server{})
	slog.Info("Server listening on", "addr", port)
	if err := s.Serve(lis); err != nil {
		slog.Error("failed to serve", "err", err)
		return err
	}

	return nil
}

type server struct {
	pb.UnimplementedGrpcbinServiceServer
}

func (s *server) Unary(ctx context.Context, req *pb.UnaryRequest) (*pb.UnaryResponse, error) {
	p, _ := peer.FromContext(ctx)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, "failed to get metadata")
	}

	headers := make(map[string]string)
	for key, values := range md {
		headers[key] = strings.Join(values, ",")
	}

	reqAttrs := req.RequestAttributes

	if reqAttrs.Delay > 0 {
		slog.Debug("Sleeping", "delay", req.RequestAttributes.Delay)
		time.Sleep(time.Duration(req.RequestAttributes.Delay) * time.Second)
	}

	if reqAttrs.HttpCode > 0 {
		slog.Debug("Returning HTTP status", "code", req.RequestAttributes.HttpCode)
		if reqAttrs.HttpCode > 400 {
			return nil, status.Error(codes.Code(req.RequestAttributes.HttpCode), "HTTP status code returned")
		}
	}

	if reqAttrs.ResponseHeaders != nil {
		for key, value := range reqAttrs.ResponseHeaders {
			grpc.SendHeader(ctx, metadata.Pairs(key, value))
		}
	}

	respAttrs := &pb.ResponseAttributes{
		RequesterIp:        p.Addr.String(),
		RequesterHost:      headers["host"],
		RequesterUserAgent: headers["user-agent"],
		RequestHeaders:     headers,
	}

	return &pb.UnaryResponse{Result: req.Data, ResponseAttributes: respAttrs}, nil
}

func (s *server) ServerStreaming(req *pb.ServerStreamingRequest, stream pb.GrpcbinService_ServerStreamingServer) error {
	// Your implementation here
	return nil
}

func (s *server) ClientStreaming(stream pb.GrpcbinService_ClientStreamingServer) error {
	// Your implementation here
	return nil
}

func (s *server) BidirectionalStreaming(stream pb.GrpcbinService_BidirectionalStreamingServer) error {
	// Your implementation here
	return nil
}
