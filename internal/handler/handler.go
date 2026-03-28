package handler

import (
	"MISPRIS/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", h.Login)
		}

		emobile := api.Group("/emobile")
		{
			emobile.GET("", h.ListEmobiles)
			emobile.GET("/:id", h.GetEmobile)
			emobile.POST("", h.CreateEmobile)
			emobile.PUT("/:id", h.UpdateEmobile)
			emobile.DELETE("/:id", h.DeleteEmobile)
		}

		battery := api.Group("/battery")
		{
			battery.GET("", h.ListBatteries)
			battery.GET("/:id", h.GetBattery)
			battery.POST("", h.CreateBattery)
			battery.PUT("/:id", h.UpdateBattery)
			battery.DELETE("/:id", h.DeleteBattery)
		}

		body := api.Group("/body")
		{
			body.GET("", h.ListBodies)
			body.GET("/:id", h.GetBody)
			body.POST("", h.CreateBody)
			body.PUT("/:id", h.UpdateBody)
			body.DELETE("/:id", h.DeleteBody)
		}

		carcass := api.Group("/carcass")
		{
			carcass.GET("", h.ListCarcasses)
			carcass.GET("/:id", h.GetCarcass)
			carcass.POST("", h.CreateCarcass)
			carcass.PUT("/:id", h.UpdateCarcass)
			carcass.DELETE("/:id", h.DeleteCarcass)
		}
		doors := api.Group("/doors")
		{
			doors.GET("", h.ListDoors)
			doors.GET("/:id", h.GetDoors)
			doors.POST("", h.CreateDoors)
			doors.PUT("/:id", h.UpdateDoors)
			doors.DELETE("/:id", h.DeleteDoors)
		}
		wings := api.Group("/wings")
		{
			wings.GET("", h.ListWings)
			wings.GET("/:id", h.GetWings)
			wings.POST("", h.CreateWings)
			wings.PUT("/:id", h.UpdateWings)
			wings.DELETE("/:id", h.DeleteWings)
		}

		chargerSystem := api.Group("/charger-system")
		{
			chargerSystem.GET("", h.ListChargerSystems)
			chargerSystem.GET("/:id", h.GetChargerSystem)
			chargerSystem.POST("", h.CreateChargerSystem)
			chargerSystem.PUT("/:id", h.UpdateChargerSystem)
			chargerSystem.DELETE("/:id", h.DeleteChargerSystem)
		}

		charger := api.Group("/charger")
		{
			charger.GET("", h.ListChargers)
			charger.GET("/:id", h.GetCharger)
			charger.POST("", h.CreateCharger)
			charger.PUT("/:id", h.UpdateCharger)
			charger.DELETE("/:id", h.DeleteCharger)
		}
		connector := api.Group("/connector")
		{
			connector.GET("", h.ListConnectors)
			connector.GET("/:id", h.GetConnector)
			connector.POST("", h.CreateConnector)
			connector.PUT("/:id", h.UpdateConnector)
			connector.DELETE("/:id", h.DeleteConnector)
		}

		chassis := api.Group("/chassis")
		{
			chassis.GET("", h.ListChassis)
			chassis.GET("/:id", h.GetChassis)
			chassis.POST("", h.CreateChassis)
			chassis.PUT("/:id", h.UpdateChassis)
			chassis.DELETE("/:id", h.DeleteChassis)
		}

		frame := api.Group("/frame")
		{
			frame.GET("", h.ListFrames)
			frame.GET("/:id", h.GetFrame)
			frame.POST("", h.CreateFrame)
			frame.PUT("/:id", h.UpdateFrame)
			frame.DELETE("/:id", h.DeleteFrame)
		}
		suspension := api.Group("/suspension")
		{
			suspension.GET("", h.ListSuspensions)
			suspension.GET("/:id", h.GetSuspension)
			suspension.POST("", h.CreateSuspension)
			suspension.PUT("/:id", h.UpdateSuspension)
			suspension.DELETE("/:id", h.DeleteSuspension)
		}
		breakSystem := api.Group("/break-system")
		{
			breakSystem.GET("", h.ListBreakSystems)
			breakSystem.GET("/:id", h.GetBreakSystem)
			breakSystem.POST("", h.CreateBreakSystem)
			breakSystem.PUT("/:id", h.UpdateBreakSystem)
			breakSystem.DELETE("/:id", h.DeleteBreakSystem)
		}

		electronics := api.Group("/electronics")
		{
			electronics.GET("", h.ListElectronics)
			electronics.GET("/:id", h.GetElectronics)
			electronics.POST("", h.CreateElectronics)
			electronics.PUT("/:id", h.UpdateElectronics)
			electronics.DELETE("/:id", h.DeleteElectronics)
		}

		controller := api.Group("/controller")
		{
			controller.GET("", h.ListControllers)
			controller.GET("/:id", h.GetController)
			controller.POST("", h.CreateController)
			controller.PUT("/:id", h.UpdateController)
			controller.DELETE("/:id", h.DeleteController)
		}
		sensor := api.Group("/sensor")
		{
			sensor.GET("", h.ListSensors)
			sensor.GET("/:id", h.GetSensor)
			sensor.POST("", h.CreateSensor)
			sensor.PUT("/:id", h.UpdateSensor)
			sensor.DELETE("/:id", h.DeleteSensor)
		}
		wiring := api.Group("/wiring")
		{
			wiring.GET("", h.ListWirings)
			wiring.GET("/:id", h.GetWiring)
			wiring.POST("", h.CreateWiring)
			wiring.PUT("/:id", h.UpdateWiring)
			wiring.DELETE("/:id", h.DeleteWiring)
		}

		powerPoint := api.Group("/power-point")
		{
			powerPoint.GET("", h.ListPowerPoints)
			powerPoint.GET("/:id", h.GetPowerPoint)
			powerPoint.POST("", h.CreatePowerPoint)
			powerPoint.PUT("/:id", h.UpdatePowerPoint)
			powerPoint.DELETE("/:id", h.DeletePowerPoint)
		}
		
		engine := api.Group("/engine")
		{
			engine.GET("", h.ListEngines)
			engine.GET("/:id", h.GetEngine)
			engine.POST("", h.CreateEngine)
			engine.PUT("/:id", h.UpdateEngine)
			engine.DELETE("/:id", h.DeleteEngine)
		}
		inverter := api.Group("/inverter")
		{
			inverter.GET("", h.ListInverters)
			inverter.GET("/:id", h.GetInverter)
			inverter.POST("", h.CreateInverter)
			inverter.PUT("/:id", h.UpdateInverter)
			inverter.DELETE("/:id", h.DeleteInverter)
		}
		gearbox := api.Group("/gearbox")
		{
			gearbox.GET("", h.ListGearboxes)
			gearbox.GET("/:id", h.GetGearbox)
			gearbox.POST("", h.CreateGearbox)
			gearbox.PUT("/:id", h.UpdateGearbox)
			gearbox.DELETE("/:id", h.DeleteGearbox)
		}
	}

	router.GET("/health", h.Health)
	return router
}
