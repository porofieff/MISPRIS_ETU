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
		AllowCredentials: false,
	}))

	router.Static("/css", "./frontend/css")
	router.Static("/js", "./frontend/js")
	router.StaticFile("/", "./frontend/index.html")

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", h.Login)
		}

		users := api.Group("/users")
		{
			users.GET("list", h.ListUsers)
			users.GET("/getUser:id", h.GetUser)
			users.POST("create", h.CreateUser)
			users.PUT("/update:id", h.UpdateUser)
			users.DELETE("/delete:id", h.DeleteUser)
		}

		emobile := api.Group("/emobile")
		{
			emobile.GET("list", h.ListEmobiles)
			emobile.GET("/getEmobile:id", h.GetEmobile)
			emobile.POST("create", h.CreateEmobile)
			emobile.PUT("/update:id", h.UpdateEmobile)
			emobile.DELETE("/delete:id", h.DeleteEmobile)
		}

		battery := api.Group("/battery")
		{
			battery.GET("list", h.ListBatteries)
			battery.GET("/getBattery:id", h.GetBattery)
			battery.POST("create", h.CreateBattery)
			battery.PUT("/update:id", h.UpdateBattery)
			battery.DELETE("/delete:id", h.DeleteBattery)
		}

		body := api.Group("/body")
		{
			body.GET("list", h.ListBodies)
			body.GET("/getBody:id", h.GetBody)
			body.POST("create", h.CreateBody)
			body.PUT("/update:id", h.UpdateBody)
			body.DELETE("/delete:id", h.DeleteBody)
		}

		carcass := api.Group("/carcass")
		{
			carcass.GET("list", h.ListCarcasses)
			carcass.GET("/getCarcass:id", h.GetCarcass)
			carcass.POST("create", h.CreateCarcass)
			carcass.PUT("/update:id", h.UpdateCarcass)
			carcass.DELETE("/delete:id", h.DeleteCarcass)
		}

		doors := api.Group("/doors")
		{
			doors.GET("list", h.ListDoors)
			doors.GET("/getDoors:id", h.GetDoors)
			doors.POST("create", h.CreateDoors)
			doors.PUT("/update:id", h.UpdateDoors)
			doors.DELETE("/delete:id", h.DeleteDoors)
		}

		wings := api.Group("/wings")
		{
			wings.GET("list", h.ListWings)
			wings.GET("/getWings:id", h.GetWings)
			wings.POST("create", h.CreateWings)
			wings.PUT("/update:id", h.UpdateWings)
			wings.DELETE("/delete:id", h.DeleteWings)
		}

		chargerSystem := api.Group("/charger-system")
		{
			chargerSystem.GET("list", h.ListChargerSystems)
			chargerSystem.GET("/getChargSystem:id", h.GetChargerSystem)
			chargerSystem.POST("create", h.CreateChargerSystem)
			chargerSystem.PUT("/update:id", h.UpdateChargerSystem)
			chargerSystem.DELETE("/delete:id", h.DeleteChargerSystem)
		}

		charger := api.Group("/charger")
		{
			charger.GET("list", h.ListChargers)
			charger.GET("/getCharger:id", h.GetCharger)
			charger.POST("create", h.CreateCharger)
			charger.PUT("/update:id", h.UpdateCharger)
			charger.DELETE("/delete:id", h.DeleteCharger)
		}

		connector := api.Group("/connector")
		{
			connector.GET("list", h.ListConnectors)
			connector.GET("/getConnector:id", h.GetConnector)
			connector.POST("create", h.CreateConnector)
			connector.PUT("/update:id", h.UpdateConnector)
			connector.DELETE("/delete:id", h.DeleteConnector)
		}

		chassis := api.Group("/chassis")
		{
			chassis.GET("list", h.ListChassis)
			chassis.GET("/getChassis:id", h.GetChassis)
			chassis.POST("create", h.CreateChassis)
			chassis.PUT("/update:id", h.UpdateChassis)
			chassis.DELETE("/delete:id", h.DeleteChassis)
		}

		frame := api.Group("/frame")
		{
			frame.GET("list", h.ListFrames)
			frame.GET("/getFrame:id", h.GetFrame)
			frame.POST("create", h.CreateFrame)
			frame.PUT("/update:id", h.UpdateFrame)
			frame.DELETE("/delete:id", h.DeleteFrame)
		}

		suspension := api.Group("/suspension")
		{
			suspension.GET("list", h.ListSuspensions)
			suspension.GET("/getSuspension:id", h.GetSuspension)
			suspension.POST("create", h.CreateSuspension)
			suspension.PUT("/update:id", h.UpdateSuspension)
			suspension.DELETE("/delete:id", h.DeleteSuspension)
		}

		breakSystem := api.Group("/break-system")
		{
			breakSystem.GET("list", h.ListBreakSystems)
			breakSystem.GET("/getBreakSystem:id", h.GetBreakSystem)
			breakSystem.POST("create", h.CreateBreakSystem)
			breakSystem.PUT("/update:id", h.UpdateBreakSystem)
			breakSystem.DELETE("/delete:id", h.DeleteBreakSystem)
		}

		electronics := api.Group("/electronics")
		{
			electronics.GET("list", h.ListElectronics)
			electronics.GET("/getElectronics:id", h.GetElectronics)
			electronics.POST("create", h.CreateElectronics)
			electronics.PUT("/update:id", h.UpdateElectronics)
			electronics.DELETE("/delete:id", h.DeleteElectronics)
		}

		controller := api.Group("/controller")
		{
			controller.GET("list", h.ListControllers)
			controller.GET("/getController:id", h.GetController)
			controller.POST("create", h.CreateController)
			controller.PUT("/update:id", h.UpdateController)
			controller.DELETE("/delete:id", h.DeleteController)
		}

		sensor := api.Group("/sensor")
		{
			sensor.GET("list", h.ListSensors)
			sensor.GET("/getSensor:id", h.GetSensor)
			sensor.POST("create", h.CreateSensor)
			sensor.PUT("/update:id", h.UpdateSensor)
			sensor.DELETE("/delete:id", h.DeleteSensor)
		}

		wiring := api.Group("/wiring")
		{
			wiring.GET("list", h.ListWirings)
			wiring.GET("/getWiring:id", h.GetWiring)
			wiring.POST("create", h.CreateWiring)
			wiring.PUT("/update:id", h.UpdateWiring)
			wiring.DELETE("/delete:id", h.DeleteWiring)
		}

		powerPoint := api.Group("/power-point")
		{
			powerPoint.GET("list", h.ListPowerPoints)
			powerPoint.GET("/getPowerPoint:id", h.GetPowerPoint)
			powerPoint.POST("create", h.CreatePowerPoint)
			powerPoint.PUT("/update:id", h.UpdatePowerPoint)
			powerPoint.DELETE("/delete:id", h.DeletePowerPoint)
		}

		engine := api.Group("/engine")
		{
			engine.GET("list", h.ListEngines)
			engine.GET("/getEngine:id", h.GetEngine)
			engine.POST("create", h.CreateEngine)
			engine.PUT("/update:id", h.UpdateEngine)
			engine.DELETE("/delete:id", h.DeleteEngine)
		}

		inverter := api.Group("/inverter")
		{
			inverter.GET("list", h.ListInverters)
			inverter.GET("/getInverter:id", h.GetInverter)
			inverter.POST("create", h.CreateInverter)
			inverter.PUT("/update:id", h.UpdateInverter)
			inverter.DELETE("/delete:id", h.DeleteInverter)
		}

		gearbox := api.Group("/gearbox")
		{
			gearbox.GET("list", h.ListGearboxes)
			gearbox.GET("/getGearbox:id", h.GetGearbox)
			gearbox.POST("create", h.CreateGearbox)
			gearbox.PUT("/update:id", h.UpdateGearbox)
			gearbox.DELETE("/delete:id", h.DeleteGearbox)
		}

		// ── ПР2: Перечисления ────────────────────────────────────────────

		enumClass := api.Group("/enum-class")
		{
			enumClass.GET("list", h.ListEnumClasses)
			enumClass.GET("/getEnumClass:id", h.GetEnumClass)
			enumClass.POST("create", h.CreateEnumClass)
			enumClass.PUT("/update:id", h.UpdateEnumClass)
			enumClass.DELETE("/delete:id", h.DeleteEnumClass)
			// Специальные методы ПР2:
			// Получить значения перечисления по порядку (SQL-функция get_enum_values)
			enumClass.GET("/values:id", h.GetEnumClassValues)
			// Проверить допустимость значения (SQL-функция validate_enum_value)
			enumClass.POST("validate", h.ValidateEnumValue)
		}

		enumPosition := api.Group("/enum-position")
		{
			enumPosition.GET("list", h.ListEnumPositions)
			enumPosition.GET("/getEnumPosition:id", h.GetEnumPosition)
			enumPosition.POST("create", h.CreateEnumPosition)
			enumPosition.PUT("/update:id", h.UpdateEnumPosition)
			enumPosition.DELETE("/delete:id", h.DeleteEnumPosition)
			// Изменить порядок позиции в перечислении
			enumPosition.POST("/reorder:id", h.ReorderEnumPosition)
		}

		// ── ПР3: Параметры изделий ──────────────────────────────────────

		parameter := api.Group("/parameter")
		{
			parameter.GET("list", h.ListParameters)
			parameter.GET("/getParameter:id", h.GetParameter)
			parameter.POST("create", h.CreateParameter)
			parameter.PUT("/update:id", h.UpdateParameter)
			parameter.DELETE("/delete:id", h.DeleteParameter)
		}

		componentParameter := api.Group("/component-parameter")
		{
			componentParameter.GET("list", h.ListComponentParameters)
			componentParameter.GET("/getComponentParameter:id", h.GetComponentParameter)
			componentParameter.POST("create", h.CreateComponentParameter)
			componentParameter.PUT("/update:id", h.UpdateComponentParameter)
			componentParameter.DELETE("/delete:id", h.DeleteComponentParameter)
			// Параметры компонента по типу (SQL-функция get_component_parameters)
			componentParameter.GET("/byType:type", h.GetComponentParametersByType)
			// Скопировать параметры от одного типа к другому (SQL-процедура copy_component_parameters)
			componentParameter.POST("copyFromType", h.CopyComponentParameters)
		}

		emobileParameter := api.Group("/emobile-parameter")
		{
			emobileParameter.GET("list", h.ListEmobileParameterValues)
			emobileParameter.GET("/getEmobileParameter:id", h.GetEmobileParameterValue)
			emobileParameter.POST("create", h.CreateEmobileParameterValue)
			emobileParameter.PUT("/update:id", h.UpdateEmobileParameterValue)
			emobileParameter.DELETE("/delete:id", h.DeleteEmobileParameterValue)
			// Все значения параметров конкретного автомобиля
			emobileParameter.GET("/byEmobile:id", h.GetEmobileParameterValuesByEmobile)
		}
	}

	router.GET("/health", h.Health)
	return router
}
