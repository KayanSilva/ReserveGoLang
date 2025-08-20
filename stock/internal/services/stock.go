package services

import (
	"fmt"
	"stock/internal/models"
	"strconv"
	"time"
)

type Stock struct {
	Items map[string]models.Item
	Logs  []models.Log
}

func NewStockService() *Stock {
	return &Stock{
		Items: make(map[string]models.Item),
		Logs:  []models.Log{},
	}
}

func (s *Stock) AddItem(item models.Item, user string) error {
	if _, exists := s.Items[strconv.Itoa(item.Id)]; exists {
		return fmt.Errorf("item %s already exists", item.Name)
	}
	if item.Quantity <= 0 {
		return fmt.Errorf("%s quantity cannot be zero or negative", item.Name)
	}

	existingItem, exists := s.Items[strconv.Itoa(item.Id)]
	if exists {
		existingItem.Quantity += item.Quantity
	}

	s.Items[strconv.Itoa(item.Id)] = item
	s.Logs = append(s.Logs, models.Log{
		Timestamp: time.Now().UTC(),
		Action:    "add",
		User:      user,
		ItemID:    item.Id,
		Quantity:  item.Quantity,
		Reason:    "Item added to stock",
	})
	return nil
}

func (s *Stock) GetItems() []models.Item {
	var itemList []models.Item
	for _, item := range s.Items {
		itemList = append(itemList, item)
	}

	return itemList
}

func (s *Stock) GetLogs() []models.Log {
	return s.Logs
}

func (s *Stock) CalculateTotalCost() float64 {
	var totalCost float64
	for _, item := range s.Items {
		totalCost += float64(item.Quantity) * item.Price
	}
	return totalCost
}

func (s *Stock) RemoveItem(itemID int, quantity int, user string) error {
	itemKey := strconv.Itoa(itemID)
	if item, exists := s.Items[itemKey]; exists {
		if item.Quantity < quantity {
			return fmt.Errorf("not enough stock for item %s", item.Name)
		}
		item.Quantity -= quantity
		if item.Quantity <= 0 {
			delete(s.Items, itemKey)
		} else {
			s.Items[itemKey] = item
		}
		s.Logs = append(s.Logs, models.Log{
			Timestamp: time.Now().UTC(),
			Action:    "remove",
			User:      user,
			ItemID:    itemID,
			Quantity:  quantity,
			Reason:    "Item removed from stock",
		})
		return nil
	}

	return fmt.Errorf("item with ID %d not found", itemID)
}

func FindBy[T any](data []T, comparator func(T) bool) ([]T, error) {
	var result []T
	for _, v := range data {
		if comparator(v) {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("not found any item")
	}
	return result, nil
}
