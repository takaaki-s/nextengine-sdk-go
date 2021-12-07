package token

import (
	"context"
	"nextengine-sdk-go/entity"
)

type MemoryTokenRepository struct {
	t *entity.Token
}

func NewMemoryTokenRepository(accessToken, refreshToken string) *MemoryTokenRepository {
	return &MemoryTokenRepository{
		t: &entity.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (tr *MemoryTokenRepository) Token(_ context.Context) (*entity.Token, error) {
	tok := &entity.Token{
		AccessToken:         tr.t.AccessToken,
		RefreshToken:        tr.t.RefreshToken,
		AccessTokenEndDate:  tr.t.AccessTokenEndDate,
		RefreshTokenEndDate: tr.t.RefreshTokenEndDate,
	}

	return tok, nil
}

func (tr *MemoryTokenRepository) Save(_ context.Context, tok *entity.Token) error {
	tr.t.AccessToken = tok.AccessToken
	tr.t.RefreshToken = tok.RefreshToken
	tr.t.AccessTokenEndDate = tok.AccessTokenEndDate
	tr.t.RefreshTokenEndDate = tok.RefreshTokenEndDate

	return nil
}
