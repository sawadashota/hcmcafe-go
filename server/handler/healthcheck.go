package handler

import "net/http"

type HealthCheck struct {
}

type PingArgs struct {
}

type PongReply struct {
	Message string `json:"message"`
}

// Ping is for health check
func (h *HealthCheck) Ping(r *http.Request, args *PingArgs, reply *PongReply) error {
	reply.Message = "pong!"
	return nil
}
