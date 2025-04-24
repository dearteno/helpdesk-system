package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dearteno/helpdesk-system/api-gateway/pkg/config"
	"github.com/dearteno/helpdesk-system/api-gateway/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set up Gin router
	router := gin.Default()

	// Add middleware
	router.Use(middleware.CORS())
	router.Use(middleware.SupabaseAuth(cfg.SupabaseURL, cfg.SupabaseKey))

	// Set up API routes
	setupRoutes(router, cfg)

	// Create server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Shutdown gracefully with timeout
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}

func setupRoutes(router *gin.Engine, cfg *config.Config) {
	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")

	// Helpdesk routes
	helpdeskRoutes := v1.Group("/helpdesk")
	helpdeskRoutes.GET("", middleware.ProxyService(cfg.HelpdeskServiceURL, "/helpdesk.service.HelpdeskService/ListHelpdesks"))
	helpdeskRoutes.GET("/:id", middleware.ProxyService(cfg.HelpdeskServiceURL, "/helpdesk.service.HelpdeskService/GetHelpdesk"))
	helpdeskRoutes.POST("", middleware.ProxyService(cfg.HelpdeskServiceURL, "/helpdesk.service.HelpdeskService/CreateHelpdesk"))
	helpdeskRoutes.PUT("/:id", middleware.ProxyService(cfg.HelpdeskServiceURL, "/helpdesk.service.HelpdeskService/UpdateHelpdesk"))
	helpdeskRoutes.DELETE("/:id", middleware.ProxyService(cfg.HelpdeskServiceURL, "/helpdesk.service.HelpdeskService/DeleteHelpdesk"))
	helpdeskRoutes.GET("/search", middleware.ProxyService(cfg.HelpdeskServiceURL, "/helpdesk.service.HelpdeskService/SearchHelpdesk"))

	// Ticket routes
	ticketRoutes := v1.Group("/ticket")
	ticketRoutes.GET("", middleware.ProxyService(cfg.TicketServiceURL, "/ticket.service.TicketService/ListTickets"))
	ticketRoutes.GET("/:id", middleware.ProxyService(cfg.TicketServiceURL, "/ticket.service.TicketService/GetTicket"))
	ticketRoutes.POST("", middleware.ProxyService(cfg.TicketServiceURL, "/ticket.service.TicketService/CreateTicket"))
	ticketRoutes.PUT("/:id", middleware.ProxyService(cfg.TicketServiceURL, "/ticket.service.TicketService/UpdateTicket"))
	ticketRoutes.DELETE("/:id", middleware.ProxyService(cfg.TicketServiceURL, "/ticket.service.TicketService/DeleteTicket"))
	ticketRoutes.GET("/search", middleware.ProxyService(cfg.TicketServiceURL, "/ticket.service.TicketService/SearchTicket"))

	// Issues routes
	issuesRoutes := v1.Group("/issues")
	issuesRoutes.GET("", middleware.ProxyService(cfg.IssuesServiceURL, "/issues.service.IssuesService/ListIssues"))
	issuesRoutes.GET("/:id", middleware.ProxyService(cfg.IssuesServiceURL, "/issues.service.IssuesService/GetIssue"))
	issuesRoutes.POST("", middleware.ProxyService(cfg.IssuesServiceURL, "/issues.service.IssuesService/CreateIssue"))
	issuesRoutes.PUT("/:id", middleware.ProxyService(cfg.IssuesServiceURL, "/issues.service.IssuesService/UpdateIssue"))
	issuesRoutes.DELETE("/:id", middleware.ProxyService(cfg.IssuesServiceURL, "/issues.service.IssuesService/DeleteIssue"))
	issuesRoutes.GET("/search", middleware.ProxyService(cfg.IssuesServiceURL, "/issues.service.IssuesService/SearchIssues"))

	// Track routes
	trackRoutes := v1.Group("/track")
	trackRoutes.GET("", middleware.ProxyService(cfg.TrackServiceURL, "/track.service.TrackService/ListTracks"))
	trackRoutes.GET("/:id", middleware.ProxyService(cfg.TrackServiceURL, "/track.service.TrackService/GetTrack"))
	trackRoutes.POST("", middleware.ProxyService(cfg.TrackServiceURL, "/track.service.TrackService/CreateTrack"))
	trackRoutes.PUT("/:id", middleware.ProxyService(cfg.TrackServiceURL, "/track.service.TrackService/UpdateTrack"))
	trackRoutes.DELETE("/:id", middleware.ProxyService(cfg.TrackServiceURL, "/track.service.TrackService/DeleteTrack"))
	trackRoutes.GET("/search", middleware.ProxyService(cfg.TrackServiceURL, "/track.service.TrackService/SearchTracks"))

	// Network routes
	networkRoutes := v1.Group("/network")
	networkRoutes.GET("", middleware.ProxyService(cfg.NetworkServiceURL, "/network.service.NetworkService/ListNetworks"))
	networkRoutes.GET("/:id", middleware.ProxyService(cfg.NetworkServiceURL, "/network.service.NetworkService/GetNetwork"))
	networkRoutes.POST("", middleware.ProxyService(cfg.NetworkServiceURL, "/network.service.NetworkService/CreateNetwork"))
	networkRoutes.PUT("/:id", middleware.ProxyService(cfg.NetworkServiceURL, "/network.service.NetworkService/UpdateNetwork"))
	networkRoutes.DELETE("/:id", middleware.ProxyService(cfg.NetworkServiceURL, "/network.service.NetworkService/DeleteNetwork"))
	networkRoutes.GET("/search", middleware.ProxyService(cfg.NetworkServiceURL, "/network.service.NetworkService/SearchNetworks"))

	// FAQ routes
	faqRoutes := v1.Group("/faq")
	faqRoutes.GET("", middleware.ProxyService(cfg.FAQServiceURL, "/faq.service.FAQService/ListFAQs"))
	faqRoutes.GET("/:id", middleware.ProxyService(cfg.FAQServiceURL, "/faq.service.FAQService/GetFAQ"))
	faqRoutes.POST("", middleware.ProxyService(cfg.FAQServiceURL, "/faq.service.FAQService/CreateFAQ"))
	faqRoutes.PUT("/:id", middleware.ProxyService(cfg.FAQServiceURL, "/faq.service.FAQService/UpdateFAQ"))
	faqRoutes.DELETE("/:id", middleware.ProxyService(cfg.FAQServiceURL, "/faq.service.FAQService/DeleteFAQ"))
	faqRoutes.GET("/search", middleware.ProxyService(cfg.FAQServiceURL, "/faq.service.FAQService/SearchFAQs"))
}
