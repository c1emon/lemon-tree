package authreq

import (
	"fmt"

	"github.com/c1emon/lemontree/pkg/cachex"
)

func codeKey(code string) string {
	return fmt.Sprintf("auth_code:%s", code)
}

func idKey(id string) string {
	return fmt.Sprintf("auth_id:%s", id)
}

func NewAuthCodeService(cacher cachex.Cacher) *AuthCodeService {
	return &AuthCodeService{cacher: cacher}
}

type AuthCodeService struct {
	cacher cachex.Cacher
}

func (s *AuthCodeService) Set(code, id string) {
	s.cacher.Set(codeKey(code), id)
	s.cacher.Set(idKey(id), code)
}

func (s *AuthCodeService) GetIdByCode(code string) (string, error) {

	val, ok := s.cacher.Get(codeKey(code))
	if !ok {
		return "", fmt.Errorf("invaild code %s", code)
	}
	if id, ok := val.(string); ok {
		return id, nil
	} else {
		return "", fmt.Errorf("invaild code %s type", code)
	}
}

func (s *AuthCodeService) GetCodeById(id string) (string, error) {

	val, ok := s.cacher.Get(idKey(id))
	if !ok {
		return "", fmt.Errorf("invaild id %s", id)
	}
	if id, ok := val.(string); ok {
		return id, nil
	} else {
		return "", fmt.Errorf("invaild id %s type", id)
	}
}

func (s *AuthCodeService) DelById(id string) {
	code, _ := s.GetCodeById(id)

	s.cacher.Del(idKey(id), codeKey(code))
}
