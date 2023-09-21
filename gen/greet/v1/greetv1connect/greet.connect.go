// protoc-gen-connect-goによって生成されたコード。編集しないでください。
//
// Source: greet/v1/greet.proto

package greetv1connect

import (
	v1 "connect-getting-started/gen/greet/v1"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect "connectrpc.com/connect"
)

// これは、生成されたファイルとconnectパッケージの互換性を保証するためのコンパイル時のアサーションです。
// この定数が定義されていないというコンパイラエラーが発生した場合、このコードはバイナリにコンパイルされているconnectよりも新しいバージョンのconnectで生成されています。
// 古いバージョンのconnectでこのコードを再生成するか、バイナリにコンパイルされているconnectのバージョンを更新することで問題を解決できます。
const _ = connect.IsAtLeastVersion0_1_0

const (
	// GreetServiceNameは、GreetServiceサービスの完全修飾名です。
	GreetServiceName = "greet.v1.GreetService"
)

// これらの定数は、このパッケージで定義されている RPC の完全修飾名です。これらは実行時に Spec.Procedure として、また HTTP ルートの最後の 2 つのセグメントとして公開されます。
// これらは google.golang.org/protobuf/reflect/protoreflect で使われる完全修飾メソッド名とは異なることに注意してください。これらの定数から
// リフレクション形式のメソッド名に変換するには、先頭のスラッシュを削除し、残りのスラッシュをピリオドに変換します。
const (
	// GreetServiceGreetProcedureは、GreetServiceのGreet RPCの完全修飾名です。
	GreetServiceGreetProcedure = "/greet.v1.GreetService/Greet"
)

// GreetServiceClientは、greete.v1.GreetServiceサービスのクライアントです。
type GreetServiceClient interface {
	Greet(context.Context, *connect.Request[v1.GreetRequest]) (*connect.Response[v1.GreetResponse], error)
}

// NewGreetServiceClientはgreet.v1.GreetServiceサービスのクライアントを構築します。
// デフォルトでは、バイナリのProtobuf CodecでConnectプロトコルを使用し、gzip圧縮された応答を要求し、非圧縮のリクエストを送信します。
// gRPCまたはgRPC-Webプロトコルを使用するには、connect.WithGRPC()またはconnect.WithGRPCWeb()オプションを指定します。
// ここで指定するURLは、ConnectまたはgRPCサーバのベースURLでなければなりません（例えば、http://api.acme.comまたはhttps://acme.com/grpc）
func NewGreetServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GreetServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &greetServiceClient{
		greet: connect.NewClient[v1.GreetRequest, v1.GreetResponse](
			httpClient,
			baseURL+GreetServiceGreetProcedure,
			opts...,
		),
	}
}

// greetServiceClient は GreetServiceClient を実装しています。
type greetServiceClient struct {
	greet *connect.Client[v1.GreetRequest, v1.GreetResponse]
}

// Greetはgreet.v1.GreetService.Greetを呼び出します。
func (c *greetServiceClient) Greet(ctx context.Context, req *connect.Request[v1.GreetRequest]) (*connect.Response[v1.GreetResponse], error) {
	return c.greet.CallUnary(ctx, req)
}

// GreetServiceHandlerは、greet.v1.GreetServiceサービスの実装です。
type GreetServiceHandler interface {
	Greet(context.Context, *connect.Request[v1.GreetRequest]) (*connect.Response[v1.GreetResponse], error)
}

// NewGreetServiceHandlerはサービス実装からHTTPハンドラを構築します。ハンドラをマウントするパスとハンドラ自身を返します。
// デフォルトでは、ハンドラはバイナリの Protobuf および JSON コーデックで Connect、gRPC、および gRPC-Web プロトコルをサポートします。
// また、gzip 圧縮もサポートしています。
func NewGreetServiceHandler(svc GreetServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	greetServiceGreetHandler := connect.NewUnaryHandler(
		GreetServiceGreetProcedure,
		svc.Greet,
		opts...,
	)
	return "/greet.v1.GreetService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GreetServiceGreetProcedure:
			greetServiceGreetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGreetServiceHandlerは、すべてのメソッドからCodeUnimplementedを返します。
type UnimplementedGreetServiceHandler struct{}

func (UnimplementedGreetServiceHandler) Greet(context.Context, *connect.Request[v1.GreetRequest]) (*connect.Response[v1.GreetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("greet.v1.GreetService.Greet is not implemented"))
}
