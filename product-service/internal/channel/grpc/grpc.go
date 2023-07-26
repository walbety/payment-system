package grpc

import (
	"context"
	"github.com/walbety/payment-system/product-service/internal/channel/grpc/impl"
	"github.com/walbety/payment-system/product-service/internal/service"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

type productServiceGrpcServer struct {
	svc service.ProductService
	impl.UnimplementedProductServiceServer
}

func Listen() error {
	server := grpc.NewServer()
	port := 31102 // TODO config - adicionar porta grpc
	productGrpc := new(productServiceGrpcServer)
	productGrpc.svc = service.ProductService{} // TODO service - implementar a inicializacao

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		// TODO log - logar fatal com logrus
		return err
	}

	impl.RegisterProductServiceServer(server, productGrpc)

	// TODO grpc - adicionar logs ao implementar logs

	return server.Serve(listener)
}

func (r productServiceGrpcServer) ListProductByCategory(request *impl.CategoryTypeRequest) (*impl.ProductList, error) {
	// TODO adicionar passagem de ctx
	ctx := context.Background()

	products, err := r.svc.ListProductByCategory(request.CategoryType)
	if err != nil {
		// TODO log error
		return &impl.ProductList{}, err
	}
	return &impl.ProductList{ProductList: {impl.Product{}}}, err
}

// TODO error - implementar tratamento de constantes de erros
