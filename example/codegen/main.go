package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/danielvladco/go-proto-gql/example/codegen/gql/constructs"
	"github.com/danielvladco/go-proto-gql/example/codegen/gql/options"
	"github.com/danielvladco/go-proto-gql/example/codegen/pb"
)

const defaultPort = "8088"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	constructsHandler := handler.NewDefaultServer(constructs.NewExecutableSchema(constructs.Config{
		Resolvers: constructsRoot{},
	}))

	optionsHandler := handler.NewDefaultServer(options.NewExecutableSchema(options.Config{
		Resolvers: optionsRoot{},
	}))
	http.Handle("/", playground.Handler("GraphQL playground", "/constructs-query"))
	http.Handle("/constructs-query", constructsHandler)
	http.Handle("/options-query", optionsHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// constructsRoot implements constructs.ResolverRoot showcasing that all resolvers were generated successfully
// we need a bit of binding like in example bellow and we are ready to go

type constructsRoot struct{}

func (r constructsRoot) Maps() constructs.MapsResolver             { return pb.MapsResolvers{} }
func (r constructsRoot) MapsInput() constructs.MapsInputResolver   { return pb.MapsInputResolvers{} }
func (r constructsRoot) OneofInput() constructs.OneofInputResolver { return pb.OneofInputResolvers{} }
func (r constructsRoot) Mutation() constructs.MutationResolver {
	return &pb.ConstructsResolvers{Service: pb.ConstructsServer(nil)}
}
func (r constructsRoot) Oneof() constructs.OneofResolver {
	return pb.OneofResolvers{}
}
func (r constructsRoot) Query() constructs.QueryResolver {
	return dummy{}
}

// dummy is generated when graphql schema doesn't have any query resolvers.
// In this case the dummy resolver will not be generated by the library and
// you should do it yourself like in example bellow
type dummy struct{}

func (d dummy) Dummy(ctx context.Context) (*bool, error) { panic("implement me") }

// optionsRoot implements options.ResolverRoot to showcase the generated resolvers
// as well as the missing ones. Some resolvers are missing because they use grpc streams
// which is too complex for graphql to deal with it by default.
//
// I might consider implementing it on the future. If you need this feature let me know by
// submitting an issue or if the issue already exists, show activity on it so I know there is real interest.
type optionsRoot struct{}

func (r optionsRoot) Mutation() options.MutationResolver {
	return &optionsMutationQueryResolver{
		ServiceResolvers: &pb.ServiceResolvers{Service: pb.ServiceServer(nil)},
		QueryResolvers:   &pb.QueryResolvers{Service: pb.QueryServer(nil)},
	}
}

func (r optionsRoot) Query() options.QueryResolver {
	return &optionsMutationQueryResolver{
		ServiceResolvers: &pb.ServiceResolvers{Service: pb.ServiceServer(nil)},
		QueryResolvers:   &pb.QueryResolvers{Service: pb.QueryServer(nil)},
	}
}

func (r optionsRoot) Subscription() options.SubscriptionResolver {
	return &optionsSubscriptionResolver{}
}

type optionsMutationQueryResolver struct {
	*pb.ServiceResolvers
	*pb.QueryResolvers
}

func (o optionsMutationQueryResolver) ServicePublish(ctx context.Context, in *pb.Data) (*pb.Data, error) {
	panic("implement me")
}
func (o optionsMutationQueryResolver) ServicePubSub1(ctx context.Context, in *pb.Data) (*pb.Data, error) {
	panic("implement me")
}
func (o optionsMutationQueryResolver) ServiceInvalidSubscribe3(ctx context.Context, in *pb.Data) (*pb.Data, error) {
	panic("implement me")
}
func (o optionsMutationQueryResolver) ServicePubSub2(ctx context.Context, in *pb.Data) (*pb.Data, error) {
	panic("implement me")
}
func (o optionsMutationQueryResolver) ServiceInvalidSubscribe1(ctx context.Context, in *pb.Data) (*pb.Data, error) {
	panic("implement me")
}

type optionsSubscriptionResolver struct{}

func (o optionsSubscriptionResolver) ServiceSubscribe(ctx context.Context, in *pb.Data) (<-chan *pb.Data, error) {
	panic("implement me")
}
func (o optionsSubscriptionResolver) ServicePubSub1(ctx context.Context, in *pb.Data) (<-chan *pb.Data, error) {
	panic("implement me")
}
func (o optionsSubscriptionResolver) ServiceInvalidSubscribe2(ctx context.Context, in *pb.Data) (<-chan *pb.Data, error) {
	panic("implement me")
}
func (o optionsSubscriptionResolver) ServiceInvalidSubscribe3(ctx context.Context, in *pb.Data) (<-chan *pb.Data, error) {
	panic("implement me")
}
func (o optionsSubscriptionResolver) ServicePubSub2(ctx context.Context, in *pb.Data) (<-chan *pb.Data, error) {
	panic("implement me")
}
func (o optionsSubscriptionResolver) QuerySubscribe(ctx context.Context, in *pb.Data) (<-chan *pb.Data, error) {
	panic("implement me")
}