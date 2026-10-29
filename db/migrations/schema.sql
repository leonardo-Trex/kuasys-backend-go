CREATE TABLE users (
    id     UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    keycloak_id VARCHAR(80) NOT NULL,
    email  TEXT UNIQUE NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

-- Com essa configuração eu consigo trabalhar com o cadastro simples e completo
-- PK e FK no mesmo campo para ter 1:1 entre o usuario e o seu perfil
CREATE TABLE user_profile (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    name TEXT NOT NULL,
    cpf VARCHAR(11) NOT NULL UNIQUE, -- TEXT? tu é doido fi
    phone TEXT NOT NULL
);

CREATE TABLE addresses (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    user_id UUID NOT NULL REFERENCES user_profile(user_id),

    cep VARCHAR(50) NOT NULL,
    number TEXT NOT NULL,
    street TEXT NOT NULL,
    complement TEXT,
    neighborhood TEXT NOT NULL,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    is_primary BOOLEAN NOT NULL DEFAULT FALSE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE

);

CREATE UNIQUE INDEX idx_addresses_user_primary
ON addresses (user_id)
WHERE is_primary = TRUE AND is_deleted = FALSE;
