package httpserver

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func New(handler http.Handler, port string) *Server {
	//создаем экземпляр http.Server
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:         port,
	}

	//создаем наш экземпляр Server указывая httpServer, канал для оповещения об ошибке и таймаут завершения обработки запросов в случае закрытыя сервера
	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: 3 * time.Second,
	}

	//запускаем сервер
	s.start()

	//возвращаем экземпляр Server
	return s
}

func (s *Server) start() {
	go func() {
		//начинаем слушать порт, если ListenAndServe отдаст ошибку (например порт уже занят) - приложение будет завершено
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	//возвращаем канал в который может быть записана ошибка
	return s.notify
}

func (s *Server) Shutdown() error {
	//создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	//закрываем сервер указывая контекст с таймаутом обработки текущих запросов
	return s.server.Shutdown(ctx)
}
