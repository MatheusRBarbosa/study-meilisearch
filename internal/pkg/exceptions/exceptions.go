package exceptions

type DomainError struct {
	Code         int
	InternalCode string
	Message      string
	Description  string
}

var (
	DEFAULT = newError(500, "0001", "Erro interno.", "Um erro interno do servidor aconteceu. Muito triste =(")

	// 1xxx - Generic validations
	UNAUTHORIZED       = newError(401, "1001", "Credenciais Inválidas!", "As credenciais que você digitou são inválidas. Veja se digitou tudo corretamente e tente novamente.")
	FORBIDDEN          = newError(403, "1002", "Sem permissão!", "Seu acesso não possui permissão para realizar esta ação.")
	URL_PARAMS_MISSING = newError(500, "1003", "Parametros inválidos!", "Um ou mais parametros de URL estão faltando.")

	// 2xxx - Email
	EMAIL_ALREADY_EXIST = newError(400, "2001", "Email já está em uso.", "Caso já possua uma conta, tente recuperar a senha.")

	// 3xxx - User
	USER_NOT_FOUND = newError(404, "3001", "Não encontrado.", "O usuário procurado não foi encontrado.")

	// 4xxx - Validation code
	CODE_INVALID      = newError(400, "4001", "Código inválido.", "Verifique se os dados foram digitados corretamente e tente novamente.")
	CODE_EXPIRED      = newError(400, "4002", "Código expirado", "Este código não está mais válido.")
	CODE_ALREADY_USED = newError(400, "4003", "Código já utilizado", "Este código já foi usado e não pode ser usado novamente.")
)

func (e *DomainError) Error() string {
	return e.Message
}

func SlimError(code int, message string) *DomainError {
	return &DomainError{
		Message: message,
		Code:    code,
	}
}

func newError(code int, internalCode string, msg string, description string) *DomainError {
	return &DomainError{
		Message:      msg,
		Code:         code,
		InternalCode: internalCode,
		Description:  description,
	}
}

func GetError(e *DomainError) *DomainError {
	return &DomainError{
		Message:     e.Message,
		Code:        e.Code,
		Description: e.Description,
	}
}
