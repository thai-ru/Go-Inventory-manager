package pkg

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"inventory-manager/api/pb"
)

type Server struct {
	pb.UnimplementedInventoryServiceServer
	Items map[string]*pb.Item // Exported field
}

func (s *Server) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	id := uuid.New().String()
	req.Item.Id = id
	s.Items[id] = req.Item
	return &pb.AddItemResponse{Id: id}, nil
}

func (s *Server) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
	if _, exists := s.Items[req.Item.Id]; !exists {
		return nil, errors.New("item not found")
	}
	s.Items[req.Item.Id] = req.Item
	return &pb.UpdateItemResponse{Success: true}, nil
}

func (s *Server) ListItems(ctx context.Context, req *pb.ListItemsRequest) (*pb.ListItemsResponse, error) {
	items := []*pb.Item{}
	for _, item := range s.Items {
		items = append(items, item)
	}
	return &pb.ListItemsResponse{Items: items}, nil
}

func (s *Server) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	if _, exists := s.Items[req.Id]; !exists {
		return nil, errors.New("item not found")
	}
	delete(s.Items, req.Id)
	return &pb.DeleteItemResponse{Success: true}, nil
}
