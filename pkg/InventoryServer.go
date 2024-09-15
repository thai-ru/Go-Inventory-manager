package pkg

import (
	"context"
	"errors"
	"inventory-manager/api/pb"
)

type server struct {
	pb.UnimplementedInventoryServiceServer
	items map[string]*pb.Item
}

func (s *server) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	id := generateID()
	req.Item.Id = id
	s.items[id] = req.Item
	return &pb.AddItemResponse{Id: id}, nil
}

func (s *server) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
	if _, exists := s.items[req.Item.Id]; !exists {
		return nil, errors.New("item not found")
	}
	s.items[req.Item.Id] = req.Item
	return &pb.UpdateItemResponse{Success: true}, nil
}

func (s *server) ListItems(ctx context.Context, req *pb.ListItemsRequest) (*pb.ListItemsResponse, error) {
	items := []*pb.Item{}
	for _, item := range s.items {
		items = append(items, item)
	}
	return &pb.ListItemsResponse{Items: items}, nil
}

func (s *server) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	if _, exists := s.items[req.Id]; !exists {
		return nil, errors.New("item not found")
	}
	delete(s.items, req.Id)
	return &pb.DeleteItemResponse{Success: true}, nil
}
