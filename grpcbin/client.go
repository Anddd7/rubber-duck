package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"strings"

	pb "github.com/Anddd7/rubber-duck/grpcbin/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func connect(serverAddr string) (pb.GrpcbinServiceClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("failed to connect", "err", err)
		return nil, err
	}
	// defer conn.Close()
	client := pb.NewGrpcbinServiceClient(conn)

	return client, nil
}

type UnaryCmd struct {
	Message string
	pb.RequestAttributes
	Headers map[string]string
}

func (cmd *UnaryCmd) Run(globals *Globals) error {
	client, err := connect(fmt.Sprintf("%s:%d", globals.Host, globals.Port))
	if err != nil {
		return err
	}

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(cmd.Headers))
	response, err := client.Unary(ctx, &pb.UnaryRequest{
		Data:              cmd.Message,
		RequestAttributes: &cmd.RequestAttributes,
	})

	if err != nil {
		slog.Error("Unary RPC failed", "err", err)
		return err
	}

	printResponse(ctx, response.Result, response.ResponseAttributes)

	return nil
}

func printResponse(ctx context.Context, result string, respAttrs *pb.ResponseAttributes) {
	slog.Info(">>---->>---->>---->>---->>---->>")

	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		log.Printf("Failed to get metadata from response context")
	}

	slog.Info("Request Inspection")
	slog.Info("-", "requester_ip", respAttrs.RequesterIp)
	slog.Info("-", "requester_user_agent", respAttrs.RequesterUserAgent)
	slog.Info("-", "requester_host", respAttrs.RequesterHost)

	slog.Info("Request Headers")
	for key, value := range respAttrs.RequestHeaders {
		slog.Info("-", key, value)
	}
	slog.Info("<<----<<----<<----<<----<<----<<")

	slog.Info("Response Headers")
	for key, values := range md {
		slog.Info("-", key, strings.Join(values, ","))
	}

	slog.Info("Response Body")
	slog.Info("-", "response", result)
}

type ServerStreamingCmd struct {
	Message string
	pb.RequestAttributes
}

func (cmd *ServerStreamingCmd) Run(globals *Globals) error {
	client, err := connect(fmt.Sprintf("%s:%d", globals.Host, globals.Port))
	if err != nil {
		return err
	}

	stream, err := client.ServerStreaming(context.Background(), &pb.ServerStreamingRequest{
		Data:              cmd.Message,
		RequestAttributes: &cmd.RequestAttributes,
	})
	if err != nil {
		slog.Error("Server Streaming RPC failed", "err", err)
		return err
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			slog.Error("Server Streaming RPC stream closed", "err", err)
			return err
		}
		slog.Info("Server Streaming Response", "response", response)
	}
}

type ClientStreamingCmd struct {
	Messages []string
	pb.RequestAttributes
}

func (cmd *ClientStreamingCmd) Run(globals *Globals) error {
	client, err := connect(fmt.Sprintf("%s:%d", globals.Host, globals.Port))
	if err != nil {
		return err
	}

	stream, err := client.ClientStreaming(context.Background())
	if err != nil {
		slog.Error("Client Streaming RPC failed", "err", err)
		return err
	}
	for _, message := range cmd.Messages {
		err := stream.Send(&pb.ClientStreamingRequest{
			Data:              []string{message},
			RequestAttributes: &cmd.RequestAttributes,
		})
		if err != nil {
			slog.Error("Failed to send client streaming request", "err", err)
			return err
		}
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		slog.Error("Failed to receive client streaming response", "err", err)
		return err
	}
	slog.Info("Client Streaming Response", "response", response)

	return nil
}

type BidirectionalStreamingCmd struct {
	Messages []string
	pb.RequestAttributes
}

func (cmd *BidirectionalStreamingCmd) Run(globals *Globals) error {
	client, err := connect(fmt.Sprintf("%s:%d", globals.Host, globals.Port))
	if err != nil {
		return err
	}

	stream, err := client.BidirectionalStreaming(context.Background())
	if err != nil {
		slog.Error("Bidirectional Streaming RPC failed", "err", err)
		return err
	}
	for _, message := range cmd.Messages {
		err := stream.Send(&pb.BidirectionalStreamingRequest{
			Data:              []string{message},
			RequestAttributes: &cmd.RequestAttributes,
		})
		if err != nil {
			slog.Error("Failed to send bidirectional streaming request", "err", err)
			return err
		}
		response, err := stream.Recv()
		if err != nil {
			slog.Error("Failed to receive bidirectional streaming response", "err", err)
			return err
		}
		slog.Info("Bidirectional Streaming Response", "response", response)
	}

	return nil
}
