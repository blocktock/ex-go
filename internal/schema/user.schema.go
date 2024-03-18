// internal/http/dto/user.schema.go

package schema

// CreateUserRequestDTO 定义了创建用户请求的DTO
type CreateUserRequestDTO struct {
	Name string `json:"name"`
}
