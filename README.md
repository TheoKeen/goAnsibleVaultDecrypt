Small tool to decrypt ansible-vault secured files without ansible-vault.


# Build instructions:

git clone https://github.com/TheoKeen/goAnsibleVaultDecrypt.git
go build -ldflags "-s -w" go_vdecrypt.go


# Usage

./go_vdecrypt ./filename "password"
