package usecase

import (
	"context"
	"go-clean-architecture-pzn/internal/entity"
	"go-clean-architecture-pzn/internal/gateway/messaging"
	"go-clean-architecture-pzn/internal/model"
	"go-clean-architecture-pzn/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AddressUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	AddressRepository *repository.AddressRepository
	ContactRepository *repository.ContactRepository
	AddressProducer   *messaging.AddressProducer
}

func NewAddressUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, addressRepository *repository.AddressRepository, contactRepository *repository.ContactRepository, addressProducer *messaging.AddressProducer) *AddressUseCase {
	return &AddressUseCase{
		DB:                db,
		Log:               logger,
		Validate:          validate,
		AddressRepository: addressRepository,
		ContactRepository: contactRepository,
		AddressProducer:   addressProducer,
	}
}

func (c *AddressUseCase) Create(ctx context.Context, request *model.CreateAddressRequest) (*model.AddressResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	contact := new(entity.Contact)
	// if err := tx.Where("id = ? AND user_id = ?", request.ContactId, user.ID).First(contact).Error; err != nil {
	if err := c.ContactRepository.FindByIdAndUserId(tx, contact, request.ContactId, request.UserId); err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return nil, fiber.ErrNotFound
	}

	address := &entity.Address{
		ID:         uuid.NewString(),
		ContactId:  contact.ID,
		Street:     request.Street,
		City:       request.City,
		Province:   request.Province,
		PostalCode: request.PostalCode,
		Country:    request.Country,
	}

	// if err := tx.Create(address).Error; err != nil {
	if err := c.AddressRepository.Create(tx, address); err != nil {
		c.Log.WithError(err).Error("failed to create address")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	response := &model.AddressResponse{
		ID:         address.ID,
		Street:     address.Street,
		City:       address.City,
		Province:   address.Province,
		PostalCode: address.PostalCode,
		Country:    address.Country,
		CreatedAt:  address.CreatedAt,
		UpdatedAt:  address.UpdatedAt,
	}

	return response, nil
}

func (c *AddressUseCase) Update(ctx context.Context, request *model.UpdateAddressRequest) (*model.AddressResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	contact := new(entity.Contact)
	// if err := tx.Where("id = ? AND user_id = ?", request.ContactId, user.ID).First(contact).Error; err != nil {
	if err := c.ContactRepository.FindByIdAndUserId(tx, contact, request.ContactId, request.UserId); err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return nil, fiber.ErrNotFound
	}

	address := new(entity.Address)
	// if err := tx.Where("contact_id = ?", contact.ID).Find(address).Error; err != nil {
	if err := c.AddressRepository.FindByIdAndContactId(tx, address, request.ID, contact.ID); err != nil {
		c.Log.WithError(err).Error("failed to find addresses")
		return nil, fiber.ErrInternalServerError
	}

	address.Street = request.Street
	address.City = request.City
	address.Province = request.Province
	address.PostalCode = request.PostalCode
	address.Country = request.Country

	// if err := tx.Save(address).Error; err != nil {
	if err := c.AddressRepository.Update(tx, address); err != nil {
		c.Log.WithError(err).Error("failed to update address")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	response := &model.AddressResponse{
		ID:         address.ID,
		Street:     address.Street,
		City:       address.City,
		Province:   address.Province,
		PostalCode: address.PostalCode,
		Country:    address.Country,
		CreatedAt:  address.CreatedAt,
		UpdatedAt:  address.UpdatedAt,
	}

	return response, nil
}

func (c *AddressUseCase) Get(ctx context.Context, request *model.GetAddressRequest) (*model.AddressResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	contact := new(entity.Contact)
	// if err := tx.Where("id = ? AND user_id = ?", contactId, user.ID).First(contact).Error; err != nil {
	if err := c.ContactRepository.FindByIdAndUserId(tx, contact, request.ContactId, request.UserId); err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return nil, fiber.ErrNotFound
	}

	address := new(entity.Address)
	// if err := tx.Where("id = ? AND contact_id = ?", addressId, contactId).First(address).Error; err != nil {
	if err := c.AddressRepository.FindByIdAndContactId(tx, address, request.ID, request.ContactId); err != nil {
		c.Log.WithError(err).Error("failed to find address")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	response := &model.AddressResponse{
		ID:         address.ID,
		Street:     address.Street,
		City:       address.City,
		Province:   address.Province,
		PostalCode: address.PostalCode,
		Country:    address.Country,
		CreatedAt:  address.CreatedAt,
		UpdatedAt:  address.UpdatedAt,
	}

	return response, nil
}

func (c *AddressUseCase) Delete(ctx context.Context, request *model.DeleteAddressRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	contact := new(entity.Contact)
	// if err := tx.Where("id = ? AND user_id = ?", contactId, user.ID).First(contact).Error; err != nil {
	if err := c.ContactRepository.FindByIdAndUserId(tx, contact, request.ContactId, request.UserId); err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return fiber.ErrNotFound
	}

	address := new(entity.Address)
	// if err := tx.Where("id = ? AND contact_id = ?", addressId, contactId).First(address).Error; err != nil {
	if err := c.AddressRepository.FindByIdAndContactId(tx, address, request.ID, request.ContactId); err != nil {
		c.Log.WithError(err).Error("failed to find address")
		return fiber.ErrNotFound
	}

	// if err := tx.Delete(address).Error; err != nil {
	if err := c.AddressRepository.Delete(tx, address); err != nil {
		c.Log.WithError(err).Error("failed to delete address")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *AddressUseCase) List(ctx context.Context, request *model.ListAddressRequest) ([]model.AddressResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	contact := new(entity.Contact)
	// if err := tx.Where("id = ? AND user_id = ?", contactId, user.ID).First(contact).Error; err != nil {
	if err := c.ContactRepository.FindByIdAndUserId(tx, contact, request.ContactId, request.UserId); err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return nil, fiber.ErrNotFound
	}

	// var addresses []entity.Address
	// if err := tx.Where("contact_id = ?", contact.ID).Find(&addresses).Error; err != nil {
	addresses, err := c.AddressRepository.FindAllByContactId(tx, contact.ID)
	if err != nil {
		c.Log.WithError(err).Error("failed to find addresses")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.AddressResponse, len(addresses))
	for i, address := range addresses {
		responses[i] = model.AddressResponse{
			ID:         address.ID,
			Street:     address.Street,
			City:       address.City,
			Province:   address.Province,
			PostalCode: address.PostalCode,
			Country:    address.Country,
			CreatedAt:  address.CreatedAt,
			UpdatedAt:  address.UpdatedAt,
		}
	}

	return responses, nil
}
