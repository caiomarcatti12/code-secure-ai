package web

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/yourusername/dependency-check-automation/internal/application"
)

type Server struct {
    engine  *gin.Engine
    vulnSvc application.VulnerabilityService
}

func NewServer(vs application.VulnerabilityService) *Server {
    server := &Server{
        engine:  gin.Default(),
        vulnSvc: vs,
    }
    server.setupRoutes()
    return server
}

func (s *Server) setupRoutes() {
    s.engine.POST("/process", s.processReport)
}

func (s *Server) processReport(c *gin.Context) {
    var payload struct {
        ReportPath string `json:"report_path"`
    }

    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
        return
    }

    if err := s.vulnSvc.ProcessVulnerabilitiesAndCreatePR(payload.ReportPath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "PR created successfully"})
}

func (s *Server) Start() error {
    return s.engine.Run(":8080")
}

