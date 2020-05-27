package app

import (
	"database/sql"
	dishhttp "github.com/efimovad/Vegan_delivery_API/internal/app/dish/delivery/http"
	dish_repo "github.com/efimovad/Vegan_delivery_API/internal/app/dish/repository"
	dishusecase "github.com/efimovad/Vegan_delivery_API/internal/app/dish/usecase"
	orderhttp "github.com/efimovad/Vegan_delivery_API/internal/app/order/delivery/http"
	order_repo "github.com/efimovad/Vegan_delivery_API/internal/app/order/repository"
	orderusecase "github.com/efimovad/Vegan_delivery_API/internal/app/order/usecase"
	placehttp "github.com/efimovad/Vegan_delivery_API/internal/app/place/delivery/http"
	"github.com/efimovad/Vegan_delivery_API/internal/app/place/repository"
	placeusecase "github.com/efimovad/Vegan_delivery_API/internal/app/place/usecase"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"strings"
)

type Server struct {
	SessionStore	sessions.Store
	Config			*Config
	Sanitizer		*bluemonday.Policy
	ServerEcho		*echo.Echo
}

func NewServer(config *Config) (*Server, error) {
	s := &Server{
		SessionStore: sessions.NewCookieStore([]byte(config.SessionKey)),
		Sanitizer:    bluemonday.UGCPolicy(),
		Config:       config,
	}
	return s, nil
}

func (s *Server) Configure() error {
	db, err := newDB(s.Config.DatabaseURL)
	if err != nil {
		return err
	}

	s.ServerEcho = echo.New()
	g := s.ServerEcho.Group("/api/v1")

	s.ServerEcho.Use(middleware.Logger())
	s.ServerEcho.Use(middleware.Recover())

	placeRepo := place_repo.NewPlaceRepository(db)
	placeUcase := placeusecase.NewPlaceUsecase(placeRepo)
	placehttp.NewHandler(g, placeUcase)

	dishRepo := dish_repo.NewDishRepository(db)
	dishUcase := dishusecase.NewDishUsecase(dishRepo)
	dishhttp.NewHandler(g, dishUcase)

	orderRepo := order_repo.NewOrderRepository(db)
	orderUcase := orderusecase.NewOrderUsecase(orderRepo)
	orderhttp.NewHandler(g, orderUcase)

	return nil
}

func Start(port string, dburl string) error {
	config := NewConfig(port, dburl)
	server, err := NewServer(config)
	if err != nil {
		return errors.Wrap(err, "creating server")
	}

	if err := server.Configure(); err != nil {
		return errors.Wrap(err, "configuring server")
	}

	log.Println("starting server at", config.BindAddr)
	return server.ServerEcho.Start(":"+port)
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(20)

	file, err := ioutil.ReadFile("./internal/database/sql/init_schema.sql")
	if err != nil {
		return nil, err
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err = db.Exec(request)
		if err != nil {
			log.Fatal(err)//return nil, err
		}
	}

	file, err = ioutil.ReadFile("./internal/database/sql/full_tables.sql")
	if err != nil {
		return nil, err
	}

	requests = strings.Split(string(file), ";")
	for _, request := range requests {
		_, err = db.Exec(request)
		if err != nil {
			log.Fatal(err)//return nil, err
		}
	}

	return db, nil
}