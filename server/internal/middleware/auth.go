package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/JimMir45/KAR/server/internal/pkg/response"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func Auth(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "missing authorization header")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Unauthorized(c, "invalid authorization format")
			c.Abort()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(parts[1], claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			response.Unauthorized(c, "invalid or expired token")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func RequireRole(roles ...string) gin.HandlerFunc {
	roleSet := make(map[string]struct{}, len(roles))
	for _, r := range roles {
		roleSet[r] = struct{}{}
	}

	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			response.Unauthorized(c, "missing role")
			c.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			response.Unauthorized(c, "invalid role type")
			c.Abort()
			return
		}

		// super_admin can access everything
		if roleStr == "super_admin" {
			c.Next()
			return
		}

		if _, allowed := roleSet[roleStr]; !allowed {
			response.Forbidden(c, "insufficient permissions")
			c.Abort()
			return
		}

		c.Next()
	}
}

func GetUserID(c *gin.Context) uint {
	id, _ := c.Get("user_id")
	uid, _ := id.(uint)
	return uid
}

func GetRole(c *gin.Context) string {
	role, _ := c.Get("role")
	r, _ := role.(string)
	return r
}
