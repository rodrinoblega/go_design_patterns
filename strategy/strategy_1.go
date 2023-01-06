package main

import "fmt"

type PasswordProtector struct {
	user          string
	passwordNaame string
	hashAlgorithm HashAlgorithm
}

type HashAlgorithm interface {
	Hash(p *PasswordProtector)
}

func NewPasswordProtector(user string, passwordNaame string, hashAlgorithm HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		user:          user,
		passwordNaame: passwordNaame,
		hashAlgorithm: hashAlgorithm,
	}
}

func (p *PasswordProtector) setHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}

func (p *PasswordProtector) Hash() {
	p.hashAlgorithm.Hash(p)
}

type SHA struct{}

func (SHA) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using SHA for %s\n", p.passwordNaame)
}

type MD5 struct{}

func (MD5) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using MD5 for %s\n", p.passwordNaame)
}

func main() {
	sha := &SHA{}
	md5 := &MD5{}

	passwordProtector := NewPasswordProtector("rodrigo", "gmail password", sha)
	passwordProtector.Hash()
	passwordProtector.setHashAlgorithm(md5)
	passwordProtector.Hash()
}
