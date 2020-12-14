package server

import (
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
)

type Exec struct {
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

func (s *Server) execHandler(c echo.Context) error {
	req := new(Exec)
	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}

	cmd := exec.Command(req.Cmd, req.Args...)
	stdout, err := cmd.Output()

	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.String(http.StatusOK, string(stdout))
}
