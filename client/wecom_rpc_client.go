/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a simple gRPC client that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It interacts with the route guide service whose definition can be found in routeguide/route_guide.booking.
package client

import (
	"context"
	"github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
	"reflect"
	"time"
)

const DEV = "dev"
const TEST = "test"
const PROD = "prod"

type RpcConfig struct {
	RpcServerPort      string
	RpcWecomHost       string
	RpcBusinessHost    string
	RpcTrackHost       string
	Tls                bool
	CertFile           string
	KeyFile            string
	CaFile             string
	ServerNameOverride string
}

var Mode string
var Config *RpcConfig

func init() {
	if args := os.Args; len(args) > 1 && args[1] == "prod" {
		Mode = PROD
	} else if args := os.Args; len(args) > 1 && args[1] == "test" {
		Mode = TEST
	} else {
		Mode = DEV
	}
	conf := new(RpcConfig)
	viper.AddConfigPath("./conf")
	viper.SetConfigName("config." + Mode)
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.Wrap(err, "failed on reading config file"))
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(errors.Wrap(err, "failed on unmarshal config file"))
	}
	Config = conf
}

func customInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// Add a custom value to the context
	ctx = metadata.AppendToOutgoingContext(ctx, "from", "kpi")
	// Perform any other logic before invoking the RPC method
	//global.Info.Info("Added custom value to context before calling method %s", method)
	// Call the RPC method with the updated context
	return invoker(ctx, method, req, reply, cc, opts...)
}

func InitWecomRpcClientConn() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if Config.Tls {
		creds, err := credentials.NewClientTLSFromFile(Config.CaFile, Config.ServerNameOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	opts = append(opts, grpc.WithUnaryInterceptor(customInterceptor))

	conn, err := grpc.NewClient(Config.RpcWecomHost, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func SetTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	var timeout time.Duration
	if Mode == "dev" {
		timeout = 114514 * time.Second
	} else {
		timeout = 10 * time.Second
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	return ctx, cancel
}

func InvokeWecomRPCMethod(ctx context.Context, methodName string, ao *wecom_rpc.RequestAO) (*wecom_rpc.ResponseVO, error) {
	conn, err := InitWecomRpcClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := wecom_rpc.NewWecomRPCClient(conn)

	ctx, cancel := SetTimeout(ctx)
	defer cancel()

	method := reflect.ValueOf(client).MethodByName(methodName)
	if !method.IsValid() {
		return nil, errors.New("method not found")
	}

	args := []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(ao)}
	result := method.Call(args)

	if result[0].IsValid() {
		vo, ok := result[0].Interface().(*wecom_rpc.ResponseVO)
		if !ok {
			return nil, errors.New("invalid response")
		}
		if result[1].IsValid() {
			err, _ := result[1].Interface().(error)
			if err != nil {
				return nil, errors.Wrap(err, "error")
			}
		}
		return vo, nil
	}
	return nil, errors.New("invalid response")
}
