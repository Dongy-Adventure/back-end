package controller

import (
	"net/http"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAppointmentController interface {
	GetAppointments(c *gin.Context)
	GetAppointmentByID(c *gin.Context)
	GetAppointmentByOrderID(c *gin.Context)
	CreateAppointment(c *gin.Context)
	UpdateAppointmentDate(c *gin.Context)
	UpdateAppointmentPlace(c *gin.Context)
}

type AppointmentController struct {
	appointmentService service.IAppointmentService
}

func NewAppointmentController(s service.IAppointmentService) IAppointmentController {
	return AppointmentController{
		appointmentService: s,
	}
}

// GetAppointments godoc
//	@Summary		Get all appointments
//	@Description	Retrieves all appointments
//	@Tags			appointment
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.SuccessResponse{data=[]dto.Appointment}
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/appointment/ [get]
func (s AppointmentController) GetAppointments(c *gin.Context) {
	res, err := s.appointmentService.GetAppointments()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No appointments",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get Appointments success",
		Data:    res,
	})
}

// GetAppointmentsByID godoc
//	@Summary		Get a appointment by ID
//	@Description	Retrieves a appointment's data by its ID
//	@Tags			appointment
//	@Accept			json
//	@Produce		json
//	@Param			appointment_id	path		string	true	"Appointment ID"
//	@Success		200				{object}	dto.SuccessResponse{data=dto.Appointment}
//	@Failure		500				{object}	dto.ErrorResponse
//	@Router			/appointment/{appointment_id} [get]
func (s AppointmentController) GetAppointmentByID(c *gin.Context) {
	appointmentIDstr := c.Param("appointment_id")
	appointmentID, err := primitive.ObjectIDFromHex(appointmentIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Invalid appointmentID format",
			Message: err.Error(),
		})
		return
	}
	res, err := s.appointmentService.GetAppointmentByID(appointmentID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No appointment with this appointmentID",
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get appointment success",
		Data:    res,
	})
}

// GetAppointmentsByOrderID godoc
//	@Summary		Get appointments by orderID
//	@Description	Retrieves each order's appointment by order ID
//	@Tags			appointment
//	@Accept			json
//	@Produce		json
//	@Param			order_id	path		string	true	"Order ID"
//	@Success		200			{object}	dto.SuccessResponse{data=dto.Appointment}
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/appointment/order/{order_id} [get]
func (s AppointmentController) GetAppointmentByOrderID(c *gin.Context) {
	orderIDstr := c.Param("order_id")
	orderID, err := primitive.ObjectIDFromHex(orderIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Invalid orderID format",
			Message: err.Error(),
		})
		return
	}
	res, err := s.appointmentService.GetAppointmentByOrderID(orderID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No appointment with this orderID",
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get appointment success",
		Data:    res,
	})
}

// CreateAppointment godoc
//	@Summary		Create a new appointment
//	@Description	Creates a new appointment in the database
//	@Tags			appointment
//	@Accept			json
//	@Produce		json
//	@Param			appointment	body		dto.AppointmentCreateRequest	true	"Appointment to create"
//	@Success		201			{object}	dto.SuccessResponse{data=dto.Appointment}
//	@Failure		400			{object}	dto.ErrorResponse
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/appointment/ [post]
func (s AppointmentController) CreateAppointment(c *gin.Context) {
	var newAppointment model.Appointment

	if err := c.BindJSON(&newAppointment); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}

	res, err := s.appointmentService.CreateAppointment(&newAppointment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Failed to insert to database",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusCreated,
		Message: "Appointment created",
		Data:    res,
	})
}

// UpdateAppointment godoc
//	@Summary		Update an appointment date by ID
//	@Description	Updates an existing appointment's date by its ID
//	@Tags			appointment
//	@Accept			json
//	@Produce		json
//	@Param			appointment_id	path		string						true	"Appointment ID"
//	@Param			appointment		body		dto.AppointmentDateRequest	true	"Appointment date to update"
//	@Success		200				{object}	dto.SuccessResponse{data=dto.Appointment}
//	@Failure		400				{object}	dto.ErrorResponse
//	@Failure		500				{object}	dto.ErrorResponse
//	@Router			/appointment/{appointment_id}/date [put]
func (s AppointmentController) UpdateAppointmentDate(c *gin.Context) {
	appointmentIDstr := c.Param("appointment_id")
	appointmentID, err := primitive.ObjectIDFromHex(appointmentIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid appointmentID format",
			Message: err.Error(),
		})
		return
	}

	var dateRequest dto.AppointmentDateRequest
	if err := c.ShouldBindJSON(&dateRequest); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}

	// Parse the date from string to time.Time
	parsedDate, err := time.Parse("2006-01-02", dateRequest.Date) // Assuming input format is "YYYY-MM-DD"
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid date format, expected YYYY-MM-DD",
			Message: err.Error(),
		})
		return
	}

	updatedAppointment := model.Appointment{
		Date:     parsedDate,
		TimeSlot: dateRequest.TimeSlot,
	}

	res, err := s.appointmentService.UpdateAppointmentDate(appointmentID, &updatedAppointment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Failed to update appointment date",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Appointment date updated successfully",
		Data:    res,
	})
}


// UpdateAppointment godoc
//	@Summary		Update an appointment place by ID
//	@Description	Updates an existing appointment's place by its ID
//	@Tags			appointment
//	@Accept			json
//	@Produce		json
//	@Param			appointment_id	path		string						true	"Appointment ID"
//	@Param			appointment		body		dto.AppointmentPlaceRequest	true	"Appointment place to update"
//	@Success		200				{object}	dto.SuccessResponse{data=dto.Appointment}
//	@Failure		400				{object}	dto.ErrorResponse
//	@Failure		500				{object}	dto.ErrorResponse
//	@Router			/appointment/{appointment_id}/place [put]
func (s AppointmentController) UpdateAppointmentPlace(c *gin.Context) {
	appointmentIDstr := c.Param("appointment_id")
	appointmentID, err := primitive.ObjectIDFromHex(appointmentIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Invalid appointmentID format",
			Message: err.Error(),
		})
		return
	}
	var updatedAppointment model.Appointment
	if err := c.BindJSON(&updatedAppointment); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}

	res, err := s.appointmentService.UpdateAppointmentPlace(appointmentID, &updatedAppointment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Failed to update appointment place",
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Update appointment place success",
		Data:    res,
	})
}

